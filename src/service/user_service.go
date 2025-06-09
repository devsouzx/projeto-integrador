package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"
	"github.com/devsouzx/projeto-integrador/src/repository"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
	emailService EmailService,
) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
		emailService:   emailService,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
	emailService   EmailService
}

type UserDomainService interface {
	LoginUserService(
		loginRequest request.LoginRequest,
	) (model.UserInterface, string, error)
	SendCodeRecoveryService(recoveyRequest request.PasswordRecoveryRequest) error
	VerifyCode(VerifyCodeRequest request.VerifyCodeRequest) error
	ResetPassword(resetPasswordRequest request.ResetPasswordRequest) error
}

func (ud *userDomainService) LoginUserService(loginRequest request.LoginRequest) (model.UserInterface, string, error) {
	user, err := ud.userRepository.FindUserByIdentifierAndPassword(
		loginRequest.Identifier,
		loginRequest.Password,
		loginRequest.Role,
	)

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (ud *userDomainService) SendCodeRecoveryService(recoveyRequest request.PasswordRecoveryRequest) error {
	code, err := generateCode(6)
	if err != nil {
		return fmt.Errorf("erro ao gerar código: %w", err)
	}

	expiresAt := time.Now().Add(10 * time.Minute)

	userEmail, err := ud.userRepository.CreateRecoveryCode(recoveyRequest.Identifier, code, expiresAt)
	if err != nil {
		return err
	}

	err = ud.emailService.SendRecoveryEmail(userEmail, code)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %w", err)
	}

	return nil
}

func (ud *userDomainService) VerifyCode(verifyCodeRequest request.VerifyCodeRequest) error {
	valid, err := ud.userRepository.VerifyRecoveryCode(verifyCodeRequest.Identifier, verifyCodeRequest.Code)

	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("código inválido ou expirado")
	}

	return nil
}

func (ud *userDomainService) ResetPassword(resetPasswordRequest request.ResetPasswordRequest) error {
    if resetPasswordRequest.NewPassword != resetPasswordRequest.ConfirmPassword {
        return fmt.Errorf("as senhas não coincidem")
    }
	
    valid, err := ud.userRepository.VerifyRecoveryCode(resetPasswordRequest.Identifier, resetPasswordRequest.Code)
    if err != nil {
        return err
    }
    
    if !valid {
        return fmt.Errorf("código inválido ou expirado")
    }
	
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPasswordRequest.NewPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("erro ao gerar hash da senha: %w", err)
    }
	
    err = ud.userRepository.UpdatePassword(resetPasswordRequest.Identifier, string(hashedPassword))
    if err != nil {
        return err
    }
    
    return nil
}

func generateCode(length int) (string, error) {
	const digits = "0123456789"
	var code []byte

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code = append(code, digits[n.Int64()])
	}

	return string(code), nil
}
