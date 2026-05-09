package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = "secretkey12345"

func GenerateToken(email string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // This is check syntax
		if !ok {
			return nil, errors.New("Unexpected signing Method")
		}

		return secretKey, nil
	})

	if err != nil {
		return errors.New("Could not parsed token")
	}

	tokenIsvalid := parsedToken.Valid

	if !tokenIsvalid {
		return errors.New("Could not parsed token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("Invalid Token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int)

	return nil
}
