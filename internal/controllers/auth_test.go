package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"

	"apiServer/internal/dto"
	mock_services "apiServer/internal/services/mock"
)

func TestAuthController_ListUsers(t *testing.T) {

	// get a mock service object
	ctrl := gomock.NewController(t)
	svc := mock_services.NewMockAuthService(ctrl)
	var userData []dto.UserData
	userData = append(userData, dto.UserData{
		SignupRequest: dto.SignupRequest{
			Name:     "test",
			Email:    "test",
			Password: "test",
		},
		Id: 1,
	})
	// this is a mock call
	svc.EXPECT().ListUsers().Return(userData, nil)

	// we have to create a route

	app := fiber.New()

	ac := NewAuthController(svc)

	app.Get("/user", ac.ListUsers)

	// making the request to this url
	request, err := http.NewRequest(http.MethodGet, "/user", nil)
	assert.Nil(t, err)

	// so this will execute request object and gets us the response
	resp, err := app.Test(request, -1)
	assert.Nil(t, err)
	all, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	log.Println("response----", string(all))
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestAuthController_Signup(t *testing.T) {
	// get a mock service object
	ctrl := gomock.NewController(t)
	svc := mock_services.NewMockAuthService(ctrl)

	sr := dto.SignupRequest{
		Name:     "shubh",
		Email:    "test",
		Password: "test",
	}

	// Anytimes will make sure that the CreateUser method gets called any number of times required
	svc.EXPECT().CreateUser(gomock.Any()).Return(uint(1), nil).AnyTimes()

	app := fiber.New()

	ac := NewAuthController(svc)

	app.Post("/user", ac.Signup)

	marshal, err := json.Marshal(sr)
	assert.Nil(t, err)

	// making the request to this url
	request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(marshal))
	request.Header.Set("Content-Type", "application/json")

	assert.Nil(t, err)

	// so this will execute request object and gets us the response
	resp, err := app.Test(request, -1)
	assert.Nil(t, err)
	all, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	log.Println("response----", string(all))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
