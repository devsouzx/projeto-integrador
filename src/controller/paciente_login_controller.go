package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/model"
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
}

func (uc *userController) LoginUser(c *gin.Context) {
	var userRequest request.UserLogin
	fmt.Println("User Request:", userRequest)

	if err := c.ShouldBind(&userRequest); err != nil {
		c.String(http.StatusBadRequest, "Error trying to validate user info", err)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Identificador,
		userRequest.Password,
	)
	fmt.Println("Domain User:", domain)

	user, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error trying to call loginUser service": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	redirectPath := fmt.Sprintf("/%s/dashboard", strings.ToLower(user.GetRole()))
	c.Redirect(http.StatusFound, redirectPath)
}
