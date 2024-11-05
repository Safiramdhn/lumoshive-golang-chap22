package services

import (
	"errors"
	"golang-beginner-22/models"
	"golang-beginner-22/repositories"
)

type UserService struct {
	UserRepo repositories.UserRepositoryDB
}

func NewUserService(userRepo repositories.UserRepositoryDB) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) LoginService(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password is required")
	}

	user, err := s.UserRepo.Login(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email or password is required")
	}

	if user.Name == "" {
		return nil, errors.New("name is required")
	}

	newUser, err := s.UserRepo.Create(*user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
