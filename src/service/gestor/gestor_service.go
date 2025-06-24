package gestor

import (
	"github.com/devsouzx/projeto-integrador/src/repository/gestor"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
)

func NewGestorService(
	gestorRepository gestor.GestorRepositoryInterface,
) GestorServiceInterface {
	return &gestorService{
		gestorRepository: gestorRepository,
	}
}

type gestorService struct {
	gestorRepository gestor.GestorRepositoryInterface
}

type GestorServiceInterface interface {
	CadastrarMedico(req request.CadastroProfissionalRequest) (*model.Medico, error)
}

func (s *gestorService) CadastrarMedico(req request.CadastroProfissionalRequest) (*model.Medico, error) {
	return s.gestorRepository.CadastrarMedico(req)
}
