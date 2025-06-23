package model

import "time"

type Medico struct {
	BaseUser
	CRM              string    `json:"crm"`
	CPF              string    `json:"cpf"`
	Especialidade    string    `json:"especialidade"`
	Telefone         string    `json:"telefone"`
	DataNascimento   time.Time `json:"data_nascimento"`
	EnderecoConsulta string    `json:"endereco_consulta"`
	Avatar           []byte    `json:"avatar"`
	UnidadeID        int       `json:"unidade_saude_id"`
}

func NewMedico(email, password, name, crm, cpf string) UserInterface {
	return &Medico{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "medico",
		},
		CRM:            crm,
		CPF:            cpf,
		DataNascimento: data_nascimento,
		Especialidade:  especialidade,
		Telefone:       telefone,
		UnidadeID:      unidadeID,
	}
}

func (m *Medico) GetUnidadeID() int {
	return m.UnidadeID
}
