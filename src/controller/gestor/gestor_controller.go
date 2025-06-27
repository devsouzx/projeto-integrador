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
	c.HTML(http.StatusOK, "dashboard-gestor.html", nil)
}

func (gc *gestorController) RenderNovoUsuario(c *gin.Context) {
	c.HTML(http.StatusOK, "novo-usuario-gestor.html", nil)
}

func (gc *gestorController) RenderUsuarios(c *gin.Context) {
	c.HTML(http.StatusOK, "usuarios-gestor.html", nil)
}

func (gc *gestorController) RenderRelatorios(c *gin.Context) {
	c.HTML(http.StatusOK, "relatorios-gestor.html", nil)
}

func (gc *gestorController) RenderExames(c *gin.Context) {
	c.HTML(http.StatusOK, "exames-gestor.html", nil)
}