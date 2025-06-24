package agente

import (
	"github.com/devsouzx/projeto-integrador/src/repository/agente"
)

func NewAgenteService(
	agenteRepository agente.AgenteRepositoryInterface,
) AgenteServiceInterface {
	return &agenteService{
		agenteRepository: agenteRepository,
	}
}

type agenteService struct {
	agenteRepository agente.AgenteRepositoryInterface
}

type AgenteServiceInterface interface {
}
