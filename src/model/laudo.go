package model

import "time"

type Laudo struct {
	ID               string `json:"id" db:"id"`
	DataEmissao      time.Time `json:"data_emissao" db:"data_emissao"`
	Diagnostico      string `json:"diagnostico" db:"diagnostico"`
	Comentarios      string `json:"comentarios" db:"comentarios"`
	Recomendacoes    string `json:"recomendacoes" db:"recomendacoes"`
	ExameID          string `json:"exame_id" db:"exame_id"`
	MedicoResponsavelID string `json:"medico_responsavel_id" db:"medico_responsavel_id"`
	EncaminhamentoID *string `json:"encaminhamento_id" db:"encaminhamento_id"`
}