package gestor

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/devsouzx/projeto-integrador/src/repository/gestor"
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
	CadastrarEnfermeiro(req request.CadastroProfissionalRequest) (*model.Enfermeiro, error)
	CadastrarAgente(req request.CadastroProfissionalRequest) (*model.AgenteComunitario, error)
	CadastrarGestor(req request.CadastroProfissionalRequest) (*model.Gestor, error)
}

func (s *gestorService) CadastrarMedico(req request.CadastroProfissionalRequest) (*model.Medico, error) {
	return s.gestorRepository.CadastrarMedico(req)
}

func (s *gestorService) CadastrarEnfermeiro(req request.CadastroProfissionalRequest) (*model.Enfermeiro, error) {
	return s.gestorRepository.CadastrarEnfermeiro(req)
}

func (s *gestorService) CadastrarAgente(req request.CadastroProfissionalRequest) (*model.AgenteComunitario, error) {
	return s.gestorRepository.CadastrarAgente(req)
}

func (s *gestorService) CadastrarGestor(req request.CadastroProfissionalRequest) (*model.Gestor, error) {
	return s.gestorRepository.CadastrarGestor(req)
}
