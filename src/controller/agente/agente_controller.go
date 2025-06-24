package agente

import (
	"github.com/devsouzx/projeto-integrador/src/service/agente"
)

func NewAgenteController(
	agenteService agente.AgenteServiceInterface,
) AgenteControllerInterface {
	return &agenteController{
		agenteService: agenteService,
	}
}

type agenteController struct {
	agenteService agente.AgenteServiceInterface
}

type AgenteControllerInterface interface {
}
