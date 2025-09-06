/*

Using JWT Tokens to authenticate requests
- JWT Token will store the ID of the user as "id"

*/

package middleware

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dotping-me/learning-go-with-rest-api/backend/data"
	"github.com/dotping-me/learning-go-with-rest-api/backend/models"
	"github.com/gin-gonic/gin"
)

func InitJWT(secret string) *jwt.GinJWTMiddleware {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "example zone",
		Key:           []byte(secret),
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

			var user models.UserProfile
			queryResult := data.DB.First(&user, "username = ?", uname)
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
			if user, ok := data.(*models.UserProfile); ok {
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
			if id, ok := claims["id"].(float64); ok {
				return uint(id)
			}
			return nil
		},
	})

	if err != nil {
		log.Fatal("JWT.Error: " + err.Error())
	}

	return jwtMiddleware
}
