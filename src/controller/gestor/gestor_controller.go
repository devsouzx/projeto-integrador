package gestor

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/service/gestor"
	"github.com/gin-gonic/gin"
)

func NewGestorController(
	gestorService gestor.GestorServiceInterface,
) GestorControllerInterface {
	return &gestorController{
		gestorService: gestorService,
	}
}

type gestorController struct {
	gestorService gestor.GestorServiceInterface
}

type GestorControllerInterface interface {
	CadastrarProfissional(c *gin.Context)
	RenderDashboard(c *gin.Context)
	RenderNovoUsuario(c *gin.Context)
	RenderUsuarios(c *gin.Context)
	RenderRelatorios(c *gin.Context)
	RenderExames(c *gin.Context)
}

func (gc *gestorController) RenderDashboard(c *gin.Context) {
	gestor, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "dashboard-gestor.html", gin.H{
		"Gestor": gestor,
	})
}

func (gc *gestorController) RenderNovoUsuario(c *gin.Context) {
	gestor, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "novo-usuario-gestor.html", gin.H{
		"Gestor": gestor,
	})
}

func (gc *gestorController) RenderUsuarios(c *gin.Context) {
	gestor, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "usuarios-gestor.html", gin.H{
		"Gestor": gestor,
	})
}

func (gc *gestorController) RenderRelatorios(c *gin.Context) {
	gestor, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "relatorios-gestor.html", gin.H{
		"Gestor": gestor,
	})
}

func (gc *gestorController) RenderExames(c *gin.Context) {
	gestor, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "exames-gestor.html", gin.H{
		"Gestor": gestor,
	})
}