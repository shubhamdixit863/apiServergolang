package main

import (
	"github.com/gofiber/fiber/v3"
	"log"

	"apiServer/internal/controllers"
	"apiServer/internal/repositories"
	"apiServer/internal/services"
)

func main() {

	app := fiber.New()
	// We will crate the controller object
	mysql := repositories.NewMongodb()
	authService := services.NewAuthService(mysql)                // dependency injection
	authController := controllers.NewAuthController(authService) // dependency injection
	app.Get("/", authController.Signup)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("error starting the server")
	}

}
