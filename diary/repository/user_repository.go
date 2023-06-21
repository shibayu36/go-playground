package repository

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(email string, name string) error {
	// TODO: Modify how to create now for testability
	now := time.Now()
	res, err := r.db.Exec(
		`INSERT INTO users (email, name, created_at, updated_at)
			VALUES (?, ?, ?, ?)`,
		email, name, now, now,
	)
	if err != nil {
		return err
	}

	// TODO: Return *model.User

	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
	return err
}
