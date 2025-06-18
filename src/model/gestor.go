package model

type Gestor struct {
	BaseUser
	CPF       string `json:"cpf"`
	Telefone  string `json:"telefone"`
	UnidadeID string `json:"unidade_saude_id,omitempty"`
}

func NewGestor(email, password, name, cpf, telefone, unidadeID string) *AgenteComunitario {
	return &AgenteComunitario{
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
