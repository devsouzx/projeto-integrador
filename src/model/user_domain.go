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

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	CPF    string `json:"cpf"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserDomainInterface interface {
	GenerateToken() (string, error)
	EncryptPassword()
	SetID(id string)
	GetEmail() string
	GetPassword() string
	GetRole() string
	GetName() string
	GetID() string
	GetCPF() string
}

func NewUserDomain(
	email, cpf, password, name, role string,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		CPF:    cpf,
		Password: password,
		Name:     name,
		Role:     role,
	}
}

func NewUserLoginDomain(
	identificador, password string,
) UserDomainInterface {
	user := &userDomain{
		Password: password,
	}
	
	if strings.Contains(identificador, "@") {
		user.Email = identificador
	} else {
		user.CPF = identificador
	}

	return user
}

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GetID() string {
	return ud.ID
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetCPF() string {
	return ud.CPF
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetRole() string {
	return ud.Role
}

func (ud *userDomain) GenerateToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.Email,
		"cpf": ud.CPF,
		"name":  ud.Name,
		"role":  ud.Role,
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

	user := userDomain{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		CPF: claims["cpf"].(string),
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

		role := user.(userDomain).Role
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

func VerifyToken(tokenValue string) (UserDomainInterface, error) {
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

	return &userDomain{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		CPF: claims["cpf"].(string),
		Name:  claims["name"].(string),
		Role:  claims["role"].(string),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func (ud *userDomain) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	ud.Password = string(hashedPassword)
}