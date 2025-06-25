package routes

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/enfermeiro"
	"github.com/gin-gonic/gin"
)

func InitEnfermeiroRoutes(rg *gin.RouterGroup, enfermeiroController enfermeiro.EnfermeiroControllerInterface) {
	enfermeira := rg.Group("/enfermeiro")
	enfermeira.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("enfermeiro"))
	{
		enfermeira.GET("/dashboard", enfermeiroController.RenderDashboardPage)
		enfermeira.GET("/nova-ficha", enfermeiroController.RenderNovaFichaPage)
		enfermeira.GET("/editar-ficha", enfermeiroController.RenderEdicaoFichaPage)
		enfermeira.GET("/consultar", enfermeiroController.RenderConsultarPacientesPage)
		enfermeira.GET("/agendamentos", enfermeiroController.RenderAgendamentosPage)
		enfermeira.GET("/relatorios", enfermeiroController.RenderRelatoriosPage)

		enfermeira.GET("/minha-unidade", enfermeiroController.GetUnidadeEnfermeiro)

		enfermeira.POST("/agendamentos", enfermeiroController.CreateAgendamento)
		enfermeira.GET("/agendamentos/listar", enfermeiroController.GetAgendamentos)
		enfermeira.PUT("/agendamentos/:id/status", enfermeiroController.UpdateAgendamentoStatus)
		enfermeira.DELETE("/agendamentos/:id", enfermeiroController.DeleteAgendamento)
	}
}