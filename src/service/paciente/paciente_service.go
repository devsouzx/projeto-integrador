package paciente

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/paciente"
)

func NewPacienteService(
	pacienteRepository paciente.PacienteRepository,
) PacienteService {
	return &pacienteService{
		repository: pacienteRepository,
	}
}

type pacienteService struct {
	repository paciente.PacienteRepository
}

type PacienteService interface {
	ListarPacientes(page, pageSize int, search string) ([]*model.Paciente, int, error)
	GetPacienteByID(id string) (*model.Paciente, error)
	GetAnamneseByPacienteID(pacienteId string) (*model.Anamnese, error)
}

func (ps *pacienteService) ListarPacientes(page, pageSize int, search string) ([]*model.Paciente, int, error) {
    if page < 1 {
        page = 1
    }
    
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    return ps.repository.ListarPacientes(page, pageSize, search)
}

func (ps *pacienteService) GetPacienteByID(id string) (*model.Paciente, error) {
	paciente, err := ps.repository.FindPacienteByID(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar paciente: %w", err)
	}

	return paciente, nil
}

func (ps *pacienteService) GetAnamneseByPacienteID(pacienteId string) (*model.Anamnese, error) {
	anamnese, err := ps.repository.FindAnamneseByPacienteID(pacienteId)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar anamnese: %w", err)
	}

	return anamnese, nil
}