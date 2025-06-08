package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/devsouzx/projeto-integrador/src/service"
	"github.com/gin-gonic/gin"
)

func NewUserController(
	userService service.UserDomainService,
) UserController {
	return &userController{
		service: userService,
	}
}

type userController struct {
	service service.UserDomainService
}

type UserController interface {
	LoginUser(c *gin.Context)
	SendCodeRecovey(c *gin.Context)
}

func (uc *userController) LoginUser(c *gin.Context) {
	var userRequest request.LoginRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.String(http.StatusBadRequest, "Error trying to validate user info", err)
		return
	}
	fmt.Println(userRequest)

	user, token, err := uc.service.LoginUserService(userRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error trying to call loginUser service": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	redirectPath := fmt.Sprintf("/%s/dashboard", strings.ToLower(user.GetRole()))
	c.Redirect(http.StatusFound, redirectPath)
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
