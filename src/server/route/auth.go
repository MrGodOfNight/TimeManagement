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
