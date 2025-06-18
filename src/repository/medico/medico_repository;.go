package medico

import (
	"database/sql"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func NewMedicoRepository(
	database *sql.DB,
) MedicoRepositoryInterface {
	return &medicoRepository{
		DB: database,
	}
}

type medicoRepository struct {
	DB *sql.DB
}

type MedicoRepositoryInterface interface {
	FindMedicoByIdentifier(identifier string) (model.UserInterface, error)
}