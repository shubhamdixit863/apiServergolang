package services

import (
	"fmt"

	"apiServer/internal/dto"
	"apiServer/internal/entities"
	"apiServer/internal/repositories"
)

type AuthService interface {
	CreateUser(signupRequest dto.SignupRequest) (uint, error)
	ListUsers() ([]dto.UserData, error)
}

type authService struct {
	repository repositories.Repository
}

func (a authService) ListUsers() ([]dto.UserData, error) {
	users, err := a.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	// we will create a slice of signup request
	var signupR []dto.UserData

	for i := 0; i < len(users); i++ {
		sign := dto.UserData{
			SignupRequest: dto.SignupRequest{
				Name:     users[i].Name,
				Email:    users[i].Email,
				Password: users[i].Password,
			},
			Id: 0,
		}

		signupR = append(signupR, sign)

	}

	return signupR, nil
}

func (a authService) CreateUser(signupRequest dto.SignupRequest) (uint, error) {
	//return a.repository.CreateUser()
	// here will convert the dto to entity
	var user entities.User
	user.Name = signupRequest.Name
	user.Email = signupRequest.Email
	user.Password = signupRequest.Password
	createUser, err := a.repository.CreateUser(user)
	if err != nil {
		return user.ID, fmt.Errorf("error creating the user %s", err)
	}

	return createUser.ID, nil

}

func NewAuthService(repository repositories.Repository) AuthService {
	return &authService{repository}
}
