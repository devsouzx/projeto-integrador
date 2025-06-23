package laudo

import (
	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/laudo"
)

type LaudoServiceInterface interface {
	GetLaudosByMedicoID(medicoID string) ([]*model.Laudo, error)
}

type laudoService struct {
	repo laudo.LaudoRepositoryInterface
}

func NewLaudoService(repo laudo.LaudoRepositoryInterface) LaudoServiceInterface {
	return &laudoService{repo: repo}
}

func (s *laudoService) GetLaudosByMedicoID(medicoID string) ([]*model.Laudo, error) {
	return s.repo.FindByMedicoID(medicoID)
}