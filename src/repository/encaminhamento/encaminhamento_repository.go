package encaminhamento

import (
	"database/sql"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
)

type EncaminhamentoRepositoryInterface interface {
	Create(encaminhamento *model.Encaminhamento) error
	FindByID(id string) (*model.Encaminhamento, error)
	FindByPacienteID(pacienteID string, page, pageSize int) ([]*model.Encaminhamento, int, error)
	FindByMedicoID(medicoID string, page, pageSize int) ([]*model.Encaminhamento, int, error)
	UpdateStatus(id string, status model.EncaminhamentoStatus, observacoes string) error
	UpdateAgendamento(id string, dataAgendamento time.Time) error
}

type encaminhamentoRepository struct {
	DB *sql.DB
}

func NewEncaminhamentoRepository(db *sql.DB) EncaminhamentoRepositoryInterface {
	return &encaminhamentoRepository{DB: db}
}

func (r *encaminhamentoRepository) Create(encaminhamento *model.Encaminhamento) error {
	query := `
		INSERT INTO encaminhamentos (
			paciente_id, medico_id, laudo_id, especialidade, urgencia, 
			justificativa, unidade_referencia, status, data_encaminhamento,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`

	encaminhamento.DataEncaminhamento = time.Now()
	encaminhamento.CreatedAt = time.Now()
	encaminhamento.UpdatedAt = time.Now()

	err := r.DB.QueryRow(
		query,
		encaminhamento.PacienteID,
		encaminhamento.MedicoID,
		encaminhamento.LaudoID,
		encaminhamento.Especialidade,
		encaminhamento.Urgencia,
		encaminhamento.Justificativa,
		encaminhamento.UnidadeReferencia,
		encaminhamento.Status,
		encaminhamento.DataEncaminhamento,
		encaminhamento.CreatedAt,
		encaminhamento.UpdatedAt,
	).Scan(&encaminhamento.ID)

	return err
}

func (r *encaminhamentoRepository) FindByID(id string) (*model.Encaminhamento, error) {
	query := `
		SELECT 
			e.id, e.paciente_id, e.medico_id, e.laudo_id, e.especialidade, e.urgencia,
			e.justificativa, e.unidade_referencia, e.status, e.data_encaminhamento,
			e.data_agendamento, e.data_conclusao, e.observacoes, e.created_at, e.updated_at,
			p.id, p.nome, p.cns,
			m.id, m.nomecompleto, m.crm,
			l.id, l.paciente_id, l.medico_id, l.ficha_id, l.data_exame, l.data_laudo,
			l.resultado, l.cid, l.adequabilidade, l.microbiologia, l.celulas_endometriais,
			l.comentarios, l.recomendacoes, l.status, l.created_at, l.updated_at
		FROM encaminhamentos e
		LEFT JOIN paciente p ON p.id = e.paciente_id
		LEFT JOIN medico m ON m.id = e.medico_id
		LEFT JOIN laudos l ON l.id = e.laudo_id
		WHERE e.id = $1
	`

	var encaminhamento model.Encaminhamento
	var paciente model.Paciente
	var medico model.Medico
	var laudo model.Laudo

	err := r.DB.QueryRow(query, id).Scan(
		&encaminhamento.ID,
		&encaminhamento.PacienteID,
		&encaminhamento.MedicoID,
		&encaminhamento.LaudoID,
		&encaminhamento.Especialidade,
		&encaminhamento.Urgencia,
		&encaminhamento.Justificativa,
		&encaminhamento.UnidadeReferencia,
		&encaminhamento.Status,
		&encaminhamento.DataEncaminhamento,
		&encaminhamento.DataAgendamento,
		&encaminhamento.DataConclusao,
		&encaminhamento.Observacoes,
		&encaminhamento.CreatedAt,
		&encaminhamento.UpdatedAt,
		&paciente.ID,
		&paciente.Name,
		&paciente.CNS,
		&medico.ID,
		&medico.Name,
		&medico.CRM,
		&laudo.ID,
		&laudo.PacienteID,
		&laudo.MedicoID,
		&laudo.FichaID,
		&laudo.DataExame,
		&laudo.DataLaudo,
		&laudo.Resultado,
		&laudo.CID,
		&laudo.Adequabilidade,
		&laudo.Microbiologia,
		&laudo.CelulasEndometriais,
		&laudo.Comentarios,
		&laudo.Recomendacoes,
		&laudo.Status,
		&laudo.CreatedAt,
		&laudo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	encaminhamento.Paciente = &paciente
	encaminhamento.Medico = &medico
	
	if laudo.ID != "" {
		encaminhamento.Laudo = &laudo
	}

	return &encaminhamento, nil
}

func (r *encaminhamentoRepository) FindByPacienteID(pacienteID string, page, pageSize int) ([]*model.Encaminhamento, int, error) {
	query := `
		SELECT 
			e.id, e.especialidade, e.urgencia, e.status, e.data_encaminhamento,
			e.data_agendamento, e.unidade_referencia, e.updated_at,
			m.id, m.nomecompleto, m.crm
		FROM encaminhamentos e
		LEFT JOIN medico m ON m.id = e.medico_id
		WHERE e.paciente_id = $1
		ORDER BY e.data_encaminhamento DESC
		LIMIT $2 OFFSET $3
	`

	countQuery := `SELECT COUNT(*) FROM encaminhamentos WHERE paciente_id = $1`

	var total int
	err := r.DB.QueryRow(countQuery, pacienteID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.DB.Query(query, pacienteID, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var encaminhamentos []*model.Encaminhamento
	for rows.Next() {
		var encaminhamento model.Encaminhamento
		var medico model.Medico

		err := rows.Scan(
			&encaminhamento.ID,
			&encaminhamento.Especialidade,
			&encaminhamento.Urgencia,
			&encaminhamento.Status,
			&encaminhamento.DataEncaminhamento,
			&encaminhamento.DataAgendamento,
			&encaminhamento.UnidadeReferencia,
			&encaminhamento.UpdatedAt,
			&medico.ID,
			&medico.Name,
			&medico.CRM,
		)
		if err != nil {
			return nil, 0, err
		}

		encaminhamento.Medico = &medico
		encaminhamentos = append(encaminhamentos, &encaminhamento)
	}

	return encaminhamentos, total, nil
}

func (r *encaminhamentoRepository) FindByMedicoID(medicoID string, page, pageSize int) ([]*model.Encaminhamento, int, error) {
	query := `
		SELECT 
			e.id, e.paciente_id, e.especialidade, e.urgencia, e.status, 
			e.data_encaminhamento, e.data_agendamento, e.unidade_referencia,
			p.id, p.nome, p.cns
		FROM encaminhamentos e
		LEFT JOIN paciente p ON p.id = e.paciente_id
		WHERE e.medico_id = $1
		ORDER BY e.data_encaminhamento DESC
		LIMIT $2 OFFSET $3
	`

	countQuery := `SELECT COUNT(*) FROM encaminhamentos WHERE medico_id = $1`

	var total int
	err := r.DB.QueryRow(countQuery, medicoID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.DB.Query(query, medicoID, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var encaminhamentos []*model.Encaminhamento
	for rows.Next() {
		var encaminhamento model.Encaminhamento
		var paciente model.Paciente

		err := rows.Scan(
			&encaminhamento.ID,
			&encaminhamento.PacienteID,
			&encaminhamento.Especialidade,
			&encaminhamento.Urgencia,
			&encaminhamento.Status,
			&encaminhamento.DataEncaminhamento,
			&encaminhamento.DataAgendamento,
			&encaminhamento.UnidadeReferencia,
			&paciente.ID,
			&paciente.Name,
			&paciente.CNS,
		)
		if err != nil {
			return nil, 0, err
		}

		encaminhamento.Paciente = &paciente
		encaminhamentos = append(encaminhamentos, &encaminhamento)
	}

	return encaminhamentos, total, nil
}

func (r *encaminhamentoRepository) UpdateStatus(id string, status model.EncaminhamentoStatus, observacoes string) error {
	query := `
		UPDATE encaminhamentos 
		SET status = $1, observacoes = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := r.DB.Exec(
		query,
		status,
		observacoes,
		time.Now(),
		id,
	)

	return err
}

func (r *encaminhamentoRepository) UpdateAgendamento(id string, dataAgendamento time.Time) error {
	query := `
		UPDATE encaminhamentos 
		SET data_agendamento = $1, status = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := r.DB.Exec(
		query,
		dataAgendamento,
		model.EncaminhamentoStatusAgendado,
		time.Now(),
		id,
	)

	return err
}