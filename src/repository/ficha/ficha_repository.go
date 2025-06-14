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
	CreateEndereco(endereco *model.Endereco) error
	// FindByID(id string) (*model.FichaCitopatologica, error)
	// FindByPacienteID(pacienteID string) ([]*model.FichaCitopatologica, error)
	// Update(ficha *model.FichaCitopatologica) error
}

func (fr *fichaRepository) Create(ficha *model.FichaCitopatologica) (*model.FichaCitopatologica, error) {
	query := `INSERT INTO ficha (protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, paciente_id, responsavel_id) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7) 
	          RETURNING id, created_at, updated_at`

	err := fr.DB.QueryRow(query, 
		ficha.Protocolo, ficha.DataColeta, ficha.ResponsavelColeta, ficha.MotivoExame, ficha.Observacoes,
		ficha.PacienteID, ficha.ResponsavelID).Scan(&ficha.ID, &ficha.CreatedAt, &ficha.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar ficha: %v", err)
	}

	return ficha, nil
}

func (fr *fichaRepository) FindByID(id string) (*model.FichaCitopatologica, error) {
	query := `SELECT id, protocolo, cnes, unidade_saude, municipio, uf, data_coleta, 
	          responsavel_coleta, motivo_exame, observacoes, paciente_id, responsavel_id, 
	          created_at, updated_at FROM fichas WHERE id = $1`

	var ficha model.FichaCitopatologica
	err := fr.DB.QueryRow(query, id).Scan(
		&ficha.ID, &ficha.Protocolo, &ficha.CNES, &ficha.UnidadeSaude, &ficha.Municipio,
		&ficha.UF, &ficha.DataColeta, &ficha.ResponsavelColeta, &ficha.MotivoExame,
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

func (fr *fichaRepository) CreateEndereco(endereco *model.Endereco) error {
	if endereco.PacienteID == "" {
        return fmt.Errorf("ID do paciente é obrigatório")
    }

	err := fr.DB.QueryRow(`INSERT INTO endereco (cep, logradouro, complemento, numero, bairro, cidade, 
	          uf, referencia, codigo_municipio, paciente_id) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	          RETURNING id`, 
		endereco.CEP, endereco.Logradouro, endereco.Complemento, endereco.Numero, endereco.Bairro, endereco.Municipio, endereco.UF, endereco.Referencia, endereco.CodigoMunicipio, endereco.PacienteID).Scan(&endereco.ID)
	if err != nil {
		return fmt.Errorf("erro ao criar dados residenciais: %v", err)
	}
	return nil
}