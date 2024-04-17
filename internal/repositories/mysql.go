package repositories

import (
	"gorm.io/gorm"

	"apiServer/internal/entities"
)

type mysql struct {
	// It will have the Db connection object
	db *gorm.DB
}

func (m mysql) GetAllUsers() ([]entities.User, error) {

	var users []entities.User
	result := m.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (m mysql) CreateUser(user entities.User) (entities.User, error) {
	result := m.db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}

func NewMysql(db *gorm.DB) Repository {
	return &mysql{
		db: db,
	}
}
