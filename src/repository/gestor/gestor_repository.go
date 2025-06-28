package gestor

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
)

func NewGestorRepository(
	db *sql.DB,
) GestorRepositoryInterface {
	return &gestorRepository{
		DB: db,
	}
}

type gestorRepository struct {
	DB *sql.DB
}

type GestorRepositoryInterface interface {
	CadastrarMedico(medico request.CadastroProfissionalRequest) (*model.Medico, error)
	CadastrarEnfermeiro(enfermeiro request.CadastroProfissionalRequest) (*model.Enfermeiro, error)
	CadastrarAgente(agente request.CadastroProfissionalRequest) (*model.AgenteComunitario, error)
	CadastrarGestor(gestor request.CadastroProfissionalRequest) (*model.Gestor, error)

	FindGestorByIdentifier(identifier string) (*model.Gestor, error)
}

func (gr *gestorRepository) FindGestorByIdentifier(identifier string) (*model.Gestor, error) {
	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")

	query := `
		SELECT id, email, senha, nomecompleto, cpf
		FROM gestor
		WHERE id = $1 OR cpf = $2
	`

	var gestor model.Gestor

	err := gr.DB.QueryRow(query, identifier, cpfFormatado).Scan(
		&gestor.ID,
		&gestor.Email,
		&gestor.Password,
		&gestor.Name,
		&gestor.CPF,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("gestor não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar gestor: %w", err)
	}

	return &gestor, nil
}

func (gr *gestorRepository) CadastrarMedico(medicoRequest request.CadastroProfissionalRequest) (*model.Medico, error) {
	query := `
    INSERT INTO medico (
        crm, email, senha, nomecompleto, cpf, especialidade, telefone, unidade_saude_id
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, created_at, updated_at`

	medico := model.NewMedico(
		medicoRequest.Email,
		medicoRequest.Password,
		medicoRequest.Name,
		medicoRequest.CRM,
		medicoRequest.CPF,
		medicoRequest.Especialidade,
		medicoRequest.Telefone,
		24222,
	)
	medico.EncryptPassword()

	crmFormatado := strings.ReplaceAll(medicoRequest.CRM, "/", "")

	err := gr.DB.QueryRow(
		query,
		crmFormatado,
		medicoRequest.Email,
		medico.Password,
		medicoRequest.Name,
		medicoRequest.CPF,
		medicoRequest.Especialidade,
		medicoRequest.Telefone,
		medico.UnidadeID,
	).Scan(&medico.ID, &medico.CreatedAt, &medico.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar medico: %v", err)
	}

	return medico, nil
}

func (gr *gestorRepository) CadastrarEnfermeiro(enfermeiroRequest request.CadastroProfissionalRequest) (*model.Enfermeiro, error) {
	query := `
    INSERT INTO enfermeiro (
        coren, email, senha, nomecompleto, cpf, telefone
    ) VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, created_at, updated_at`

	enfermeiro := model.NewEnfermeiro(
		enfermeiroRequest.Email,
		enfermeiroRequest.Password,
		enfermeiroRequest.Name,
		enfermeiroRequest.COREN,
		enfermeiroRequest.CPF,
		enfermeiroRequest.Telefone,
	)
	enfermeiro.EncryptPassword()

	corenFormatado := strings.ReplaceAll(enfermeiroRequest.COREN, "/", "")

	err := gr.DB.QueryRow(
		query,
		corenFormatado,
		enfermeiroRequest.Email,
		enfermeiro.Password,
		enfermeiroRequest.Name,
		enfermeiroRequest.CPF,
		enfermeiroRequest.Telefone,
	).Scan(&enfermeiro.ID, &enfermeiro.CreatedAt, &enfermeiro.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar enfermeiro: %v", err)
	}

	return enfermeiro, nil
}

func (gr *gestorRepository) CadastrarAgente(agenteRequest request.CadastroProfissionalRequest) (*model.AgenteComunitario, error) {
	query := `
	INSERT INTO agente_comunitario (
    email, senha, nomecompleto, cpf, telefone, unidade_saude_id
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, updated_at`

	agente := model.NewAgenteComunitario(
		agenteRequest.Email,
		agenteRequest.Password,
		agenteRequest.Name,
		agenteRequest.CPF,
		agenteRequest.Telefone,
		agenteRequest.UnidadeID,
	)
	agente.EncryptPassword()

	err := gr.DB.QueryRow(
		query,
		agenteRequest.Email,
		agente.Password,
		agenteRequest.Name,
		agenteRequest.CPF,
		agenteRequest.Telefone,
		24222,
	).Scan(&agente.ID, &agente.CreatedAt, &agente.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar agente: %v", err)
	}

	return agente, nil
}

func (gr *gestorRepository) CadastrarGestor(gestorRequest request.CadastroProfissionalRequest) (*model.Gestor, error) {
	query := `
    INSERT INTO gestor (
        email, senha, nomecompleto, cpf, telefone, unidade_saude_id
    ) VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, created_at, updated_at`

	gestor := model.NewGestor(
		gestorRequest.Email,
		gestorRequest.Password,
		gestorRequest.Name,
		gestorRequest.CPF,
		gestorRequest.Telefone,
		gestorRequest.UnidadeID,
	)
	gestor.EncryptPassword()

	err := gr.DB.QueryRow(
		query,
		gestorRequest.Email,
		gestor.Password,
		gestorRequest.Name,
		gestorRequest.CPF,
		gestorRequest.Telefone,
		24222,
	).Scan(&gestor.ID, &gestor.CreatedAt, &gestor.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar gestor: %v", err)
	}

	return gestor, nil
}
