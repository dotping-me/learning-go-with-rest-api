package main

import (
	"encoding/json"
	"net/http"
)

// Acts like a simple router in this example
func handleUserProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserProfile(w, r) // Redirects to Controller function

	case http.MethodPost:
		registerUserProfile(w, r)

	case http.MethodPatch:
		updateUserProfile(w, r)

	default:
		http.Error(w, "HTTP Method not allowed!", http.StatusMethodNotAllowed)
	}
}

func getUserProfile(w http.ResponseWriter, r *http.Request) {

	// Query user from database
	var user UserProfile
	var reqId = r.URL.Query().Get("id")

	queryResult := DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Sends JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func registerUserProfile(w http.ResponseWriter, r *http.Request) {
	var payload UserProfile
	var err error = json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close() // Frees up resources

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if payload.Email == "" || payload.Username == "" {
		http.Error(w, "JSON Body missing important fields", http.StatusBadRequest)
		return
	}

	// Makes new database entry
	DB.Create(&UserProfile{
		Email:    payload.Email,
		Username: payload.Username,
	})

	// Sends Response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registered Successfully!",
	})
}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {

	// Query user from database
	var user UserProfile
	var reqId = r.URL.Query().Get("id")

	queryResult := DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Reads JSON body
	var payload UserProfile
	var err error = json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if payload.Email == "" || payload.Username == "" {
		http.Error(w, "JSON Body missing important fields", http.StatusBadRequest)
		return
	}

	// Updates Profile
	DB.Save(&UserProfile{
		Email:    payload.Email,
		Username: payload.Username,
	})

	w.WriteHeader(http.StatusOK)
}
