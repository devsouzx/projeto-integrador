package model

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)

	tokenValue, err := c.Cookie("token")

	if err != nil || tokenValue == "" {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	user := BaseUser{
		ID:       claims["id"].(string),
		Email:    claims["email"].(string),
		Name:     claims["name"].(string),
		Role:     claims["role"].(string),
		Verified: claims["verified"].(bool),
	}

	if user.Role == "paciente" && !user.Verified {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.Abort()
			return
		}

		role := user.(BaseUser).Role
		for _, allowed := range allowedRoles {
			if role == allowed {
				fmt.Println("User role is allowed:", role)
				return
			}
		}

		RedirectToDashboard(c, role)
		c.Abort()
	}
}