package user

import (
	"net/http"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/gin-gonic/gin"
)

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

func (uc *userController) VerifyAccount(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.Redirect(http.StatusFound, "/login?error=token_required")
		return
	}

	user, err := uc.service.VerifyUserToken(token)
	if err != nil {
		c.Redirect(http.StatusFound, "/login?error=invalid_token")
		return
	}

	user.SetVerified(true)
	user.SetVerifyToken("")
	user.SetTokenExpiresAt(time.Time{})

	if err := uc.service.UpdateUser(user); err != nil {
		c.Redirect(http.StatusFound, "/login?error=verification_failed")
		return
	}

	c.Redirect(http.StatusFound, "/login?verified=true")
}
