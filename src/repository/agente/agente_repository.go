package agente

import (
	"database/sql"

	"github.com/devsouzx/projeto-integrador/src/service/agente"
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