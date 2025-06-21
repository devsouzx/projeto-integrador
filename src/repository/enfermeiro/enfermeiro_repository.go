package enfermeiro

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func NewEnfermeiroRepository(
	database *sql.DB,
) EnfermeiroRepositoryInterface {
	return &enfermeiroRepository{
		DB: database,
	}
}

type enfermeiroRepository struct {
	DB *sql.DB
}

type EnfermeiroRepositoryInterface interface {
	FindEnfermeiroByIdentifier(identifier string) (model.UserInterface, error)
}

func (er *enfermeiroRepository) FindEnfermeiroByIdentifier(identifier string) (model.UserInterface, error) {
	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")
	clean := strings.ToUpper(strings.ReplaceAll(identifier, "/", ""))
	
	query := `
			SELECT id, email, senha, nomecompleto, cpf, coren, unidade_saude_id
			FROM enfermeiro
			WHERE id = $1 OR cpf = $2 OR coren = $3
		`

	var enfermeiro model.Enfermeiro

	err := er.DB.QueryRow(query, identifier, cpfFormatado, clean).Scan(
		&enfermeiro.ID,
		&enfermeiro.Email,
		&enfermeiro.Password,
		&enfermeiro.Name,
		&enfermeiro.CPF,
		&enfermeiro.COREN,
		&enfermeiro.UnidadeID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("enfermeiro não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar enfermeiro: %w", err)
	}

	return &enfermeiro, nil
}
