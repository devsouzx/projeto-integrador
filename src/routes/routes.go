package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	auth := r.Group("/")
	{
		auth.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login-paciente.html", nil)
		})

		auth.GET("/login-profissionais", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login-profissionais.html", nil)
		})

		auth.GET("/cadastro", func(c *gin.Context) {
			c.HTML(http.StatusOK, "cadastro-paciente.html", nil)
		})

		auth.GET("/recuperar-senha", func(c *gin.Context) {
			c.HTML(http.StatusOK, "recuperar-senha.html", nil)
		})
	}

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

		medico.GET("/encaminhamentos/novo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "encaminhamento-medico.html", nil)
		})

		medico.GET("/laudos/novo", func(c *gin.Context) {
			c.HTML(http.StatusOK, "laudo-medico.html", nil)
		})

		medico.GET("/medicos/pacientes/123456", func(c *gin.Context) {
			c.HTML(http.StatusOK, "paciente-medico.html", nil)
		})
	}	

	enfermeira := r.Group("/enfermeiro")
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
	}

	paciente := r.Group("/paciente")
	{
		paciente.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-paciente.html", nil)
		})

		paciente.GET("/agendar", func(c *gin.Context) {
			c.HTML(http.StatusOK, "agendamento-paciente.html", nil)
		})

		paciente.GET("/historico_exames", func(c *gin.Context) {
			c.HTML(http.StatusOK, "historico-exames-paciente.html", nil)
		})

		paciente.GET("/localizar_ubs", func(c *gin.Context) {
			c.HTML(http.StatusOK, "localizar-ubs-paciente.html", nil)
		})

		paciente.GET("/notificacoes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "notificacoes-paciente.html", nil)
		})

		paciente.GET("/orientacoes", func(c *gin.Context) {
			c.HTML(http.StatusOK, "orientacoes-paciente.html", nil)
		})
	}

	gestor := r.Group("/gestor")
	{
		gestor.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard-gestor.html", nil)
		})

		gestor.GET("/novo-usuario", func(c *gin.Context) {
			c.HTML(http.StatusOK, "novo-usuario-gestor.html", nil)
		})

		gestor.GET("/usuarios", func(c *gin.Context) {	
			c.HTML(http.StatusOK, "usuarios-gestor.html", nil)
		})

		gestor.GET("/relatorios", func(c *gin.Context) {
			c.HTML(http.StatusOK, "relatorios-gestor.html", nil)
		})

		gestor.GET("/exames", func(c *gin.Context) {
			c.HTML(http.StatusOK, "exames-gestor.html", nil)
		})
	}
}
