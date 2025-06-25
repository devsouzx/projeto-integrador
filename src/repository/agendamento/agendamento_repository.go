package agendamento

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
)

type AgendamentoRepositoryInterface interface {
	Create(agendamento *model.Agendamento) error
	FindByID(id string) (*model.Agendamento, error)
	FindByProfissional(profissionalID, profissionalTipo string) ([]*model.Agendamento, error)
	FindByProfissionalAndDate(profissionalID, profissionalTipo string, date time.Time) ([]*model.Agendamento, error)
	UpdateStatus(id, status string) error
	Delete(id string) error
}

type agendamentoRepository struct {
	DB *sql.DB
}

func NewAgendamentoRepository(db *sql.DB) AgendamentoRepositoryInterface {
	return &agendamentoRepository{DB: db}
}

func (r *agendamentoRepository) Create(agendamento *model.Agendamento) error {
	query := `
		INSERT INTO agendamento (paciente_id, profissional_id, profissional_tipo, data_hora, tipo_consulta, status, observacoes)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	
	err := r.DB.QueryRow(
		query,
		agendamento.PacienteID,
		agendamento.ProfissionalID,
		agendamento.ProfissionalTipo,
		agendamento.DataHora,
		agendamento.TipoConsulta,
		agendamento.Status,
		agendamento.Observacoes,
	).Scan(&agendamento.ID, &agendamento.CreatedAt, &agendamento.UpdatedAt)
	
	if err != nil {
		return fmt.Errorf("erro ao criar agendamento: %w", err)
	}
	
	return nil
}

func (r *agendamentoRepository) FindByID(id string) (*model.Agendamento, error) {
	query := `
		SELECT a.id, a.paciente_id, a.profissional_id, a.profissional_tipo, a.data_hora, 
		       a.tipo_consulta, a.status, a.observacoes, a.created_at, a.updated_at,
		       p.nome as paciente_nome, p.cns as paciente_cns
		FROM agendamento a
		JOIN paciente p ON a.paciente_id = p.id
		WHERE a.id = $1
	`
	
	var agendamento model.Agendamento
	err := r.DB.QueryRow(query, id).Scan(
		&agendamento.ID,
		&agendamento.PacienteID,
		&agendamento.ProfissionalID,
		&agendamento.ProfissionalTipo,
		&agendamento.DataHora,
		&agendamento.TipoConsulta,
		&agendamento.Status,
		&agendamento.Observacoes,
		&agendamento.CreatedAt,
		&agendamento.UpdatedAt,
		&agendamento.PacienteNome,
		&agendamento.PacienteCNS,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("agendamento não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar agendamento: %w", err)
	}
	
	return &agendamento, nil
}

func (r *agendamentoRepository) FindByProfissional(profissionalID, profissionalTipo string) ([]*model.Agendamento, error) {
	query := `
		SELECT a.id, a.paciente_id, a.profissional_id, a.profissional_tipo, a.data_hora, 
		       a.tipo_consulta, a.status, a.observacoes, a.created_at, a.updated_at,
		       p.nome as paciente_nome, p.cns as paciente_cns
		FROM agendamento a
		JOIN paciente p ON a.paciente_id = p.id
		WHERE a.profissional_id = $1 AND a.profissional_tipo = $2
		ORDER BY a.data_hora ASC
	`
	
	rows, err := r.DB.Query(query, profissionalID, profissionalTipo)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar agendamentos: %w", err)
	}
	defer rows.Close()
	
	var agendamentos []*model.Agendamento
	for rows.Next() {
		var agendamento model.Agendamento
		err := rows.Scan(
			&agendamento.ID,
			&agendamento.PacienteID,
			&agendamento.ProfissionalID,
			&agendamento.ProfissionalTipo,
			&agendamento.DataHora,
			&agendamento.TipoConsulta,
			&agendamento.Status,
			&agendamento.Observacoes,
			&agendamento.CreatedAt,
			&agendamento.UpdatedAt,
			&agendamento.PacienteNome,
			&agendamento.PacienteCNS,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler agendamento: %w", err)
		}
		agendamentos = append(agendamentos, &agendamento)
	}
	
	return agendamentos, nil
}

func (r *agendamentoRepository) FindByProfissionalAndDate(profissionalID, profissionalTipo string, date time.Time) ([]*model.Agendamento, error) {
	query := `
		SELECT a.id, a.paciente_id, a.profissional_id, a.profissional_tipo, a.data_hora, 
		       a.tipo_consulta, a.status, a.observacoes, a.created_at, a.updated_at,
		       p.nome as paciente_nome, p.cns as paciente_cns
		FROM agendamento a
		JOIN paciente p ON a.paciente_id = p.id
		WHERE a.profissional_id = $1 AND a.profissional_tipo = $2 
		      AND DATE(a.data_hora) = DATE($3)
		ORDER BY a.data_hora ASC
	`
	
	rows, err := r.DB.Query(query, profissionalID, profissionalTipo, date)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar agendamentos: %w", err)
	}
	defer rows.Close()
	
	var agendamentos []*model.Agendamento
	for rows.Next() {
		var agendamento model.Agendamento
		err := rows.Scan(
			&agendamento.ID,
			&agendamento.PacienteID,
			&agendamento.ProfissionalID,
			&agendamento.ProfissionalTipo,
			&agendamento.DataHora,
			&agendamento.TipoConsulta,
			&agendamento.Status,
			&agendamento.Observacoes,
			&agendamento.CreatedAt,
			&agendamento.UpdatedAt,
			&agendamento.PacienteNome,
			&agendamento.PacienteCNS,
		)
		fmt.Println(err)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler agendamento: %w", err)
		}
		agendamentos = append(agendamentos, &agendamento)
	}
	
	return agendamentos, nil
}

func (r *agendamentoRepository) UpdateStatus(id, status string) error {
	query := `
		UPDATE agendamento
		SET status = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`
	
	_, err := r.DB.Exec(query, status, id)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status do agendamento: %w", err)
	}
	
	return nil
}

func (r *agendamentoRepository) Delete(id string) error {
	query := `DELETE FROM agendamento WHERE id = $1`
	
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar agendamento: %w", err)
	}
	
	return nil
}
