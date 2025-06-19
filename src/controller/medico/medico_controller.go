package medico

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

func NewMedicoController(
	service medicoService.MedicoServiceInterface, 
	cnesService datasus.CNESService,
	pacienteService paciente.PacienteService,
) MedicoControllerInterface {
	return &medicoController{
		service: service,
		cnesService: cnesService,
		pacienteService: pacienteService,
	}
}

type medicoController struct {
	service medicoService.MedicoServiceInterface
	cnesService datasus.CNESService
	pacienteService paciente.PacienteService
}

type MedicoControllerInterface interface {
	getMedicoLogado(c *gin.Context) (model.UserInterface, bool)
	RenderDashboard(c *gin.Context)
	RenderNovaFichaPage(c *gin.Context)
	ExibirListaPacientes(c *gin.Context)
	ExibirLaudos(c *gin.Context)
	RenderPacientePage(c *gin.Context)
	ExibirEncaminhamentos(c *gin.Context)
	ExibirEstatisticas(c *gin.Context)
	RenderNovoEncaminhamentoPage(c *gin.Context)
	RenderNovoLaudoPage(c *gin.Context)
	GetUnidadeMedico(c *gin.Context)
}