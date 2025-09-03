package main

import (
	c "context"
	"fmt"
	"net/http"
	"strings"
)

// Takes in a handler function and returns another handler function
func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Validate user id
		var userId = r.URL.Query().Get("id")
		userProf, ok := database[userId]
		if !ok || userId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Validate Auth Token
		token := r.Header.Get("Authorization")
		fmt.Printf("Auth Token: %v \n", token)

		if !isValidToken(userProf, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Stores user profile
		ctx := c.WithValue(r.Context(), "userProf", userProf)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) // Calls next handler function
	}
}

// Helper function
func isValidToken(userProf UserProfile, token string) bool {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ") == userProf.Token
	}

	return false
}
