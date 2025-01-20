package store

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Storage struct {
	Posts Posts
	Users Users
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
