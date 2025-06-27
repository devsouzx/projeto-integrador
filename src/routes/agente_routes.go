package routes

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/agente"
	"github.com/gin-gonic/gin"
)

func InitAgenteRoutes(rg *gin.RouterGroup, agenteController agente.AgenteControllerInterface) {
	agente := rg.Group("/agente")
	agente.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("agente"))
	{
		agente.GET("/dashboard", agenteController.RenderDashboard)
		agente.GET("/pacientes", agenteController.RenderPacientes)
		agente.GET("/visitas", agenteController.RenderVisitas)
		agente.GET("/territorio", agenteController.RenderTerritorio)
		agente.GET("/alertas", agenteController.RenderAlertas)
	}
}