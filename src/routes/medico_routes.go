package routes

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitMedicoRoutes(rg *gin.RouterGroup) {
	medico := rg.Group("/medico")
	medico.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico"))
	{
		medico.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-medico.html", nil)
		})

		medico.GET("/nova-ficha", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nova-ficha-medico.html", nil)
		})

		medico.GET("/pacientes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "pacientes-lista-medico.html", nil)
		})

		medico.GET("/laudos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "laudos-lista-medico.html", nil)
		})

		medico.GET("/estatisticas", func(c *gin.Context) {
			c.HTML(http.StatusOK, "estatisticas-medico.html", nil)
		})

		medico.GET("/encaminhamentos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "encaminhamentos-lista-medico.html", nil)
		})

		medico.GET("/encaminhamentos/novo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "encaminhamento-medico.html", nil)
		})

		medico.GET("/laudos/novo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "laudo-medico.html", nil)
		})

		medico.GET("/medicos/pacientes/123456", func(c *gin.Context) {
			c.HTML(http.StatusOK, "paciente-medico.html", nil)
		})
	}
}