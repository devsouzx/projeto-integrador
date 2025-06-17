package datasus

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
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
	Total            int            `json:"total"`
}

func (s *CNESService) BuscarUnidadesPorMunicipio(codigoIBGE string, limit, offset int) ([]UnidadeSaude, int, error) {
    url := fmt.Sprintf("%sestabelecimentos?codigo_municipio=%s&codigo_tipo_unidade=2&limit=%d&offset=%d", 
        s.BaseURL, codigoIBGE, limit, offset)
    
    resp, err := s.Client.Get(url)
    if err != nil {
        return nil, 0, fmt.Errorf("erro na requisição: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, 0, fmt.Errorf("erro ao ler resposta: %v", err)
    }

    var apiResponse ResponseAPI
    if err := json.Unmarshal(body, &apiResponse); err != nil {
        return nil, 0, fmt.Errorf("erro ao decodificar JSON: %v", err)
    }

    return apiResponse.Estabelecimentos, apiResponse.Total, nil
}

func (s *CNESService) BuscarUnidadePorCNES(codigoCNES int) (*UnidadeSaude, error) {
    url := fmt.Sprintf("%sestabelecimentos/%d", s.BaseURL, codigoCNES)
    
    resp, err := s.Client.Get(url)
    if err != nil {
        return nil, fmt.Errorf("erro na requisição: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return nil, fmt.Errorf("unidade com CNES %d não encontrada", codigoCNES)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("erro ao ler resposta: %v", err)
    }

    var unidade UnidadeSaude
    if err := json.Unmarshal(body, &unidade); err != nil {
        return nil, fmt.Errorf("erro ao decodificar JSON: %v", err)
    }

    return &unidade, nil
}

func (s *CNESService) BuscarUnidadeMaisProxima(codigoIBGE string, lat, long float64) (*UnidadeSaude, error) {
    codigoMunicipio, err := strconv.Atoi(codigoIBGE)
    if err != nil {
        return nil, fmt.Errorf("código IBGE inválido: %v", err)
    }

    const pageSize = 100
    var unidadesValidas []UnidadeSaude
	
    for offset := 0; ; offset += pageSize {
        unidades, total, err := s.BuscarUnidadesPorMunicipio(codigoIBGE, pageSize, offset)
        if err != nil {
            return nil, fmt.Errorf("erro ao buscar unidades: %v", err)
        }
		
        for _, u := range unidades {
            if u.CodigoMunicipio == codigoMunicipio && u.Latitude != 0 && u.Longitude != 0 {
                unidadesValidas = append(unidadesValidas, u)
            }
        }
        
        if offset+pageSize >= total {
            break
        }
    }

    if len(unidadesValidas) == 0 {
        return nil, fmt.Errorf("nenhuma unidade válida encontrada para o município %s", codigoIBGE)
    }
	
    var maisProxima *UnidadeSaude
    menorDistancia := math.MaxFloat64

    for i, u := range unidadesValidas {
        dist := calcularDistancia(lat, long, u.Latitude, u.Longitude)
        if dist < menorDistancia {
            menorDistancia = dist
            maisProxima = &unidadesValidas[i]
        }
    }

    return maisProxima, nil
}

func calcularDistancia(lat1, lon1, lat2, lon2 float64) float64 {
    const R = 6371.0
	
    dLat := (lat2 - lat1) * math.Pi / 180
    dLon := (lon2 - lon1) * math.Pi / 180
	
    a := math.Sin(dLat/2)*math.Sin(dLat/2) +
        math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
            math.Sin(dLon/2)*math.Sin(dLon/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

    return R * c
}