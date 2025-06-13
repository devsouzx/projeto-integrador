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
	GetDashboard(c *gin.Context)
}

func (mc *medicoController) GetDashboard(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Abort()
		return
	}
	
	identifier := user.(model.BaseUser).ID

	medico, err := mc.service.FindMedicoByIdentifier(identifier)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.HTML(http.StatusOK, "dashboard-medico.html", gin.H{
		"Medico": medico,
	})
}