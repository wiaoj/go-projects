package utils

import (
	"golang_projects/config"
	"golang_projects/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(config.GetJwtSecretKey())
var expirationMinutes = config.GetJwtExpirationMinutes()

func GenerateJwtToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationMinutes)).Unix()
	claims["level"] = 1
	claims["userId"] = user.ID

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
