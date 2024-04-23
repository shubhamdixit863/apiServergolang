package repositories

import (
	"gorm.io/gorm"

	"apiServer/internal/entities"
)

type mysql struct {
	// It will have the Db connection object
	db *gorm.DB
}

func (m mysql) GetUserByUserName(userName string) (entities.User, error) {
	var user entities.User

	tx := m.db.First(&user, "email = ?", userName)
	if tx.Error != nil {
		return user, tx.Error
	}

	return user, nil
}

func (m mysql) GetUser(userId uint) (entities.User, error) {
	var user entities.User
	user.ID = userId
	tx := m.db.First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func (m mysql) UpdateUser(user entities.User) error {
	tx := m.db.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m mysql) DeleteUser(userId uint) error {
	tx := m.db.Delete(&entities.User{}, userId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

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

//https://github.com/robfig/cron
