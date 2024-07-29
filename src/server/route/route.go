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
