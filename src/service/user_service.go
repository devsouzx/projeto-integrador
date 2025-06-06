package service

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	LoginUserService(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, string, error)
}

func (ud *userDomainService)  LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, error) {
	user, err := ud.userRepository.FindPacienteByEmailOrCPFAndPassword(
		userDomain.GetEmail(),
		userDomain.GetCPF(),
		userDomain.GetPassword(),
	)
	fmt.Println("User found:", user)
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}
	
	return user, token, nil
}
