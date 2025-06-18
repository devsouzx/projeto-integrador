package ficha

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
	EnderecoExists(pacienteId string) (bool, error)
	UpsertEndereco(endereco *model.Endereco, pacienteId string) error
	// FindByID(id string) (*model.FichaCitopatologica, error)
	// FindByPacienteID(pacienteID string) ([]*model.FichaCitopatologica, error)
	// Update(ficha *model.FichaCitopatologica) error
}

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

func (fr *fichaRepository) CreateEndereco(endereco *model.Endereco, pacienteId string) error {
	if pacienteId == "" {
        return fmt.Errorf("ID do paciente é obrigatório")
    }

	err := fr.DB.QueryRow(`INSERT INTO endereco (cep, logradouro, complemento, numero, bairro, cidade, 
	          uf, paciente_id) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	          RETURNING id`, 
		endereco.CEP, endereco.Logradouro, endereco.Complemento, endereco.Numero, endereco.Bairro, endereco.Cidade, endereco.UF, pacienteId).Scan(&endereco.ID)
	if err != nil {
		return fmt.Errorf("erro ao criar dados residenciais: %v", err)
	}
	return nil
}

func (fr *fichaRepository) UpsertEndereco(endereco *model.Endereco, pacienteId string) error {
    var existingId string
    err := fr.DB.QueryRow(`
        SELECT id FROM endereco WHERE paciente_id = $1 LIMIT 1
    `, pacienteId).Scan(&existingId)

    if err == nil {
        _, err = fr.DB.Exec(`
            UPDATE endereco SET
                cep = $1, logradouro = $2, complemento = $3, numero = $4,
                bairro = $5, cidade = $6, uf = $7, updated_at = NOW()
            WHERE id = $8
        `, endereco.CEP, endereco.Logradouro, endereco.Complemento, endereco.Numero,
           endereco.Bairro, endereco.Cidade, endereco.UF, existingId)
    } else if errors.Is(err, sql.ErrNoRows) {
        _, err = fr.DB.Exec(`
            INSERT INTO endereco 
                (cep, logradouro, complemento, numero, bairro, cidade, uf, paciente_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        `, endereco.CEP, endereco.Logradouro, endereco.Complemento, endereco.Numero,
           endereco.Bairro, endereco.Cidade, endereco.UF, pacienteId)
    }
    return err
}

func (fr *fichaRepository) EnderecoExists(pacienteId string) (bool, error) {
    var exists bool
    query := `SELECT EXISTS(SELECT 1 FROM endereco WHERE paciente_id = $1)`
    err := fr.DB.QueryRow(query, pacienteId).Scan(&exists)
    if err != nil {
        return false, fmt.Errorf("erro ao verificar endereço: %v", err)
    }
    return exists, nil
}

func (fr *fichaRepository) UpdateEndereco(endereco *model.Endereco, pacienteId string) error {
    query := `UPDATE endereco 
              SET cep = $1, logradouro = $2, complemento = $3, numero = $4, 
                  bairro = $5, cidade = $6, uf = $7
              WHERE paciente_id = $8`
    _, err := fr.DB.Exec(query,
        endereco.CEP, endereco.Logradouro, endereco.Complemento, endereco.Numero,
        endereco.Bairro, endereco.Cidade, endereco.UF, pacienteId)
    if err != nil {
        return fmt.Errorf("erro ao atualizar endereço: %v", err)
    }
    return nil
}