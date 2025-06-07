package model

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type BaseUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type User interface {
	GenerateToken() (string, error)
	EncryptPassword()
	GetEmail() string
	GetPassword() string
	GetRole() string
	SetRole(role string)
	GetName() string
	GetID() string
}

func NewUserLoginDomain(
	identifier string,
	password string,
	role string,
) User {
	return &BaseUser{
		Email:    identifier,
		Password: password,
		Role:     role,
	}
}

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (u *BaseUser) GetID() string {
	return u.ID
}

func (u *BaseUser) GetEmail() string {
	return u.Email
}

func (u *BaseUser) GetPassword() string {
	return u.Password
}

func (u *BaseUser) GetName() string {
	return u.Name
}

func (u *BaseUser) GetRole() string {
	return u.Role
}

func (u *BaseUser) SetRole(role string) {
	u.Role = role
}

func (u *BaseUser) GenerateToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"role":  u.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		Role:  claims["role"].(string),
	}

	c.Set("user", user)
	c.Next()
}

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		role := user.(BaseUser).Role
		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.String(http.StatusForbidden, "Acesso não autorizado")
		c.Abort()
	}
}

func VerifyToken(tokenValue string) (User, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errors.New("invalid token")
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &BaseUser{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		Role:  claims["role"].(string),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func (u *BaseUser) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
}