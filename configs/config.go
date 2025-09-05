/*

Configuration settings for the program

*/

package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DSN       string
	JWTSecret string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env: ", err)
	}

	// Gets port on which server will listen
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set in .env! Using Default (8080)")
		port = "8080"
	}

	// Gets Postgres URI
	dsn := os.Getenv("DB_URI")
	if dsn == "" {
		log.Fatal("DB_URI not set in .env!")
	}

	// Gets JWT Key
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env!")
	}

	return Config{
		Port: port,
		DSN:  dsn,
	}
}
