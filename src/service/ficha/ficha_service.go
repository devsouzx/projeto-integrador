package ficha

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
)

func NewFichaService(
	fichaRepository fichaRepository.FichaRepositoryInterface,
	userRepository userRepository.UserRepository,
) FichaServiceInterface {
	return &fichaService{
		fichaRepository: fichaRepository,
		userRepository:  userRepository,
	}
}

type fichaService struct {
	fichaRepository fichaRepository.FichaRepositoryInterface
	userRepository  userRepository.UserRepository
}

type FichaServiceInterface interface {
	CreateFicha(request *request.FichaRequest) (*model.FichaCitopatologica, error)
}

func (fs *fichaService) CreateFicha(request *request.FichaRequest) (*model.FichaCitopatologica, error) {
	paciente, err := fs.userRepository.FindUserByIdentifier(request.Paciente.CPF, "paciente")
	if err != nil {
		paciente, err = fs.userRepository.CreatePaciente(&request.Paciente)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar paciente: %v", err)
		}
	} else {
		err := fs.userRepository.UpdatePaciente(paciente.(*model.Paciente))
		if err != nil {
			return nil, fmt.Errorf("erro ao atualizar paciente: %v", err)
		}
	}

	pacienteId := paciente.GetID()
	err = fs.fichaRepository.UpsertEndereco(&request.DadosResidenciais, pacienteId)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar/atualizar dados residenciais: %v", err)
	}

	request.Ficha.PacienteID = paciente.GetID()
	ficha, err := fs.fichaRepository.Create(&request.Ficha)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar ficha: %v", err)
	}

	request.Anamnese.FichaID = ficha.ID
	err = fs.fichaRepository.UpsertAnamnase(&request.Anamnese, pacienteId)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar/atualizar dados da anamnase: %v", err)
	}

	return ficha, nil
}
