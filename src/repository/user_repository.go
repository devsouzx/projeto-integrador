package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"golang.org/x/crypto/bcrypt"
)

func NewUserRepository(
	database *sql.DB,
) UserRepository {
	return &userRepository{
		DB: database,
	}
}

type userRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	FindUserByIdentifierAndPassword(
		indentifier string,
		password string,
		role string,
	) (model.UserInterface, error)
	CreateRecoveryCode(identifier string, code string, expiresAt time.Time) (string, error)
}

func (ur *userRepository) FindUserByIdentifierAndPassword(identifier, password, role string) (model.UserInterface, error) {
	var query string
	var user model.UserInterface

	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")
	clean := strings.ToUpper(strings.ReplaceAll(identifier, "/", ""))

	switch role {
	case "paciente":
		query = `
			SELECT id, email, senha, nomecompleto, cpf
			FROM paciente
			WHERE email = $1 OR cpf = $2
		`
		var paciente model.Paciente

		err := ur.DB.QueryRow(query, identifier, cpfFormatado).Scan(
			&paciente.ID,
			&paciente.Email,
			&paciente.Password,
			&paciente.Name,
			&paciente.CPF,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("usuário ou senha inválidos")
			}
			return nil, fmt.Errorf("erro ao buscar paciente: %w", err)
		}

		user = &paciente
	case "medico":
		query = `
			SELECT id, email, senha, nomecompleto, cpf, crm
			FROM medico
			WHERE crm = $1 OR cpf = $2
		`
		var medico model.Medico

		err := ur.DB.QueryRow(query, clean, cpfFormatado).Scan(
			&medico.ID,
			&medico.Email,
			&medico.Password,
			&medico.Name,
			&medico.CRM,
			&medico.CPF,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("usuário ou senha inválidos")
			}
			return nil, fmt.Errorf("erro ao buscar médico: %w", err)
		}

		user = &medico
	case "enfermeiro":
		query = `
			SELECT id, email, senha, nomecompleto, cpf, coren, telefone
			FROM enfermeiro
		    WHERE coren = $1 OR cpf = $2
	    `
		var enfermeiro model.Enfermeiro

		err := ur.DB.QueryRow(query, clean, cpfFormatado).Scan(
			&enfermeiro.ID,
			&enfermeiro.Email,
			&enfermeiro.Password,
			&enfermeiro.Name,
			&enfermeiro.CPF,
			&enfermeiro.COREN,
			&enfermeiro.Telefone,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("usuário ou senha inválidos")
			}
			return nil, fmt.Errorf("erro ao buscar enfermeiro: %w", err)
		}

		user = &enfermeiro
	case "agente":
		query = `
			SELECT id, email, senha, nomecompleto, cpf, telefone, unidade_saude_id
			FROM agente_comunitario
			WHERE email = $1 OR cpf = $2
		`
		var agente model.AgenteComunitario

		err := ur.DB.QueryRow(query, identifier, cpfFormatado).Scan(
			&agente.ID,
			&agente.Email,
			&agente.Password,
			&agente.Name,
			&agente.CPF,
			&agente.Telefone,
			&agente.UnidadeID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("usuário ou senha inválidos")
			}
			return nil, fmt.Errorf("erro ao buscar agente comunitário: %w", err)
		}

		user = &agente
	case "gestor":
		query = `
			SELECT id, email, senha, nomecompleto, cpf, telefone, unidade_saude_id
			FROM gestor
			WHERE email = $1 OR cpf = $2
		`
		var gestor model.Gestor

		err := ur.DB.QueryRow(query, identifier, cpfFormatado).Scan(
			&gestor.ID,
			&gestor.Email,
			&gestor.Password,
			&gestor.Name,
			&gestor.CPF,
			&gestor.Telefone,
			&gestor.UnidadeID,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("usuário ou senha inválidos")
			}
			return nil, fmt.Errorf("erro ao buscar gestor: %w", err)
		}

		user = &gestor
	default:
		return nil, fmt.Errorf("tipo de usuário inválido")
	}

	// 2a$10$MJsoWkTYzYBR1elzP13y8eR1cfOJ9KO7yXvErBzXPQW8MrnZTWR6q
	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(password)); err != nil {
		return nil, fmt.Errorf("usuário ou senha inválidos")
	}

	user.SetRole(role)

	return user, nil
}

func (ur *userRepository) CreateRecoveryCode(identifier string, code string, expiresAt time.Time) (string, error) {
    userEmail, err := ur.findUserEmail(identifier)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", fmt.Errorf("usuário não encontrado")
        }
        return "", fmt.Errorf("erro ao buscar usuário: %w", err)
    }

    _, err = ur.DB.Exec(`
        INSERT INTO password_recovery (identifier, code, expires_at)
        VALUES ($1, $2, $3)
        ON CONFLICT (identifier) 
        DO UPDATE SET code = $2, expires_at = $3, created_at = NOW()
    `, identifier, code, expiresAt)

    if err != nil {
        return "", fmt.Errorf("erro ao criar código de recuperação: %w", err)
    }
    
    return userEmail, nil
}

func (ur *userRepository) findUserEmail(identifier string) (string, error) {
    cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", "")
    
    tables := []string{"paciente", "medico", "enfermeiro", "agente_comunitario", "gestor"}
    
    for _, table := range tables {
        var email string
        err := ur.DB.QueryRow(fmt.Sprintf(`
            SELECT email FROM %s WHERE email = $1 OR cpf = $2 LIMIT 1
        `, table), identifier, cpfFormatado).Scan(&email)
        
        if err == nil {
            return email, nil
        } else if err != sql.ErrNoRows {
            return "", err
        }
    }
    
    return "", sql.ErrNoRows
}