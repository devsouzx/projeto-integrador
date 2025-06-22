package enfermeiro

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	"github.com/gin-gonic/gin"
)

func NewEnfermeiroController(
	enfermeiroService enfermeiro.EnfermeiroServiceInterface,
	cnesService datasus.CNESService,
) EnfermeiroControllerInterface {
	return &enfermeiroController{
		enfermeiroService: enfermeiroService,
		cnesService:       cnesService,
	}
}

type enfermeiroController struct {
	enfermeiroService enfermeiro.EnfermeiroServiceInterface
	cnesService       datasus.CNESService
}

type EnfermeiroControllerInterface interface {
	GetUnidadeEnfermeiro(c *gin.Context)

	RenderDashboardPage(c *gin.Context)
	RenderNovaFichaPage(c *gin.Context)
	RenderEdicaoFichaPage(c *gin.Context)
	RenderConsultarPacientesPage(c *gin.Context)
	RenderAgendamentosPage(c *gin.Context)
	RenderRelatoriosPage(c *gin.Context)
}

func (ec *enfermeiroController) GetUnidadeEnfermeiro(c *gin.Context) {
	enfermeiro, exists := ec.getEnfermeiroLogado(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado ou sessão expirada"})
		return
	}

	unidade, err := ec.cnesService.BuscarUnidadePorCNES(enfermeiro.UnidadeID)
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

func (mc *enfermeiroController) getEnfermeiroLogado(c *gin.Context) (*model.Enfermeiro, bool) {
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

	enfermeiro, err := mc.enfermeiroService.FindEnfermeiroByIdentifier(baseUser.ID)
	if err != nil {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return nil, false
	}

	return enfermeiro, true
}

func (mc *enfermeiroController) RenderDashboardPage(c *gin.Context) {
	enfermeiro, ok := mc.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	c.HTML(http.StatusOK, "dashboard-enfermeira.html", gin.H{
		"Enfermeiro": enfermeiro,
	})
}

func (mc *enfermeiroController) RenderNovaFichaPage(c *gin.Context) {
	enfermeiro, ok := mc.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	c.HTML(http.StatusOK, "nova-ficha-enfermeira.html", gin.H{
		"Enfermeiro": enfermeiro,
	})
}

func (mc *enfermeiroController) RenderEdicaoFichaPage(c *gin.Context) {
	enfermeiro, ok := mc.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	c.HTML(http.StatusOK, "editar-ficha-enfermeira.html", gin.H{
		"Enfermeiro": enfermeiro,
	})
}

func (mc *enfermeiroController) RenderConsultarPacientesPage(c *gin.Context) {
	enfermeiro, ok := mc.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	c.HTML(http.StatusOK, "consultar-enfermeira.html", gin.H{
		"Enfermeiro": enfermeiro,
	})
}

func (mc *enfermeiroController) RenderAgendamentosPage(c *gin.Context) {
	enfermeiro, ok := mc.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	c.HTML(http.StatusOK, "agendamentos-enfermeira.html", gin.H{
		"Enfermeiro": enfermeiro,
	})
}

func (mc *enfermeiroController) RenderRelatoriosPage(c *gin.Context) {
	enfermeiro, ok := mc.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	c.HTML(http.StatusOK, "relatorios-enfermeira.html", gin.H{
		"Enfermeiro": enfermeiro,
	})
}