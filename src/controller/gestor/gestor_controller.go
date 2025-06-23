package agente

import (
	"github.com/devsouzx/projeto-integrador/src/service/gestor"
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
	
}