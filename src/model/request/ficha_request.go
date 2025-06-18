package request

import "github.com/devsouzx/projeto-integrador/src/model"

type FichaRequest struct {
	Ficha             model.FichaCitopatologica `json:"ficha"`
	Paciente          model.Paciente              `json:"paciente"`
	DadosResidenciais model.Endereco              `json:"dados_residenciais"`
	Anamnese          model.Anamnese              `json:"anamnese"`
	ExameClinico      model.ExameClinico          `json:"exame_clinico"`
}