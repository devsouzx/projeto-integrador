package repository

import (
	"database/sql"
	"fmt"
	"strings"

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
	FindUserByIndetifierAndPassword(
		indentifier string,
		password string,
		role string,
	) (model.User, error)
}

func (ur *userRepository) FindUserByIndetifierAndPassword(indetifier, password, role string) (model.User, error) {
	var query string
	var user model.User

	fmt.Println("Finding user with identifier:", indetifier, "and role:", role)

	cpfFormatado := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(indetifier, ".", ""), "-", ""), " ", "")

	switch role {
	case "paciente":
		fmt.Println("Searching for paciente with email or cpf:", indetifier, cpfFormatado)
		query = `
			SELECT paciente_id, email, senha, nomecompleto, cpf
			FROM paciente
			WHERE email = $1 OR cpf = $2
		`
		var paciente model.Paciente

		err := ur.DB.QueryRow(query, indetifier, cpfFormatado).Scan(
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
			SELECT id, email, cpf, nomecompleto, senha, crm
			FROM medico
			WHERE coren = $1 OR cpf = $2
		`
	default:
		return nil, fmt.Errorf("tipo de usuário inválido")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(password)); err != nil {
		return nil, fmt.Errorf("usuário ou senha inválidos")
	}

	user.SetRole(role)

	return user, nil
}
