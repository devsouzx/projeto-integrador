package laudo

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/laudo"
	"github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

type LaudoControllerInterface interface {
	CreateLaudo(c *gin.Context)
	GetLaudosByMedico(c *gin.Context)
}

type laudoController struct {
	laudoService    laudo.LaudoServiceInterface
	pacienteService paciente.PacienteService
	medicoService   medico.MedicoServiceInterface
}

func NewLaudoController(
	laudoService laudo.LaudoServiceInterface,
	pacienteService paciente.PacienteService,
	medicoService medico.MedicoServiceInterface,
) LaudoControllerInterface {
	return &laudoController{
		laudoService:    laudoService,
		pacienteService: pacienteService,
		medicoService:   medicoService,
	}
}

func (lc *laudoController) CreateLaudo(c *gin.Context) {
	log.Println("Início CreateLaudo")
	defer log.Println("Fim CreateLaudo")
	var laudo model.Laudo
	if err := c.ShouldBindJSON(&laudo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	laudo.MedicoID = medico.(model.BaseUser).ID
	laudo.Status = model.LaudoStatusConcluido
	laudo.DataLaudo = time.Now()

	ficha, err := lc.pacienteService.FindUltimaFichaByPacienteID(laudo.PacienteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	laudo.FichaID = ficha.ID

	if err := lc.laudoService.CreateLaudo(&laudo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Laudo criado com sucesso",
		"data":    laudo,
	})
}

func (lc *laudoController) GetLaudosByMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	resultadoFilter := c.Query("resultado")
	search := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	laudos, total, err := lc.laudoService.GetLaudosByMedicoID(
		medico.(model.BaseUser).ID,
		resultadoFilter,
		search,
		page,
		pageSize,
	)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": laudos,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  int(math.Ceil(float64(total) / float64(pageSize))),
			"hasNext":     page*pageSize < total,
			"hasPrev":     page > 1,
		},
	})
}
