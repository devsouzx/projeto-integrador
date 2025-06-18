package paciente

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/devsouzx/projeto-integrador/src/repository/user"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
)

func NewPacienteController(
	pacienteService paciente.PacienteService, 
    userRepository user.UserRepository,
) PacienteControllerInterface {
	return &pacienteController{
		pacienteService: pacienteService,
        userRepository: userRepository,
	}
}

type pacienteController struct {
	pacienteService paciente.PacienteService
    userRepository user.UserRepository
}

type PacienteControllerInterface interface {
	BuscarPacientePorCPF(c *gin.Context)
    ListarPacientes(c *gin.Context)
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
            "error": "Erro ao listar pacientes",
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
            "mae":           p.NomeMae,
            "cns":           p.CNS,
            "cpf":           p.CPF,
            "nascimento":    p.NascimentoTime.Format("2006-01-02"),
            "idade":         calcularIdade(p.NascimentoTime),
            "telefone":      p.Telefone,
            "email":         p.Email,
            "raca":         p.RacaCor,
            "nacionalidade": p.Nacionalidade,
            "escolaridade":  p.Escolaridade,
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
            "hasNext":    page < totalPages,
            "hasPrev":    page > 1,
        },
    })
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