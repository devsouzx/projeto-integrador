package dependency

import (
	"database/sql"

	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	pacienteController "github.com/devsouzx/projeto-integrador/src/controller/paciente"
	unidadeController "github.com/devsouzx/projeto-integrador/src/controller/unidade"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	"github.com/devsouzx/projeto-integrador/src/repository/paciente"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	pacienteService "github.com/devsouzx/projeto-integrador/src/service/paciente"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController userController.UserController
	MedicoController medicoController.MedicoControllerInterface
	FichaController fichaController.FichaController
	UnidadeController unidadeController.UnidadeController
	PacienteController pacienteController.PacienteControllerInterface
}

func InitContainer(db *sql.DB) *Container {
	emailService := emailService.NewEmailService()
	userRepository := userRepository.NewUserRepository(db)
	
	userService := userService.NewUserDomainService(userRepository, emailService)
	userController := userController.NewUserController(userService, emailService)

	unidadeService := datasus.NewCNESService()
	unidadeController := unidadeController.NewUnidadeController(unidadeService)

	medicoRepository := medicoRepository.NewMedicoRepository(db)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	medicoController := medicoController.NewMedicoController(medicoService, *unidadeService)

	fichRepository := fichaRepository.NewFichaRepository(db)
	fichaService := fichaService.NewFichaService(fichRepository, userRepository)
	fichaController := fichaController.NewFichaController(fichaService)

	pacienteRepository := paciente.NewPacienteRepository(db)
	pacienteService := pacienteService.NewPacienteService(pacienteRepository)
	pacienteController := pacienteController.NewPacienteController(pacienteService, userRepository)

	return &Container{
		UserController: userController,
		MedicoController: medicoController,
		FichaController: fichaController,
		UnidadeController: unidadeController,
		PacienteController: pacienteController,
	}
}
