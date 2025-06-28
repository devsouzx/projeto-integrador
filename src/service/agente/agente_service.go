package agente

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
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
	FindAgenteByIdentifier(identifier string) (*model.AgenteComunitario, error)
}

func (as *agenteService) FindAgenteByIdentifier(identifier string) (*model.AgenteComunitario, error) {
	agente, err := as.agenteRepository.FindAgenteByIdentifier(identifier)
	if err != nil {
		fmt.Println("Erro ao buscar agente:", err)
		return nil, fmt.Errorf("erro ao buscar o agente")
	}

	return agente, nil
}
