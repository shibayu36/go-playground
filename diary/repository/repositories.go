package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	User *UserRepository

	db *sqlx.DB
}

func NewRepositories(dsn string) (*Repositories, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Opening mysql failed: %v", err)
	}
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

func (r *Repositories) Close() error {
	return r.db.Close()
}
