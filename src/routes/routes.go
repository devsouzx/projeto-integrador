package routes

import (
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, container *dependency.Container) {
	InitAuthRoutes(rg, container.UserController)
	InitPacienteRoutes(rg)
	InitMedicoRoutes(rg, container.MedicoController)
	InitEnfermeiroRoutes(rg)
	InitAgenteRoutes(rg)
	InitGestorRoutes(rg)
}