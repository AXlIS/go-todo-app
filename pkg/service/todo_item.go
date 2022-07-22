package service

import (
	"github.com/AXlIS/go-todo-app"
	"github.com/AXlIS/go-todo-app/pkg/repository"
)

type TodoItemService struct {
	repos     repository.TodoItem
	listRepos repository.TodoList
}

func NewTodoItemService(repos repository.TodoItem, listRepos repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repos:     repos,
		listRepos: listRepos,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepos.GetById(userId, listId)
	if err != nil {
		// list does not exist or does not belong to user
		return 0, err
	}
	return s.repos.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repos.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repos.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repos.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repos.Update(userId, itemId, input)
}
