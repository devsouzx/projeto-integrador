package routes

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/gestor"
	"github.com/gin-gonic/gin"
)

func InitGestorRoutes(rg *gin.RouterGroup, gestorController gestor.GestorControllerInterface) {
	gestor := rg.Group("/gestor")
	gestor.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("gestor"))
	{
		gestor.GET("/dashboard", gestorController.RenderDashboard)
		gestor.GET("/novo-usuario", gestorController.RenderNovoUsuario)
		gestor.GET("/usuarios", gestorController.RenderUsuarios)
		gestor.GET("/relatorios", gestorController.RenderRelatorios)
		gestor.GET("/exames", gestorController.RenderExames)

		gestor.POST("/novo-profissional", gestorController.CadastrarProfissional)
	}
}