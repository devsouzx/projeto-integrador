package model

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func RedirectToDashboard(c *gin.Context, role string) {
	switch role {
	case "medico":
		c.Redirect(http.StatusFound, "/medico/dashboard")
	case "enfermeiro":
		c.Redirect(http.StatusFound, "/enfermeiro/dashboard")
	case "agente":
		c.Redirect(http.StatusFound, "/agente/dashboard")
	case "gestor":
		c.Redirect(http.StatusFound, "/gestor/dashboard")
	case "paciente":
		c.Redirect(http.StatusFound, "/paciente/dashboard")
	default:
		c.Redirect(http.StatusFound, "/login")
	}
}