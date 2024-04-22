package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"apiServer/internal/controllers"
	"apiServer/internal/repositories"
	"apiServer/internal/services"
)

func DbConnect(username, password, host, port, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func Config() error {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../")    // path to look for the config file in
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func main() {

	err := Config()
	if err != nil {
		log.Fatal("error getting the config", err)
	}

	dbObject, err := DbConnect(viper.GetString("db.DB_USERNAME"), viper.GetString("db.DB_PASSWORD"),
		viper.GetString("db.DB_HOST"), viper.GetString("db.DB_PORT"), viper.GetString("db.DB_NAME"))

	if err != nil {
		log.Fatal("error connecting with db", err)

	}

	log.Println("Connected with the Database....")

	app := fiber.New()
	// We will crate the controller object
	mysql := repositories.NewMysql(dbObject)
	authService := services.NewAuthService(mysql)                // dependency injection
	authController := controllers.NewAuthController(authService) // dependency injection
	app.Post("/user", authController.Signup)

	app.Get("/user", authController.ListUsers)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("error starting the server")
	}

}
