package ficha

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
)

func NewFichaService(
	fichaRepository fichaRepository.FichaRepositoryInterface,
	userRepository userRepository.UserRepository,
) FichaServiceInterface {
	return &fichaService{
		fichaRepository: fichaRepository,
		userRepository: userRepository,
	}
}

type fichaService struct {
	fichaRepository fichaRepository.FichaRepositoryInterface
	userRepository userRepository.UserRepository
}

type FichaServiceInterface interface {
	CreateFicha(request *model.FichaRequest) (*model.FichaCitopatologica, error)
	// GetFichaByID(id string) (*model.Ficha, error)
	// GetFichasByPaciente(pacienteID string) ([]*model.Ficha, error)
	// UpdateFicha(ficha *model.Ficha) error
}

func (fs *fichaService) CreateFicha(request *model.FichaRequest) (*model.FichaCitopatologica, error) {
	paciente, err := fs.userRepository.FindUserByIdentifier(request.Paciente.CPF, "paciente")
	if err != nil {
		paciente, err = fs.userRepository.CreatePaciente(&request.Paciente)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar paciente: %v", err)
		}
	}

	request.DadosResidenciais.PacienteID = paciente.GetID()
	err = fs.fichaRepository.CreateEndereco(&request.DadosResidenciais)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar dados residenciais: %v", err)
	}

	request.Ficha.PacienteID = paciente.GetID()
	ficha, err := fs.fichaRepository.Create(&request.Ficha)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar ficha: %v", err)
	}
	
	return ficha, nil
}

