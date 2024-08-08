/*
	MIT License

	Copyright (c) 2024 Ushakov Igor

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

*/

package model

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Key for JWT
var jwtKey = []byte("my_secret_key")

// Function for generating JWT token
func GenerateToken(username string) (string, int64, error) {
	// Set expiration time (12 hours)
	expirationTime := time.Now().Add(12 * time.Hour).Unix()

	// Create claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime,
		Subject:   username,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

// Function for checking if token is valid
func CheckToken(tokenString string) bool {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false
	}

	// Get claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return false
	}

	// Get expiration time and check if token is valid
	expirationTime := claims.ExpiresAt
	return time.Now().Unix() < expirationTime
}

func GetUsernameFromToken(tokenString string) (string, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	// Get claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	return claims.Subject, nil
}
