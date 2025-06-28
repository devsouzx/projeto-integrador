package medico

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"github.com/devsouzx/projeto-integrador/src/service/encaminhamento"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

func NewMedicoController(
	service medicoService.MedicoServiceInterface, 
	cnesService datasus.CNESService,
	pacienteService paciente.PacienteService,
	encaminhamentoService encaminhamento.EncaminhamentoServiceInterface,
) MedicoControllerInterface {
	return &medicoController{
		service: service,
		cnesService: cnesService,
		pacienteService: pacienteService,
		encaminhamentoService: encaminhamentoService,
	}
}

type medicoController struct {
	service medicoService.MedicoServiceInterface
	cnesService datasus.CNESService
	pacienteService paciente.PacienteService
	encaminhamentoService encaminhamento.EncaminhamentoServiceInterface
}

type MedicoControllerInterface interface {
	GetUnidadeMedico(c *gin.Context)
	GetPacienteByID(c *gin.Context)

	RenderDashboard(c *gin.Context)
	RenderNovaFichaPage(c *gin.Context)
	RenderEncaminhamentosPage(c *gin.Context)
	RenderNovoEncaminhamentoPage(c *gin.Context)
	RenderListaPacientesPage(c *gin.Context)
	RenderPacientePage(c *gin.Context)
	RenderNovoLaudoPage(c *gin.Context)
	RenderLaudosPage(c *gin.Context)
	RenderEstatisticasPage(c *gin.Context)
	RenderEncaminhamentoPage(c *gin.Context)
}

func (mc *medicoController) GetUnidadeMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	medicoObj, err := mc.service.FindMedicoByIdentifier(medico.(model.BaseUser).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar médico"})
		return
	}

	unidade, err := mc.cnesService.BuscarUnidadePorCNES(medicoObj.(*model.Medico).UnidadeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar unidade de saúde"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unidade": gin.H{
			"CodigoCNES":      unidade.CodigoCNES,
			"NomeFantasia":    unidade.NomeFantasia,
			"NomeRazaoSocial": unidade.NomeRazaoSocial,
			"Endereco":        unidade.Endereco,
			"CodigoUF":        unidade.CodigoUF,
		},
	})
}

func (mc *medicoController) getMedicoLogado(c *gin.Context) (model.UserInterface, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, false
	}

	baseUser, ok := user.(model.BaseUser)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil, false
	}

	medico, err := mc.service.FindMedicoByIdentifier(baseUser.ID)
	if err != nil {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return nil, false
	}

	return medico, true
}

func (mc *medicoController) GetPacienteByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do paciente é obrigatório"})
		return
	}

	paciente, err := mc.pacienteService.GetPacienteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Paciente encontrado com sucesso",
		"data": paciente,
	})
}

func (mc *medicoController) RenderDashboard(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "dashboard-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderNovaFichaPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "nova-ficha-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderEncaminhamentosPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "encaminhamentos-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderNovoEncaminhamentoPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "encaminhamento-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderListaPacientesPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "pacientes-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderPacientePage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id := c.Param("id")
	if id == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "ID do paciente é obrigatório"})
		return
	}

	paciente, err := mc.pacienteService.GetPacienteByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": err.Error()})
		return
	}

	anamnese, err := mc.pacienteService.GetAnamneseByPacienteID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "paciente-medico.html", gin.H{
		"Medico":   medico,
		"Paciente": paciente,
		"Anamnese": anamnese,
	})
}

func (mc *medicoController) RenderNovoLaudoPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "laudo-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderLaudosPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "laudos-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderEstatisticasPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "estatisticas-medico.html", gin.H{
		"Medico": medico,
	})
}

func (ec *medicoController) RenderEncaminhamentoPage(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id := c.Param("id")
	
	encaminhamento, err := ec.encaminhamentoService.GetEncaminhamentoByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Encaminhamento não encontrado",
		})
		return
	}
	
	c.HTML(http.StatusOK, "encaminhamento.html", gin.H{
		"Medico":        medico,
		"Encaminhamento": encaminhamento,
	})
}