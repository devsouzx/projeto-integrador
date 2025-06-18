package medico

import (
	"net/http"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/gin-gonic/gin"
)

func (mc *medicoController) getMedicoLogado(c *gin.Context) (model.UserInterface, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, false
	}

	baseUser, ok := user.(model.BaseUser)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil, false
	}

	medico, err := mc.service.FindMedicoByIdentifier(baseUser.ID)
	if err != nil {
		c.Redirect(http.StatusFound, "/login-profissionais")
		return nil, false
	}

	return medico, true
}