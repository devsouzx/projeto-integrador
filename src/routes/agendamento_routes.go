package routes

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/enfermeiro"
	"github.com/gin-gonic/gin"
)

func InitAgendamentoRoutes(rg *gin.RouterGroup, enfermeiroController enfermeiro.EnfermeiroControllerInterface) {
	agendamento := rg.Group("/agendamentos")
	agendamento.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("efermeiro"))
	{
		
	}
}