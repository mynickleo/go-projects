package repository

import (
	"database/sql"
	"errors"
	"restApi/internal/models"
)

type UserRepository interface {
	GetUserByID(id int64) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) CreateUser(user *models.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}
