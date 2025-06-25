package model

import "time"

type Enfermeiro struct {
	BaseUser
	CPF       string `json:"cpf"`
	COREN     string `json:"coren"`
	Telefone  string `json:"telefone"`
	UnidadeID int    `json:"unidade_saude_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Agendamento struct {
	ID               string    `json:"id"`
	PacienteID       string    `json:"paciente_id"`
	ProfissionalID   string    `json:"profissional_id"`
	ProfissionalTipo string    `json:"profissional_tipo"`
	DataHora         time.Time `json:"data_hora"`
	TipoConsulta     string    `json:"tipo_consulta"`
	Status           string    `json:"status"`
	Observacoes      string    `json:"observacoes"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	PacienteNome string `json:"paciente_nome,omitempty"`
	PacienteCNS  string `json:"paciente_cns,omitempty"`
}

func NewEnfermeiro(email, password, name, cpf, coren, telefone string) *Enfermeiro {
	return &Enfermeiro{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "enfermeiro",
		},
		CPF:      cpf,
		COREN:    coren,
		Telefone: telefone,
	}
}

func (e *Enfermeiro) GetUnidadeID() int {
	return e.UnidadeID
}
