package gestor

import (
	"github.com/devsouzx/projeto-integrador/src/service/gestor"
	"github.com/gin-gonic/gin"
)

func NewGestorController(
	gestorService gestor.GestorServiceInterface,
) GestorControllerInterface {
	return &gestorController{
		gestorService: gestorService,
	}
}

type gestorController struct {
	gestorService gestor.GestorServiceInterface
}

type GestorControllerInterface interface {
	CadastrarProfissional(c *gin.Context)
}
