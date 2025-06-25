package encaminhamento

import (
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/encaminhamento"
)

type EncaminhamentoServiceInterface interface {
	CreateEncaminhamento(encaminhamento *model.Encaminhamento) error
	GetEncaminhamentoByID(id string) (*model.Encaminhamento, error)
	GetEncaminhamentosByPacienteID(pacienteID string, page, pageSize int) ([]*model.Encaminhamento, int, error)
	GetEncaminhamentosByMedicoID(medicoID string, page, pageSize int) ([]*model.Encaminhamento, int, error)
	UpdateStatus(id string, status model.EncaminhamentoStatus, observacoes string) error
	UpdateAgendamento(id string, dataAgendamento time.Time) error
}

type encaminhamentoService struct {
	repo encaminhamento.EncaminhamentoRepositoryInterface
}

func NewEncaminhamentoService(repo encaminhamento.EncaminhamentoRepositoryInterface) EncaminhamentoServiceInterface {
	return &encaminhamentoService{repo: repo}
}

func (s *encaminhamentoService) CreateEncaminhamento(encaminhamento *model.Encaminhamento) error {
	encaminhamento.Status = model.EncaminhamentoStatusPendente
	return s.repo.Create(encaminhamento)
}

func (s *encaminhamentoService) GetEncaminhamentoByID(id string) (*model.Encaminhamento, error) {
	return s.repo.FindByID(id)
}

func (s *encaminhamentoService) GetEncaminhamentosByPacienteID(pacienteID string, page, pageSize int) ([]*model.Encaminhamento, int, error) {
	return s.repo.FindByPacienteID(pacienteID, page, pageSize)
}

func (s *encaminhamentoService) GetEncaminhamentosByMedicoID(medicoID string, page, pageSize int) ([]*model.Encaminhamento, int, error) {
	return s.repo.FindByMedicoID(medicoID, page, pageSize)
}

func (s *encaminhamentoService) UpdateStatus(id string, status model.EncaminhamentoStatus, observacoes string) error {
	return s.repo.UpdateStatus(id, status, observacoes)
}

func (s *encaminhamentoService) UpdateAgendamento(id string, dataAgendamento time.Time) error {
	return s.repo.UpdateAgendamento(id, dataAgendamento)
}