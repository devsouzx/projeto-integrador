package routes

import (
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	"github.com/devsouzx/projeto-integrador/src/controller/paciente"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func InitMedicoRoutes(rg *gin.RouterGroup, medicoController medicoController.MedicoControllerInterface, pacienteController paciente.PacienteControllerInterface) {
	medico := rg.Group("/medico")
	medico.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico"))
	{
		medico.GET("/dashboard", medicoController.RenderDashboard)
		medico.GET("/nova-ficha", medicoController.RenderNovaFichaPage)
		medico.GET("/pacientes", medicoController.ExibirListaPacientes)
		medico.GET("/laudos", medicoController.ExibirLaudos)
		medico.GET("/estatisticas", medicoController.ExibirEstatisticas)
		medico.GET("/encaminhamentos", medicoController.ExibirEncaminhamentos)
		medico.GET("/encaminhamentos/novo", medicoController.RenderNovoEncaminhamentoPage)
		medico.GET("/laudos/novo", medicoController.RenderNovoLaudoPage)

		medico.GET("/paciente/:id", medicoController.RenderPacientePage)
        
		medico.GET("/minha-unidade", medicoController.GetUnidadeMedico)
	}
}