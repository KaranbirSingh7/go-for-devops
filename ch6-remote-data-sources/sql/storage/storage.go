package storage

import (
	"context"
	"database/sql"
)

type Storage struct {
	conn        *sql.DB
	getUserStmt *sql.Stmt //prepared statement
}

type UserRec struct {
	User        string
	ID          int
	DisplayName string
}

// create a new Storage object
func NewStorage(ctx context.Context, conn *sql.DB) *Storage {
	prepGetUserStmt, _ := conn.PrepareContext(ctx, `SELECT "User","DisplayName" FROM users WHERE "ID" = $1`)
	return &Storage{
		conn:        conn,
		getUserStmt: prepGetUserStmt,
	}
}

// GetUser get user by id
func (s *Storage) GetUser(ctx context.Context, id int) (UserRec, error) {
	u := UserRec{ID: 1}
	err := s.getUserStmt.QueryRow(id).Scan(&u)
	return u, err
}
