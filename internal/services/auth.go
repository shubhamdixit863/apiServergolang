package services

import "apiServer/internal/repositories"

type AuthService interface {
	CreateUser() string
}

type authService struct {
	repository repositories.Repository
}

func (a authService) CreateUser() string {
	return a.repository.CreateUser()
}

func NewAuthService(repository repositories.Repository) AuthService {
	return &authService{repository}
}
