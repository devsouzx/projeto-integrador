package model

import "time"

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