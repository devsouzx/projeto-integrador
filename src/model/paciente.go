package model

import (
	"time"
)

type Paciente struct {
	BaseUser
	Apelido        string    `json:"apelido" db:"apelido"`
	NomeMae        string    `json:"mae" db:"nome_mae"`
	CNS            string    `json:"cns" db:"cns"`
	CPF            string    `json:"cpf" db:"cpf"`
	DataNascimento string    `json:"nascimento" db:"-"`
	NascimentoTime time.Time `json:"-" db:"data_nascimento"`
	Nacionalidade  string    `json:"nacionalidade" db:"nacionalidade"`
	RacaCor        string    `json:"cor" db:"raca_cor"`
	Escolaridade   string    `json:"escolaridade" db:"escolaridade"`
	Endereco       *Endereco  `json:"endereco"`
	Telefone       string    `json:"telefone" db:"telefone"`
	UltimoExame *ExameClinico `json:"ultima_exame"`
	Fichas []*FichaCitopatologica `json:"fichas"`
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

func (p *Paciente) GetPhone() string {
	return p.Telefone
}

func (p *Paciente) GetID() string {
	return p.ID
}

func (p *Paciente) GetEmail() string { 
	return p.Email 
}

func (p *Paciente) GetInspecaoColo() string {
	if p.UltimoExame == nil {
		return ""
	}
	return p.UltimoExame.InspecaoColo
}

func (p *Paciente) GetDataUltimoExame() time.Time {
    if p.UltimoExame == nil {
        return time.Time{}
    }
    return p.UltimoExame.CreatedAt
}

func (p *Paciente) GetIdade() int {
    if p.NascimentoTime.IsZero() {
        return 0
    }
    now := time.Now()
    years := now.Year() - p.NascimentoTime.Year()
    if now.YearDay() < p.NascimentoTime.YearDay() {
        years--
    }
    return years
}