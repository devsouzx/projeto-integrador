package routes

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitPacienteRoutes(rg *gin.RouterGroup) {
	paciente := rg.Group("/paciente")
	paciente.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("paciente"))
	{
		paciente.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-paciente.html", nil)
		})

		paciente.GET("/agendar", func(c *gin.Context) {
			c.HTML(http.StatusOK, "agendamento-paciente.html", nil)
		})

		paciente.GET("/historico_exames", func(c *gin.Context) {
			c.HTML(http.StatusOK, "historico-exames-paciente.html", nil)
		})

		paciente.GET("/localizar_ubs", func(c *gin.Context) {
			c.HTML(http.StatusOK, "localizar-ubs-paciente.html", nil)
		})
		
		paciente.GET("/notificacoes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "notificacoes-paciente.html", nil)
		})

		paciente.GET("/orientacoes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "orientacoes-paciente.html", nil)
		})
	}
}