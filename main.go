package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Applies Middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middleware{
	TokenAuthMiddleware,
}

func main() {

	// Loads .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Wrap function in middleware
	var handler http.HandlerFunc = handleUserProfile
	for _, mdwr := range middlewares {
		handler = mdwr(handler) // Calls middleware function
	}

	port := os.Getenv("PORT")
	http.HandleFunc("/user/profile", handler)

	log.Println("Server is listening on localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
