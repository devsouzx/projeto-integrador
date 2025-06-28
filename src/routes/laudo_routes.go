package routes

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/laudo"
	"github.com/gin-gonic/gin"
)

func InitLaudoRoutes(rg *gin.RouterGroup, laudoController laudo.LaudoControllerInterface) {
	laudos := rg.Group("/laudos")
	laudos.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("medico"))
	{
		laudos.POST("", laudoController.CreateLaudo)
		laudos.GET("/medico/meus-laudos", laudoController.GetLaudosByMedico)
		laudos.GET("/:id/download", laudoController.DownloadLaudo)
	}
}