package model

import (
	"time"
)

type EncaminhamentoStatus string

const (
	EncaminhamentoStatusPendente   EncaminhamentoStatus = "pendente"
	EncaminhamentoStatusAgendado   EncaminhamentoStatus = "agendado"
	EncaminhamentoStatusConcluido  EncaminhamentoStatus = "concluido"
	EncaminhamentoStatusCancelado  EncaminhamentoStatus = "cancelado"
)

type NivelUrgencia string

const (
	UrgenciaRotina     NivelUrgencia = "rotina"
	UrgenciaPrioritario NivelUrgencia = "prioritario"
	UrgenciaUrgente    NivelUrgencia = "urgente"
)

type Encaminhamento struct {
	ID                string               `json:"id" db:"id"`
	PacienteID        string               `json:"paciente_id" db:"paciente_id"`
	MedicoID          string               `json:"medico_id" db:"medico_id"`
	LaudoID           *string              `json:"laudo_id,omitempty" db:"laudo_id"`
	Especialidade     string               `json:"especialidade" db:"especialidade"`
	Urgencia          NivelUrgencia        `json:"urgencia" db:"urgencia"`
	Justificativa     string               `json:"justificativa" db:"justificativa"`
	UnidadeReferencia string               `json:"unidade_referencia" db:"unidade_referencia"`
	Status            EncaminhamentoStatus `json:"status" db:"status"`
	DataEncaminhamento time.Time           `json:"data_encaminhamento" db:"data_encaminhamento"`
	DataAgendamento   *time.Time           `json:"data_agendamento,omitempty" db:"data_agendamento"`
	DataConclusao     *time.Time           `json:"data_conclusao,omitempty" db:"data_conclusao"`
	Observacoes       *string              `json:"observacoes,omitempty" db:"observacoes"`
	CreatedAt         time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at" db:"updated_at"`
	
	Paciente *Paciente `json:"paciente,omitempty"`
	Medico   *Medico   `json:"medico,omitempty"`
	Laudo    *Laudo    `json:"laudo,omitempty"`
}