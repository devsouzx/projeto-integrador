package model

import "time"

type Endereco struct {
	ID          string    `json:"id"`
	CEP         string    `json:"cep"`
	Logradouro  string    `json:"logradouro"`
	Numero      string    `json:"numero"`
	Complemento string    `json:"complemento"`
	Bairro      string    `json:"bairro"`
	Cidade      string    `json:"cidade"`
	UF          string    `json:"uf"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}