package model

type AgenteComunitario struct {
	BaseUser
	CPF       string `json:"cpf"`
	Telefone  string `json:"telefone"`
	UnidadeID string `json:"unidade_saude_id,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewAgenteComunitario(email, password, name, cpf, telefone, unidadeID string) *AgenteComunitario {
	return &AgenteComunitario{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "agente",
		},
		CPF:       cpf,
		Telefone:  telefone,
		UnidadeID: unidadeID,
	}
}
