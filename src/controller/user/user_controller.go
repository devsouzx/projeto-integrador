package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model/request"
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
	CadastrarUsuario(c *gin.Context) //cadastro de usuário
	VerifyAccount(c *gin.Context)
}

func (uc *userController) LoginUser(c *gin.Context) {
	var userRequest request.LoginRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao validar dados do usuário",
			"details": err.Error(),
		})
		return
	}

	fmt.Println("req", userRequest)
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

func (uc *userController) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.Redirect(http.StatusFound, "/login")
}

// cadastro de usuário
func (uc *userController) CadastrarUsuario(c *gin.Context) {
	var cadRequest request.CadastroRequest

	if err := c.ShouldBind(&cadRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro ao validar dados do usuário",
			"details": err.Error(),
		})
		return
	}

	usuario, err := uc.service.CadastrarUsuario(cadRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"details": err.Error(),
		})
		return
	}

	newToken, err := usuario.GenerateVerifyToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	if err := uc.service.UpdateUser(usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	verificationLink := fmt.Sprintf("http://localhost:8080/verify?token=%s", newToken)
	if err := uc.emailService.SendVerificationEmail(usuario.GetEmail(), usuario.GetName(), verificationLink, newToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar email"})
		return
	}

	response := gin.H{
		"message": "Usuario cadastrado com sucesso",
		"data": gin.H{
			"id":                usuario.ID,
			"nomecompleto":      usuario.Name,
			"cpf":               usuario.CPF,
			"cartaosus":         usuario.CNS,
			"nomecompletodamae": usuario.NomeMae,
			"datadenascimento":  usuario.DataNascimento,
			"email":             usuario.Email,
			"telefone":          usuario.Telefone,
		},
	}

	if usuario.Endereco != nil {
		response["data"].(gin.H)["endereco"] = gin.H{
			"cep":         usuario.Endereco.CEP,
			"logradouro":  usuario.Endereco.Logradouro,
			"numero":      usuario.Endereco.Numero,
			"complemento": usuario.Endereco.Complemento,
			"bairro":      usuario.Endereco.Bairro,
			"cidade":      usuario.Endereco.Cidade,
			"uf":          usuario.Endereco.UF,
		}
	}

	c.JSON(http.StatusCreated, response)
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
