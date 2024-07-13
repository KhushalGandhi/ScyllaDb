package repositories

import (
	"github.com/gocql/gocql"
	"scylladb/db"
	"scylladb/models"
	"time"
)

type TODORepository struct{}

func (r *TODORepository) Create(todo *models.TODO) error {
	session, err := db.Cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	todo.ID = gocql.TimeUUID()
	todo.Created = time.Now().Unix()
	todo.Updated = time.Now().Unix()

	return session.Query(`INSERT INTO todos (id, user_id, title, description, status, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		todo.ID, todo.UserID, todo.Title, todo.Description, todo.Status, todo.Created, todo.Updated).Exec()
}

func (r *TODORepository) GetByID(userID, id gocql.UUID) (*models.TODO, error) {
	session, err := db.Cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var todo models.TODO
	if err := session.Query(`SELECT id, user_id, title, description, status, created, updated FROM todos WHERE user_id = ? AND id = ?`,
		userID, id).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TODORepository) List(userID gocql.UUID, status string, limit int, pageState []byte, sortBy string) ([]models.TODO, []byte, error) {
	session, err := db.Cluster.CreateSession()
	if err != nil {
		return nil, nil, err
	}
	defer session.Close()

	query := `SELECT id, user_id, title, description, status, created, updated FROM todos WHERE user_id = ?`
	var todos []models.TODO

	if status != "" {
		query += " AND status = ?"
	}

	if sortBy == "desc" {
		query += " ORDER BY created DESC"
	} else {
		query += " ORDER BY created ASC"
	}

	query += " LIMIT ?"

	iter := session.Query(query, userID, status, limit).PageState(pageState).Iter()

	var todo models.TODO
	for iter.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
		todos = append(todos, todo)
	}

	newPageState := iter.PageState()
	if err := iter.Close(); err != nil {
		return nil, nil, err
	}

	return todos, newPageState, nil
}

func (r *TODORepository) Update(todo *models.TODO) error {
	session, err := db.Cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	todo.Updated = time.Now().Unix()

	return session.Query(`UPDATE todos SET title = ?, description = ?, status = ?, updated = ? WHERE id = ? AND user_id = ?`,
		todo.Title, todo.Description, todo.Status, todo.Updated, todo.ID, todo.UserID).Exec()
}

func (r *TODORepository) Delete(userID, id gocql.UUID) error {
	session, err := db.Cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Query(`DELETE FROM todos WHERE id = ? AND user_id = ?`, id, userID).Exec()
}
