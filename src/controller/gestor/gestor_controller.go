package gestor

import (
	"github.com/devsouzx/projeto-integrador/src/service/gestor"
	"github.com/gin-gonic/gin"
)

func NewGestorController(
	gestorService gestor.GestorSevice,
) GestorControllerInterface {
	return &gestorController{
		gestorService: gestorService,
	}
}

type gestorController struct {
	gestorService gestor.GestorService
}

type GestorControllerInterface interface {
	CadastrarMedico(c *gin.Context)
}
