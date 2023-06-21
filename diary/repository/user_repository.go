package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shibayu36/go-playground/diary/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(email string, name string) (*model.User, error) {
	// TODO: Modify how to create now for testability
	now := time.Now()
	res, err := r.db.Exec(
		`INSERT INTO users (email, name, created_at, updated_at)
			VALUES (?, ?, ?, ?)`,
		email, name, now, now,
	)
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	user := &model.User{
		UserID:    id,
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return user, nil
}
