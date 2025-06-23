package laudo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
)

type LaudoRepositoryInterface interface {
    Create(laudo *model.Laudo) error
    FindByMedicoID(medicoID string, resultadoFilter string, search string, page int, pageSize int) ([]*model.Laudo, int, error)
}

type laudoRepository struct {
	DB *sql.DB
}

func NewLaudoRepository(db *sql.DB) LaudoRepositoryInterface {
	return &laudoRepository{DB: db}
}

func (r *laudoRepository) Create(laudo *model.Laudo) error {
	query := `
		INSERT INTO laudos (
			paciente_id, medico_id, ficha_id, data_exame, data_laudo, 
			resultado, cid, adequabilidade, microbiologia, 
			celulas_endometriais, comentarios, recomendacoes, status
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	laudo.CID = laudo.GetCID()
	laudo.CreatedAt = time.Now()
	laudo.UpdatedAt = time.Now()

	_, err := r.DB.Exec(
		query,
		laudo.PacienteID,
		laudo.MedicoID,
		laudo.FichaID,
		laudo.DataExame,
		laudo.DataLaudo,
		laudo.Resultado,
		laudo.CID,
		laudo.Adequabilidade,
		laudo.Microbiologia,
		laudo.CelulasEndometriais,
		laudo.Comentarios,
		laudo.Recomendacoes,
		laudo.Status,
	)

	return err
}

func (r *laudoRepository) FindByMedicoID(medicoID string, resultadoFilter string, search string, page int, pageSize int) ([]*model.Laudo, int, error) {
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
    `

    args := []interface{}{medicoID}
    argCounter := 2
	
    if resultadoFilter != "" {
        query += fmt.Sprintf(" AND l.resultado = $%d", argCounter)
        args = append(args, resultadoFilter)
        argCounter++
    }

    if search != "" {
        query += fmt.Sprintf(" AND (p.nome ILIKE $%d OR p.cns ILIKE $%d OR l.cid ILIKE $%d)", 
            argCounter, argCounter+1, argCounter+2)
        args = append(args, "%"+search+"%", "%"+search+"%", "%"+search+"%")
        argCounter += 3
    }
	
    query += " ORDER BY l.data_exame DESC"
    query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
    args = append(args, pageSize, (page-1)*pageSize)
	
    rows, err := r.DB.Query(query, args...)
    if err != nil {
        return nil, 0, err
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
            return nil, 0, err
        }

        laudo.Paciente = &paciente
        laudos = append(laudos, &laudo)
    }
	
    countQuery := "SELECT COUNT(*) FROM laudos l INNER JOIN paciente p ON p.id = l.paciente_id WHERE l.medico_id = $1"
    countArgs := []interface{}{medicoID}

    if resultadoFilter != "" {
        countQuery += " AND l.resultado = $2"
        countArgs = append(countArgs, resultadoFilter)
    }

    if search != "" {
        countQuery += " AND (p.nome ILIKE $2 OR p.cns ILIKE $3 OR l.cid ILIKE $4)"
        if resultadoFilter != "" {
            countQuery = strings.Replace(countQuery, "$2", "$3", 1)
            countQuery = strings.Replace(countQuery, "$3", "$4", 1)
            countQuery = strings.Replace(countQuery, "$4", "$5", 1)
        }
        countArgs = append(countArgs, "%"+search+"%", "%"+search+"%", "%"+search+"%")
    }

    var total int
    err = r.DB.QueryRow(countQuery, countArgs...).Scan(&total)
    if err != nil {
        return nil, 0, err
    }

    return laudos, total, nil
}
