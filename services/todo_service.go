package services

import (
	"errors"
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

func (s *TodoService) GetTodosByUserId(token string) ([]models.Todos, error) {
	return s.TodoRepo.GetTodos(token)
}

func (s *TodoService) GetTodoCount(token string) (*models.Todos, error) {
	return s.TodoRepo.GetCount(token)
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.TodoRepo.Delete(id)
}

func (s *TodoService) UpdateTodo(todo *models.Todos) (*models.Todos, error) {
	return s.TodoRepo.Update(todo)
}
