package routes

import (
	"net/http"

	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitMedicoRoutes(rg *gin.RouterGroup, medicoController medicoController.MedicoControllerInterface) {
	medico := rg.Group("/medico")
	medico.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico"))
	{
		medico.GET("/dashboard", medicoController.GetDashboard)

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

		medico.GET("/pacientes/123456", func(c *gin.Context) {
			c.HTML(http.StatusOK, "paciente-medico.html", nil)
		})
	}
}