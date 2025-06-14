package model

import (
	"time"
)

type FichaCitopatologica struct {
	ID                string    `json:"id" db:"id"`
	Protocolo         string    `json:"protocolo" db:"protocolo"`
	CNES              string    `json:"cnes" db:"cnes"`
	UnidadeSaude      string    `json:"unidade_saude" db:"unidade_saude"`
	Municipio         string    `json:"municipio" db:"municipio"`
	UF                string    `json:"uf" db:"uf"`
	DataColeta        time.Time `json:"data_coleta" db:"data_coleta"`
	ResponsavelColeta string    `json:"responsavel_coleta" db:"responsavel_coleta"`
	MotivoExame       string    `json:"motivo_exame" db:"motivo_exame"`
	Observacoes       string    `json:"observacoes" db:"observacoes"`
	PacienteID        string    `json:"paciente_id" db:"paciente_id"`
	ResponsavelID          string    `json:"responsavel_id" db:"responsavel_id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type Anamnese struct {
	ID                   string `json:"id" db:"id"`
	MotivoExame          string `json:"motivo_exame" db:"motivo_exame"`
	FezExame             string `json:"fez_exame" db:"fez_exame"`
	UltimoExameAno       string `json:"ultimo_exame_ano" db:"ultimo_exame_ano"`
	UsaDIU               string `json:"usa_diu" db:"usa_diu"`
	Graviada             string `json:"gravida" db:"gravida"`
	Anticoncepcional     string `json:"anticoncepcional" db:"anticoncepcional"`
	HormonioMenopausa    string `json:"hormonio_menopausa" db:"hormonio_menopausa"`
	Radioterapia         string `json:"radioterapia" db:"radioterapia"`
	DUM                  string `json:"dum" db:"dum"`
	NaoLembraDUM         bool   `json:"nao_lembra_dum" db:"nao_lembra_dum"`
	SangramentoPosCoito  string `json:"sangramento_pos_coito" db:"sangramento_pos_coito"`
	SangramentoPosMenopausa string `json:"sangramento_pos_menopausa" db:"sangramento_pos_menopausa"`
	PacienteID           string `json:"paciente_id" db:"paciente_id"`
	FichaID              string `json:"ficha_id" db:"ficha_id"`
}

type Endereco struct {
	ID               string `json:"id" db:"id"`
	Logradouro       string `json:"logradouro" db:"logradouro"`
	Numero           string `json:"numero" db:"numero"`
	Complemento      string `json:"complemento" db:"complemento"`
	Bairro           string `json:"bairro" db:"bairro"`
	CodigoMunicipio  string `json:"codigo_municipio" db:"codigo_municipio"`
	Municipio        string `json:"municipio" db:"municipio"`
	UF        string `json:"uf" db:"uf"`
	CEP              string `json:"cep" db:"cep"`
	Referencia       string `json:"referencia" db:"referencia"`
	Telefone         string `json:"telefone" db:"telefone"`
	PacienteID       string `json:"paciente_id" db:"paciente_id"`
}

type FichaRequest struct {
    Ficha            FichaCitopatologica `json:"ficha"`
    Paciente         Paciente            `json:"paciente"`
    DadosResidenciais Endereco           `json:"dados_residenciais"`
    Anamnese         Anamnese            `json:"anamnese"`
    ExameClinico     ExameClinico       `json:"exame_clinico"`
}