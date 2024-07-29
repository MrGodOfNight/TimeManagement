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
	loginRoutes(mux)
}

func loginRoutes(mux *http.ServeMux) {
	// For login, we'll use the LoginHandler function because we don't need to protect it
	mux.HandleFunc("/login", LoginHandler)
	// Create a new ServeMux for the register route
	loginRouter := http.NewServeMux()
	loginRouter.HandleFunc("/register", RegisterHandler)
	// Use the TokenAuthMiddleware to protect the register route
	mux.Handle("/register", TokenAuthMiddleware(loginRouter))
}

// Middleware for token validation
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Decode the token from the header
		token := r.Header.Get("token")
		if token == "" {
			http.Error(w, "Missing token", http.StatusBadRequest)
			return
		}
		if !model.CheckToken(token) {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, pass control to the next handler
		next.ServeHTTP(w, r)
	})
}
