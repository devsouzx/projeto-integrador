package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"github.com/gin-gonic/gin"
)

type UnidadeController struct {
	cnesService *datasus.CNESService
}

func NewUnidadeController(cnesService *datasus.CNESService) *UnidadeController {
	return &UnidadeController{
		cnesService: cnesService,
	}
}

func (uc *UnidadeController) ListarUnidades(c *gin.Context) {
	codigoIBGE := c.Query("municipio")
	if codigoIBGE == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'municipio' é obrigatório"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	if perPage < 1 {
		perPage = 10
	}

	if perPage > 100 {
		perPage = 100
	}

	unidades, total, err := uc.cnesService.BuscarUnidadesPorMunicipio(codigoIBGE, perPage, (page-1)*perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Falha ao buscar unidades",
			"details": err.Error(),
		})
		return
	}

	totalPages := 0
	if total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(perPage)))
	}

	var nextPage, prevPage interface{}
	if page < totalPages {
		nextPage = page + 1
	} else {
		nextPage = nil
	}

	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"data": unidades,
		"pagination": gin.H{
			"total":         total,
			"per_page":      perPage,
			"current_page":  page,
			"total_pages":   totalPages,
			"next_page":     nextPage,
			"previous_page": prevPage,
		},
	})
}
