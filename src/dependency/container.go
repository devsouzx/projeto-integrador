package dependency

import (
	"database/sql"

	enfermeiroController "github.com/devsouzx/projeto-integrador/src/controller/enfermeiro"
	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	gestorController "github.com/devsouzx/projeto-integrador/src/controller/gestor"
	laudoController "github.com/devsouzx/projeto-integrador/src/controller/laudo"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	pacienteController "github.com/devsouzx/projeto-integrador/src/controller/paciente"
	unidadeController "github.com/devsouzx/projeto-integrador/src/controller/unidade"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	enfermeiroRepository "github.com/devsouzx/projeto-integrador/src/repository/enfermeiro"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	laudoRepository "github.com/devsouzx/projeto-integrador/src/repository/laudo"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	pacienteRepository "github.com/devsouzx/projeto-integrador/src/repository/paciente"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	gestorRepository "github.com/devsouzx/projeto-integrador/src/repository/gestor"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	enfermeiroService "github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	gestorService "github.com/devsouzx/projeto-integrador/src/service/gestor"
	laudoService "github.com/devsouzx/projeto-integrador/src/service/laudo"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/notifications"
	pacienteService "github.com/devsouzx/projeto-integrador/src/service/paciente"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
	agendamentoRepository "github.com/devsouzx/projeto-integrador/src/repository/agendamento"
	agendamentoService "github.com/devsouzx/projeto-integrador/src/service/agendamento"
	encaminhamentoService "github.com/devsouzx/projeto-integrador/src/service/encaminhamento"
	encaminhamentoRepository "github.com/devsouzx/projeto-integrador/src/repository/encaminhamento"
	encaminhamentoController "github.com/devsouzx/projeto-integrador/src/controller/encaminhamento"
)

type Container struct {
	UserController       userController.UserController
	MedicoController     medicoController.MedicoControllerInterface
	FichaController      fichaController.FichaController
	UnidadeController    unidadeController.UnidadeController
	PacienteController   pacienteController.PacienteControllerInterface
	EnfermeiroController enfermeiroController.EnfermeiroControllerInterface
	LaudoController      laudoController.LaudoControllerInterface
	GestorController     gestorController.GestorControllerInterface
	EncaminhamentoController encaminhamentoController.EncaminhamentoControllerInterface
}

func InitContainer(db *sql.DB) *Container {
	// Incializa Repositories
	userRepository := userRepository.NewUserRepository(db)
	pacienteRepository := pacienteRepository.NewPacienteRepository(db)
	medicoRepository := medicoRepository.NewMedicoRepository(db)
	fichRepository := fichaRepository.NewFichaRepository(db)
	enfermeiroRepository := enfermeiroRepository.NewEnfermeiroRepository(db)
	laudoRepository := laudoRepository.NewLaudoRepository(db)
	gestorRepository := gestorRepository.NewGestorRepository(db)
	agendamentoRepo := agendamentoRepository.NewAgendamentoRepository(db)
	encaminhamentoRepository := encaminhamentoRepository.NewEncaminhamentoRepository(db)

	// Inicializa Services
	emailService := emailService.NewEmailService()
	unidadeService := datasus.NewCNESService()
	smsService := notifications.NewSMSService()
	pacienteService := pacienteService.NewPacienteService(pacienteRepository)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	enfermeiroService := enfermeiroService.NewEnfermeiroService(enfermeiroRepository)
	notificationsService := notifications.NewNotificationService(smsService)
	laudoService := laudoService.NewLaudoService(laudoRepository)
	agendamentoService := agendamentoService.NewAgendamentoService(agendamentoRepo)
	gestorService := gestorService.NewGestorService(gestorRepository)
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
	encaminhamentoService := encaminhamentoService.NewEncaminhamentoService(encaminhamentoRepository)

	// Inicializa Controllers
	unidadeController := unidadeController.NewUnidadeController(unidadeService)
	pacienteController := pacienteController.NewPacienteController(pacienteService)
	fichaController := fichaController.NewFichaController(fichaService)
	gestorController := gestorController.NewGestorController(gestorService)
	enfermeiroController := enfermeiroController.NewEnfermeiroController(
		enfermeiroService,
		*unidadeService,
		agendamentoService,
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
	laudoController := laudoController.NewLaudoController(
		laudoService,
		pacienteService,
		medicoService,
	)
	enacaminhamentoController := encaminhamentoController.NewEncaminhamentoController(
		encaminhamentoService,
		pacienteService,
		medicoService,
	)

	return &Container{
		UserController:       userController,
		MedicoController:     medicoController,
		FichaController:      fichaController,
		UnidadeController:    unidadeController,
		PacienteController:   pacienteController,
		EnfermeiroController: enfermeiroController,
		LaudoController:      laudoController,
		GestorController: gestorController,
		EncaminhamentoController: enacaminhamentoController,
	}
}
