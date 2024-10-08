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

package route

import (
	"TimeManagement/src/server/model"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Structs for request body
type ReqBodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ReqBodyRegister struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	End         string `json:"end"`
	NewUsername string `json:"new_username"`
	Admin       *int64 `json:"admin,omitempty"`
}
type User struct {
	Login    string
	Password string
	End      string
	LastTime string
	Admin    int64
}

// Function for login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	var req ReqBodyLogin
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Check if user exists
	row := model.QueryRowSQL("SELECT password, admin FROM \"user\" WHERE login = $1", req.Username)
	// Get hash from database
	var hash, admin string
	if err := row.Scan(&hash, &admin); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(req.Password, hash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate token
	token, _, err := model.GenerateToken(req.Username)
	if err != nil {
		logger.Error(true, "Error generating token: \n", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	// Update last login time
	_, err = model.ExecSQL("UPDATE \"user\" SET last_time = $1 WHERE login = $2", time.Now().Format("2006-01-02 15:04:05"), req.Username)
	if err != nil {
		logger.Error(true, "Error updating last time: \n", err)
		http.Error(w, "Error updating last time", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"admin": admin,
	})
}

// Function for register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	var req ReqBodyRegister
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.Password == "" || req.End == "" || req.Admin == nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Get username from token
	username, err := model.GetUsernameFromToken(r.Header.Get("Token"))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Check if user exists and if user level is sufficient
	if admin, err := model.QueryValueSQL("SELECT admin FROM \"user\" WHERE login = $1", username); err != nil {
		logger.Error(true, "Error checking user level: \n", err)
		http.Error(w, "Error checking user level", http.StatusInternalServerError)
		return
	} else {
		if admin.(int64) < 1 || admin.(int64) <= *req.Admin {
			http.Error(w, "Insufficient user level", http.StatusUnauthorized)
			return
		}
	}

	// Hashing password
	hash, err := hashPassword(req.Password)
	if err != nil {
		logger.Error(true, "Error hashing password: \n", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	// Insert user into database
	model.ExecSQL("INSERT INTO \"user\" (login, password, end_date, admin) VALUES ($1, $2, $3, $4)", req.Username, hash, req.End, req.Admin)
}

func ChangeUserHandler(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	var req ReqBodyRegister
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Get username from token
	username, err := model.GetUsernameFromToken(r.Header.Get("Token"))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	row := model.QueryRowSQL("SELECT login, password, end_date, admin, last_time FROM \"user\" WHERE login = $1", req.Username)
	var user User
	err = row.Scan(&user.Login, &user.Password, &user.End, &user.Admin, &user.LastTime)
	if err != nil {
		logger.Error(true, "Error getting user: \n", err)
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	// Check if user exists and if user level is sufficient
	if admin, err := model.QueryValueSQL("SELECT admin FROM \"user\" WHERE login = $1", username); err != nil {
		logger.Error(true, "Error checking user level: \n", err)
		http.Error(w, "Error checking user level", http.StatusInternalServerError)
		return
	} else {
		if admin.(int64) < 1 {
			http.Error(w, "Insufficient user level", http.StatusUnauthorized)
			return
		}
		if req.Admin != nil && *req.Admin > admin.(int64) || user.Admin > admin.(int64) {
			http.Error(w, "Insufficient user level", http.StatusUnauthorized)
			return
		}
	}

	// Update admin level
	if req.Admin != nil {
		_, err = model.ExecSQL("UPDATE \"user\" SET admin = $1 WHERE login = $2", req.Admin, user.Login)
		if err != nil {
			logger.Error(true, "Error updating admin level: \n", err)
			http.Error(w, "Error updating admin level", http.StatusInternalServerError)
			return
		}
	}
	// Update end date
	if req.End != "" {
		_, err = model.ExecSQL("UPDATE \"user\" SET end_date = $1 WHERE login = $2", req.End, user.Login)
		if err != nil {
			logger.Error(true, "Error updating end date: \n", err)
			http.Error(w, "Error updating end date", http.StatusInternalServerError)
			return
		}
	}
	// Update password
	if req.Password != "" {
		hash, err := hashPassword(req.Password)
		if err != nil {
			logger.Error(true, "Error hashing password: \n", err)
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		_, err = model.ExecSQL("UPDATE \"user\" SET password = $1 WHERE login = $2", hash, user.Login)
		if err != nil {
			logger.Error(true, "Error updating password: \n", err)
			http.Error(w, "Error updating password", http.StatusInternalServerError)
			return
		}
	}
	// Update username
	if req.NewUsername != "" {
		_, err = model.ExecSQL("UPDATE \"user\" SET login = $1 WHERE login = $2", req.NewUsername, user.Login)
		if err != nil {
			logger.Error(true, "Error updating username: \n", err)
			http.Error(w, "Error updating username", http.StatusInternalServerError)
			return
		}
	}
	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func Admin(login string, password string) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	// Hashing password
	hash, err := hashPassword(password)
	if err != nil {
		logger.Error(true, "Error hashing password: \n", err)
		return
	}
	// Insert user into database
	ans, err := model.ExecSQL("INSERT INTO \"user\" (login, password, end_date, admin) VALUES ($1, $2, $3, $4)", login, hash, "01.01.2099", 9999)
	if err != nil {
		logger.Error(true, "Error inserting user: \n", err)
		logger.Error(true, "Answer: \n", ans)
		return
	}
	logger.Info(false, "User created", ans)
}

// Function for hashing password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Function for checking password hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
