package model

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (u *BaseUser) GenerateToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":       u.ID,
		"email":    u.Email,
		"name":     u.Name,
		"role":     u.Role,
		"verified": u.Role == "paciente" && u.Verified,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
		ID:       claims["id"].(string),
		Email:    claims["email"].(string),
		Name:     claims["name"].(string),
		Role:     claims["role"].(string),
		Verified: claims["verified"].(bool),
	}, nil
}

func (u *BaseUser) GenerateVerifyToken() (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)

	expirationTime := time.Now().UTC().Add(1 * time.Hour)

	claims := jwt.MapClaims{
		"id":       u.ID,
		"email":    u.Email,
		"exp":      expirationTime.Unix(),
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