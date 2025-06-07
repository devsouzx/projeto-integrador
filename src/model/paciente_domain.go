package model

import "time"

type Paciente struct {
	BaseUser
	CPF                 string    `json:"cpf"`
	CartaoSUS           string    `json:"cartaosus"`
	NomeCompletoDaMae   string    `json:"nomecompletodamae"`
	DataDeNascimento    time.Time `json:"datadenascimento"`
	Telefone            string    `json:"telefone"`
	Apelido             string    `json:"apelido,omitempty"`
	Raca                string    `json:"raca,omitempty"`
	Nacionalidade       string    `json:"nacionalidade,omitempty"`
	Escolaridade        string    `json:"escolaridade,omitempty"`
}

func NewPaciente(
	email, password, name, cpf, cartaoSUS, nomeMae string, nascimento time.Time,
	telefone, apelido, raca, nacionalidade, escolaridade string,
) *Paciente {
	return &Paciente{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "paciente",
		},
		CPF:               cpf,
		CartaoSUS:         cartaoSUS,
		NomeCompletoDaMae: nomeMae,
		DataDeNascimento:  nascimento,
		Telefone:          telefone,
		Apelido:           apelido,
		Raca:              raca,
		Nacionalidade:     nacionalidade,
		Escolaridade:      escolaridade,
	}
}

func (p *Paciente) GetCPF() string {
	return p.CPF
}
