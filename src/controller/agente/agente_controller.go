package agente

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/service/agente"
	"github.com/gin-gonic/gin"
)

func NewAgenteController(
	agenteService agente.AgenteServiceInterface,
) AgenteControllerInterface {
	return &agenteController{
		agenteService: agenteService,
	}
}

type agenteController struct {
	agenteService agente.AgenteServiceInterface
}

type AgenteControllerInterface interface {
	RenderDashboard(c *gin.Context)
	RenderPacientes(c *gin.Context)
	RenderVisitas(c *gin.Context)
	RenderTerritorio(c *gin.Context)
	RenderAlertas(c *gin.Context)
}

func (ac *agenteController) RenderDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard-agente.html", nil)
}

func (ac *agenteController) RenderPacientes(c *gin.Context) {
	c.HTML(http.StatusOK, "cadastro-paciente-agente.html", nil)
}

func (ac *agenteController) RenderVisitas(c *gin.Context) {
	c.HTML(http.StatusOK, "visitas-agente.html", nil)
}

func (ac *agenteController) RenderTerritorio(c *gin.Context) {
	c.HTML(http.StatusOK, "territorio-agente.html", nil)
}

func (ac *agenteController) RenderAlertas(c *gin.Context) {
	c.HTML(http.StatusOK, "alertas-agente.html", nil)
}