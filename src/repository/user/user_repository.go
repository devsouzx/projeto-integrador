package user

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
	VerifyRecoveryCode(identifier string, code string) (bool, error)
	UpdatePassword(identifier string, newPassword string) error
	FindUserByIdentifier(identifier string, role string) (model.UserInterface, error)
	CreatePaciente(paciente *model.Paciente) (*model.Paciente, error)
	FindByVerificationToken(token string) (model.UserInterface, error)
	UpdateUser(user model.UserInterface) error
	FindByID(id string, role string) (model.UserInterface, error)
	UpdatePaciente(paciente *model.Paciente) error
}

func (ur *userRepository) FindUserByIdentifierAndPassword(identifier, password, role string) (model.UserInterface, error) {
	user, err := ur.FindUserByIdentifier(identifier, role)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuario: %s", err)
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

func (ur *userRepository) VerifyRecoveryCode(identifier string, code string) (bool, error) {
	var valid bool
	err := ur.DB.QueryRow(`
        SELECT code = $2 AND expires_at > NOW()
        FROM password_recovery
        WHERE identifier = $1
    `, identifier, code).Scan(&valid)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("erro ao verificar código: %w", err)
	}

	return valid, nil
}

func (ur *userRepository) UpdatePassword(identifier string, newPassword string) error {
	tables := []struct {
		name  string
		query string
	}{
		{"paciente", "UPDATE paciente SET senha = $1 WHERE email = $2 OR cpf = $3"},
		{"medico", "UPDATE medico SET senha = $1 WHERE email = $2 OR cpf = $3"},
		{"enfermeiro", "UPDATE enfermeiro SET senha = $1 WHERE email = $2 OR cpf = $3"},
		{"agente_comunitario", "UPDATE agente_comunitario SET senha = $1 WHERE email = $2 OR cpf = $3"},
		{"gestor", "UPDATE gestor SET senha = $1 WHERE email = $2 OR cpf = $3"},
	}

	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")

	for _, table := range tables {
		result, err := ur.DB.Exec(table.query, newPassword, identifier, cpfFormatado)
		if err != nil {
			return fmt.Errorf("erro ao atualizar senha na tabela %s: %w", table.name, err)
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected > 0 {
			_, err = ur.DB.Exec("DELETE FROM password_recovery WHERE identifier = $1", identifier)
			if err != nil {
				return fmt.Errorf("erro ao limpar código de recuperação: %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("usuário não foi encontrado em nenhuma tabela")
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

func (ur *userRepository) FindUserByIdentifier(identifier string, role string) (model.UserInterface, error) {
	var query string
	var user model.UserInterface

	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(identifier, ".", ""), "-", ""), " ", "")
	clean := strings.ToUpper(strings.ReplaceAll(identifier, "/", ""))

	switch role {
	case "paciente":
		query = `
			SELECT id, 
            email, 
            senha, 
            nome AS name,
            cpf, 
            cns, 
            nome_mae, 
            data_nascimento,
            telefone, 
            apelido, 
            raca AS raca_cor, 
            nacionalidade,
            escolaridade,
            created_at, 
            updated_at,
			is_verified
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
			&paciente.CNS,
			&paciente.NomeMae,
			&paciente.NascimentoTime,
			&paciente.Telefone,
			&paciente.Apelido,
			&paciente.RacaCor,
			&paciente.Nacionalidade,
			&paciente.Escolaridade,
			&paciente.CreatedAt,
			&paciente.UpdatedAt,
			&paciente.Verified,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("usuário ou senha inválidos")
			}
			return nil, fmt.Errorf("erro ao buscar paciente: %w", err)
		}

		endereco, err := ur.buscarEnderecoPorPacienteID(paciente.ID)
		if err == nil {
			paciente.Endereco = endereco
		} else {
			fmt.Printf("Erro ao buscar endereço: %v\n", err)
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

	return user, nil
}

func (r *userRepository) CreatePaciente(paciente *model.Paciente) (*model.Paciente, error) {
	if paciente.DataNascimento != "" {
		nascimento, err := time.Parse("2006-01-02", paciente.DataNascimento)
		if err != nil {
			return nil, fmt.Errorf("formato de data inválido: %v", err)
		}
		paciente.NascimentoTime = nascimento
	}

	paciente.EncryptPassword()
	apelido := strings.Split(paciente.Name, " ")[0]

	query := `
    INSERT INTO paciente (
        email, senha, nome, 
        apelido, nome_mae, cns, cpf, 
        data_nascimento, nacionalidade, 
        raca, escolaridade, telefone
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    RETURNING id, created_at, updated_at`

	err := r.DB.QueryRow(
		query,
		paciente.Email,
		paciente.Password,
		paciente.Name,
		apelido,
		paciente.NomeMae,
		paciente.CNS,
		paciente.CPF,
		paciente.NascimentoTime,
		paciente.Nacionalidade,
		paciente.RacaCor,
		paciente.Escolaridade,
		paciente.Telefone,
	).Scan(&paciente.ID, &paciente.CreatedAt, &paciente.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("erro ao criar paciente: %v", err)
	}

	if paciente.Endereco != nil {
		queryEndereco := `
        INSERT INTO endereco (
            cep, logradouro, numero, complemento, 
            bairro, cidade, uf, paciente_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, created_at, updated_at`

		err = r.DB.QueryRow(
			queryEndereco,
			paciente.Endereco.CEP,
			paciente.Endereco.Logradouro,
			paciente.Endereco.Numero,
			paciente.Endereco.Complemento,
			paciente.Endereco.Bairro,
			paciente.Endereco.Cidade,
			paciente.Endereco.UF,
			paciente.ID,
		).Scan(
			&paciente.Endereco.ID,
			&paciente.Endereco.CreatedAt,
			&paciente.Endereco.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("erro ao criar endereço: %v", err)
		}
	}

	return paciente, nil
}

func (r *userRepository) UpdatePaciente(paciente *model.Paciente) error {
    query := `UPDATE paciente SET nome = $1, 
        apelido = $2, nome_mae = $3, cns = $4, cpf = $5, 
        data_nascimento = $6, nacionalidade = $7, 
        raca = $8, escolaridade = $9, telefone = $10 WHERE id = $11`
    _, err := r.DB.Exec(query, 
		paciente.Name,
		paciente.Apelido,
		paciente.NomeMae,
		paciente.CNS,
		paciente.CPF,
		paciente.NascimentoTime,
		paciente.Nacionalidade,
		paciente.RacaCor,
		paciente.Escolaridade,
		paciente.Telefone,
        paciente.ID)
    return err
}

func (ur *userRepository) FindByID(id string, role string) (model.UserInterface, error) {
	var paciente model.Paciente
	err := ur.DB.QueryRow(`
            SELECT id, email, senha, nome
            FROM paciente 
            WHERE id = $1
        `, id).Scan(
		&paciente.ID,
		&paciente.Email,
		&paciente.Password,
		&paciente.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("paciente não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar paciente: %w", err)
	}

	return &paciente, nil
}

func (ur *userRepository) FindByVerificationToken(token string) (model.UserInterface, error) {
	var user model.BaseUser
	err := ur.DB.QueryRow(`
        SELECT id, email, nome, is_verified, verification_token, token_expires_at AT TIME ZONE 'UTC'
        FROM paciente 
        WHERE verification_token = $1
    `, token).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Verified,
		&user.VerifyToken,
		&user.TokenExpiresAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("token não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return &user, nil
}

func (ur *userRepository) UpdateUser(user model.UserInterface) error {
	_, err := ur.DB.Exec(`
        UPDATE paciente SET
            is_verified = $1,
            verification_token = $2,
            token_expires_at = $3,
            updated_at = NOW()
        WHERE id = $4
    `,
		user.IsVerified(),
		user.GetVerifyToken(),
		user.GetTokenExpiresAt(),
		user.GetID(),
	)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

func (r *userRepository) buscarEnderecoPorPacienteID(pacienteID string) (*model.Endereco, error) {
    query := `
        SELECT 
            id, cep, logradouro, numero, complemento,
            bairro, cidade, uf
        FROM endereco
        WHERE paciente_id = $1
    `

    var endereco model.Endereco
    err := r.DB.QueryRow(query, pacienteID).Scan(
        &endereco.ID,
        &endereco.CEP,
        &endereco.Logradouro,
        &endereco.Numero,
        &endereco.Complemento,
        &endereco.Bairro,
        &endereco.Cidade,
        &endereco.UF,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return &model.Endereco{}, nil
        }
        return nil, fmt.Errorf("erro ao buscar endereço: %w", err)
    }

    return &endereco, nil
}