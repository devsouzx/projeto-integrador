package agente

import (
	"github.com/devsouzx/projeto-integrador/src/repository/gestor"
)

func NewGestorService(
	gestorRepository gestor.AgenteRepository,
) GestorServiceInterface {
	return &gestorService{
		gestorRepository: gestorRepository,
	}
}

type gestorService struct {
	gestorRepository gestor.GestorRepository
}

type GestorServiceInterface interface {
	
}