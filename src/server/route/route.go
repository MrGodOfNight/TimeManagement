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
	"net/http"
)

func Routes(mux *http.ServeMux) {
	// For login, we'll use the LoginHandler function because we don't need to protect it
	mux.HandleFunc("/login", LoginHandler)
	workRoutes(mux)
	statisticsRoutes(mux)
	adminRoutes(mux)
}

func workRoutes(mux *http.ServeMux) {
	// Create a new ServeMux for the work routes
	NewServeMux := http.NewServeMux()
	NewServeMux.HandleFunc("/work/start", StartWork)
	NewServeMux.HandleFunc("/work/stop", StopWork)
	NewServeMux.HandleFunc("/work/report", Report)
	NewServeMux.HandleFunc("/break/start", StartBreak)
	NewServeMux.HandleFunc("/break/stop", StopBreak)
	// Use the TokenAuthMiddleware to protect the work routes
	mux.Handle("/work/start", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/work/stop", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/work/report", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/break/start", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/break/stop", TokenAuthMiddleware(NewServeMux))
}

func statisticsRoutes(mux *http.ServeMux) {
	// Create a new ServeMux for the statistics routes
	NewServeMux := http.NewServeMux()
	NewServeMux.HandleFunc("/statistics/day/", UserDayStatistics)
	NewServeMux.HandleFunc("/statistics/month/", UserMonthStatistics)
	// Use the TokenAuthMiddleware to protect the statistics routes
	mux.Handle("/statistics/day/", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/statistics/month/", TokenAuthMiddleware(NewServeMux))
}

func adminRoutes(mux *http.ServeMux) {
	// Create a new ServeMux for the register route
	NewServeMux := http.NewServeMux()
	NewServeMux.HandleFunc("/admin/register", RegisterHandler)
	NewServeMux.HandleFunc("/admin/changeuser", ChangeUserHandler)
	// NewServeMux.HandleFunc("/admin/deleteuser", DeleteUserHandler)
	// NewServeMux.HandleFunc("/admin/checkreport", CheckReportHandler)
	NewServeMux.HandleFunc("/admin/getusers", GetUsersHandler)
	// Use the TokenAuthMiddleware to protect the register route
	mux.Handle("/admin/register", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/admin/changeuser", TokenAuthMiddleware(NewServeMux))
	// mux.Handle("/admin/deleteuser", TokenAuthMiddleware(NewServeMux))
	// mux.Handle("/admin/checkreport", TokenAuthMiddleware(NewServeMux))
	mux.Handle("/admin/getusers", TokenAuthMiddleware(NewServeMux))
}

// Middleware for token validation
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Decode the token from the header
		token := r.Header.Get("Token")
		if token == "" {
			http.Error(w, "Missing token", http.StatusBadRequest)
			return
		}
		// Check if token is valid
		if !model.CheckToken(token) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, pass control to the next handler
		next.ServeHTTP(w, r)
	})
}
