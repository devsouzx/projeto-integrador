package medico

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *medicoController) ExibirListaPacientes(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "pacientes-lista-medico.html", gin.H{
		"Medico": medico,
	})
}

func (mc *medicoController) ExibirPaciente(c *gin.Context) {
	medico, ok := mc.getMedicoLogado(c)
	if !ok {
		return
	}

	c.HTML(http.StatusOK, "paciente-medico.html", gin.H{
		"Medico": medico,
	})
}