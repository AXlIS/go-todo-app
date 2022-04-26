package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/AXlIS/go-todo-app"
	"github.com/AXlIS/go-todo-app/pkg/repository"
)

const salt = "inpx98ere9r9asdf"

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
