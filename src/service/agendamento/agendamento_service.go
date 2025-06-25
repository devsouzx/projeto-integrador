package agendamento

import (
	"fmt"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/agendamento"
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
	repo agendamento.AgendamentoRepositoryInterface
}

func NewAgendamentoService(repo agendamento.AgendamentoRepositoryInterface) AgendamentoServiceInterface {
	return &agendamentoService{repo: repo}
}

func (s *agendamentoService) CreateAgendamento(agendamento *model.Agendamento) error {
	if agendamento.PacienteID == "" {
		return fmt.Errorf("paciente é obrigatório")
	}
	if agendamento.DataHora.Before(time.Now()) {
		return fmt.Errorf("data/hora não pode ser no passado")
	}
	
	return s.repo.Create(agendamento)
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
	
	return s.repo.UpdateStatus(id, status)
}

func (s *agendamentoService) DeleteAgendamento(id string) error {
	return s.repo.Delete(id)
}