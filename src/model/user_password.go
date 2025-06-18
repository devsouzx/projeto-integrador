package model

import "golang.org/x/crypto/bcrypt"

func (u *BaseUser) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.GetPassword()), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
}