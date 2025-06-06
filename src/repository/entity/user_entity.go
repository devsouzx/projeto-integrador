package entity

type UserEntity struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CPF    string `json:"cpf"`
	Password string `json:"password"`
	Role     string `json:"role"`
}