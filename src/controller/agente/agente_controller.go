package agente

import (
	"github.com/devsouzx/projeto-integrador/src/service/agente"
)

func NewAgenteController(
	agenteService agente.AgenteSevice,
) AgenteControllerInterface {
	return &agenteController{
		agenteService: agenteService,
	}
}

type agenteController struct {
	agenteService agente.AgenteService
}

type AgenteControllerInterface interface {
}
