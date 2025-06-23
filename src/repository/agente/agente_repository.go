package agente

import (
	"database/sql"
)

func NewAgenteRepository(
	db *sql.DB,
) AgenteRepositoryInterface {
	return &agenteRepository{
		DB: db,
	}
}

type agenteRepository struct {
	DB *sql.DB
}

type AgenteRepositoryInterface interface {
}
