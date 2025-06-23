package agente

import (
	"github.com/devsouzx/projeto-integrador/src/repository/agente"
)

func NewAgenteService(
	agenteRepository agente.AgenteRepository,
) AgenteServiceInterface {
	return &agenteService{
		agenteRepository: agenteRepository,
	}
}

type agenteService struct {
	agenteRepository agente.AgenteRepository
}

type AgenteServiceInterface interface {
	
}