package model

import "time"

type Medico struct {
	BaseUser
	CRM              string    `json:"crm"`
	CPF              string    `json:"cpf"`
	Especialidade    string    `json:"especialidade"`
	Telefone         string    `json:"telefone"`
	EnderecoConsulta string    `json:"endereco_consulta"`
	Avatar           []byte    `json:"avatar"`
	UnidadeID        int       `json:"unidade_saude_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

func NewMedico(email, password, name, crm, cpf, especialidade, telefone string, unidadeID int) *Medico {
	return &Medico{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "medico",
		},
		CRM:           crm,
		CPF:           cpf,
		Especialidade: especialidade,
		Telefone:      telefone,
		UnidadeID:     unidadeID,
	}
}

func (m *Medico) GetUnidadeID() int {
	return m.UnidadeID
}
