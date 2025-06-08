package routes

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitAgenteRoutes(rg *gin.RouterGroup) {
	agente := rg.Group("/agente")
	agente.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("agente"))
	{
		agente.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-agente.html", nil)
		})

		agente.GET("/pacientes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "cadastro-paciente-agente.html", nil)
		})

		agente.GET("/visitas", func(c *gin.Context) {
			c.HTML(http.StatusOK, "visitas-agente.html", nil)
		})

		agente.GET("/territorio", func(c *gin.Context) {
			c.HTML(http.StatusOK, "territorio-agente.html", nil)
		})

		agente.GET("/alertas", func(c *gin.Context) {
			c.HTML(http.StatusOK, "alertas-agente.html", nil)
		})
	}
}