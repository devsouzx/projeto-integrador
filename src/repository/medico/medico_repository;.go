package medico

import (
	"database/sql"
	"fmt"
	"strings"

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

func (mr *medicoRepository) FindMedicoByIdentifier(identifier string) (model.UserInterface, error) {
	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")
	clean := strings.ToUpper(strings.ReplaceAll(identifier, "/", ""))

	query := `
			SELECT id, email, senha, nomecompleto, cpf, crm
			FROM medico
			WHERE id = $1 OR cpf = $2 OR crm = $3
		`

	var medico model.Medico

	err := mr.DB.QueryRow(query, identifier, cpfFormatado, clean).Scan(
		&medico.ID,
		&medico.Email,
		&medico.Password,
		&medico.Name,
		&medico.CPF,
		&medico.CRM,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("medico não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar médico: %w", err)
	}

	return &medico, nil
}
