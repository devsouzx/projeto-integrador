package medico

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/gin-gonic/gin"
)

func NewMedicoController(
	service medicoService.MedicoServiceInterface,
) MedicoControllerInterface {
	return &medicoController{
		service: service,
	}
}

type medicoController struct {
	service medicoService.MedicoServiceInterface
}

type MedicoControllerInterface interface {
	getMedicoLogado(c *gin.Context) (model.UserInterface, bool)
	GetDashboard(c *gin.Context)
	RenderNovaFichaPage(c *gin.Context)
	ExibirListaPacientes(c *gin.Context)
	ExibirLaudos(c *gin.Context)
	ExibirPaciente(c *gin.Context)
	ExibirEncaminhamentos(c *gin.Context)
	ExibirEstatisticas(c *gin.Context)
	RenderNovoEncaminhamentoPage(c *gin.Context)
	RenderNovoLaudoPage(c *gin.Context)
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados do médico"})
		return nil, false
	}

	return medico, true
}

func (mc *medicoController) GetDashboard(c *gin.Context) {
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

func (mc *medicoController) ExibirListaPacientes(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "pacientes-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) ExibirLaudos(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}
	
	c.HTML(http.StatusOK, "laudos-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) ExibirPaciente(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "paciente-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) ExibirEncaminhamentos(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "encaminhamentos-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) ExibirEstatisticas(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "estatisticas-medico.html", gin.H{
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

func (mc *medicoController) RenderNovoLaudoPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "laudo-medico.html", gin.H{
		"Medico": medico,
	})
}
