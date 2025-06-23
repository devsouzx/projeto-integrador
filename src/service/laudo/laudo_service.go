package laudo

import (
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
}

func (s *laudoService) CreateLaudo(laudo *model.Laudo) error {
	return s.repo.Create(laudo)
}

func (s *laudoService) GetLaudosByMedicoID(medicoID string, resultadoFilter string, search string, page int, pageSize int) ([]*model.Laudo, int, error) {
    return s.repo.FindByMedicoID(medicoID, resultadoFilter, search, page, pageSize)
}