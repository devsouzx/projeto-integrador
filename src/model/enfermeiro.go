package model

type Enfermeiro struct {
	BaseUser
	CPF   string `json:"cpf"`
	COREN string `json:"coren"`
	Telefone string `json:"telefone"`
}

func NewEnfermeiro(email, password, name, cpf, coren, telefone string) *Enfermeiro {
	return &Enfermeiro{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "enfermeiro",
		},
		CPF:   cpf,
		COREN: coren,
		Telefone: telefone,
	}
}
