package main

import (
	"log"
	"strconv"
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

			// Passwords doesn't match
			if login.Password != user.Password {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},

		// State what the token will contain
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*UserProfile); ok {
				return jwt.MapClaims{
					"id": user.ID, // Successful
				}
			}

			return jwt.MapClaims{} // Failed
		},

		// Unautherized
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			idRaw, ok := claims["id"]
			if !ok || idRaw == nil {
				return nil
			}

			switch v := idRaw.(type) {
			case float64:
				return uint(v)

			case string:
				idInt, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					return nil
				}

				return uint(idInt)
			}

			return nil
		},
	})

	if err != nil {
		log.Fatal("JWT.Error: " + err.Error())
	}

	return jwtMiddleware
}
