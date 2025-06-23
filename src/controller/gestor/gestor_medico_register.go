package gestor

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/gin-gonic/gin"
)

func (gc *gestorController) CadastrarUsuario(c *gin.Context) {
	var cadRequest request.CadastroUsuarioRequest

	if err := c.ShouldBindJSON(&cadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao validar dados do usuário",
			"details": err.Error(),
		})
		return
	}

	switch cadRequest.TipoUsuario {
	case "medico":
		medico, err := gc.service.CadastrarMedico(cadRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erro ao cadastrar médico",
				"details": err.Error(),
			})
			return
		}
		response := gin.H{
			"message": "Médico cadastrado com sucesso",
			"data": gin.H{
				"id":               medico.ID,
				"nomecompleto":     medico.Name,
				"crm":              medico.CRM,
				"especialidade":    medico.Especialidade,
				"cpf":              medico.CPF,
				"datadenascimento": medico.DataNascimento,
				"email":            medico.Email,
				"telefone":         medico.Telefone,
			},
		}
		c.JSON(http.StatusCreated, response)
	case "enfermeiro":
		// A implementação para enfermeiro deve ser feita aqui
	case "agente":
		// A implementação para agente deve ser feita aqui
	case "gestor":
		// A implementação para gestor deve ser feita aqui
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuário inválido"})
	}
}
