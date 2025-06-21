package enfermeiro

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	"github.com/gin-gonic/gin"
)

func NewEnfermeiroController(
	enfermeiroService enfermeiro.EnfermeiroServiceInterface,
	cnesService datasus.CNESService,
) EnfermeiroControllerInterface {
	return &enfermeiroController{
		enfermeiroService: enfermeiroService,
		cnesService: cnesService,
	}
}

type enfermeiroController struct {
	enfermeiroService enfermeiro.EnfermeiroServiceInterface
	cnesService datasus.CNESService
}

type EnfermeiroControllerInterface interface {
	GetUnidadeEnfermeiro(c *gin.Context)
}

func (ec *enfermeiroController) GetUnidadeEnfermeiro(c *gin.Context) {
	enfermeiro, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	enfermeiroObj, err := ec.enfermeiroService.FindEnfermeiroByIdentifier(enfermeiro.(model.BaseUser).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar enfermeiro"})
		return
	}

	unidade, err := ec.cnesService.BuscarUnidadePorCNES(enfermeiroObj.(*model.Enfermeiro).UnidadeID)
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
