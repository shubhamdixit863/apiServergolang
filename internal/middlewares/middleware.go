package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	"apiServer/internal/dto"
)

func AuthorizedRequest() fiber.Handler {
	return func(c fiber.Ctx) error {
		// we have to get the token from the header

		mp := c.GetReqHeaders()
		token := mp["Authorization"][0]

		// validate this token

		var mySigningKey = []byte(viper.GetString("secretKey"))

		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			return c.JSON(dto.ErrorResponse("invalid token", err.Error()))
		}

		/*
			if _, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {

				// check for role based access and all
				return nil
			}

		*/

		return c.Next()

	}
}
