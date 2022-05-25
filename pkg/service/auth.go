package service

import (
	"crypto/sha1"
	"fmt"
	"todo-app"
	"todo-app/pkg/repository"
)

const salt = "10983gnuyc197chf183sgu109v2k4yu98gbc23yb0munyc2h3jdf"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
