package laudo

import (
	"database/sql"

	"github.com/devsouzx/projeto-integrador/src/model"
)

type LaudoRepositoryInterface interface {
	FindByMedicoID(medicoID string) ([]*model.Laudo, error)
}

type laudoRepository struct {
	DB *sql.DB
}

func NewLaudoRepository(db *sql.DB) LaudoRepositoryInterface {
	return &laudoRepository{DB: db}
}

func (r *laudoRepository) FindByMedicoID(medicoID string) ([]*model.Laudo, error) {
	query := `
		SELECT 
			l.id, l.paciente_id, l.medico_id, l.ficha_id, l.data_exame, l.data_laudo,
			l.resultado, l.cid, l.adequabilidade, l.microbiologia, 
			l.celulas_endometriais, l.comentarios, l.recomendacoes, l.status,
			l.created_at, l.updated_at,
			p.id, p.nome, p.cns
		FROM laudos l
		INNER JOIN paciente p ON p.id = l.paciente_id
		WHERE l.medico_id = $1
		ORDER BY l.data_exame DESC
	`

	rows, err := r.DB.Query(query, medicoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laudos []*model.Laudo
	for rows.Next() {
		var laudo model.Laudo
		var paciente model.Paciente

		err := rows.Scan(
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
			&paciente.ID,
			&paciente.Name,
			&paciente.CNS,
		)
		if err != nil {
			return nil, err
		}

		laudo.Paciente = &paciente
		laudos = append(laudos, &laudo)
	}

	return laudos, nil
}
