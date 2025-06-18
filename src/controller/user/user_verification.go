package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
