package encaminhamento

import (
	"net/http"
	"strconv"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/encaminhamento"
	"github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

type EncaminhamentoControllerInterface interface {
	CreateEncaminhamento(c *gin.Context)
	GetEncaminhamento(c *gin.Context)
	GetEncaminhamentosByPaciente(c *gin.Context)
	GetEncaminhamentosByMedico(c *gin.Context)
	UpdateStatus(c *gin.Context)
}

type encaminhamentoController struct {
	encaminhamentoService encaminhamento.EncaminhamentoServiceInterface
	pacienteService       paciente.PacienteService
	medicoService         medico.MedicoServiceInterface
}

func NewEncaminhamentoController(
	encaminhamentoService encaminhamento.EncaminhamentoServiceInterface,
	pacienteService paciente.PacienteService,
	medicoService medico.MedicoServiceInterface,
) EncaminhamentoControllerInterface {
	return &encaminhamentoController{
		encaminhamentoService: encaminhamentoService,
		pacienteService:       pacienteService,
		medicoService:         medicoService,
	}
}

func (ec *encaminhamentoController) CreateEncaminhamento(c *gin.Context) {
	var encaminhamento model.Encaminhamento
	if err := c.ShouldBindJSON(&encaminhamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	encaminhamento.MedicoID = medico.(model.BaseUser).ID

	if err := ec.encaminhamentoService.CreateEncaminhamento(&encaminhamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Encaminhamento criado com sucesso",
		"data":    encaminhamento,
	})
}

func (ec *encaminhamentoController) GetEncaminhamento(c *gin.Context) {
	id := c.Param("id")

	encaminhamento, err := ec.encaminhamentoService.GetEncaminhamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encaminhamento não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": encaminhamento,
	})
}

func (ec *encaminhamentoController) GetEncaminhamentosByPaciente(c *gin.Context) {
	pacienteID := c.Param("pacienteId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	encaminhamentos, total, err := ec.encaminhamentoService.GetEncaminhamentosByPacienteID(pacienteID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": encaminhamentos,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  (total + pageSize - 1) / pageSize,
			"hasNext":     page*pageSize < total,
			"hasPrev":     page > 1,
		},
	})
}

func (ec *encaminhamentoController) GetEncaminhamentosByMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	encaminhamentos, total, err := ec.encaminhamentoService.GetEncaminhamentosByMedicoID(medico.(model.BaseUser).ID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": encaminhamentos,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  (total + pageSize - 1) / pageSize,
			"hasNext":     page*pageSize < total,
			"hasPrev":     page > 1,
		},
	})
}

func (ec *encaminhamentoController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Status      model.EncaminhamentoStatus `json:"status"`
		Observacoes string                     `json:"observacoes"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.encaminhamentoService.UpdateStatus(id, request.Status, request.Observacoes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status do encaminhamento atualizado com sucesso",
	})
}