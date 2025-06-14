package model

import "time"

type Medico struct {
	BaseUser
	CRM              string `json:"crm"`
	CPF              string `json:"cpf"`
	Especialidade    string `json:"especialidade"`
	Telefone         string `json:"telefone"`
	DataNascimento   time.Time `json:"data_nascimento"`
	EnderecoConsulta string `json:"endereco_consulta"`
	Avatar       []byte `json:"avatar"`
}

func NewMedico(email, password, name, crm, cpf string) UserInterface {
	return &Medico{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "medico",
		},
		CRM: crm,
		CPF: cpf,
	}
}