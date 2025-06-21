package routes

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/controller/enfermeiro"
	"github.com/gin-gonic/gin"
)

func InitEnfermeiroRoutes(rg *gin.RouterGroup, enfermeiroController enfermeiro.EnfermeiroControllerInterface) {
	enfermeira := rg.Group("/enfermeiro")
	enfermeira.Use(model.VerifyTokenMiddleware, model.AuthorizeRole("enfermeiro"))
	{
		enfermeira.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-enfermeira.html", nil)
		})

		enfermeira.GET("/nova-ficha", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nova-ficha-enfermeira.html", nil)
		})

		enfermeira.GET("/editar-ficha", func(c *gin.Context) {
			c.HTML(http.StatusOK, "editar-ficha-enfermeira.html", nil)
		})

		enfermeira.GET("/consultar", func(c *gin.Context) {
			c.HTML(http.StatusOK, "consultar-enfermeira.html", nil)
		})

		enfermeira.GET("/agendamentos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "agendamentos-enfermeira.html", nil)
		})

		enfermeira.GET("/relatorios", func(c *gin.Context) {
			c.HTML(http.StatusOK, "relatorios-enfermeira.html", nil)
		})

		enfermeira.GET("/minha-unidade", enfermeiroController.GetUnidadeEnfermeiro)
	}
}