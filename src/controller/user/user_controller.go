package user

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/agente"
	"github.com/devsouzx/projeto-integrador/src/service/email"
	"github.com/devsouzx/projeto-integrador/src/service/enfermeiro"
	"github.com/devsouzx/projeto-integrador/src/service/gestor"
	"github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
	"github.com/gin-gonic/gin"
)

func NewUserController(
	userService userService.UserDomainService,
	emailService email.EmailService,
    pacienteService paciente.PacienteService,
    medicoService medico.MedicoServiceInterface,
    enfermeiroService enfermeiro.EnfermeiroServiceInterface,
    agenteService agente.AgenteServiceInterface,
    gestorService gestor.GestorServiceInterface,
) UserController {
	return &userController{
		service:      userService,
		emailService: emailService,
		pacienteService: pacienteService,
		medicoService:   medicoService,
		enfermeiroService: enfermeiroService,
		agenteService:   agenteService,
		gestorService:   gestorService,
	}
}

type Mensagem struct {
    Tipo  string `json:"tipo"`
    Texto string `json:"texto"`
}

type userController struct {
	service      userService.UserDomainService
	emailService email.EmailService
    pacienteService paciente.PacienteService
    medicoService   medico.MedicoServiceInterface
    enfermeiroService enfermeiro.EnfermeiroServiceInterface
    agenteService   agente.AgenteServiceInterface
    gestorService   gestor.GestorServiceInterface
}

type UserController interface {
	LoginUser(c *gin.Context)
	SendCodeRecovey(c *gin.Context)
	VerifyCode(c *gin.Context)
	ResetPassword(c *gin.Context)
	Logout(c *gin.Context)
	CadastrarUsuario(c *gin.Context)
	VerifyAccount(c *gin.Context)

	RenderPerfil(c *gin.Context)
	RenderConfiguracoes(c *gin.Context)
}

func (uc *userController) RenderPerfil(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists {
        c.Redirect(http.StatusFound, "/login")
        return
    }

    baseUser := user.(model.BaseUser)

    var detalhes interface{}
    var err error
    
    switch baseUser.Role {
    case "paciente":
        detalhes, err = uc.pacienteService.GetPacienteByID(baseUser.ID)
    case "medico":
        detalhes, err = uc.medicoService.FindMedicoByIdentifier(baseUser.ID)
    case "enfermeiro":
        detalhes, err = uc.enfermeiroService.FindEnfermeiroByIdentifier(baseUser.ID)
    case "agente":
        detalhes, err = uc.agenteService.FindAgenteByIdentifier(baseUser.ID)
    case "gestor":
        detalhes, err = uc.gestorService.FindGestorByIdentifier(baseUser.ID)
    }
    
    if err != nil {
        c.HTML(http.StatusOK, "perfil.html", gin.H{
            "Usuario": baseUser,
            "Mensagem": Mensagem{Tipo: "erro", Texto: "Erro ao carregar dados do perfil"},
        })
        return
    }
    
    c.HTML(http.StatusOK, "perfil.html", gin.H{
        "Usuario": baseUser,
        "Detalhes": detalhes,
    })
}

func (uc *userController) RenderConfiguracoes(c *gin.Context) {
    user, exists := c.Get("user")
    if !exists {
        c.Redirect(http.StatusFound, "/login")
        return
    }

    c.HTML(http.StatusOK, "configuracoes.html", gin.H{
        "Usuario": user,
    })
}