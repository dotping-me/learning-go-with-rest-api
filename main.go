/*

Program Entry Point!

*/

package main

import (
	"log"

	"github.com/dotping-me/learning-go-with-rest-api/api"
	"github.com/dotping-me/learning-go-with-rest-api/api/middleware"
	"github.com/dotping-me/learning-go-with-rest-api/configs"
	"github.com/dotping-me/learning-go-with-rest-api/data"
	"github.com/dotping-me/learning-go-with-rest-api/web"
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

	// Loads Resources
	router.LoadHTMLGlob("./web/templates/*")
	router.Static("/static", "./web/static")
	web.RegisterWebRoutes(router) // Web

	router.Run(":" + cfg.Port) // Runs Server
}
