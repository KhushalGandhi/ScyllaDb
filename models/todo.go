package models

import "github.com/gocql/gocql"

type TODO struct {
	ID          gocql.UUID `json:"id"`
	UserID      gocql.UUID `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created     int64      `json:"created"`
	Updated     int64      `json:"updated"`
}
