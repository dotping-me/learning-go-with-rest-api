package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/dotping-me/learning-go-with-rest-api/web/templates"
	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	component := templates.Main("World!")

	c.Status(http.StatusOK)
	templ.Handler(component).ServeHTTP(c.Writer, c.Request)
}
