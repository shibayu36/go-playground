package repository

import (
	"github.com/Songmu/flextime"
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
	now := flextime.Now()
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

func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
