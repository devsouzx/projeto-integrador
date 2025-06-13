package model

import "time"

type Laudo struct {
	ID          string
	ExameID     string
	MedicoID    string
	Conteudo    string
	DataEmissao time.Time
}