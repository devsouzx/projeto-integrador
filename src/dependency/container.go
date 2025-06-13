package dependency

import (
	"database/sql"

	medicoRepository "github.com/devsouzx/projeto-integrador/src/repository/medico"
	medicoService "github.com/devsouzx/projeto-integrador/src/service/medico"
	medicoController "github.com/devsouzx/projeto-integrador/src/controller/medico"
	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController userController.UserController
	MedicoController medicoController.MedicoControllerInterface
}

func InitContainer(db *sql.DB) *Container {
	emailService := emailService.NewEmailService()
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserDomainService(userRepository, emailService)
	userController := userController.NewUserController(userService)

	medicoRepository := medicoRepository.NewMedicoRepository(db)
	medicoService := medicoService.NewMedicoService(medicoRepository)
	medicoController := medicoController.NewMedicoController(medicoService)

	return &Container{
		UserController: userController,
		MedicoController: medicoController,
	}
}
