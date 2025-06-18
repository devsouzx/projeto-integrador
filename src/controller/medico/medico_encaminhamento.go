package medico

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *medicoController) ExibirEncaminhamentos(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "encaminhamentos-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) RenderNovoEncaminhamentoPage(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "encaminhamento-medico.html", gin.H{
		"Medico": medico,
	})
}