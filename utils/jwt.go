package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// [TODO] fix later
const secretKey = "SUPER_SECRET_KEY"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
