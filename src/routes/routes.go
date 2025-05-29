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
	r.POST("/login", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/user")
	})
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "userpage.html", nil)
	})
}