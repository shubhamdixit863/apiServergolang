package repositories

import "apiServer/internal/entities"

type Repository interface {
	CreateUser(user entities.User) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
}
