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
	InitMedicoRoutes(rg, container.MedicoController, container.PacienteController, container.EncaminhamentoController)
	InitEnfermeiroRoutes(rg, container.EnfermeiroController)
	InitAgenteRoutes(rg, container.AgenteController)
	InitGestorRoutes(rg, container.GestorController)
	InitLaudoRoutes(rg, container.LaudoController)

	rg.Use(model.VerifyTokenMiddleware)
    {
        rg.GET("/perfil", container.UserController.RenderPerfil)
        rg.GET("/configuracoes", container.UserController.RenderConfiguracoes)
    }

	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	rg.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico", "enfermeiro"))
	{
		rg.GET("/buscar-paciente", container.PacienteController.BuscarPacientePorCPF)
		rg.POST("/nova-ficha", container.FichaController.Create)
		rg.GET("/pacientes-lista", container.PacienteController.ListarPacientes)
		rg.GET("/api/pacientes/:id", container.PacienteController.GetPaciente)
	}

	// Inicializa Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	
}
