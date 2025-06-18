package ficha

import (
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	"github.com/gin-gonic/gin"
)

type fichaController struct {
	service fichaService.FichaServiceInterface
}

func NewFichaController(service fichaService.FichaServiceInterface) FichaController {
	return &fichaController{service: service}
}

type FichaController interface {
	Create(c *gin.Context)
}