package routes

import (
	"net/http"

	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(rg *gin.RouterGroup, userController userController.UserController) {
	auth := rg.Group("/")
	{
		auth.GET("/login", func(c *gin.Context) {
			if _, err := c.Cookie("token"); err == nil {
				user, exists := c.Get("user")
				if exists {
					model.RedirectToDashboard(c, user.(model.BaseUser).Role)
					return
				}
			}
			c.HTML(http.StatusOK, "login-paciente.html", nil)
		})

		auth.POST("/login", userController.LoginUser)

		auth.GET("/login-profissionais", func(c *gin.Context) {
			if _, err := c.Cookie("token"); err == nil {
				user, exists := c.Get("user")
				if exists {
					model.RedirectToDashboard(c, user.(model.BaseUser).Role)
					return
				}
			}
			c.HTML(http.StatusOK, "login-profissionais.html", nil)
		})

		auth.GET("/cadastro", func(c *gin.Context) {
			c.HTML(http.StatusOK, "cadastro-paciente.html", nil)
		})

		auth.GET("/recuperar-senha", func(c *gin.Context) {
			c.HTML(http.StatusOK, "recuperar-senha.html", nil)
		})

		auth.POST("/recuperar-senha", userController.SendCodeRecovey)
		auth.POST("/verificar-codigo", userController.VerifyCode)
		auth.POST("/redefinir-senha", userController.ResetPassword)

		auth.GET("/logout", userController.Logout)
	}
}