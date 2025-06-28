package routes

import (
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	"github.com/devsouzx/projeto-integrador/src/controller/paciente"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/encaminhamento"
	"github.com/gin-gonic/gin"
)

func InitMedicoRoutes(rg *gin.RouterGroup, medicoController medicoController.MedicoControllerInterface, pacienteController paciente.PacienteControllerInterface, encaminhamentoController encaminhamento.EncaminhamentoControllerInterface) {
	medico := rg.Group("/medico")
	medico.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico"))
	{
		medico.GET("/dashboard", medicoController.RenderDashboard)
		medico.GET("/nova-ficha", medicoController.RenderNovaFichaPage)
		medico.GET("/pacientes", medicoController.RenderListaPacientesPage)
		medico.GET("/laudos", medicoController.RenderLaudosPage)
		medico.GET("/estatisticas", medicoController.RenderEstatisticasPage)
		medico.GET("/encaminhamentos", medicoController.RenderEncaminhamentosPage)
		medico.GET("/encaminhamentos/:id", medicoController.RenderEncaminhamentoPage)
		medico.GET("/encaminhamentos/novo", medicoController.RenderNovoEncaminhamentoPage)
		medico.GET("/laudos/novo", medicoController.RenderNovoLaudoPage)
		medico.GET("/paciente/:id", medicoController.RenderPacientePage)
		medico.GET("/api/paciente/:id", medicoController.GetPacienteByID)

		medico.GET("/minha-unidade", medicoController.GetUnidadeMedico)

		medico.GET("/api/encaminhamentos", encaminhamentoController.GetEncaminhamentosByMedico)
		medico.GET("/api/encaminhamentos/:id", encaminhamentoController.GetEncaminhamento)
		medico.GET("/api/encaminhamentos/paciente/:pacienteId", encaminhamentoController.GetEncaminhamentosByPaciente)
		medico.PUT("/api/encaminhamentos/update/:id", encaminhamentoController.UpdateStatus)
		medico.POST("/api/encaminhamentos", encaminhamentoController.CreateEncaminhamento)
	}
}
