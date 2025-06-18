package paciente

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func NewPacienteRepository(
	db *sql.DB,
) PacienteRepository {
	return &pacienteRepository{
		DB: db,
	}
}

type pacienteRepository struct {
	DB *sql.DB
}

type PacienteRepository interface {
	ListarPacientes(page, pageSize int, search string) ([]*model.Paciente, int, error)
}

func (pr *pacienteRepository) ListarPacientes(page, pageSize int, search string) ([]*model.Paciente, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
    
	countQuery := "SELECT COUNT(*) FROM paciente"
	countArgs := []interface{}{}

	if search != "" {
		countQuery += " WHERE nome ILIKE $1 OR cpf ILIKE $1 OR cns ILIKE $1"
		countArgs = append(countArgs, "%"+search+"%")
	}

	var total int
	err := pr.DB.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao contar pacientes: %w", err)
	}
    
	query := `
		SELECT 
			id, nome, email, cpf, cns, nome_mae, data_nascimento, 
			telefone, apelido, raca, nacionalidade, escolaridade, 
			created_at, updated_at
		FROM paciente`

	args := []interface{}{}
	argPos := 1

	if search != "" {
		query += fmt.Sprintf(" WHERE nome ILIKE $%d OR cpf ILIKE $%d OR cns ILIKE $%d", argPos, argPos, argPos)
		args = append(args, "%"+search+"%")
		argPos++
	}
    
	query += " ORDER BY nome ASC"
    
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, pageSize, (page-1)*pageSize)
    
	rows, err := pr.DB.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao buscar pacientes: %w", err)
	}
	defer rows.Close()

	var pacientes []*model.Paciente
	for rows.Next() {
		var p model.Paciente
		var nascimentoTime time.Time

		err := rows.Scan(
			&p.ID, &p.Name, &p.Email, &p.CPF, &p.CNS, &p.NomeMae, &nascimentoTime,
			&p.Telefone, &p.Apelido, &p.RacaCor, &p.Nacionalidade, &p.Escolaridade,
			&p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("erro ao scanear paciente: %w", err)
		}

		p.NascimentoTime = nascimentoTime
		pacientes = append(pacientes, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("erro após iterar pacientes: %w", err)
	}

	return pacientes, total, nil
}