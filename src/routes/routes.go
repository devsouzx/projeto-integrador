package routes

import (
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, container *dependency.Container) {
	InitAuthRoutes(rg, container.UserController)
	InitPacienteRoutes(rg, container.PacienteController)
	InitMedicoRoutes(rg, container.MedicoController, container.FichaController, container.PacienteController)
	InitEnfermeiroRoutes(rg)
	InitAgenteRoutes(rg)
	InitGestorRoutes(rg)
}