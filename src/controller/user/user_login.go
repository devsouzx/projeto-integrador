package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/devsouzx/projeto-integrador/src/model/response"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Login de usuário
// @Description Realiza o login de um usuário e retorna um token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body request.LoginRequest true "Credenciais de login"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} object "Erro na validação dos dados"
// @Failure 401 {object} object "Credenciais inválidas"
// @Failure 403 {object} object "Conta não verificada"
// @Failure 500 {object} object "Erro interno do servidor"
// @Router /login [post]
func (uc *userController) LoginUser(c *gin.Context) {
	var userRequest request.LoginRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao validar dados do usuário",
			"details": err.Error(),
		})
		return
	}

	user, token, err := uc.service.LoginUserService(userRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"details": err.Error(),
		})
		return
	}

	if user.GetRole() == "paciente" && !user.IsVerified() {
		if user.GetVerifyToken() != "" && user.GetTokenExpiresAt().Before(time.Now().UTC()) {
			newToken, err := user.GenerateVerifyToken()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
				return
			}

			if err := uc.service.UpdateUser(user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
				return
			}

			verificationLink := fmt.Sprintf("https://localhost:8080/verify?token=%s", newToken)
			if err := uc.emailService.SendVerificationEmail(user.GetEmail(), user.GetName(), verificationLink, newToken); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar email"})
				return
			}

			c.JSON(http.StatusForbidden, gin.H{
				"error": "Seu link de verificação expirou. Enviamos um novo email para ativação",
				"code":  "new_verification_sent",
			})
			return
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "Conta não verificada. Por favor, verifique seu email",
			"code":  "account_not_verified",
		})
		return
	}

	response := response.LoginResponse{
		Message: "Login realizado com sucesso",
		Token:   token,
		User: response.UserResponse{
			ID:    user.GetID(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
			Role:  user.GetRole(),
		},
		Redirect: fmt.Sprintf("/%s/dashboard", strings.ToLower(user.GetRole())),
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, response)
}

func (uc *userController) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.Redirect(http.StatusFound, "/login")
}