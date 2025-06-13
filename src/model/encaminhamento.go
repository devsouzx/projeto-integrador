package model

import "time"

type Encaminhamento struct {
	ID            string
	PacienteID    string
	MedicoID      string
	Especialidade string
	Motivo        string
	Prioridade    string
	Data          time.Time
	Status        string
}