package ficha

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	fichaService "github.com/devsouzx/projeto-integrador/src/service/ficha"
	"github.com/gin-gonic/gin"
)

type FichaController interface {
	Create(c *gin.Context)
	//GetByID(c *gin.Context)
	//GetByPaciente(c *gin.Context)
	// Update(c *gin.Context)
}

type fichaController struct {
	service fichaService.FichaServiceInterface
}

func NewFichaController(service fichaService.FichaServiceInterface) FichaController {
	return &fichaController{service: service}
}

func (fc *fichaController) Create(c *gin.Context) {
    var request model.FichaRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Erro ao processar os dados",
            "details": err.Error(),
        })
        return
    }
	
    medico, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
        return
    }
    request.Ficha.ResponsavelID = medico.(model.BaseUser).ID

    createdFicha, err := fc.service.CreateFicha(&request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdFicha)
}