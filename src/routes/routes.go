package routes

import (
	"github.com/devsouzx/projeto-integrador/docs"
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(rg *gin.RouterGroup, container *dependency.Container) {
	InitAuthRoutes(rg, container.UserController)
	InitPacienteRoutes(rg, container.PacienteController)
	InitMedicoRoutes(rg, container.MedicoController, container.FichaController, container.PacienteController)
	InitEnfermeiroRoutes(rg)
	InitAgenteRoutes(rg)
	InitGestorRoutes(rg)

	// Inicializa Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
