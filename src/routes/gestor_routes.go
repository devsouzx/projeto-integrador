package routes

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/gestor"
	"github.com/gin-gonic/gin"
)

func InitGestorRoutes(rg *gin.RouterGroup, gestorController gestor.GestorControllerInterface) {
	gestor := rg.Group("/gestor")
	gestor.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("gestor"))
	{
		gestor.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-gestor.html", nil)
		})

		gestor.GET("/novo-usuario", func(c *gin.Context) {
			c.HTML(http.StatusOK, "novo-usuario-gestor.html", nil)
		})

		gestor.GET("/usuarios", func(c *gin.Context) {
			c.HTML(http.StatusOK, "usuarios-gestor.html", nil)
		})

		gestor.GET("/relatorios", func(c *gin.Context) {
			c.HTML(http.StatusOK, "relatorios-gestor.html", nil)
		})

		gestor.GET("/exames", func(c *gin.Context) {
			c.HTML(http.StatusOK, "exames-gestor.html", nil)
		})

		gestor.POST("/novo-profissional", gestorController.CadastrarProfissional)
	}
}