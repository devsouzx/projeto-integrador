package user

import (
	"github.com/devsouzx/projeto-integrador/src/service/email"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
	"github.com/gin-gonic/gin"
)

func NewUserController(
	userService userService.UserDomainService,
	emailService email.EmailService,
) UserController {
	return &userController{
		service:      userService,
		emailService: emailService,
	}
}

type userController struct {
	service      userService.UserDomainService
	emailService email.EmailService
}

type UserController interface {
	LoginUser(c *gin.Context)
	SendCodeRecovey(c *gin.Context)
	VerifyCode(c *gin.Context)
	ResetPassword(c *gin.Context)
	Logout(c *gin.Context)
	CadastrarUsuario(c *gin.Context)
	VerifyAccount(c *gin.Context)
}