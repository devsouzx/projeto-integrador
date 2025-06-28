package agente

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/model"
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
	FindAgenteByIdentifier(identifier string) (*model.AgenteComunitario, error)
}

func (ar *agenteRepository) FindAgenteByIdentifier(identifier string) (*model.AgenteComunitario, error) {
	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")

	query := `
			SELECT id, email, senha, nomecompleto, cpf
			FROM agente_comunitario
			WHERE id = $1 OR cpf = $2
		`

	var agente model.AgenteComunitario

	err := ar.DB.QueryRow(query, identifier, cpfFormatado).Scan(
		&agente.ID,
		&agente.Email,
		&agente.Password,
		&agente.Name,
		&agente.CPF,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("agente não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar agente: %w", err)
	}

	return &agente, nil
}
