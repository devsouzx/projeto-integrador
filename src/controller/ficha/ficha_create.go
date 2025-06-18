package ficha

import (
	"fmt"
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/gin-gonic/gin"
)

func (fc *fichaController) Create(c *gin.Context) {
	var request request.FichaRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"details": err.Error(),
		})
		return
	}

	fmt.Printf("Incoming request data:\nFicha: %+v\nPaciente: %+v\nDadosResidenciais: %+v\n",
		request.Ficha, request.Paciente, request.DadosResidenciais)

	if request.Paciente.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome do paciente é obrigatório"})
		return
	}
	if request.Paciente.CPF == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF do paciente é obrigatório"})
		return
	}
	if request.Ficha.DataColeta == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data da coleta é obrigatória"})
		return
	}

	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	request.Ficha.ResponsavelID = medico.(model.BaseUser).ID
	createdFicha, err := fc.service.CreateFicha(&request)
	if err != nil {
		fmt.Printf("Error creating ficha: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdFicha)
}