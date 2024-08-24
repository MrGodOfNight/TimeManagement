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
)

// Struct for work and break
type ReqStartWorkBody struct {
	Time string `json:"time"`
}
type ReqStopWorkBody struct {
	Time string `json:"time"`
	ID   int    `json:"id"`
}

// Struct for report
type ReqReportBody struct {
	Text string `json:"text"`
	ID   int    `json:"id"`
}

func StartWork(w http.ResponseWriter, r *http.Request) {
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
	var req ReqStartWorkBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Write request in console
	logger.Debug(false, req)
	row, err := model.QueryValueSQL(`INSERT INTO work_time (start_time, user_id)
	VALUES ($1, (SELECT id FROM "user" WHERE login = $2)) RETURNING id`, req.Time, username)
	if err != nil {
		logger.Error(true, "Error inserting work time: \n", err)
		http.Error(w, "Error inserting work time", http.StatusInternalServerError)
		return
	}
	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"time_id": row,
	})
}

func StopWork(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	// Check if token is valid
	if !model.CheckToken(r.Header.Get("Token")) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	var req ReqStopWorkBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Write request in console
	logger.Debug(false, req)
	_, err = model.ExecSQL(`UPDATE work_time SET end_time = $1 WHERE id = $2`, req.Time, req.ID)
	if err != nil {
		logger.Error(true, "Error inserting work time: \n", err)
		http.Error(w, "Error inserting work time", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func StartBreak(w http.ResponseWriter, r *http.Request) {
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
	var req ReqStartWorkBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Write request in console
	logger.Debug(false, req)
	row, err := model.QueryValueSQL(`INSERT INTO break_time (start_time, user_id) 
	VALUES ($1, (SELECT id FROM "user" WHERE login = $2)) RETURNING id`, req.Time, username)
	if err != nil {
		logger.Error(true, "Error inserting break time: \n", err)
		http.Error(w, "Error inserting break time", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"time_id": row,
	})
}

func StopBreak(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	// Check if token is valid
	if !model.CheckToken(r.Header.Get("Token")) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	var req ReqStopWorkBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Write request in console
	logger.Debug(false, req)
	_, err = model.ExecSQL(`UPDATE break_time SET end_time = $1 WHERE id = $2`, req.Time, req.ID)
	if err != nil {
		logger.Error(true, "Error inserting break time: \n", err)
		http.Error(w, "Error inserting break time", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func Report(w http.ResponseWriter, r *http.Request) {
	logger, err := model.NewLogger()
	if err != nil {
		log.Println("Error creating logger: \n", err)
		return
	}
	// Check if token is valid
	if !model.CheckToken(r.Header.Get("Token")) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	var req ReqReportBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Write request in console
	logger.Debug(false, req)
	_, err = model.ExecSQL(`INSERT INTO report (text, time_id) VALUES ($1, $2)`, req.Text, req.ID)
	if err != nil {
		logger.Error(true, "Error inserting break time: \n", err)
		http.Error(w, "Error inserting break time", http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}
