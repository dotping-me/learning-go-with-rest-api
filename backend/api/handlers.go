package api

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dotping-me/learning-go-with-rest-api/backend/data"
	"github.com/dotping-me/learning-go-with-rest-api/backend/models"
	"github.com/dotping-me/learning-go-with-rest-api/frontend/templates"
	"github.com/gin-gonic/gin"
)

// ----------------- API Handlers -----------------

// Registers a new user
func registerUserProfile(mw *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.UserProfile
		c.SetCookie("jwt", "", -1, "/", "", false, true) // Deletes any old cookies

		// Reads JSON Body
		var err error = c.ShouldBindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		if payload.Username == "" || payload.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Body missing some fields"})
			return
		}

		// Makes new database entry
		err = data.DB.Create(&models.UserProfile{
			Username: payload.Username,
			Password: payload.Password,
		}).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register User"})
			return
		}

		newUser := models.UserProfile{
			Username: payload.Username,
			Password: payload.Password,
		}

		token, expire, _ := mw.TokenGenerator(&newUser)

		c.SetCookie("jwt", token, 0, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message":  "User created",
			"username": newUser.Username,
			"token":    token,
			"expire":   expire.Format(time.RFC3339),
		})
	}
}

// Gets the profile of a user
func getUserProfile(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing User ID"})
		return
	}

	// Query user from database
	var user models.UserProfile
	queryResult := data.DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Updates data for an already existing user
func updateUserProfile(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing User ID"})
		return
	}

	// Query user from database
	var user models.UserProfile
	queryResult := data.DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Reads JSON body
	var payload models.UserProfile
	var err error = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Username == "" || payload.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Body missing some fields"})
		return
	}

	// Updates Profile
	user.Username = payload.Username
	user.Password = payload.Password

	err = data.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})
}

// Deletes a user
func deleteUserProfile(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing User ID"})
		return
	}

	// Query user from database
	var user models.UserProfile
	queryResult := data.DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Deletes user
	var err error = data.DB.Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

// Creates a post
func createPost(c *gin.Context) {

	// Gets id set in JWT token
	claims := jwt.ExtractClaims(c)
	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId := uint(userIdFloat)

	var payload models.Post
	var err error = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post cannot be empty"})
		return
	}

	// Creates new post
	post := models.Post{
		Content:       payload.Content,
		UserProfileID: userId,
	}

	err = data.DB.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Post"})
		return
	}

	// Fetch newly created post from database with UserProfile joined
	// TODO: Optimize later!
	if err := data.DB.Preload("UserProfile").First(&post, post.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch new Post"})
		return
	}

	// Renders templ component here!
	var buffer bytes.Buffer
	if err := templates.Post(post, post.UserProfile.Username).Render(c.Request.Context(), &buffer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
		"html": buffer.String(),
	})
}

// Gets a post
func getPost(c *gin.Context) {
	reqId := c.Param("pid")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Query user from database
	var p models.Post
	queryResult := data.DB.First(&p, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, p)
}

// Gets all posts
func getPostsAll(c *gin.Context) {

	// Queries database
	var AllPosts []models.Post
	queryResult := data.DB.Find(&AllPosts)
	if queryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all Posts"})
		return
	}

	c.JSON(http.StatusOK, AllPosts)
}

// Deletes Post
func deletePost(c *gin.Context) {

	// Gets id set in JWT token
	claims := jwt.ExtractClaims(c)
	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId := uint(userIdFloat)

	// Retrieves Post ID
	postId := c.Param("pid")
	if postId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Queries database
	var p models.Post
	queryResult := data.DB.First(&p, "id = ?", postId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Authenticate that this user is the author of this post
	if p.UserProfileID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	// Deletes post
	var err error = data.DB.Delete(&p).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

// Creates a comment
func createComment(c *gin.Context) {

	// Gets id set in JWT token
	claims := jwt.ExtractClaims(c)
	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId := uint(userIdFloat)

	// Retrieving and Casting Post ID to uint
	postIdStr := c.Param("pid")
	postID, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Handling JSON for Comment
	var payload models.Comment
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Comment cannot be empty"})
		return
	}

	// Creates new Comment
	comment := models.Comment{
		Content:       payload.Content,
		PostID:        uint(postID),
		UserProfileID: userId,
	}

	err = data.DB.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// Gets a comment
func getComment(c *gin.Context) {

	/*

		// Is this part really necessary??

		// Retrieving Post ID
		postIdStr := c.Param("pid")
		postID, err := strconv.Atoi(postIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
			return
		}

	*/

	// Retrieving Comment ID
	commentIdStr := c.Param("cid")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Queries
	var comment models.Comment
	queryResult := data.DB.First(&comment, "id = ?", commentId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// Gets all comments
func getCommentsAll(c *gin.Context) {

	// Retrieving and Casting Post ID to uint
	postIdStr := c.Param("pid")
	postID, err := strconv.Atoi(postIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Queries database
	var allComments []models.Comment
	queryResult := data.DB.Where("post_id = ?", postID).Find(&allComments)
	if queryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all Comments"})
		return
	}

	c.JSON(http.StatusOK, allComments)
}

// Deletes a comment
func deleteComment(c *gin.Context) {

	// Gets id set in JWT token
	claims := jwt.ExtractClaims(c)
	userIdFloat, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId := uint(userIdFloat)

	// Retrieving Comment ID
	commentIdStr := c.Param("cid")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Queries database
	var comment models.Comment
	queryResult := data.DB.First(&comment, "id = ?", commentId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Authenticate that this user is the author of this comment
	if comment.UserProfileID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this comment"})
		return
	}

	// Deletes comment
	queryResult = data.DB.Delete(&comment)
	if queryResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully!"})
}

// ----------------- Web Handlers -----------------

// Helper function to get username from Cookie or JWT token

// Returns home page template
func HomePage(jwtMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := "" // Default value

		if u, exists := c.Get("username"); exists {
			if uStr, ok := u.(string); ok {
				username = uStr
			}
		}

		// Gets all posts (Custom handler)

		// Queries database (Joins UserProfileID with Username)
		var AllPosts []models.Post
		queryResult := data.DB.Preload("UserProfile").Order("posted_at DESC").Find(&AllPosts)
		if queryResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all Posts"})
			return
		}

		home := templates.Home(AllPosts, username)
		templates.Main("Home", username, "", home).Render(c.Request.Context(), c.Writer)
	}
}

// Returns login page template
func LoginPage(c *gin.Context) {
	login := templates.Login()
	templates.Main("Log In", "", "signup", login).Render(c, c.Writer)
}

func SignupPage(c *gin.Context) {
	signup := templates.Signup()
	templates.Main("Sign Up", "", "login", signup).Render(c, c.Writer)
}
