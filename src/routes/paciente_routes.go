package routes

import (
	pacienteController "github.com/devsouzx/projeto-integrador/src/controller/paciente"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitPacienteRoutes(rg *gin.RouterGroup, pacienteController pacienteController.PacienteControllerInterface) {
	paciente := rg.Group("/paciente")
	paciente.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("paciente"))
	{
		paciente.GET("/dashboard", pacienteController.RenderDashboard)
		paciente.GET("/agendar", pacienteController.RenderAgendar)
		paciente.GET("/historico_exames", pacienteController.RenderHistoricoExames)
		paciente.GET("/localizar_ubs", pacienteController.RenderLocalizarUBS)
		paciente.GET("/notificacoes", pacienteController.RenderNotificacoes)
		paciente.GET("/orientacoes", pacienteController.RenderOrientacoes)
	}
}
