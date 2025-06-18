package paciente

import (
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