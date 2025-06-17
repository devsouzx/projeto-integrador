package model

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type BaseUser struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Name           string    `json:"name"`
	Role           string    `json:"role"`
	Verified       bool      `json:"is_verified"`
	VerifyToken    string    `json:"verify_token"`
	TokenExpiresAt time.Time `json:"token_expires_at"`
}

type UserInterface interface {
	GenerateToken() (string, error)
	EncryptPassword()
	GetEmail() string
	GetPassword() string
	GetRole() string
	SetRole(role string)
	GetName() string
	GetID() string
	IsVerified() bool
	SetVerified(status bool)
	GetVerifyToken() string
	SetVerifyToken(token string)
	GetTokenExpiresAt() time.Time
	SetTokenExpiresAt(expiry time.Time)
	GenerateVerifyToken() (string, error)
}

func NewUserLoginDomain(
	identifier string,
	password string,
	role string,
) UserInterface {
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
		"verified": u.Verified,
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
		ID:       claims["id"].(string),
		Email:    claims["email"].(string),
		Name:     claims["name"].(string),
		Role:     claims["role"].(string),
		Verified: claims["verified"].(bool),
	}

	if !user.Verified {
		c.Redirect(http.StatusFound, "/verify-account")
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
				c.Next()
				return
			}
		}

		RedirectToDashboard(c, role)
		c.Abort()
	}
}

func VerifyToken(tokenValue string) (UserInterface, error) {
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
		Verified: claims["verified"].(bool),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func (u *BaseUser) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.GetPassword()), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
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

func (u *BaseUser) IsVerified() bool {
	return u.Verified
}

func (u *BaseUser) SetVerified(status bool) {
	u.Verified = status
}

func (u *BaseUser) GetVerifyToken() string {
	return u.VerifyToken
}

func (u *BaseUser) SetVerifyToken(token string) {
	u.VerifyToken = token
}

func (u *BaseUser) GetTokenExpiresAt() time.Time {
	return u.TokenExpiresAt
}

func (u *BaseUser) SetTokenExpiresAt(expiry time.Time) {
	u.TokenExpiresAt = expiry
}

func (u *BaseUser) GenerateVerifyToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	expirationTime := time.Now().UTC().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   expirationTime.Unix(),
		"verified": u.Verified,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	u.VerifyToken = tokenString
	u.TokenExpiresAt = expirationTime

	return tokenString, nil
}
