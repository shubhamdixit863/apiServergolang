package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"apiServer/internal/dto"
	"apiServer/internal/entities"
	"apiServer/internal/repositories"
	"apiServer/internal/utils"
)

type AuthService interface {
	CreateUser(signupRequest dto.SignupRequest) (uint, error)
	ListUsers() ([]dto.UserData, error)
	GetUserByEmail(request dto.SignInRequest) (string, error)
	GetUserById(id string) (dto.UserData, error)
	UpdateUser(signupRequest dto.SignupRequest, id string) (uint, error)
	DeleteUser(id string) error
}

type authService struct {
	repository repositories.Repository
}

func (a authService) DeleteUser(id string) error {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = a.repository.DeleteUser(uint(atoi))
	if err != nil {
		return err
	}

	return nil
}

func (a authService) UpdateUser(signupRequest dto.SignupRequest, id string) (uint, error) {

	atoi, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	var updateUser entities.User
	updateUser.Name = signupRequest.Name
	updateUser.Email = signupRequest.Email
	updateUser.ID = uint(atoi)
	updateUser.UpdatedAt = time.Now()
	updateUser.CreatedAt = time.Now()
	password, err := utils.GeneratehashPassword(signupRequest.Password)
	if err != nil {
		return 0, err
	}
	updateUser.Password = password

	err = a.repository.UpdateUser(updateUser)
	if err != nil {
		return 0, err
	}
	return updateUser.ID, nil
}

func (a authService) GetUserByEmail(request dto.SignInRequest) (string, error) {
	var token string
	user, err := a.repository.GetUserByUserName(request.Email)
	if err != nil {
		return token, err
	}
	// Comparing the password for business logic will go here

	if utils.CheckPasswordHash(request.Password, user.Password) {
		// Passwords match
		jwt, err := utils.GenerateJWT(request.Email)
		if err != nil {
			return token, err
		}
		return jwt, nil
	}
	return token, errors.New("username or Password Wrong")
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
	password, err := utils.GeneratehashPassword(signupRequest.Password)
	if err != nil {
		return 0, err
	}
	user.Password = password
	// we will encrypt the password
	createUser, err := a.repository.CreateUser(user)
	if err != nil {
		return user.ID, fmt.Errorf("error creating the user %s", err)
	}

	return createUser.ID, nil

}

func NewAuthService(repository repositories.Repository) AuthService {
	return &authService{repository}
}

func (a authService) GetUserById(id string) (dto.UserData, error) {
	// type conversions
	var userData dto.UserData
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return userData, err
	}

	user, err := a.repository.GetUser(uint(atoi))
	if err != nil {
		return userData, err
	}

	log.Println(user)

	userData.Id = user.ID
	userData.Name = user.Name
	userData.Email = user.Email
	return userData, nil
}
