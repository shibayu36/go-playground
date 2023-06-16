package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	User *UserRepository
}

func NewRepositories(dsn string) (*Repositories, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Opening mysql failed: %v", err)
	}
	return &Repositories{
		User: NewUserRepository(db),
	}, nil
}

func (r *Repositories) Close() error {
	return r.User.db.Close()
}
