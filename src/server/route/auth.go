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
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Struct for login and register
type ReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	End      string `json:"end"`
	Admin    int    `json:"admin"`
}

// Function for login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req ReqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	row := model.QueryRowSQL("SELECT password FROM users WHERE login = $1", req.Username)
	var hash string
	err := row.Scan(&hash)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(req.Password, hash) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, expirationTime, err := model.GenerateToken()
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":          token,
		"expirationTime": expirationTime,
	})
}

// Function for register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req ReqBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if admin, err := model.QueryValueSQL("SELECT admin FROM users WHERE login = $1", req.Username); err != nil {
		http.Error(w, "Error checking user level", http.StatusInternalServerError)
		return
	} else {
		if admin.(int) < 1 {
			http.Error(w, "Insufficient user level", http.StatusUnauthorized)
			return
		}
	}

	hash, err := hashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	model.ExecSQL("INSERT INTO users (login, password, end_date, admin) VALUES ($1, $2, $3, $4)", req.Username, hash, req.End, req.Admin)
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
