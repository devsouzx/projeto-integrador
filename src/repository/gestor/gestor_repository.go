package gestor

import (
	"database/sql"

	"github.com/devsouzx/projeto-integrador/src/service/gestor"
)

func NewGestorRepository(
	db *sql.DB,
) GestorRepositoryInterface {
	return &gestorRepository{
		DB: db,
	}
}

type gestorRepository struct {
	DB *sql.DB
}

type GestorRepositoryInterface interface {
	
}