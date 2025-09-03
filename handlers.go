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

	case http.MethodPatch:
		UpdateUserProfile(w, r)

	default:
		http.Error(w, "HTTP Method not allowed!", http.StatusMethodNotAllowed)
	}
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userProf := r.Context().Value("userProf").(UserProfile) // Casting to UserProfule type at the end

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
	newId := "USER" + strconv.Itoa(len(database)+1)
	database[newId] = UserProfile{
		Id:       newId,
		Email:    payload.Email,
		Username: payload.Username,
		Token:    "",
	}

	// Sends Response
	fmt.Println(database)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registered Successfully!",
	})
}

func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	userProf := r.Context().Value("userProf").(UserProfile) // Casting to UserProfule type at the end

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
	userProf.Email = payload.Email
	userProf.Username = payload.Username
	database[userProf.Id] = userProf

	fmt.Println(database[userProf.Id]) // Checks if changes were applied
	w.WriteHeader(http.StatusOK)
}
