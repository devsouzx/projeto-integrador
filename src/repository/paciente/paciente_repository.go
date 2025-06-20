package paciente

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/datasus"
	"strconv"
)

func NewPacienteRepository(
	db *sql.DB,
) PacienteRepository {
	return &pacienteRepository{
		DB: db,
	}
}

type pacienteRepository struct {
	DB *sql.DB
}

type PacienteRepository interface {
	ListarPacientes(page, pageSize int, search string) ([]*model.Paciente, int, error)
	FindPacienteByID(id string) (*model.Paciente, error)
	FindAnamneseByPacienteID(pacienteId string) (*model.Anamnese, error)
	FindExamesByPacienteID(pacienteID string) ([]*model.ExameClinico, error)
	FindFichasByPacienteID(pacienteID string) ([]*model.FichaCitopatologica, error)
}

func (pr *pacienteRepository) ListarPacientes(page, pageSize int, search string) ([]*model.Paciente, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	countQuery := "SELECT COUNT(*) FROM paciente"
	countArgs := []interface{}{}

	if search != "" {
		countQuery += " WHERE nome ILIKE $1 OR cpf ILIKE $1 OR cns ILIKE $1"
		countArgs = append(countArgs, "%"+search+"%")
	}

	var total int
	err := pr.DB.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao contar pacientes: %w", err)
	}

	query := `
		SELECT 
			id, nome, email, cpf, cns, nome_mae, data_nascimento, 
			telefone, apelido, raca, nacionalidade, escolaridade, 
			created_at, updated_at
		FROM paciente`

	args := []interface{}{}
	argPos := 1

	if search != "" {
		query += fmt.Sprintf(" WHERE nome ILIKE $%d OR cpf ILIKE $%d OR cns ILIKE $%d", argPos, argPos, argPos)
		args = append(args, "%"+search+"%")
		argPos++
	}

	query += " ORDER BY nome ASC"

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, pageSize, (page-1)*pageSize)

	rows, err := pr.DB.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao buscar pacientes: %w", err)
	}
	defer rows.Close()

	var pacientes []*model.Paciente
	for rows.Next() {
		var p model.Paciente
		var nascimentoTime time.Time

		err := rows.Scan(
			&p.ID, &p.Name, &p.Email, &p.CPF, &p.CNS, &p.NomeMae, &nascimentoTime,
			&p.Telefone, &p.Apelido, &p.RacaCor, &p.Nacionalidade, &p.Escolaridade,
			&p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("erro ao scanear paciente: %w", err)
		}

		p.NascimentoTime = nascimentoTime

		exames, err := pr.FindExamesByPacienteID(p.ID)
		if err != nil {
			return nil, 0, fmt.Errorf("erro ao buscar exames: %w", err)
		}

		if len(exames) > 0 {
			p.UltimaInspecao = exames[0].InspecaoColo
		} else {
			p.UltimaInspecao = ""
		}

		pacientes = append(pacientes, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("erro após iterar pacientes: %w", err)
	}

	return pacientes, total, nil
}

func (pr *pacienteRepository) FindPacienteByID(id string) (*model.Paciente, error) {
	query := `
        SELECT 
            id, nome, email, cpf, cns, nome_mae, data_nascimento, 
            telefone, apelido, raca, nacionalidade, escolaridade,
            created_at, updated_at
        FROM paciente
        WHERE id = $1
    `

	var paciente model.Paciente
	var nascimentoTime time.Time

	err := pr.DB.QueryRow(query, id).Scan(
		&paciente.ID,
		&paciente.Name,
		&paciente.Email,
		&paciente.CPF,
		&paciente.CNS,
		&paciente.NomeMae,
		&nascimentoTime,
		&paciente.Telefone,
		&paciente.Apelido,
		&paciente.RacaCor,
		&paciente.Nacionalidade,
		&paciente.Escolaridade,
		&paciente.CreatedAt,
		&paciente.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("paciente não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar paciente: %w", err)
	}

	paciente.NascimentoTime = nascimentoTime

	endereco, err := pr.buscarEnderecoPorPacienteID(paciente.ID)
	if err == nil {
		paciente.Endereco = endereco
	}

	return &paciente, nil
}

func (pr *pacienteRepository) buscarEnderecoPorPacienteID(pacienteId string) (*model.Endereco, error) {
	query := `
        SELECT 
            id, cep, logradouro, numero, complemento, bairro, cidade, 
            uf, created_at, updated_at
        FROM endereco
        WHERE paciente_id = $1
    `

	var endereco model.Endereco

	err := pr.DB.QueryRow(query, pacienteId).Scan(
		&endereco.ID,
		&endereco.CEP,
		&endereco.Logradouro,
		&endereco.Numero,
		&endereco.Complemento,
		&endereco.Bairro,
		&endereco.Cidade,
		&endereco.UF,
		&endereco.CreatedAt,
		&endereco.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("endereco não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar endereco: %w", err)
	}

	return &endereco, nil
}

func (pr *pacienteRepository) FindAnamneseByPacienteID(pacienteId string) (*model.Anamnese, error) {
	query := `
        SELECT 
            id, motivo_exame, fez_exame, usa_diu, 
            gravida, anticoncepcional, hormonio_menopausa, radioterapia,
            dum, nao_lembra_dum, sangramento_pos_coito, 
            sangramento_pos_menopausa, paciente_id, ficha_id
        FROM anamnese
        WHERE paciente_id = $1
    `

	var anamnese model.Anamnese
	var dum sql.NullTime

	err := pr.DB.QueryRow(query, pacienteId).Scan(
		&anamnese.ID,
		&anamnese.MotivoExame,
		&anamnese.FezExame,
		&anamnese.UsaDIU,
		&anamnese.Gravida,
		&anamnese.Anticoncepcional,
		&anamnese.HormonioMenopausa,
		&anamnese.Radioterapia,
		&dum,
		&anamnese.NaoLembraDUM,
		&anamnese.SangramentoPosCoito,
		&anamnese.SangramentoPosMenopausa,
		&anamnese.PacienteID,
		&anamnese.FichaID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("anamnese não encontrada")
		}
		return nil, fmt.Errorf("erro ao buscar anamnese: %w", err)
	}

	anamnese.DUM = dum

	return &anamnese, nil
}

func (pr *pacienteRepository) FindExamesByPacienteID(pacienteID string) ([]*model.ExameClinico, error) {
	query := `
        SELECT 
            id, inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id, created_at
        FROM exame
        WHERE paciente_id = $1
        ORDER BY created_at DESC
    `

	rows, err := pr.DB.Query(query, pacienteID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar exames: %w", err)
	}
	defer rows.Close()

	var exames []*model.ExameClinico
	for rows.Next() {
		var exame model.ExameClinico

		err := rows.Scan(
			&exame.ID,
			&exame.InspecaoColo,
			&exame.SinaisDST,
			&exame.Observacoes,
			&exame.PacienteID,
			&exame.FichaID,
			&exame.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao scanear exame: %w", err)
		}

		exames = append(exames, &exame)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro após iterar exames: %w", err)
	}

	return exames, nil
}

func (pr *pacienteRepository) FindFichasByPacienteID(pacienteID string) ([]*model.FichaCitopatologica, error) {
	query := `
        SELECT 
            id, protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, created_at, unidade_id
        FROM ficha
        WHERE paciente_id = $1
        ORDER BY data_coleta DESC
    `

	rows, err := pr.DB.Query(query, pacienteID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar fichas: %w", err)
	}
	defer rows.Close()

	var fichas []*model.FichaCitopatologica
	for rows.Next() {
		var ficha model.FichaCitopatologica

		err := rows.Scan(
			&ficha.ID,
			&ficha.Protocolo,
			&ficha.DataColeta,
			&ficha.ResponsavelColeta,
			&ficha.MotivoExame,
			&ficha.Observacoes,
			&ficha.CreatedAt,
			&ficha.UnidadeID,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao scanear exame: %w", err)
		}

		unidadeIDInt, err := strconv.Atoi(ficha.UnidadeID)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter UnidadeID para int: %w", err)
		}
		unidade, err := datasus.NewCNESService().BuscarUnidadePorCNES(unidadeIDInt)
		if err != nil {
			return nil, fmt.Errorf("erro ao buscar unidade: %w", err)
		}
		ficha.Unidade = unidade.NomeRazaoSocial

		fichas = append(fichas, &ficha)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro após iterar fichas: %w", err)
	}

	return fichas, nil
}