package controllers

import (
	"github.com/gofiber/fiber/v3"

	"apiServer/internal/services"
)

type AuthController struct {
	// We can easily do dependency inject
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (ac *AuthController) Signup(c fiber.Ctx) error {
	return c.SendString(ac.service.CreateUser())
}
