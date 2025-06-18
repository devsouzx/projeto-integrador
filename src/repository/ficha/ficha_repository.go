package ficha

import (
	"database/sql"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func NewFichaRepository(
	database *sql.DB,
) FichaRepositoryInterface {
	return &fichaRepository{
		DB: database,
	}
}

type fichaRepository struct {
	DB *sql.DB
}

type FichaRepositoryInterface interface {
	Create(ficha *model.FichaCitopatologica) (*model.FichaCitopatologica, error)
	FindByID(id string) (*model.FichaCitopatologica, error)
	CreateEndereco(endereco *model.Endereco, pacienteId string) error
	UpdateEndereco(endereco *model.Endereco, pacienteId string) error
	EnderecoExists(pacienteId string) (bool, error)
	UpsertEndereco(endereco *model.Endereco, pacienteId string) error
}