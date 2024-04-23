package repositories

import "apiServer/internal/entities"

type Repository interface {
	CreateUser(user entities.User) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUser(userId uint) (entities.User, error)
	GetUserByUserName(userName string) (entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(userId uint) error
}
