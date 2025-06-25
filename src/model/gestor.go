package model

type Gestor struct {
	BaseUser
	CPF       string `json:"cpf"`
	Telefone  string `json:"telefone"`
	UnidadeID string `json:"unidade_saude_id,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewGestor(email, password, name, cpf, telefone, unidadeID string) *Gestor {
	return &Gestor{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "gestor",
		},
		CPF:       cpf,
		Telefone:  telefone,
		UnidadeID: unidadeID,
	}
}
