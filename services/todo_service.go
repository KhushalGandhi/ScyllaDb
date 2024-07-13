package services

import (
	"github.com/gocql/gocql"
	"scylladb/models"
	"scylladb/repositories"
)

type TODOService struct {
	Repository *repositories.TODORepository
}

func (s *TODOService) Create(todo *models.TODO) error {
	return s.Repository.Create(todo)
}

func (s *TODOService) GetByID(userID, id gocql.UUID) (*models.TODO, error) {
	return s.Repository.GetByID(userID, id)
}

func (s *TODOService) List(userID gocql.UUID, status string, limit int, pageState []byte, sortBy string) ([]models.TODO, []byte, error) {
	return s.Repository.List(userID, status, limit, pageState, sortBy)
}

func (s *TODOService) Update(todo *models.TODO) error {
	return s.Repository.Update(todo)
}

func (s *TODOService) Delete(userID, id gocql.UUID) error {
	return s.Repository.Delete(userID, id)
}
