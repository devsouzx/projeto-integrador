package dependency

import (
	"database/sql"

	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	pacienteController "github.com/devsouzx/projeto-integrador/src/controller/paciente"
	unidadeController "github.com/devsouzx/projeto-integrador/src/controller/unidade"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	"github.com/devsouzx/projeto-integrador/src/repository/paciente"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	pacienteService "github.com/devsouzx/projeto-integrador/src/service/paciente"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController     userController.UserController
	MedicoController   medicoController.MedicoControllerInterface
	FichaController    fichaController.FichaController
	UnidadeController  unidadeController.UnidadeController
	PacienteController pacienteController.PacienteControllerInterface
}

func InitContainer(db *sql.DB) *Container {
	// Incializa Repositories
	userRepository := userRepository.NewUserRepository(db)
	pacienteRepository := paciente.NewPacienteRepository(db)
	medicoRepository := medicoRepository.NewMedicoRepository(db)
	fichRepository := fichaRepository.NewFichaRepository(db)

	// Inicializa Services
	emailService := emailService.NewEmailService()
	unidadeService := datasus.NewCNESService()
	pacienteService := pacienteService.NewPacienteService(pacienteRepository)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	userService := userService.NewUserDomainService(
		userRepository,
		emailService,
		pacienteRepository,
	)
	fichaService := fichaService.NewFichaService(
		fichRepository,
		userRepository,
		pacienteRepository,
	)

	// Inicializa Controllers
	unidadeController := unidadeController.NewUnidadeController(unidadeService)
	pacienteController := pacienteController.NewPacienteController(pacienteService)
	fichaController := fichaController.NewFichaController(fichaService)
	userController := userController.NewUserController(
		userService, 
		emailService,
	)
	medicoController := medicoController.NewMedicoController(
		medicoService, 
		*unidadeService, 
		pacienteService,
	)

	return &Container{
		UserController:     userController,
		MedicoController:   medicoController,
		FichaController:    fichaController,
		UnidadeController:  unidadeController,
		PacienteController: pacienteController,
	}
}
