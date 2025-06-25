package enfermeiro

import (
	"net/http"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/agendamento"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	"github.com/gin-gonic/gin"
)

func NewEnfermeiroController(
	enfermeiroService enfermeiro.EnfermeiroServiceInterface,
	cnesService datasus.CNESService,
	agendamentoService agendamento.AgendamentoServiceInterface,
) EnfermeiroControllerInterface {
	return &enfermeiroController{
		enfermeiroService:  enfermeiroService,
		cnesService:        cnesService,
		agendamentoService: agendamentoService,
	}
}

type enfermeiroController struct {
	enfermeiroService enfermeiro.EnfermeiroServiceInterface
	cnesService       datasus.CNESService
	agendamentoService agendamento.AgendamentoServiceInterface
}

type EnfermeiroControllerInterface interface {
	GetUnidadeEnfermeiro(c *gin.Context)

	CreateAgendamento(c *gin.Context)
	GetAgendamentos(c *gin.Context)
	UpdateAgendamentoStatus(c *gin.Context)
	DeleteAgendamento(c *gin.Context)

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

func (ec *enfermeiroController) RenderAgendamentosPage(c *gin.Context) {
	enfermeiro, ok := ec.getEnfermeiroLogado(c)
	if !ok {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return
	}

	hoje := time.Now()
	agendamentos, err := ec.agendamentoService.ListAgendamentosByDate(enfermeiro.ID, "enfermeiro", hoje)
	if err != nil {
		agendamentos = []*model.Agendamento{}
	}

	c.HTML(http.StatusOK, "agendamentos-enfermeira.html", gin.H{
		"Enfermeiro":   enfermeiro,
		"Agendamentos": agendamentos,
		"DataAtual":    hoje.Format("02/01/2006"),
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

func (ec *enfermeiroController) CreateAgendamento(c *gin.Context) {
	enfermeiro, ok := ec.getEnfermeiroLogado(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var agendamento model.Agendamento
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	agendamento.ProfissionalID = enfermeiro.ID
	agendamento.ProfissionalTipo = "enfermeiro"
	agendamento.Status = "pendente"

	if err := ec.agendamentoService.CreateAgendamento(&agendamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, agendamento)
}

func (ec *enfermeiroController) GetAgendamentos(c *gin.Context) {
	enfermeiro, ok := ec.getEnfermeiroLogado(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	dateStr := c.Query("date")
	var date time.Time
	var err error

	if dateStr != "" {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido. Use YYYY-MM-DD"})
			return
		}
	} else {
		date = time.Now()
	}

	agendamentos, err := ec.agendamentoService.ListAgendamentosByDate(enfermeiro.ID, "enfermeiro", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar agendamentos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        date.Format("02/01/2006"),
		"agendamentos": agendamentos,
	})
}

func (ec *enfermeiroController) UpdateAgendamentoStatus(c *gin.Context) {
	enfermeiro, ok := ec.getEnfermeiroLogado(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	id := c.Param("id")
	var request struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}
	
	agendamento, err := ec.agendamentoService.GetAgendamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	if agendamento.ProfissionalID != enfermeiro.ID || agendamento.ProfissionalTipo != "enfermeiro" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Você não tem permissão para alterar este agendamento"})
		return
	}

	if err := ec.agendamentoService.UpdateAgendamentoStatus(id, request.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso"})
}

func (ec *enfermeiroController) DeleteAgendamento(c *gin.Context) {
	enfermeiro, ok := ec.getEnfermeiroLogado(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	id := c.Param("id")

	agendamento, err := ec.agendamentoService.GetAgendamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	if agendamento.ProfissionalID != enfermeiro.ID || agendamento.ProfissionalTipo != "enfermeiro" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Você não tem permissão para deletar este agendamento"})
		return
	}

	if err := ec.agendamentoService.DeleteAgendamento(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agendamento deletado com sucesso"})
}