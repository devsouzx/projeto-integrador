package medico

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func (mc *medicoController) GetUnidadeMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	medicoObj, err := mc.service.FindMedicoByIdentifier(medico.(model.BaseUser).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar médico"})
		return
	}

	unidade, err := mc.cnesService.BuscarUnidadePorCNES(medicoObj.(*model.Medico).UnidadeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar unidade de saúde"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unidade": gin.H{
			"CodigoCNES":      unidade.CodigoCNES,
			"NomeFantasia":    unidade.NomeFantasia,
			"NomeRazaoSocial": unidade.NomeRazaoSocial,
			"Endereco":        unidade.Endereco,
			"CodigoUF":        unidade.CodigoUF,
		},
	})
}