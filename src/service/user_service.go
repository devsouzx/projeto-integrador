package service

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
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
		loginRequest request.LoginRequest,
	) (model.User, string, error)
}

func (ud *userDomainService) LoginUserService(loginRequest request.LoginRequest) (model.User, string, error) {
	user, err := ud.userRepository.FindUserByIndetifierAndPassword(
		loginRequest.Identifier,
		loginRequest.Password,
		loginRequest.Role,
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
