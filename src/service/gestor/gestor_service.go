package agente

import (
	"projeto-integrador/internal/model"
	"projeto-integrador/internal/repository/gestor"

	"github.com/devsouzx/projeto-integrador/src/model/request"
)

func NewGestorService(
	gestorRepository gestor.AgenteRepository,
) GestorServiceInterface {
	return &gestorService{
		gestorRepository: gestorRepository,
	}
}

type gestorService struct {
	gestorRepository gestor.GestorRepository
}

type GestorServiceInterface interface {
	CadastrarMedico(req request.CadastroMedicoRequest) (*model.Medico, error)
}

func (s *gestorService) CadastrarMedico(req request.CadastroMedicoRequest) (*model.Medico, error) {
	return s.gestorRepository.CadastrarMedico(medico)
}
