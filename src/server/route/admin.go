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
	"fmt"
	"log"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	// Get username from token
	username, err := model.GetUsernameFromToken(r.Header.Get("Token"))
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	row, err := model.QueryValueSQL(fmt.Sprintf(`SELECT "admin" FROM "user" WHERE login = '%s'`, username))
	if err != nil {
		logger.Error(true, "Error checking user level: \n", err)
		http.Error(w, "Error checking user level", http.StatusInternalServerError)
		return
	}
	if row.(int64) < 1 {
		http.Error(w, "Insufficient user level", http.StatusUnauthorized)
		return
	}
	// Get all users
	rows, err := model.QuerySQL(`SELECT login, password, end_date, last_time, "admin" FROM "user"`)
	if err != nil {
		logger.Error(true, "Error getting users: \n", err)
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	var data []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Login, &user.Password, &user.End, &user.LastTime, &user.Admin)
		if err != nil {
			logger.Error(true, "Error scanning row: \n", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		data = append(data, user)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error(true, "Error marshalling JSON: \n", err)
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
