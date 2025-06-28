package laudo

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/laudo"
	"github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type LaudoControllerInterface interface {
	CreateLaudo(c *gin.Context)
	GetLaudosByMedico(c *gin.Context)
	DownloadLaudo(c *gin.Context)
}

type laudoController struct {
	laudoService    laudo.LaudoServiceInterface
	pacienteService paciente.PacienteService
	medicoService   medico.MedicoServiceInterface
}

func NewLaudoController(
	laudoService laudo.LaudoServiceInterface,
	pacienteService paciente.PacienteService,
	medicoService medico.MedicoServiceInterface,
) LaudoControllerInterface {
	return &laudoController{
		laudoService:    laudoService,
		pacienteService: pacienteService,
		medicoService:   medicoService,
	}
}

func (lc *laudoController) CreateLaudo(c *gin.Context) {
	var laudo model.Laudo
	if err := c.ShouldBindJSON(&laudo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	laudo.MedicoID = medico.(model.BaseUser).ID
	laudo.Status = model.LaudoStatusConcluido
	laudo.DataLaudo = time.Now()

	ficha, err := lc.pacienteService.FindUltimaFichaByPacienteID(laudo.PacienteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	laudo.FichaID = ficha.ID

	if err := lc.laudoService.CreateLaudo(&laudo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Laudo criado com sucesso",
		"data":    laudo,
	})
}

func (lc *laudoController) GetLaudosByMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	resultadoFilter := c.Query("resultado")
	search := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	laudos, total, err := lc.laudoService.GetLaudosByMedicoID(
		medico.(model.BaseUser).ID,
		resultadoFilter,
		search,
		page,
		pageSize,
	)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": laudos,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  int(math.Ceil(float64(total) / float64(pageSize))),
			"hasNext":     page*pageSize < total,
			"hasPrev":     page > 1,
		},
	})
}

func (lc *laudoController) DownloadLaudo(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID do laudo é obrigatório"})
        return
    }

    laudo, err := lc.laudoService.GetLaudoByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    paciente, err := lc.pacienteService.GetPacienteByID(laudo.PacienteID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar paciente"})
        return
    }

    medico, err := lc.medicoService.FindMedicoByIdentifier(laudo.MedicoID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar médico"})
        return
    }

    // Configuração do PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    tr := pdf.UnicodeTranslatorFromDescriptor("")
    
    // Definição das cores (RGB)
    primaryR, primaryG, primaryB := 155, 89, 182   // #9b59b6
    secondaryR, secondaryG, secondaryB := 142, 68, 173 // #8e44ad
    lightPurpleR, lightPurpleG, lightPurpleB := 210, 180, 222
    textR, textG, textB := 51, 51, 51
    
    // Margens
    marginLeft := 15.0
    pdf.SetLeftMargin(marginLeft)
    pdf.SetRightMargin(marginLeft)
	
    // Título alinhado à direita do logo
    pdf.SetY(20)
    pdf.SetFont("Arial", "B", 18)
    pdf.SetTextColor(primaryR, primaryG, primaryB)
    pdf.CellFormat(0, 10, tr("LAUDO CITOPATOLÓGICO"), "", 1, "R", false, 0, "")
    
    // Informações da clínica
    pdf.SetFont("Arial", "I", 10)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.CellFormat(0, 5, tr("UBS - Citopatologia"), "", 1, "R", false, 0, "")
    pdf.CellFormat(0, 5, tr("CNES: XXXXXX"), "", 1, "R", false, 0, "")
    
    // Linha divisória
    pdf.SetDrawColor(secondaryR, secondaryG, secondaryB)
    pdf.Line(marginLeft, 45, 210-marginLeft, 45)
    pdf.Ln(12)

    // --- INFORMAÇÕES DO PACIENTE ---
    pdf.SetFont("Arial", "B", 12)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.SetFillColor(lightPurpleR, lightPurpleG, lightPurpleB)
    pdf.CellFormat(0, 8, tr(" DADOS DO PACIENTE "), "", 1, "L", true, 0, "")
    
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(textR, textG, textB)
    
    // Tabela de informações
    colWidth := 90.0
    lineHeight := 7.0
    
    // Linha 1
    createLabelValueRow(pdf, tr("Paciente:"), tr(paciente.Name), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    // Linha 2
    createLabelValueRow(pdf, tr("Idade:"), tr(fmt.Sprintf("%d anos", paciente.GetIdade())), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    // Linha 3
    createLabelValueRow(pdf, tr("Cartão SUS:"), tr(paciente.CNS), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    // Linha 4
    createLabelValueRow(pdf, tr("Data do Exame:"), tr(laudo.DataExame.Format("02/01/2006")), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    
    pdf.Ln(10)

    // --- RESULTADOS ---
    pdf.SetFont("Arial", "B", 12)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.SetFillColor(lightPurpleR, lightPurpleG, lightPurpleB)
    pdf.CellFormat(0, 8, tr(" RESULTADOS DO EXAME "), "", 1, "L", true, 0, "")
    
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(textR, textG, textB)
    
    // Tabela de resultados
    createLabelValueRow(pdf, tr("Adequabilidade:"), tr(laudo.Adequabilidade), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("Resultado:"), tr(string(laudo.Resultado)), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("CID:"), tr(laudo.CID), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("Microbiologia:"), tr(laudo.Microbiologia), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("Células Endometriais:"), tr(fmt.Sprintf("%v", laudo.CelulasEndometriais)), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    
    pdf.Ln(10)

    // --- COMENTÁRIOS ---
    if laudo.Comentarios != "" {
        pdf.SetFont("Arial", "B", 12)
        pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
        pdf.SetFillColor(lightPurpleR, lightPurpleG, lightPurpleB)
        pdf.CellFormat(0, 8, tr(" COMENTÁRIOS "), "", 1, "L", true, 0, "")
        
        pdf.SetFont("Arial", "", 11)
        pdf.SetTextColor(textR, textG, textB)
        pdf.MultiCell(0, 6, tr(laudo.Comentarios), "", "L", false)
        pdf.Ln(8)
    }

    // --- RECOMENDAÇÕES ---
    if laudo.Recomendacoes != "" {
        pdf.SetFont("Arial", "B", 12)
        pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
        pdf.SetFillColor(lightPurpleR, lightPurpleG, lightPurpleB)
        pdf.CellFormat(0, 8, tr(" RECOMENDAÇÕES "), "", 1, "L", true, 0, "")
        
        pdf.SetFont("Arial", "", 11)
        pdf.SetTextColor(textR, textG, textB)
        pdf.MultiCell(0, 6, tr(laudo.Recomendacoes), "", "L", false)
        pdf.Ln(8)
    }

    // --- RODAPÉ ---
    pdf.SetDrawColor(200, 200, 200)
    pdf.Line(marginLeft, pdf.GetY(), 210-marginLeft, pdf.GetY())
    pdf.Ln(5)
    
    // Data de emissão
    pdf.SetFont("Arial", "I", 9)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.CellFormat(90, 6, tr("Emitido em: ")+tr(time.Now().Format("02/01/2006 15:04")), "", 0, "L", false, 0, "")
    
    // Número do protocolo 
    pdf.CellFormat(90, 6, tr("Protocolo: ")+laudo.ID, "", 1, "R", false, 0, "")
    
    // Assinatura
    pdf.Ln(15)
    pdf.SetDrawColor(secondaryR, secondaryG, secondaryB)
    pdf.Line(120, pdf.GetY(), 180, pdf.GetY())
    pdf.SetFont("Arial", "I", 10)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.CellFormat(0, 5, tr("Dr.(a) ")+tr(medico.(*model.Medico).Name)+tr(" - CRM: ")+tr(medico.(*model.Medico).CRM), "", 1, "R", false, 0, "")
    
    // Contato da clínica
    pdf.Ln(5)
    pdf.SetFont("Arial", "I", 8)
    pdf.SetTextColor(100, 100, 100)
    pdf.CellFormat(0, 4, tr("Unidade Basica de Saude - contato@sobrevidas.com.br - (XX) XXXX-XXXX"), "", 1, "C", false, 0, "")

    var buf bytes.Buffer
    err = pdf.Output(&buf)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar PDF"})
        return
    }

    c.Header("Content-Type", "application/pdf")
    c.Header("Content-Disposition", "attachment; filename=laudo-"+laudo.ID+".pdf")
    c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

func createLabelValueRow(pdf *gofpdf.Fpdf, label, value string, colWidth, lineHeight float64, 
                        labelR, labelG, labelB int, valueR, valueG, valueB int) {
    pdf.SetFont("Arial", "B", 11)
    pdf.SetTextColor(labelR, labelG, labelB)
    pdf.CellFormat(colWidth, lineHeight, label, "", 0, "L", false, 0, "")
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(valueR, valueG, valueB)
    pdf.CellFormat(colWidth, lineHeight, value, "", 1, "L", false, 0, "")
}