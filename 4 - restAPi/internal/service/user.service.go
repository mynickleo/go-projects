package service

import (
	"errors"
	"restApi/internal/models"
	"restApi/internal/repository"
)

type UserService interface {
	GetUser(id int64) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUser(id int64) (*models.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("invalid input data")
	}
	return s.repo.CreateUser(user)
}
