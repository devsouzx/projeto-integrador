package paciente

import (
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