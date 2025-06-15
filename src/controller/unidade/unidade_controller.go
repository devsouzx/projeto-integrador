package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
)

type UnidadeController struct {
	cnesService *datasus.CNESService
}

func NewUnidadeController(cnesService *datasus.CNESService) *UnidadeController {
	return &UnidadeController{
		cnesService: cnesService,
	}
}

func (uc *UnidadeController) ListarUnidades(c *gin.Context) {
    codigoIBGE := c.Query("municipio")
    unidades, err := uc.cnesService.BuscarUnidadesPorMunicipio(codigoIBGE)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Falha ao buscar unidades",
            "details": err.Error(),
        })
        return
    }

    if len(unidades) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "Nenhuma unidade de saúde encontrada para este município",
            "data": []datasus.UnidadeSaude{},
        })
        return
    }

    c.JSON(http.StatusOK, unidades)
}