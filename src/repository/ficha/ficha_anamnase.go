package ficha

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func (fr *fichaRepository) UpsertAnamnase(anamnase *model.Anamnese, pacienteId string) error {
    if pacienteId == "" {
        return fmt.Errorf("ID do paciente é obrigatório")
    }

    var existingId string
    err := fr.DB.QueryRow(`
        SELECT id FROM anamnese WHERE paciente_id = $1 LIMIT 1
    `, pacienteId).Scan(&existingId)

    var dum interface{}
    if anamnase.NaoLembraDUM || anamnase.DUM.Time.IsZero() {
        dum = nil
    } else {
        dum = anamnase.DUM.Time
    }

    if err == nil {
        _, err = fr.DB.Exec(`
            UPDATE anamnese SET
                motivo_exame = $1, fez_exame = $2, ultimo_exame_ano = $3,
                usa_diu = $4, gravida = $5, anticoncepcional = $6,
                hormonio_menopausa = $7, radioterapia = $8, dum = $9,
                nao_lembra_dum = $10, sangramento_pos_coito = $11,
                sangramento_pos_menopausa = $12, updated_at = NOW()
            WHERE id = $13
        `, anamnase.MotivoExame, anamnase.FezExame, anamnase.UltimoExameAno,
            anamnase.UsaDIU, anamnase.Gravida, anamnase.Anticoncepcional,
            anamnase.HormonioMenopausa, anamnase.Radioterapia, dum,
            anamnase.NaoLembraDUM, anamnase.SangramentoPosCoito,
            anamnase.SangramentoPosMenopausa, existingId)
    } else if errors.Is(err, sql.ErrNoRows) {
        _, err = fr.DB.Exec(`
            INSERT INTO anamnese 
                (motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida,
                anticoncepcional, hormonio_menopausa, radioterapia, dum,
                nao_lembra_dum, sangramento_pos_coito, sangramento_pos_menopausa,
                paciente_id, ficha_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        `, anamnase.MotivoExame, anamnase.FezExame, anamnase.UltimoExameAno,
            anamnase.UsaDIU, anamnase.Gravida, anamnase.Anticoncepcional,
            anamnase.HormonioMenopausa, anamnase.Radioterapia, dum,
            anamnase.NaoLembraDUM, anamnase.SangramentoPosCoito,
            anamnase.SangramentoPosMenopausa, pacienteId, anamnase.FichaID)
    }
    
    if err != nil {
        return fmt.Errorf("erro ao criar/atualizar anamnese: %v", err)
    }
    return nil
}