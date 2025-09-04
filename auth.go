package main

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func initJWTMiddleware() *jwt.GinJWTMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "example zone",
		Key:           []byte("secret key"), // change in production
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,

		// Auth on Login
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}

			var err error = c.ShouldBindJSON(&login)
			if err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			// Query database
			uname := login.Username

			var user UserProfile
			queryResult := DB.First(&user, "username = ?", uname)
			if queryResult.Error != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			// Compares passwords
			pswrd := login.Password
			if pswrd == user.Password {
				return &login, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},

		// State what the token will contain
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*struct {
				Username string
				Password string
			}); ok {
				return jwt.MapClaims{
					"id": v.Username,
				}
			}
			return jwt.MapClaims{}
		},

		// Unautherized
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
	})

	if err != nil {
		log.Fatal("JWT.Error: " + err.Error())
	}

	return jwtMiddleware
}
