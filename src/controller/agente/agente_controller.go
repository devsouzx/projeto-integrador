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
	agente, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "dashboard-agente.html", gin.H{
		"Agente": agente,
	})
}

func (ac *agenteController) RenderPacientes(c *gin.Context) {
	agente, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "cadastro-paciente-agente.html", gin.H{
		"Agente": agente,
	})
}

func (ac *agenteController) RenderVisitas(c *gin.Context) {
	agente, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "visitas-agente.html", gin.H{
		"Agente": agente,
	})
}

func (ac *agenteController) RenderTerritorio(c *gin.Context) {
	agente, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "territorio-agente.html", gin.H{
		"Agente": agente,
	})
}

func (ac *agenteController) RenderAlertas(c *gin.Context) {
	agente, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "alertas-agente.html", gin.H{
		"Agente": agente,
	})
}