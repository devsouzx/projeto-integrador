package gestor

import (
	"database/sql"
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
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
	CadastrarMedico(medico request.CadastroProfissionalRequest) (*model.Medico, error)
}

func (gr *gestorRepository) CadastrarMedico(medicoRequest request.CadastroProfissionalRequest) (*model.Medico, error) {
	query := `
    INSERT INTO medico (
        crm, email, senha, nomecompleto, cpf, especialidade,
        telefone
    ) VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, created_at, updated_at`

	medico := model.NewMedico(
        medicoRequest.Email,
        medicoRequest.Password,
        medicoRequest.Name,
        medicoRequest.CRM,
        medicoRequest.CPF,
        medicoRequest.Especialidade,
        medicoRequest.Telefone,
        24222,
    )

	err := gr.DB.QueryRow(
		query,
		medicoRequest.CRM,
		medicoRequest.Email,
		medicoRequest.Password,
		medicoRequest.Name,
		medicoRequest.CPF,
		medicoRequest.Especialidade,
		medicoRequest.Telefone,
	).Scan(&medico.ID, &medico.CreatedAt, &medico.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar medico: %v", err)
	}

	return medico, nil
}
