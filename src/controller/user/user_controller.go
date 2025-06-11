package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	userService "github.com/devsouzx/projeto-integrador/src/service/user"
	"github.com/gin-gonic/gin"
)

func NewUserController(
	userService userService.UserDomainService,
) UserController {
	return &userController{
		service: userService,
	}
}

type userController struct {
	service userService.UserDomainService
}

type UserController interface {
	LoginUser(c *gin.Context)
	SendCodeRecovey(c *gin.Context)
	VerifyCode(c *gin.Context)
	ResetPassword(c *gin.Context)
}

func (uc *userController) LoginUser(c *gin.Context) {
	var userRequest request.LoginRequest

    if err := c.ShouldBind(&userRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Erro ao validar dados do usuário",
            "details": err.Error(),
        })
        return
    }

    user, token, err := uc.service.LoginUserService(userRequest)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Login/senha invalidos!",
            "details": err.Error(),
        })
        return
    }

    c.SetCookie("token", token, 3600, "/", "", false, true)
    c.JSON(http.StatusOK, gin.H{
        "message": "Login realizado com sucesso",
        "token": token,
        "user": gin.H{
            "id": user.GetID(),
            "name": user.GetName(),
            "email": user.GetEmail(),
            "role": user.GetRole(),
        },
        "redirect": fmt.Sprintf("/%s/dashboard", strings.ToLower(user.GetRole())),
    })
}

func (uc *userController) SendCodeRecovey(c *gin.Context) {
	var recoveyRequest request.PasswordRecoveryRequest

	if err := c.ShouldBindJSON(&recoveyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
        return
	}

	err := uc.service.SendCodeRecoveryService(recoveyRequest)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Código de recuperação enviado com sucesso"})
}

func (uc *userController) VerifyCode(c *gin.Context) {
	var verifyCodeRequest request.VerifyCodeRequest

	if err := c.ShouldBindJSON(&verifyCodeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
	}

	err := uc.service.VerifyCode(verifyCodeRequest)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "Código verificado com sucesso"})
}

func (uc *userController) ResetPassword(c *gin.Context) {
    var request request.ResetPasswordRequest
    
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }
    
    err := uc.service.ResetPassword(request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Senha redefinida com sucesso"})
}
