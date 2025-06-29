package dependency

import (
	"database/sql"

	encaminhamentoController "github.com/devsouzx/projeto-integrador/src/controller/encaminhamento"
	enfermeiroController "github.com/devsouzx/projeto-integrador/src/controller/enfermeiro"
	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	gestorController "github.com/devsouzx/projeto-integrador/src/controller/gestor"
	laudoController "github.com/devsouzx/projeto-integrador/src/controller/laudo"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	pacienteController "github.com/devsouzx/projeto-integrador/src/controller/paciente"
	unidadeController "github.com/devsouzx/projeto-integrador/src/controller/unidade"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	agendamentoRepository "github.com/devsouzx/projeto-integrador/src/repository/agendamento"
	agenteRepository "github.com/devsouzx/projeto-integrador/src/repository/agente"
	agenteController "github.com/devsouzx/projeto-integrador/src/controller/agente"
	encaminhamentoRepository "github.com/devsouzx/projeto-integrador/src/repository/encaminhamento"
	enfermeiroRepository "github.com/devsouzx/projeto-integrador/src/repository/enfermeiro"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	gestorRepository "github.com/devsouzx/projeto-integrador/src/repository/gestor"
	laudoRepository "github.com/devsouzx/projeto-integrador/src/repository/laudo"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	pacienteRepository "github.com/devsouzx/projeto-integrador/src/repository/paciente"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	agendamentoService "github.com/devsouzx/projeto-integrador/src/service/agendamento"
	agenteService "github.com/devsouzx/projeto-integrador/src/service/agente"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	encaminhamentoService "github.com/devsouzx/projeto-integrador/src/service/encaminhamento"
	enfermeiroService "github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	gestorService "github.com/devsouzx/projeto-integrador/src/service/gestor"
	laudoService "github.com/devsouzx/projeto-integrador/src/service/laudo"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/notifications"
	pacienteService "github.com/devsouzx/projeto-integrador/src/service/paciente"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController           userController.UserController
	MedicoController         medicoController.MedicoControllerInterface
	FichaController          fichaController.FichaController
	UnidadeController        unidadeController.UnidadeController
	PacienteController       pacienteController.PacienteControllerInterface
	EnfermeiroController     enfermeiroController.EnfermeiroControllerInterface
	LaudoController          laudoController.LaudoControllerInterface
	GestorController         gestorController.GestorControllerInterface
	EncaminhamentoController encaminhamentoController.EncaminhamentoControllerInterface
	AgenteController         agenteController.AgenteControllerInterface
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
	agenteRepository := agenteRepository.NewAgenteRepository(db)

	// Inicializa Services
	emailService := emailService.NewEmailService()
	unidadeService := datasus.NewCNESService()
	smsService := notifications.NewSMSService()
	pacienteService := pacienteService.NewPacienteService(pacienteRepository)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	enfermeiroService := enfermeiroService.NewEnfermeiroService(enfermeiroRepository)
	notificationsService := notifications.NewNotificationService(smsService)
	laudoService := laudoService.NewLaudoService(laudoRepository)
	gestorService := gestorService.NewGestorService(gestorRepository)
	encaminhamentoService := encaminhamentoService.NewEncaminhamentoService(encaminhamentoRepository)
	agenteService := agenteService.NewAgenteService(agenteRepository)
	userService := userService.NewUserDomainService(
		userRepository,
		emailService,
		pacienteRepository,
		smsService,
	)
	fichaService := fichaService.NewFichaService(
		fichRepository,
		userRepository,
		pacienteRepository,
		*notificationsService,
	)
	agendamentoService := agendamentoService.NewAgendamentoService(
		agendamentoRepo, 
		userService,
		smsService,
		emailService,
	)

	// Inicializa Controllers
	unidadeController := unidadeController.NewUnidadeController(unidadeService)
	pacienteController := pacienteController.NewPacienteController(pacienteService, agendamentoService)
	fichaController := fichaController.NewFichaController(fichaService)
	gestorController := gestorController.NewGestorController(gestorService)
	agenteController := agenteController.NewAgenteController(agenteService)
	enfermeiroController := enfermeiroController.NewEnfermeiroController(
		enfermeiroService,
		*unidadeService,
		agendamentoService,
	)
	userController := userController.NewUserController(
		userService,
		emailService,
		pacienteService,
		medicoService,
		enfermeiroService,
		agenteService,
		gestorService,
	)
	medicoController := medicoController.NewMedicoController(
		medicoService,
		*unidadeService,
		pacienteService,
		encaminhamentoService,
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
		UserController:           userController,
		MedicoController:         medicoController,
		FichaController:          fichaController,
		UnidadeController:        unidadeController,
		PacienteController:       pacienteController,
		EnfermeiroController:     enfermeiroController,
		LaudoController:          laudoController,
		GestorController:         gestorController,
		EncaminhamentoController: enacaminhamentoController,
		AgenteController: 	   agenteController,
	}
}
