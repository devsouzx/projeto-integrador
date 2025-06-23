package routes

import (
	"github.com/devsouzx/projeto-integrador/docs"
	"github.com/devsouzx/projeto-integrador/src/dependency"
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(rg *gin.RouterGroup, container *dependency.Container) {
	InitAuthRoutes(rg, container.UserController)
	InitPacienteRoutes(rg, container.PacienteController)
	InitMedicoRoutes(rg, container.MedicoController, container.PacienteController)
	InitEnfermeiroRoutes(rg, container.EnfermeiroController)
	InitAgenteRoutes(rg)
	InitGestorRoutes(rg)
	InitLaudoRoutes(rg, container.LaudoController)

	rg.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico", "enfermeiro"))
	{
		rg.GET("/buscar-paciente", container.PacienteController.BuscarPacientePorCPF)
		rg.POST("/nova-ficha", container.FichaController.Create)
		rg.GET("/pacientes-lista", container.PacienteController.ListarPacientes)
		rg.GET("/api/pacientes/:id", container.PacienteController.GetPaciente)
	}

	// Inicializa Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
