package service

import (
	"github.com/AXlIS/go-todo-app"
	"github.com/AXlIS/go-todo-app/pkg/repository"
)

type TodoListService struct {
	repos repository.TodoList
}

func NewTodoListService(repos repository.TodoList) *TodoListService {
	return &TodoListService{
		repos: repos,
	}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repos.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repos.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repos.GetById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repos.Update(userId, listId, input)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repos.Delete(userId, listId)
}
