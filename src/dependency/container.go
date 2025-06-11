package dependency

import (
	"database/sql"

	userController "github.com/devsouzx/projeto-integrador/src/controller/user"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	emailService"github.com/devsouzx/projeto-integrador/src/service/email"
	userService"github.com/devsouzx/projeto-integrador/src/service/user"
)

type Container struct {
	UserController userController.UserController
}

func InitContainer(db *sql.DB) *Container {
	emailService := emailService.NewEmailService()
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserDomainService(userRepository, emailService)
	userController := userController.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
