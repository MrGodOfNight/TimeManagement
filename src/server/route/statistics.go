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
	"strings"
	"time"
)

// Struct for statistics
type ReqStatisticsBody struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

// Struct for statistics
type TimeForDay struct {
	Hours   float64
	Minutes float64
	Seconds float64
}
type TimeForMonth struct {
	Date string
	Time string
}

func UserDayStatistics(w http.ResponseWriter, r *http.Request) {
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
	// Get user ID from URL
	userID := strings.TrimPrefix(r.URL.Path, "/statistics/day/")

	// Check if user ID is valid
	if userID == "" {
		http.Error(w, "User ID is missing", http.StatusBadRequest)
		return
	}

	if username != userID {
		row, err := model.QueryValueSQL(`SELECT admin FROM "user" WHERE login = $1`, username)
		if err != nil {
			logger.Error(true, "Error checking user level: \n", err)
			http.Error(w, "Error checking user level", http.StatusInternalServerError)
			return
		}
		if row.(int) < 1 {
			http.Error(w, "Insufficient user level", http.StatusUnauthorized)
			return
		}
	}

	rows, err := model.QuerySQL(`WITH user_id_cte AS (
    SELECT id
    FROM "user"
    WHERE login = $1
),
work_sessions AS (
    SELECT 
        start_time, 
        end_time,
        EXTRACT(EPOCH FROM (end_time - start_time)) AS seconds_worked
    FROM work_time
    WHERE user_id = (SELECT id FROM user_id_cte)
      AND "date" = $2
),
total_work AS (
    SELECT 
        SUM(seconds_worked) AS total_seconds
    FROM work_sessions
)
SELECT
    FLOOR(total_seconds / 3600) AS hours,
    FLOOR((total_seconds % 3600) / 60) AS minutes,
    total_seconds % 60 AS seconds
FROM total_work;
`, userID, time.Now().Format("2006-01-02"))
	if err != nil {
		logger.Error(true, "Error getting work time: \n", err)
		http.Error(w, "Error getting work time", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var timeForDay TimeForDay
		err := rows.Scan(&timeForDay.Hours, &timeForDay.Minutes, &timeForDay.Seconds)
		if err != nil {
			logger.Error(true, "Error scanning row: \n", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		// Send response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timeForDay)
	}
}

func UserMonthStatistics(w http.ResponseWriter, r *http.Request) {
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
	// Get user ID from URL
	userID := strings.TrimPrefix(r.URL.Path, "/statistics/month/")

	// Check if user ID is valid
	if userID == "" {
		http.Error(w, "User ID is missing", http.StatusBadRequest)
		return
	}

	if username != userID {
		row, err := model.QueryValueSQL(`SELECT admin FROM "user" WHERE login = $1`, username)
		if err != nil {
			logger.Error(true, "Error checking user level: \n", err)
			http.Error(w, "Error checking user level", http.StatusInternalServerError)
			return
		}
		if row.(int) < 1 {
			http.Error(w, "Insufficient user level", http.StatusUnauthorized)
			return
		}
	}

	var req ReqStatisticsBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	rows, err := model.QuerySQL(`WITH user_id_cte AS (
    SELECT id
    FROM "user"
    WHERE login = $1
),
work_sessions AS (
    SELECT 
        "date",
        EXTRACT(EPOCH FROM (end_time - start_time)) AS seconds_worked
    FROM work_time
    WHERE user_id = (SELECT id FROM user_id_cte)
      AND EXTRACT(MONTH FROM "date") = $2
      AND EXTRACT(YEAR FROM "date") = $3
),
daily_totals AS (
    SELECT
        "date",
        SUM(seconds_worked) AS total_seconds
    FROM work_sessions
    GROUP BY "date"
)
SELECT
    "date",
    CONCAT(
        FLOOR(total_seconds / 3600), ':',
        FLOOR((total_seconds % 3600) / 60), ':',
        CAST(total_seconds % 60 AS INT)
    ) AS time
FROM daily_totals
ORDER BY "date";
`, userID, req.Month, req.Year)
	if err != nil {
		logger.Error(true, "Error getting work time: \n", err)
		http.Error(w, "Error getting work time", http.StatusInternalServerError)
		return
	}

	var data []TimeForMonth
	for rows.Next() {
		var timeForMonth TimeForMonth
		err := rows.Scan(&timeForMonth.Date, &timeForMonth.Time)
		if err != nil {
			logger.Error(true, "Error scanning row: \n", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		data = append(data, timeForMonth)
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
