package repositories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"

	"apiServer/internal/entities"
)

func DbConnect(username, password, host, port, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
	return db, err
}

func TestMysql_CreateUser(t *testing.T) {
	connect, err := DbConnect("avnadmin", "AVNS_ZT_X586MDd1cQPJueei", "mysql-3ca6f9a1-shubhamdixit863-a24d.aivencloud.com", "14287", "defaultdb")

	assert.Nil(t, err)
	assert.NotNil(t, connect)

	err = connect.AutoMigrate(&entities.User{})
	assert.Nil(t, err)

	// Here we will test our method
	mysqL := NewMysql(connect)
	userData := entities.User{
		Name:      "test",
		Email:     "local@email.com",
		Password:  "some randompasswr",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := mysqL.CreateUser(userData)
	assert.Nil(t, err)
	log.Println(user)
	assert.GreaterOrEqual(t, user.ID, uint(1))

}

func TestMysql_GetAllUsers(t *testing.T) {

	connect, err := DbConnect("avnadmin", "AVNS_ZT_X586MDd1cQPJueei", "mysql-3ca6f9a1-shubhamdixit863-a24d.aivencloud.com", "14287", "defaultdb")

	assert.Nil(t, err)
	assert.NotNil(t, connect)

	err = connect.AutoMigrate(&entities.User{})
	assert.Nil(t, err)

	mysqL := NewMysql(connect)

	users, err := mysqL.GetAllUsers()
	assert.Nil(t, err)
	log.Println(users)

	assert.Greater(t, len(users), 1)

}

func TestMysql_GetUser(t *testing.T) {

	connect, err := DbConnect("avnadmin", "AVNS_ZT_X586MDd1cQPJueei", "mysql-3ca6f9a1-shubhamdixit863-a24d.aivencloud.com", "14287", "defaultdb")

	assert.Nil(t, err)
	assert.NotNil(t, connect)

	err = connect.AutoMigrate(&entities.User{})
	assert.Nil(t, err)

	mysqL := NewMysql(connect)

	user, err := mysqL.GetUser(3)
	assert.Nil(t, err)
	assert.Equal(t, "shubham", user.Name)

}

func TestMysql_UpdateUser(t *testing.T) {
	connect, err := DbConnect("avnadmin", "AVNS_ZT_X586MDd1cQPJueei", "mysql-3ca6f9a1-shubhamdixit863-a24d.aivencloud.com", "14287", "defaultdb")

	assert.Nil(t, err)
	assert.NotNil(t, connect)

	err = connect.AutoMigrate(&entities.User{})
	assert.Nil(t, err)

	// Here we will test our method
	mysqL := NewMysql(connect)
	userData := entities.User{
		Name:      "john test",
		ID:        3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = mysqL.UpdateUser(userData)
	assert.Nil(t, err)

}

func TestMysql_DeleteUser(t *testing.T) {

	connect, err := DbConnect("avnadmin", "AVNS_ZT_X586MDd1cQPJueei", "mysql-3ca6f9a1-shubhamdixit863-a24d.aivencloud.com", "14287", "defaultdb")

	assert.Nil(t, err)
	assert.NotNil(t, connect)

	err = connect.AutoMigrate(&entities.User{})
	assert.Nil(t, err)

	mysqL := NewMysql(connect)

	err = mysqL.DeleteUser(3)
	assert.Nil(t, err)

}
