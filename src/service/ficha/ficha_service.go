package ficha

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	"github.com/devsouzx/projeto-integrador/src/repository/paciente"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
)

func NewFichaService(
	fichaRepository fichaRepository.FichaRepositoryInterface,
	userRepository userRepository.UserRepository,
	pacienteRepository paciente.PacienteRepository,
) FichaServiceInterface {
	return &fichaService{
		fichaRepository: fichaRepository,
		userRepository:  userRepository,
		pacienteRepository: pacienteRepository,
	}
}

type fichaService struct {
	fichaRepository fichaRepository.FichaRepositoryInterface
	userRepository userRepository.UserRepository
	pacienteRepository paciente.PacienteRepository
}

type FichaServiceInterface interface {
	CreateFicha(request *request.FichaRequest) (*model.FichaCitopatologica, error)
}

func (fs *fichaService) CreateFicha(request *request.FichaRequest) (*model.FichaCitopatologica, error) {
	paciente, err := fs.userRepository.FindUserByIdentifier(request.Paciente.CPF, "paciente")
    if err != nil {
        paciente, err = fs.pacienteRepository.CreatePacienteficha(&request.Paciente)
        if err != nil {
            return nil, fmt.Errorf("erro ao criar paciente: %v", err)
        }
    } else {
        pacienteModel := paciente.(*model.Paciente)
        pacienteModel.Name = request.Paciente.Name
        pacienteModel.DataNascimento = request.Paciente.DataNascimento
        err := fs.pacienteRepository.UpdatePaciente(pacienteModel)
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

	request.ExameClinico.FichaID = ficha.ID
	err = fs.fichaRepository.RegistrarExame(&request.ExameClinico, pacienteId)
	if err != nil {
		return nil, fmt.Errorf("erro ao registrar dados do exame: %v", err)
	}

	return ficha, nil
}
