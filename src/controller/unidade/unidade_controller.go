package unidade

import (
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"github.com/gin-gonic/gin"
)

type unidadeController struct {
	cnesService *datasus.CNESService
}

func NewUnidadeController(cnesService *datasus.CNESService) *unidadeController {
	return &unidadeController{
		cnesService: cnesService,
	}
}

type UnidadeController interface {
	ListarUnidades(c *gin.Context)
}