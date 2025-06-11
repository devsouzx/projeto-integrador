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
}

type Exame struct {
	ID          string
	PacienteID  string
	DataColeta  time.Time
	Tipo        string
	Resultado   string
	CID         string
	Observacoes string
	Status      string
}

type Laudo struct {
    ID         string
    ExameID    string
    MedicoID   string
    Conteudo   string
    DataEmissao time.Time
}

type Encaminhamento struct {
    ID          string
    PacienteID  string
    MedicoID    string
    Especialidade string
    Motivo      string
    Prioridade  string
    Data        time.Time
    Status      string
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