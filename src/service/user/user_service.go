package user

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/model/request"

	"github.com/devsouzx/projeto-integrador/src/repository/paciente"
	userRepository "github.com/devsouzx/projeto-integrador/src/repository/user"
	emailService "github.com/devsouzx/projeto-integrador/src/service/email"
	"github.com/devsouzx/projeto-integrador/src/service/notifications"
	"golang.org/x/crypto/bcrypt"
)

func NewUserDomainService(
	userRepository userRepository.UserRepository,
	emailService emailService.EmailService,
	pacienteRepository paciente.PacienteRepository,
	smsService     *notifications.SMSService,
) UserDomainService {
	return &userDomainService{
		userRepository:     userRepository,
		emailService:       emailService,
		pacienteRepository: pacienteRepository,
		smsService:     smsService,
	}
}

type userDomainService struct {
	userRepository     userRepository.UserRepository
	pacienteRepository paciente.PacienteRepository
	emailService       emailService.EmailService
	smsService     *notifications.SMSService
}

type UserDomainService interface {
	LoginUserService(
		loginRequest request.LoginRequest,
	) (model.UserInterface, string, error)
	SendCodeRecoveryService(recoveyRequest request.PasswordRecoveryRequest) error
	VerifyCode(VerifyCodeRequest request.VerifyCodeRequest) error
	ResetPassword(resetPasswordRequest request.ResetPasswordRequest) error
	CadastrarUsuario(req request.CadastroRequest) (*model.Paciente, error)
	VerifyUserToken(token string) (model.UserInterface, error)
	UpdateUser(user model.UserInterface) error
	FindUserEmailAndPhone(identifier string) (string, string, error)
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

	userEmail, userPhone, err := ud.userRepository.CreateRecoveryCode(recoveyRequest.Identifier, code, expiresAt)
	if err != nil {
		return err
	}

	errChan := make(chan error, 2)
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        if userEmail != "" {
            if err := ud.emailService.SendRecoveryEmail(userEmail, code); err != nil {
                errChan <- fmt.Errorf("erro ao enviar e-mail: %w", err)
                return
            }
        }
        errChan <- nil
    }()
	
    wg.Add(1)
    go func() {
        defer wg.Done()
        if userPhone != "" {
            phone := "+55" + userPhone
			fmt.Println(phone)
            message := fmt.Sprintf("Seu código de recuperação é: %s. Expira em 10 minutos.", code)
            if err := ud.smsService.SendSMS(phone, message); err != nil {
                errChan <- fmt.Errorf("erro ao enviar SMS: %w", err)
                return
            }
        }
        errChan <- nil
    }()
	
    go func() {
        wg.Wait()
        close(errChan)
    }()
	
    var errors []error
    for err := range errChan {
        if err != nil {
            errors = append(errors, err)
        }
    }

    if len(errors) > 0 {
        return fmt.Errorf("erros ao enviar códigos: %v", errors)
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

func (s *userDomainService) CadastrarUsuario(req request.CadastroRequest) (*model.Paciente, error) {
	if req.Email == "" || req.Senha == "" || req.Cpf == "" || req.Cartaosus == "" {
		return nil, errors.New("todos os campos obrigatórios devem ser preenchidos")
	}

	exists, _ := s.userRepository.FindUserByIdentifier(req.Email, "paciente")
	if exists != nil {
		return nil, errors.New("e-mail já cadastrado")
	}
	exists, _ = s.userRepository.FindUserByIdentifier(req.Cpf, "paciente")
	if exists != nil {
		return nil, errors.New("cpf já cadastrado")
	}

	usuario := &model.Paciente{
		BaseUser: model.BaseUser{
			Email:    req.Email,
			Password: req.Senha,
			Name:     req.NomeCompleto,
		},
		CPF:            req.Cpf,
		CNS:            req.Cartaosus,
		NomeMae:        req.NomeCompletoDaMae,
		DataNascimento: req.DataDeNascimento,
		Telefone:       req.Telefone,
	}

	if req.Endereco.Cep != "" || req.Endereco.Logradouro != "" || req.Endereco.Numero != "" {
		usuario.Endereco = &model.Endereco{
			CEP:         req.Endereco.Cep,
			Logradouro:  req.Endereco.Logradouro,
			Numero:      req.Endereco.Numero,
			Complemento: req.Endereco.Complemento,
			Bairro:      req.Endereco.Bairro,
			Cidade:      req.Endereco.Cidade,
			UF:          req.Endereco.Uf,
		}
	}

	usuario, err := s.pacienteRepository.CreatePaciente(usuario)
	if err != nil {
		return nil, fmt.Errorf("erro ao cadastrar paciente: %w", err)
	}

	return usuario, nil
}

func (us *userDomainService) VerifyUserToken(token string) (model.UserInterface, error) {
	user, err := us.userRepository.FindByVerificationToken(token)
	if err != nil {
		return nil, errors.New("token inválido")
	}

	if user.GetTokenExpiresAt().Before(time.Now().UTC()) {
		return nil, fmt.Errorf("token expirado")
	}

	return user, nil
}

func (us *userDomainService) UpdateUser(user model.UserInterface) error {
	if user.GetID() == "" {
		return errors.New("ID do usuário é obrigatório")
	}

	userExistente, err := us.userRepository.FindByID(user.GetID(), "paciente")
	if err != nil {
		return fmt.Errorf("usuário não encontrado: %w", err)
	}

	userExistente.SetVerified(user.IsVerified())
	userExistente.SetVerifyToken(user.GetVerifyToken())
	userExistente.SetTokenExpiresAt(user.GetTokenExpiresAt())

	return us.userRepository.UpdateUser(userExistente)
}

func (us *userDomainService) FindUserEmailAndPhone(identifier string) (string, string, error) {
	userEmail, userPhone, err := us.userRepository.FindUserEmailAndPhone(identifier)
	if err != nil {
		return "", "", fmt.Errorf("erro ao buscar e-mail e telefone do usuário: %w", err)
	}

	if userEmail == "" && userPhone == "" {
		return "", "", fmt.Errorf("usuário não encontrado")
	}

	return userEmail, userPhone, nil
}