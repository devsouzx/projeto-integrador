package gestor

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/gin-gonic/gin"
)

func (gc *gestorController) CadastrarProfissional(c *gin.Context) {
	var cadRequest request.CadastroProfissionalRequest

	if err := c.ShouldBindJSON(&cadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao validar dados do usuário",
			"details": err.Error(),
		})
		return
	}

	switch cadRequest.Role {
	case "medico":
		medico, err := gc.gestorService.CadastrarMedico(cadRequest)
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
				"id":            medico.ID,
				"nomecompleto":  medico.Name,
				"crm":           medico.CRM,
				"especialidade": medico.Especialidade,
				"cpf":           medico.CPF,
				"email":         medico.Email,
				"telefone":      medico.Telefone,
			},
		}
		c.JSON(http.StatusCreated, response)
	case "enfermeiro":
		enfermeiro, err := gc.gestorService.CadastrarEnfermeiro(cadRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erro ao cadastrar enfermeiro",
				"details": err.Error(),
			})
			return
		}
		response := gin.H{
			"message": "Enfermeiro cadastrado com sucesso",
			"data": gin.H{
				"id":           enfermeiro.ID,
				"nomecompleto": enfermeiro.Name,
				"coren":        enfermeiro.COREN,
				"cpf":          enfermeiro.CPF,
				"email":        enfermeiro.Email,
				"telefone":     enfermeiro.Telefone,
			},
		}
		c.JSON(http.StatusCreated, response)
	case "agente":
		agente, err := gc.gestorService.CadastrarAgente(cadRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erro ao cadastrar agente comunitário",
				"details": err.Error(),
			})
			return
		}
		response := gin.H{
			"message": "Agente comunitário cadastrado com sucesso",
			"data": gin.H{
				"id":           agente.ID,
				"nomecompleto": agente.Name,
				"cpf":          agente.CPF,
				"email":        agente.Email,
				"telefone":     agente.Telefone,
			},
		}
		c.JSON(http.StatusCreated, response)
	case "gestor":
		gestor, err := gc.gestorService.CadastrarGestor(cadRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Erro ao cadastrar gestor",
				"details": err.Error(),
			})
			return
		}
		response := gin.H{
			"message": "Gestor cadastrado com sucesso",
			"data": gin.H{
				"id":           gestor.ID,
				"nomecompleto": gestor.Name,
				"cpf":          gestor.CPF,
				"email":        gestor.Email,
				"telefone":     gestor.Telefone,
			},
		}
		c.JSON(http.StatusCreated, response)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de usuário inválido"})
	}
}
