package dependency

import (
	"database/sql"

	unidadeController "github.com/devsouzx/projeto-integrador/src/controller/unidade"
	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController userController.UserController
	MedicoController medicoController.MedicoControllerInterface
	FichaController fichaController.FichaController
	UnidadeController unidadeController.UnidadeController
}

func InitContainer(db *sql.DB) *Container {
	emailService := emailService.NewEmailService()
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserDomainService(userRepository, emailService)
	userController := userController.NewUserController(userService)

	medicoRepository := medicoRepository.NewMedicoRepository(db)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	medicoController := medicoController.NewMedicoController(medicoService)

	fichRepository := fichaRepository.NewFichaRepository(db)
	fichaService := fichaService.NewFichaService(fichRepository, userRepository)
	fichaController := fichaController.NewFichaController(fichaService)

	unidadeService := datasus.NewCNESService()
	unidadeController := unidadeController.NewUnidadeController(unidadeService)

	return &Container{
		UserController: userController,
		MedicoController: medicoController,
		FichaController: fichaController,
		UnidadeController: *unidadeController,
	}
}
