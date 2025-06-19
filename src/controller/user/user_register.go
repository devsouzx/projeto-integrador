package user

import (
	"fmt"
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/gin-gonic/gin"
)

func (uc *userController) CadastrarUsuario(c *gin.Context) {
	var cadRequest request.CadastroRequest

	if err := c.ShouldBind(&cadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao validar dados do usuário",
			"details": err.Error(),
		})
		return
	}

	usuario, err := uc.service.CadastrarUsuario(cadRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao cadastrar usuario",
			"details": err.Error(),
		})
		return
	}

	newToken, err := usuario.GenerateVerifyToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	if err := uc.service.UpdateUser(usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	verificationLink := fmt.Sprintf("http://localhost:8080/verify?token=%s", newToken)
	if err := uc.emailService.SendVerificationEmail(usuario.GetEmail(), usuario.GetName(), verificationLink, newToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar email"})
		return
	}

	response := gin.H{
		"message": "Usuario cadastrado com sucesso",
		"data": gin.H{
			"id":                usuario.ID,
			"nomecompleto":      usuario.Name,
			"cpf":               usuario.CPF,
			"cartaosus":         usuario.CNS,
			"nomecompletodamae": usuario.NomeMae,
			"datadenascimento":  usuario.DataNascimento,
			"email":             usuario.Email,
			"telefone":          usuario.Telefone,
		},
	}

	if usuario.Endereco != nil {
		response["data"].(gin.H)["endereco"] = gin.H{
			"cep":         usuario.Endereco.CEP,
			"logradouro":  usuario.Endereco.Logradouro,
			"numero":      usuario.Endereco.Numero,
			"complemento": usuario.Endereco.Complemento,
			"bairro":      usuario.Endereco.Bairro,
			"cidade":      usuario.Endereco.Cidade,
			"uf":          usuario.Endereco.UF,
		}
	}

	c.JSON(http.StatusCreated, response)
}