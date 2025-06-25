package medico

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
)

func NewMedicoService(
	medicoRepository medicoRepository.MedicoRepositoryInterface,
) MedicoServiceInterface {
	return &medicoService{
		repository: medicoRepository,
	}
}

type medicoService struct {
	repository medicoRepository.MedicoRepositoryInterface
}

type MedicoServiceInterface interface {
	FindMedicoByIdentifier(identifier string) (model.UserInterface, error)
}

func (ms *medicoService) FindMedicoByIdentifier(identifier string) (model.UserInterface, error) {
	medico, err := ms.repository.FindMedicoByIdentifier(identifier)
	if err != nil {
		fmt.Println("Erro ao buscar médico:", err)
		return nil, fmt.Errorf("erro ao busccar o medico")
	}

	return medico, nil
}