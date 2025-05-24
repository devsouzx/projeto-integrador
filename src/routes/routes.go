package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
) {
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
}