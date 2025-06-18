package ficha

import (
	"database/sql"
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func (fr *fichaRepository) FindByID(id string) (*model.FichaCitopatologica, error) {
	query := `SELECT id, protocolo, cnes, unidade_id, data_coleta, 
	          responsavel_coleta, motivo_exame, observacoes, paciente_id, responsavel_id, 
	          created_at, updated_at FROM fichas WHERE id = $1`

	var ficha model.FichaCitopatologica
	err := fr.DB.QueryRow(query, id).Scan(
		&ficha.ID, &ficha.Protocolo, &ficha.UnidadeID, &ficha.DataColeta, &ficha.ResponsavelColeta, &ficha.MotivoExame,
		&ficha.Observacoes, &ficha.PacienteID, &ficha.ResponsavelID, &ficha.CreatedAt, &ficha.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ficha não encontrada")
		}
		return nil, fmt.Errorf("erro ao buscar ficha: %v", err)
	}

	return &ficha, nil
}