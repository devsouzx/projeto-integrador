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

func (pc *medicoController) RenderPacientePage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "ID do paciente é obrigatório"})
		return
	}

	paciente, err := pc.pacienteService.GetPacienteByID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": err.Error()})
		return
	}

	anamnase, err := pc.pacienteService.GetAnamneseByPacienteID(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "paciente-medico.html", gin.H{
		"Paciente": paciente,
	    "Anamnese": anamnase,
	})
}