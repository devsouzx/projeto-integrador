package laudo

import (
	"net/http"
	
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/laudo"
	"github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

type LaudoControllerInterface interface {
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

func (lc *laudoController) GetLaudosByMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	laudos, err := lc.laudoService.GetLaudosByMedicoID(medico.(model.BaseUser).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Laudos encontrados com sucesso",
		"data":    laudos,
	})
}
