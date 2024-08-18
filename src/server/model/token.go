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
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// Key for JWT
var jwtKey = []byte("asd@@!sd233asd1fWSDFASR#fasd235E@ds1dsQWE%%^GFd61>?L:?L")

// Function for generating JWT token
func GenerateToken(username string) (string, int64, error) {
	logger, err := NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return "", 0, err
	}
	minutes, err := strconv.Atoi(os.Getenv("TOKEN_TIME"))
	if err != nil {
		logger.Error(true, "Error converting token time: \n", err)
		return "", 0, err
	}
	// Set expiration time
	expirationTime := time.Now().Add(time.Duration(minutes) * time.Minute).Unix()

	// Create claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime,
		Subject:   username,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger.Error(true, "Error generating token: \n", err)
		return "", 0, err
	}

	return tokenString, expirationTime, nil
}

// Function for checking if token is valid
func CheckToken(tokenString string) bool {
	logger, err := NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return false
	}
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		logger.Error(true, "Error parsing token: \n", err)
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
	logger, err := NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return "", err
	}
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		logger.Error(true, "Error parsing token: \n", err)
		return "", err
	}

	// Get claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	return claims.Subject, nil
}
