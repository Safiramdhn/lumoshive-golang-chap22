package services

import (
	"errors"
	"fmt"
	"golang-beginner-22/models"
	"golang-beginner-22/repositories"
)

type TodoService struct {
	TodoRepo repositories.TodoRepositoryDB
}

func NewTodoService(todoRepo repositories.TodoRepositoryDB) *TodoService {
	return &TodoService{TodoRepo: todoRepo}
}

func (s *TodoService) CreateTodo(todo *models.Todos, token string) (*models.Todos, error) {
	if todo.Description == "" {
		return nil, errors.New("description is required")
	}

	newTodo, err := s.TodoRepo.Create(todo, token)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

func (s *TodoService) GetTodosByToken(token string) ([]models.Todos, error) {
	todos, err := s.TodoRepo.GetTodos(token)
	if err != nil {
		return nil, err
	}
	fmt.Printf("todos %v\n", todos)
	return todos, nil
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.TodoRepo.Delete(id)
}

func (s *TodoService) UpdateTodo(id int, newStatus string) error {
	return s.TodoRepo.Update(id, newStatus)
}
