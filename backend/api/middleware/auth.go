/*

Using JWT Tokens to authenticate requests
- JWT Token will store the ID of the user as "id"

*/

package middleware

import (
	"log"
	"net/http"
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

			c.SetCookie("username", user.Username, 3600, "/", "", false, true)
			return &user, nil
		},

		// State what the token will contain
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*models.UserProfile); ok {
				return jwt.MapClaims{
					"id":       user.ID, // Successful
					"username": user.Username,
				}
			}

			return jwt.MapClaims{} // Failed
		},

		// Customizes response when loging in
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.SetCookie("jwt", token, 0, "/", "", true, true)
			username, err := c.Cookie("username")

			// Fallback
			if err != nil || username == "" {
				c.JSON(http.StatusOK, gin.H{
					"code":   code,
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":     code,
				"token":    token,
				"expire":   expire.Format(time.RFC3339),
				"username": username,
			})
		},

		// What happens if user tries to access a page without being logged in
		Unauthorized: func(c *gin.Context, code int, message string) {
			if c.ContentType() == "application/json" {
				c.JSON(code, gin.H{"error": message})
			} else {
				c.Redirect(http.StatusFound, "/login")
			}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			username, ok := claims["username"].(string)
			if !ok {
				return nil
			}

			id, ok := claims["id"].(float64)
			if !ok {
				return nil
			}

			return &models.UserProfile{
				ID:       uint(id),
				Username: username,
			}
		},
	})

	if err != nil {
		log.Fatal("JWT.Error: " + err.Error())
	}

	return jwtMiddleware
}

// For pages that can be accessed by unregistered users and
// registered users but customized for the later
func OptionalAuth(mw *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil || token == "" {
			// Username stays as "Guest"
		}

		tokenObj, err := mw.ParseTokenString(token)
		if err != nil {
			// Username stays as "Guest"
		}

		claims := jwt.ExtractClaimsFromToken(tokenObj)
		if username, ok := claims["username"].(string); ok && username != "" {
			c.Set("username", username)
		}

		c.Next()
	}
}
