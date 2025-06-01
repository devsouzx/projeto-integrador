package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
  r.GET("/login", func(c *gin.Context) {
    c.HTML(http.StatusOK, "login-paciente.html", nil)
  })

  r.GET("/login-profissionais", func(c *gin.Context) {
    c.HTML(http.StatusOK, "login-profissionais.html", nil)
  })

  r.GET("/cadastro", func(c *gin.Context) {
    c.HTML(http.StatusOK, "cadastro-paciente.html", nil)
  })

  r.GET("/recuperar-senha", func(c *gin.Context) {
    c.HTML(http.StatusOK, "recuperar-senha.html", nil)
  })

  medico := r.Group("/medico")
	{
		medico.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-medico.html", nil)
		})

		medico.GET("/nova-ficha", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nova-ficha-medico.html", nil)
		})

		medico.GET("/pacientes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "pacientes-lista-medico.html", nil)
		})

		medico.GET("/laudos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "laudos-lista-medico.html", nil)
		})

		medico.GET("/estatisticas", func(c *gin.Context) {
			c.HTML(http.StatusOK, "estatisticas-medico.html", nil)
		})

		medico.GET("/encaminhamentos", func(c *gin.Context) {
			c.HTML(http.StatusOK, "encaminhamentos-lista-medico.html", nil)
		})
	}
}