package dependency

import (
	"database/sql"

	fichaController "github.com/devsouzx/projeto-integrador/src/controller/ficha"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	fichaRepository "github.com/devsouzx/projeto-integrador/src/repository/ficha"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController userController.UserController
	MedicoController medicoController.MedicoControllerInterface
	FichaController fichaController.FichaController
}

func InitContainer(db *sql.DB) *Container {
	emailService := emailService.NewEmailService()
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserDomainService(userRepository, emailService)
	userController := userController.NewUserController(userService, emailService)

	medicoRepository := medicoRepository.NewMedicoRepository(db)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	medicoController := medicoController.NewMedicoController(medicoService)

	fichRepository := fichaRepository.NewFichaRepository(db)
	fichaService := fichaService.NewFichaService(fichRepository, userRepository)
	fichaController := fichaController.NewFichaController(fichaService)

	return &Container{
		UserController: userController,
		MedicoController: medicoController,
		FichaController: fichaController,
	}
}
