package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// Function for generating JWT token
func GenerateToken() (string, int64, error) {
	expirationTime := time.Now().Add(12 * time.Hour).Unix()

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

// Function for checking if token is valid
func CheckToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return false
	}

	expirationTime := claims.ExpiresAt
	return time.Now().Unix() < expirationTime
}
