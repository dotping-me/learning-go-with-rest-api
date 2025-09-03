package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/profile", handleUserProfile)

	log.Println("Server is listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
