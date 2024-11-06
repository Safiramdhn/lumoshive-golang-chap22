package services

import (
	"errors"
	// "fmt"
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

func (s *UserService) CreateUser(userInput models.User) (*models.User, error) {
	if userInput.Email == "" || userInput.Password == "" {
		return nil, errors.New("email or password is required")
	}

	if userInput.Name == "" {
		return nil, errors.New("name is required")
	}

	newUser, err := s.UserRepo.Create(userInput)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.UserRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := s.UserRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByToken(token string) (string, error) {
	userToken, err := s.UserRepo.GetByToken(token)
	if err != nil {
		return "", err
	}
	return userToken, nil
}
