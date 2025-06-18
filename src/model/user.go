package model

import (
	"time"
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
