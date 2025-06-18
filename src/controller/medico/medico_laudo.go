package medico

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *medicoController) RenderNovoLaudoPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "laudo-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) ExibirLaudos(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "laudos-lista-medico.html", gin.H{
		"Medico": medico,
	})
}