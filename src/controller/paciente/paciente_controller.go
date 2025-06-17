package paciente

import (
	"fmt"
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	"github.com/gin-gonic/gin"
)

func NewPacienteController(
	userRepository userRepository.UserRepository,
) PacienteControllerInterface {
	return &pacienteController{
		userRepository: userRepository,
	}
}

type pacienteController struct {
	userRepository userRepository.UserRepository
}

type PacienteControllerInterface interface {
	BuscarPacientePorCPF(c *gin.Context)
}

func (pc *pacienteController) BuscarPacientePorCPF(c *gin.Context) {
    cpf := c.Query("cpf")
    if cpf == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "CPF não informado"})
        return
    }

    pacienteInterface, err := pc.userRepository.FindUserByIdentifier(cpf, "paciente")
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
        return
    }
	fmt.Println(pacienteInterface)
	
    paciente, ok := pacienteInterface.(*model.Paciente)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar dados do paciente"})
        return
    }
	fmt.Println(paciente)
	
    dataNascimento := ""
    if !paciente.NascimentoTime.IsZero() {
        dataNascimento = paciente.NascimentoTime.Format("2006-01-02")
    }
	
    response := gin.H{
        "Name":           paciente.Name,
        "Apelido":        paciente.Apelido,
        "NomeMae":        paciente.NomeMae,
        "CNS":           paciente.CNS,
        "CPF":           paciente.CPF,
        "DataNascimento": dataNascimento,
        "Nacionalidade":  paciente.Nacionalidade,
        "RacaCor":       paciente.RacaCor,
        "Escolaridade":  paciente.Escolaridade,
        "Telefone":      paciente.Telefone,
    }
	
    if paciente.Endereco != nil {
        response["Endereco"] = gin.H{
            "Logradouro":  paciente.Endereco.Logradouro,
            "Numero":      paciente.Endereco.Numero,
            "Complemento": paciente.Endereco.Complemento,
            "Bairro":      paciente.Endereco.Bairro,
            "Cidade":      paciente.Endereco.Cidade,
            "CEP":         paciente.Endereco.CEP,
            "UF":          paciente.Endereco.UF,
        }
    }

	fmt.Println(response)
    c.JSON(http.StatusOK, response)
}