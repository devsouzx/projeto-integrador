package medico

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *medicoController) ExibirEstatisticas(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "estatisticas-medico.html", gin.H{
		"Medico": medico,
	})
}