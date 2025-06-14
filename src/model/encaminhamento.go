package model

import "time"

type Encaminhamento struct {
	ID               string `json:"id" db:"id"`
	Especialidade    string `json:"especialidade" db:"especialidade"`
	Urgencia         string `json:"urgencia" db:"urgencia"`
	Justificativa    string `json:"justificativa" db:"justificativa"`
	UnidadeReferencia string `json:"unidade_referencia" db:"unidade_referencia"`
	Status           string `json:"status" db:"status"`
	DataEncaminhamento time.Time `json:"data_encaminhamento" db:"data_encaminhamento"`
	DataConclusao    *string `json:"data_conclusao" db:"data_conclusao"`
	PacienteID       string `json:"paciente_id" db:"paciente_id"`
	MedicoID         string `json:"medico_id" db:"medico_id"`
	LaudoID          string `json:"laudo_id" db:"laudo_id"`
}