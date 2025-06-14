package routes

import (
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	"github.com/devsouzx/projeto-integrador/src/model"
	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	"github.com/gin-gonic/gin"
)

func InitMedicoRoutes(rg *gin.RouterGroup, medicoController medicoController.MedicoControllerInterface, fichaController fichaController.FichaController) {
	medico := rg.Group("/medico")
	medico.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico"))
	{
		medico.GET("/dashboard", medicoController.GetDashboard)
		medico.GET("/nova-ficha", medicoController.RenderNovaFichaPage)
		medico.GET("/pacientes", medicoController.ExibirListaPacientes)
		medico.GET("/laudos", medicoController.ExibirLaudos)
		medico.GET("/estatisticas", medicoController.ExibirEstatisticas)
		medico.GET("/encaminhamentos", medicoController.ExibirEncaminhamentos)
		medico.GET("/encaminhamentos/novo", medicoController.RenderNovoEncaminhamentoPage)
		medico.GET("/laudos/novo", medicoController.RenderNovoLaudoPage)
		medico.GET("/pacientes/123456", medicoController.ExibirPaciente)

		medico.POST("/nova-ficha", fichaController.Create)
	}
}