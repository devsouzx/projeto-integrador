package enfermeiro

import (
	"fmt"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/repository/enfermeiro"
)

func NewEnfermeiroService(
	enfermeiroRepository enfermeiro.EnfermeiroRepositoryInterface,
) EnfermeiroServiceInterface {
	return &enfermeitoService{
		repository: enfermeiroRepository,
	}
}

type enfermeitoService struct {
	repository enfermeiro.EnfermeiroRepositoryInterface
}

type EnfermeiroServiceInterface interface {
	FindEnfermeiroByIdentifier(identifier string) (*model.Enfermeiro, error)
}

func (es *enfermeitoService) FindEnfermeiroByIdentifier(identifier string) (*model.Enfermeiro, error) {
	enfermeiro, err := es.repository.FindEnfermeiroByIdentifier(identifier)
	if err != nil {
		return nil, fmt.Errorf("erro ao busccar o enfermeiro")
	}

	return enfermeiro, nil
}