package ficha

import (
	"fmt"
	"strconv"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func (fr *fichaRepository) Create(ficha *model.FichaCitopatologica) (*model.FichaCitopatologica, error) {
	unidadeID, _ := strconv.Atoi(ficha.UnidadeID)

	query := `INSERT INTO ficha (protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, paciente_id, responsavel_id, unidade_id) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	          RETURNING id, created_at, updated_at`

	err := fr.DB.QueryRow(query,
		ficha.Protocolo, ficha.DataColeta, ficha.ResponsavelColeta, ficha.MotivoExame, ficha.Observacoes,
		ficha.PacienteID, ficha.ResponsavelID, unidadeID).Scan(&ficha.ID, &ficha.CreatedAt, &ficha.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar ficha: %v", err)
	}

	return ficha, nil
}