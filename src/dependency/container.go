package dependency

import (
	"database/sql"

	"github.com/devsouzx/projeto-integrador/src/controller"
	"github.com/devsouzx/projeto-integrador/src/repository"
	"github.com/devsouzx/projeto-integrador/src/service"
)

type Container struct {
	UserController controller.UserController
}

func InitContainer(db *sql.DB) *Container {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserDomainService(userRepository)
	userController := controller.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
