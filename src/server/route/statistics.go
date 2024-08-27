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
	Month string `json:"month"`
	Year  string `json:"year"`
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

	var req ReqStatisticsBody
	// Decode request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	logger.Debug(false, "Request body: ", req)

	//ERROR: Don't work
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
        FLOOR(total_seconds / 3600), 'h ',
        FLOOR((total_seconds % 3600) / 60), 'm ',
        total_seconds % 60, 's'
    ) AS time
FROM daily_totals
ORDER BY "date";
`, userID, req.Month, req.Year)
	if err != nil {
		logger.Error(true, "Error getting work time: \n", err)
		http.Error(w, "Error getting work time", http.StatusInternalServerError)
		return
	}
	logger.Debug(false, "Rows: ", rows)

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
	logger.Debug(false, "JSON data: ", jsonData)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
