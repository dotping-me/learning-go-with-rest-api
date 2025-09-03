package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Acts like a simple router in this example
func handleUserProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUserProfile(w, r) // Redirects to Controller function

	case http.MethodPost:
		RegisterUserProfile(w, r)

	default:
		http.Error(w, "HTTP Method not allowed!", http.StatusMethodNotAllowed)
	}
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("id")
	userProf, ok := database[userId] // Database variable is globally available across package main

	// Bad Request
	if !ok || userId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Prepares Response
	w.Header().Set("Content-Type", "application/json")
	res := UserProfile{
		Id:       userProf.Id,
		Email:    userProf.Email,
		Username: userProf.Username,
	}

	json.NewEncoder(w).Encode(res) // Sends JSON response
}

func RegisterUserProfile(w http.ResponseWriter, r *http.Request) {
	var userProf UserProfile
	var err error = json.NewDecoder(r.Body).Decode(&userProf)
	defer r.Body.Close() // Frees up resources

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if userProf.Email == "" || userProf.Username == "" {
		http.Error(w, "JSON Body missing important fields", http.StatusBadRequest)
		return
	}

	// Makes new database entry
	newId := "USER" + strconv.Itoa(len(database)+1)
	database[newId] = UserProfile{
		Id:       newId,
		Email:    userProf.Email,
		Username: userProf.Username,
		Token:    "",
	}

	// Sends Response
	fmt.Println(database)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registered Successfully!",
	})
}
