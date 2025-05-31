package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
) {
	// Login paciente
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login-paciente.html", nil)
	})

	// Login profissionais
	r.GET("/login-profissionais", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login-profissionais.html", nil)
	})

	// Cadastro Paciente
	r.GET("/cadastro", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cadastro.html", nil)
	})

	// Recupera senha
	r.GET("/recuperar-senha", func(c *gin.Context) {
		c.HTML(http.StatusOK, "recuperar-senha.html", nil)
	})
}