package gestor

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
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
	CadastrarMedico(medico *model.Medico) (*model.Medico, error)
}

func (gr *gestorRepository) CadastrarMedico(medico *model.Medico) (*model.Medico, error) {
	if medico.DataNascimento != "" {
		nascimento, err := time.Parse("2006-01-02", medico.DataNascimento)
		if err != nil {
			return nil, fmt.Errorf("formato de data inválido: %v", err)
		}
		medico.NascimentoTime = nascimento
	}

	query := `
    INSERT INTO medico (
        crm, email, senha, nome, cpf, data_nascimento, especialidade,
        telefone
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, created_at, updated_at`

	err := gr.DB.QueryRow(
		query,
		medico.CRM,
		medico.Email,
		medico.Password,
		medico.Name,
		medico.CPF,
		medico.NascimentoTime,
		medico.Especialidade,
		medico.Telefone,
	).Scan(&medico.ID, &medico.CreatedAt, &medico.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar medico: %v", err)
	}

	return medico, nil
}
