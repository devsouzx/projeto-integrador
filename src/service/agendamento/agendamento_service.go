package agendamento

import (
	"fmt"
	"sync"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/agendamento"
	"github.com/devsouzx/projeto-integrador/src/service/email"
	"github.com/devsouzx/projeto-integrador/src/service/notifications"
	"github.com/devsouzx/projeto-integrador/src/service/user"
)

type AgendamentoServiceInterface interface {
	CreateAgendamento(agendamento *model.Agendamento) error
	GetAgendamentoByID(id string) (*model.Agendamento, error)
	ListAgendamentosByProfissional(profissionalID, profissionalTipo string) ([]*model.Agendamento, error)
	ListAgendamentosByDate(profissionalID, profissionalTipo string, date time.Time) ([]*model.Agendamento, error)
	UpdateAgendamentoStatus(id, status string) error
	DeleteAgendamento(id string) error
}

type agendamentoService struct {
	repo         agendamento.AgendamentoRepositoryInterface
	userService  user.UserDomainService
	smsService   *notifications.SMSService
	emailService email.EmailService
}

func NewAgendamentoService(repo agendamento.AgendamentoRepositoryInterface, userService user.UserDomainService, smsService *notifications.SMSService, emailService email.EmailService) AgendamentoServiceInterface {
	return &agendamentoService{repo: repo, userService: userService, smsService: smsService, emailService: emailService}
}

func (s *agendamentoService) CreateAgendamento(agendamento *model.Agendamento) error {
	if agendamento.PacienteID == "" {
		return fmt.Errorf("paciente é obrigatório")
	}
	if agendamento.DataHora.Before(time.Now()) {
		return fmt.Errorf("data/hora não pode ser no passado")
	}

	err := s.repo.Create(agendamento)
	if err != nil {
		return fmt.Errorf("erro ao criar agendamento: %w", err)
	}

	return s.notifyAgendamentoCreated(agendamento)
}

func (s *agendamentoService) notifyAgendamentoCreated(agendamento *model.Agendamento) error {
	userEmail, userPhone, err := s.userService.FindUserEmailAndPhone(agendamento.PacienteID)
	if err != nil {
		return fmt.Errorf("erro ao buscar contato do paciente: %w", err)
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	dataHora := agendamento.DataHora.In(loc).Format("02/01/2006 às 15:04")

	errChan := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if userEmail != "" {
			if err := s.emailService.SendAgendamentoCreatedEmail(userEmail, agendamento); err != nil {
				errChan <- fmt.Errorf("erro ao enviar e-mail: %w", err)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if userPhone != "" {
			message := fmt.Sprintf(
				"Agendamento %s para %s criado com sucesso! Status: %s. Aguarde confirmação.",
				agendamento.TipoConsulta,
				dataHora,
				agendamento.Status,
			)
			if err := s.smsService.SendSMS("+55"+userPhone, message); err != nil {
				errChan <- fmt.Errorf("erro ao enviar SMS: %w", err)
			}
		}
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
		return fmt.Errorf("erros na notificação: %v", errors)
	}

	return nil
}

func (s *agendamentoService) GetAgendamentoByID(id string) (*model.Agendamento, error) {
	return s.repo.FindByID(id)
}

func (s *agendamentoService) ListAgendamentosByProfissional(profissionalID, profissionalTipo string) ([]*model.Agendamento, error) {
	return s.repo.FindByProfissional(profissionalID, profissionalTipo)
}

func (s *agendamentoService) ListAgendamentosByDate(profissionalID, profissionalTipo string, date time.Time) ([]*model.Agendamento, error) {
	return s.repo.FindByProfissionalAndDate(profissionalID, profissionalTipo, date)
}

func (s *agendamentoService) UpdateAgendamentoStatus(id, status string) error {
	validStatus := map[string]bool{
		"pendente":   true,
		"confirmado": true,
		"cancelado":  true,
		"concluido":  true,
	}

	if !validStatus[status] {
		return fmt.Errorf("status inválido")
	}

	err := s.repo.UpdateStatus(id, status)
	if err != nil {
		return fmt.Errorf("erro ao atualizar status: %w", err)
	}

	agendamento, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar agendamento: %w", err)
	}

	return s.notifyStatusUpdate(agendamento, status)
}

func (s *agendamentoService) notifyStatusUpdate(agendamento *model.Agendamento, status string) error {
	userEmail, userPhone, err := s.userService.FindUserEmailAndPhone(agendamento.PacienteID)
	if err != nil {
		return fmt.Errorf("erro ao buscar contato do paciente: %w", err)
	}

	loc, _ := time.LoadLocation("America/Sao_Paulo")
	dataHora := agendamento.DataHora.In(loc).Format("02/01/2006 às 15:04")

	errChan := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if userEmail != "" {
			var emailErr error
			switch status {
			case "confirmado":
				emailErr = s.emailService.SendAgendamentoConfirmedEmail(userEmail, agendamento)
			case "cancelado":
				emailErr = s.emailService.SendAgendamentoCanceledEmail(userEmail, agendamento)
			case "concluido":
				emailErr = s.emailService.SendAgendamentoCompletedEmail(userEmail, agendamento)
			}
			if emailErr != nil {
				errChan <- fmt.Errorf("erro ao enviar e-mail: %w", emailErr)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if userPhone != "" {
			var message string
			switch status {
			case "confirmado":
				message = fmt.Sprintf(
					"Agendamento confirmado! %s em %s. Chegue 15min antes. Qualquer dúvida entre em contato.",
					agendamento.TipoConsulta,
					dataHora,
				)
			case "cancelado":
				message = fmt.Sprintf(
					"Agendamento %s em %s foi cancelado. Entre em contato para reagendar.",
					agendamento.TipoConsulta,
					dataHora,
				)
			case "concluido":
				message = fmt.Sprintf(
					"Consulta %s concluída!",
					agendamento.TipoConsulta,
				)
			}

			if message != "" {
				if err := s.smsService.SendSMS("+55"+userPhone, message); err != nil {
					errChan <- fmt.Errorf("erro ao enviar SMS: %w", err)
				}
			}
		}
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
		return fmt.Errorf("erros na notificação: %v", errors)
	}

	return nil
}

func (s *agendamentoService) DeleteAgendamento(id string) error {
	agendamento, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("erro ao buscar agendamento: %w", err)
	}

	err = s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("erro ao deletar agendamento: %w", err)
	}

	return s.notifyStatusUpdate(agendamento, "cancelado")
}
