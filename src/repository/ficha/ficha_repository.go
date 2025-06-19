package ficha

import (
	"database/sql"
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
)

func NewFichaRepository(
	database *sql.DB,
) FichaRepositoryInterface {
	return &fichaRepository{
		DB: database,
	}
}

type fichaRepository struct {
	DB *sql.DB
}

type FichaRepositoryInterface interface {
	Create(ficha *model.FichaCitopatologica) (*model.FichaCitopatologica, error)
	FindByID(id string) (*model.FichaCitopatologica, error)
	CreateEndereco(endereco *model.Endereco, pacienteId string) error
	UpdateEndereco(endereco *model.Endereco, pacienteId string) error
	UpsertEndereco(endereco *model.Endereco, pacienteId string) error
	UpsertAnamnase(anamnase *model.Anamnese, pacienteId string) error
	RegistrarExame(exame *model.ExameClinico, pacienteId string) error
}

func (fr *fichaRepository) RegistrarExame(exame *model.ExameClinico, pacienteId string) error {
	query := `INSERT INTO exame (inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id) 
	          VALUES ($1, $2, $3, $4, $5)`

	if pacienteId == "" {
		return fmt.Errorf("PacienteID está vazio (UUID inválido)")
	}
	if exame.FichaID == "" {
		return fmt.Errorf("FichaID está vazio (UUID inválido)")
	}

	_, err := fr.DB.Exec(
		query,
		exame.InspecaoColo,
		exame.SinaisDST,
		exame.Observacoes,
		pacienteId,
		exame.FichaID,
	)

	if err != nil {
		return fmt.Errorf("erro ao registrar exame: %v", err)
	}

	return nil
}
