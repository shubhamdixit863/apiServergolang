package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthController_ListUsers(t *testing.T) {
	app := fiber.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}
