package model

import "time"

type Paciente struct {
	BaseUser
	ID             string    `json:"id" db:"id"`
	Apelido        string    `json:"apelido" db:"apelido"`
	NomeCompleto   string    `json:"nome_completo" db:"nome_completo"`
	Email          string    `json:"email" db:"email"`
	Senha          string    `json:"-" db:"senha"`
	NomeMae        string    `json:"mae" db:"nome_mae"`
	CNS            string    `json:"cns" db:"cns"`
	CPF            string    `json:"cpf" db:"cpf"`
	DataNascimento string    `json:"nascimento" db:"-"`
	NascimentoTime time.Time `json:"-" db:"data_nascimento"`
	Nacionalidade  string    `json:"nacionalidade" db:"nacionalidade"`
	RacaCor        string    `json:"cor" db:"raca_cor"`
	Escolaridade   string    `json:"escolaridade" db:"escolaridade"`
	Endereco       Endereco  `json:"-"`
	Telefone       string    `json:"telefone" db:"telefone"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
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
		CPF:            cpf,
		CNS:            cartaoSUS,
		NomeMae:        nomeMae,
		NascimentoTime: nascimento,
		Telefone:       telefone,
		Apelido:        apelido,
		RacaCor:        raca,
		Nacionalidade:  nacionalidade,
		Escolaridade:   escolaridade,
	}
}

func (p *Paciente) GetCPF() string {
	return p.CPF
}

func (p *Paciente) GetID() string {
	return p.ID
}
