package medico

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *medicoController) RenderDashboard(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "dashboard-medico.html", gin.H{
		"Medico": medico,
	})
}