package dependency

import (
	"database/sql"
	
	enfermeiroController "github.com/devsouzx/projeto-integrador/src/controller/enfermeiro"
	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	pacienteController "github.com/devsouzx/projeto-integrador/src/controller/paciente"
	unidadeController "github.com/devsouzx/projeto-integrador/src/controller/unidade"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	enfermeiroRepository "github.com/devsouzx/projeto-integrador/src/repository/enfermeiro"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	pacienteRepository "github.com/devsouzx/projeto-integrador/src/repository/paciente"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	enfermeiroService "github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/notifications"
	pacienteService "github.com/devsouzx/projeto-integrador/src/service/paciente"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController       userController.UserController
	MedicoController     medicoController.MedicoControllerInterface
	FichaController      fichaController.FichaController
	UnidadeController    unidadeController.UnidadeController
	PacienteController   pacienteController.PacienteControllerInterface
	EnfermeiroController enfermeiroController.EnfermeiroControllerInterface
}

func InitContainer(db *sql.DB) *Container {
	// Incializa Repositories
	userRepository := userRepository.NewUserRepository(db)
	pacienteRepository := pacienteRepository.NewPacienteRepository(db)
	medicoRepository := medicoRepository.NewMedicoRepository(db)
	fichRepository := fichaRepository.NewFichaRepository(db)
	enfermeiroRepository := enfermeiroRepository.NewEnfermeiroRepository(db)

	// Inicializa Services
	emailService := emailService.NewEmailService()
	unidadeService := datasus.NewCNESService()
	smsService := notifications.NewSMSService()
	pacienteService := pacienteService.NewPacienteService(pacienteRepository)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	enfermeiroService := enfermeiroService.NewEnfermeiroService(enfermeiroRepository)
	notificationsService := notifications.NewNotificationService(smsService)
	userService := userService.NewUserDomainService(
		userRepository,
		emailService,
		pacienteRepository,
	)
	fichaService := fichaService.NewFichaService(
		fichRepository,
		userRepository,
		pacienteRepository,
		*notificationsService,
	)

	// Inicializa Controllers
	unidadeController := unidadeController.NewUnidadeController(unidadeService)
	pacienteController := pacienteController.NewPacienteController(pacienteService)
	fichaController := fichaController.NewFichaController(fichaService)
	enfermeiroController := enfermeiroController.NewEnfermeiroController(
		enfermeiroService,
		*unidadeService,
	)
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
		UserController:       userController,
		MedicoController:     medicoController,
		FichaController:      fichaController,
		UnidadeController:    unidadeController,
		PacienteController:   pacienteController,
		EnfermeiroController: enfermeiroController,
	}
}
