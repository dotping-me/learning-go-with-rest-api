/*

Placeholder routes
- Will change later when I start working on the frontend

*/

package web

import "github.com/gin-gonic/gin"

func RegisterWebRoutes(router *gin.Engine) {
	router.GET("/", homeHandler)
}
