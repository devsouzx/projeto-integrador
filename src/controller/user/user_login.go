package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/gin-gonic/gin"
)

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

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"token":   token,
		"user": gin.H{
			"id":    user.GetID(),
			"name":  user.GetName(),
			"email": user.GetEmail(),
			"role":  user.GetRole(),
		},
		"redirect": fmt.Sprintf("/%s/dashboard", strings.ToLower(user.GetRole())),
	})
}

func (uc *userController) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.Redirect(http.StatusFound, "/login")
}