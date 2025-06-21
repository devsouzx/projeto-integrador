package paciente

import (
	"math"
	"net/http"
	"strconv"
	"time"
	
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

func NewPacienteController(
	pacienteService paciente.PacienteService,
) PacienteControllerInterface {
	return &pacienteController{
		pacienteService: pacienteService,
	}
}

type pacienteController struct {
	pacienteService paciente.PacienteService
}

type PacienteControllerInterface interface {
	BuscarPacientePorCPF(c *gin.Context)
	ListarPacientes(c *gin.Context)
	GetPaciente(c *gin.Context)
}

func (pc *pacienteController) BuscarPacientePorCPF(c *gin.Context) {
	cpf := c.Query("cpf")
	if cpf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF não informado"})
		return
	}

	paciente, err := pc.pacienteService.GetPacienteByCPF(cpf)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	dataNascimento := ""
	if !paciente.NascimentoTime.IsZero() {
		dataNascimento = paciente.NascimentoTime.Format("2006-01-02")
	}

	response := gin.H{
		"Name":           paciente.Name,
		"Apelido":        paciente.Apelido,
		"NomeMae":        paciente.NomeMae,
		"CNS":            paciente.CNS,
		"CPF":            paciente.CPF,
		"DataNascimento": dataNascimento,
		"Nacionalidade":  paciente.Nacionalidade,
		"RacaCor":        paciente.RacaCor,
		"Escolaridade":   paciente.Escolaridade,
		"Telefone":       paciente.Telefone,
	}

	if paciente.Endereco != nil {
		response["Endereco"] = gin.H{
			"Logradouro":  paciente.Endereco.Logradouro,
			"Numero":      paciente.Endereco.Numero,
			"Complemento": paciente.Endereco.Complemento,
			"Bairro":      paciente.Endereco.Bairro,
			"Cidade":      paciente.Endereco.Cidade,
			"CEP":         paciente.Endereco.CEP,
			"UF":          paciente.Endereco.UF,
		}
	}
	
	c.JSON(http.StatusOK, response)
}

func (pc *pacienteController) ListarPacientes(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	pacientes, total, err := pc.pacienteService.ListarPacientes(page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao listar pacientes",
			"details": err.Error(),
		})
		return
	}

	var response []gin.H
	for _, p := range pacientes {
		pacienteData := gin.H{
			"id":             p.ID,
			"nome":           p.Name,
			"apelido":        p.Apelido,
			"mae":            p.NomeMae,
			"cns":            p.CNS,
			"cpf":            p.CPF,
			"nascimento":     p.NascimentoTime.Format("2006-01-02"),
			"idade":          calcularIdade(p.NascimentoTime),
			"telefone":       p.Telefone,
			"email":          p.Email,
			"raca":           p.RacaCor,
			"nacionalidade":  p.Nacionalidade,
			"escolaridade":   p.Escolaridade,
			"ultimoExameData": p.GetDataUltimoExame(),
			"ultimaInspecao": p.GetInspecaoColo(),
		}
		response = append(response, pacienteData)
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  totalPages,
			"hasNext":     page < totalPages,
			"hasPrev":     page > 1,
		},
	})
}

func (pc *pacienteController) GetPaciente(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do paciente é obrigatório"})
		return
	}

	paciente, err := pc.pacienteService.GetPacienteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paciente)
}

func calcularIdade(nascimento time.Time) int {
	now := time.Now()
	years := now.Year() - nascimento.Year()

	if now.Month() < nascimento.Month() ||
		(now.Month() == nascimento.Month() && now.Day() < nascimento.Day()) {
		years--
	}

	return years
}