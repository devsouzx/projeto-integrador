package datasus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CNESService struct {
	BaseURL string
	Client  *http.Client
}

func NewCNESService() *CNESService {
	return &CNESService{
		BaseURL: "https://apidadosabertos.saude.gov.br/cnes/",
		Client:  &http.Client{Timeout: 30 * time.Second},
	}
}

type UnidadeSaude struct {
	CodigoCNES                    int     `json:"codigo_cnes"`
	CNPJ                          string  `json:"numero_cnpj"`
	NomeRazaoSocial               string  `json:"nome_razao_social"`
	NomeFantasia                  string  `json:"nome_fantasia"`
	TipoGestao                    string  `json:"tipo_gestao"`
	CodigoTipoUnidade             int     `json:"codigo_tipo_unidade"`
	CEP                           string  `json:"codigo_cep_estabelecimento"`
	Endereco                      string  `json:"endereco_estabelecimento"`
	Numero                        string  `json:"numero_estabelecimento"`
	Bairro                        string  `json:"bairro_estabelecimento"`
	Telefone                      string  `json:"numero_telefone_estabelecimento"`
	Latitude                      float64 `json:"latitude_estabelecimento_decimo_grau"`
	Longitude                     float64 `json:"longitude_estabelecimento_decimo_grau"`
	Email                         string  `json:"endereco_email_estabelecimento"`
	TurnoAtendimento              string  `json:"descricao_turno_atendimento"`
	AtendeSUS                     string  `json:"estabelecimento_faz_atendimento_ambulatorial_sus"`
	CodigoEstabelecimentoSaude    string  `json:"codigo_estabelecimento_saude"`
	CodigoUF                      int     `json:"codigo_uf"`
	CodigoMunicipio               int     `json:"codigo_municipio"`
	NaturezaJuridica              string  `json:"descricao_natureza_juridica_estabelecimento"`
	PossuiServicoApoio            int     `json:"estabelecimento_possui_servico_apoio"`
	PossuiAtendimentoAmbulatorial int     `json:"estabelecimento_possui_atendimento_ambulatorial"`
	DataAtualizacao               string  `json:"data_atualizacao"`
}

type Municipio struct {
	CodigoIBGE string `json:"codigo_ibge"`
	Nome       string `json:"nome"`
}

type ResponseAPI struct {
	Estabelecimentos []UnidadeSaude `json:"estabelecimentos"`
}

func (s *CNESService) BuscarUnidadesPorMunicipio(codigoIBGE string) ([]UnidadeSaude, error) {
    url := fmt.Sprintf("%sestabelecimentos?codigo_municipio=%s&status=2&limit=100000&offset=0", s.BaseURL, codigoIBGE)
    
    resp, err := s.Client.Get(url)
    if err != nil {
        return nil, fmt.Errorf("erro na requisição: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("erro ao ler resposta: %v", err)
    }

    // Debug: logue a resposta bruta
    fmt.Printf("Resposta da API: %s\n", string(body))

    var apiResponse ResponseAPI
    if err := json.Unmarshal(body, &apiResponse); err != nil {
        return nil, fmt.Errorf("erro ao decodificar JSON: %v", err)
    }

    return apiResponse.Estabelecimentos, nil
}