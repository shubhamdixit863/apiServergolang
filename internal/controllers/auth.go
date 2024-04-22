package controllers

import (
	"github.com/gofiber/fiber/v3"

	"apiServer/internal/dto"
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
	// we will be receiving the json data
	var signupRequest dto.SignupRequest
	err := c.Bind().Body(&signupRequest)
	if err != nil {
		return err
	}

	id, err := ac.service.CreateUser(signupRequest)
	if err != nil {
		return c.JSON(dto.ErrorResponse(err, nil))
	}
	return c.JSON(dto.SuccessResponse("user created successfully", id))

}

func (ac *AuthController) ListUsers(c fiber.Ctx) error {
	// we will be receiving the json data

	users, err := ac.service.ListUsers()
	if err != nil {
		return c.JSON(dto.ErrorResponse(err, nil))
	}
	return c.JSON(dto.SuccessResponse("user created successfully", users))

}
