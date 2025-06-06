package repository

import (
	"database/sql"
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/entity"
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
	FindPacienteByEmailOrCPFAndPassword(
		email string,
		cpf string,
		password string,
	) (model.UserDomainInterface, error)
}

func (ur *userRepository) FindPacienteByEmailOrCPFAndPassword(email string, cpf string, password string) (model.UserDomainInterface, error) {
	query := `
		SELECT id, email, cpf, name, user_password, role
		FROM users
		WHERE email = $1 OR cpf = $2
	`

	userEntity := &entity.UserEntity{}
	err := ur.DB.QueryRow(query, email, cpf).Scan(
		&userEntity.ID,
		&userEntity.Email,
		&userEntity.CPF,
		&userEntity.Name,
		&userEntity.Password,
		&userEntity.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário ou senha inválidos")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("usuário ou senha inválidos")
	}

	domain := model.NewUserDomain(
		userEntity.Email,
		userEntity.CPF,
		userEntity.Password,
		userEntity.Name,
		userEntity.Role,
	)

	domain.SetID(userEntity.ID)

	return domain, nil
}

