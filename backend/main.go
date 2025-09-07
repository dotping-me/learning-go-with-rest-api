/*

Program Entry Point!

*/

package main

import (
	"log"

	"github.com/dotping-me/learning-go-with-rest-api/backend/api"
	"github.com/dotping-me/learning-go-with-rest-api/backend/api/middleware"
	"github.com/dotping-me/learning-go-with-rest-api/backend/configs"
	"github.com/dotping-me/learning-go-with-rest-api/backend/data"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig() // Loads .env
	data.ConnetToDB(cfg.DSN)    // Connects to database

	// Starts Gin router
	router := gin.Default()
	jwtMiddleware := middleware.InitJWT(cfg.JWTSecret)

	// Set trusted IPs
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	// Register routes
	api.RegisterAPIRoutes(router, jwtMiddleware) // API

	router.Static("/static", "./frontend/static")
	api.RegisterWebRoutes(router, jwtMiddleware) // Web

	router.Run(":" + cfg.Port) // Runs Server
}
