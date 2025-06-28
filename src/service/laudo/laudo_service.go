package laudo

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/laudo"
)

type laudoService struct {
	repo laudo.LaudoRepositoryInterface
}

func NewLaudoService(repo laudo.LaudoRepositoryInterface) LaudoServiceInterface {
	return &laudoService{repo: repo}
}

type LaudoServiceInterface interface {
	CreateLaudo(laudo *model.Laudo) error
    GetLaudosByMedicoID(medicoID string, resultadoFilter string, search string, page int, pageSize int) ([]*model.Laudo, int, error)
	GetLaudoByID(id string) (*model.Laudo, error)
}

func (s *laudoService) CreateLaudo(laudo *model.Laudo) error {
	return s.repo.Create(laudo)
}

func (s *laudoService) GetLaudosByMedicoID(medicoID string, resultadoFilter string, search string, page int, pageSize int) ([]*model.Laudo, int, error) {
    return s.repo.FindByMedicoID(medicoID, resultadoFilter, search, page, pageSize)
}

func (s *laudoService) GetLaudoByID(id string) (*model.Laudo, error) {
	if id == "" {
		return nil, fmt.Errorf("ID do laudo é obrigatório")
	}

	laudo, err := s.repo.FindByIDWithRelations(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar laudo: %w", err)
	}

	return laudo, nil
}