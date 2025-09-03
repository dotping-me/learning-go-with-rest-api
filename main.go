package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Loads .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	http.HandleFunc("/user/profile", handleUserProfile)

	log.Println("Server is listening on localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
