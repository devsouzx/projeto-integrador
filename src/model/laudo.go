package model

import (
	"time"
)

type LaudoStatus string

const (
	LaudoStatusRascunho   LaudoStatus = "rascunho"
	LaudoStatusPendente   LaudoStatus = "pendente"
	LaudoStatusConcluido  LaudoStatus = "concluido"
	LaudoStatusCancelado  LaudoStatus = "cancelado"
)

type ResultadoCitopatologico string

const (
	ResultadoNormal        ResultadoCitopatologico = "normal"
	ResultadoASCUS         ResultadoCitopatologico = "ascus"
	ResultadoASCH          ResultadoCitopatologico = "asch"
	ResultadoLSIL          ResultadoCitopatologico = "lsil"
	ResultadoHSIL          ResultadoCitopatologico = "hsil"
	ResultadoCarcinoma     ResultadoCitopatologico = "carcinoma"
	ResultadoAdenocarcinoma ResultadoCitopatologico = "adenocarcinoma"
)

type Laudo struct {
	ID                string                `json:"id" db:"id"`
	PacienteID        string                `json:"paciente_id" db:"paciente_id"`
	MedicoID          string                `json:"medico_id" db:"medico_id"`
	FichaID           string                `json:"ficha_id" db:"ficha_id"`
	DataExame         time.Time             `json:"data_exame" db:"data_exame"`
	DataLaudo         time.Time             `json:"data_laudo" db:"data_laudo"`
	Resultado         ResultadoCitopatologico `json:"resultado" db:"resultado"`
	CID               string                `json:"cid" db:"cid"`
	Adequabilidade    string                `json:"adequabilidade" db:"adequabilidade"`
	Microbiologia     string                `json:"microbiologia" db:"microbiologia"`
	CelulasEndometriais bool                `json:"celulas_endometriais" db:"celulas_endometriais"`
	Comentarios       string                `json:"comentarios" db:"comentarios"`
	Recomendacoes     string                `json:"recomendacoes" db:"recomendacoes"`
	Status            LaudoStatus           `json:"status" db:"status"`
	CreatedAt         time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at" db:"updated_at"`
	
	Paciente *Paciente `json:"paciente,omitempty"`
	Medico   *Medico   `json:"medico,omitempty"`
	Ficha    *FichaCitopatologica `json:"ficha,omitempty"`
}

func (l *Laudo) GetCID() string {
	switch l.Resultado {
	case ResultadoNormal:
		return "Z01.4"
	case ResultadoASCUS:
		return "R87.61"
	case ResultadoASCH:
		return "R87.61"
	case ResultadoLSIL:
		return "R87.61"
	case ResultadoHSIL:
		return "D06.9"
	case ResultadoCarcinoma:
		return "C53.9"
	case ResultadoAdenocarcinoma:
		return "C53.9"
	default:
		return "R87.61"
	}
}