package encaminhamento

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/devsouzx/projeto-integrador/src/model"
	"github.com/devsouzx/projeto-integrador/src/service/encaminhamento"
	"github.com/devsouzx/projeto-integrador/src/service/medico"
	"github.com/devsouzx/projeto-integrador/src/service/paciente"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type EncaminhamentoControllerInterface interface {
	CreateEncaminhamento(c *gin.Context)
	GetEncaminhamento(c *gin.Context)
	GetEncaminhamentosByPaciente(c *gin.Context)
	GetEncaminhamentosByMedico(c *gin.Context)
	UpdateStatus(c *gin.Context)
	DownloadEncaminhamento(c *gin.Context)
}

type encaminhamentoController struct {
	encaminhamentoService encaminhamento.EncaminhamentoServiceInterface
	pacienteService       paciente.PacienteService
	medicoService         medico.MedicoServiceInterface
}

func NewEncaminhamentoController(
	encaminhamentoService encaminhamento.EncaminhamentoServiceInterface,
	pacienteService paciente.PacienteService,
	medicoService medico.MedicoServiceInterface,
) EncaminhamentoControllerInterface {
	return &encaminhamentoController{
		encaminhamentoService: encaminhamentoService,
		pacienteService:       pacienteService,
		medicoService:         medicoService,
	}
}

func (ec *encaminhamentoController) CreateEncaminhamento(c *gin.Context) {
	var encaminhamento model.Encaminhamento
	if err := c.ShouldBindJSON(&encaminhamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	encaminhamento.MedicoID = medico.(model.BaseUser).ID

	if err := ec.encaminhamentoService.CreateEncaminhamento(&encaminhamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Encaminhamento criado com sucesso",
		"data":    encaminhamento,
	})
}

func (ec *encaminhamentoController) GetEncaminhamento(c *gin.Context) {
	id := c.Param("id")

	encaminhamento, err := ec.encaminhamentoService.GetEncaminhamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encaminhamento não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": encaminhamento,
	})
}

func (ec *encaminhamentoController) GetEncaminhamentosByPaciente(c *gin.Context) {
	pacienteID := c.Param("pacienteId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	encaminhamentos, total, err := ec.encaminhamentoService.GetEncaminhamentosByPacienteID(pacienteID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": encaminhamentos,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  (total + pageSize - 1) / pageSize,
			"hasNext":     page*pageSize < total,
			"hasPrev":     page > 1,
		},
	})
}

func (ec *encaminhamentoController) GetEncaminhamentosByMedico(c *gin.Context) {
	medico, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	encaminhamentos, total, err := ec.encaminhamentoService.GetEncaminhamentosByMedicoID(medico.(model.BaseUser).ID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": encaminhamentos,
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalItems":  total,
			"totalPages":  (total + pageSize - 1) / pageSize,
			"hasNext":     page*pageSize < total,
			"hasPrev":     page > 1,
		},
	})
}

func (ec *encaminhamentoController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var request struct {
		Status      model.EncaminhamentoStatus `json:"status"`
		Observacoes string                     `json:"observacoes"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.encaminhamentoService.UpdateStatus(id, request.Status, request.Observacoes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status do encaminhamento atualizado com sucesso",
	})
}

func (ec *encaminhamentoController) DownloadEncaminhamento(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID do encaminhamento é obrigatório"})
        return
    }

    encaminhamento, err := ec.encaminhamentoService.GetEncaminhamentoByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    // Configuração do PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    tr := pdf.UnicodeTranslatorFromDescriptor("")
    
    // Definição das cores (RGB)
    primaryR, primaryG, primaryB := 155, 89, 182   // #9b59b6
    secondaryR, secondaryG, secondaryB := 142, 68, 173 // #8e44ad
    lightBlueR, lightBlueG, lightBlueB := 210, 180, 222
    textR, textG, textB := 51, 51, 51

    // Margens
    marginLeft := 15.0
    pdf.SetLeftMargin(marginLeft)
    pdf.SetRightMargin(marginLeft)
    
    // Título alinhado à direita do logo
    pdf.SetY(20)
    pdf.SetFont("Arial", "B", 18)
    pdf.SetTextColor(primaryR, primaryG, primaryB)
    pdf.CellFormat(0, 10, tr("ENCAMINHAMENTO MÉDICO"), "", 1, "R", false, 0, "")
    
    // Informações da unidade
    pdf.SetFont("Arial", "I", 10)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.CellFormat(0, 5, tr("Unidade Básica de Saúde"), "", 1, "R", false, 0, "")
    pdf.CellFormat(0, 5, tr("CNES: XXXXXX"), "", 1, "R", false, 0, "")
    
    // Linha divisória
    pdf.SetDrawColor(secondaryR, secondaryG, secondaryB)
    pdf.Line(marginLeft, 45, 210-marginLeft, 45)
    pdf.Ln(12)

    // --- INFORMAÇÕES DO PACIENTE ---
    pdf.SetFont("Arial", "B", 12)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.SetFillColor(lightBlueR, lightBlueG, lightBlueB)
    pdf.CellFormat(0, 8, tr(" DADOS DO PACIENTE "), "", 1, "L", true, 0, "")
    
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(textR, textG, textB)
    
    // Tabela de informações
    colWidth := 90.0
    lineHeight := 7.0
    
    // Linha 1
    createLabelValueRow(pdf, tr("Paciente:"), tr(encaminhamento.Paciente.Name), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    // Linha 2
    createLabelValueRow(pdf, tr("Cartão SUS:"), tr(encaminhamento.Paciente.CNS), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    // Linha 3
    createLabelValueRow(pdf, tr("Data do Encaminhamento:"), tr(encaminhamento.DataEncaminhamento.Format("02/01/2006")), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    
    pdf.Ln(10)

    // --- DADOS DO ENCAMINHAMENTO ---
    pdf.SetFont("Arial", "B", 12)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.SetFillColor(lightBlueR, lightBlueG, lightBlueB)
    pdf.CellFormat(0, 8, tr(" DADOS DO ENCAMINHAMENTO "), "", 1, "L", true, 0, "")
    
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(textR, textG, textB)
    
    // Tabela de dados
    createLabelValueRow(pdf, tr("Especialidade:"), tr(encaminhamento.Especialidade), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("Urgência:"), tr(string(encaminhamento.Urgencia)), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("Unidade de Referência:"), tr(encaminhamento.UnidadeReferencia), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    createLabelValueRow(pdf, tr("Status:"), tr(string(encaminhamento.Status)), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    
    if encaminhamento.DataAgendamento != nil {
        createLabelValueRow(pdf, tr("Data Agendada:"), tr(encaminhamento.DataAgendamento.Format("02/01/2006")), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
    }
    
    pdf.Ln(10)

    // --- JUSTIFICATIVA ---
    pdf.SetFont("Arial", "B", 12)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.SetFillColor(lightBlueR, lightBlueG, lightBlueB)
    pdf.CellFormat(0, 8, tr(" JUSTIFICATIVA "), "", 1, "L", true, 0, "")
    
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(textR, textG, textB)
    pdf.MultiCell(0, 6, tr(encaminhamento.Justificativa), "", "L", false)
    pdf.Ln(8)

    // --- OBSERVAÇÕES ---
    if encaminhamento.Observacoes != nil && *encaminhamento.Observacoes != "" {
        pdf.SetFont("Arial", "B", 12)
        pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
        pdf.SetFillColor(lightBlueR, lightBlueG, lightBlueB)
        pdf.CellFormat(0, 8, tr(" OBSERVAÇÕES "), "", 1, "L", true, 0, "")
        
        pdf.SetFont("Arial", "", 11)
        pdf.SetTextColor(textR, textG, textB)
        pdf.MultiCell(0, 6, tr(*encaminhamento.Observacoes), "", "L", false)
        pdf.Ln(8)
    }

    // --- LAUDO ASSOCIADO ---
    if encaminhamento.Laudo != nil {
        pdf.SetFont("Arial", "B", 12)
        pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
        pdf.SetFillColor(lightBlueR, lightBlueG, lightBlueB)
        pdf.CellFormat(0, 8, tr(" LAUDO ASSOCIADO "), "", 1, "L", true, 0, "")
        
        pdf.SetFont("Arial", "", 11)
        pdf.SetTextColor(textR, textG, textB)
        
        createLabelValueRow(pdf, tr("Resultado:"), tr(string(encaminhamento.Laudo.Resultado)), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
        createLabelValueRow(pdf, tr("CID:"), tr(encaminhamento.Laudo.CID), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
        createLabelValueRow(pdf, tr("Data do Exame:"), tr(encaminhamento.Laudo.DataExame.Format("02/01/2006")), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
        createLabelValueRow(pdf, tr("Data do Laudo:"), tr(encaminhamento.Laudo.DataLaudo.Format("02/01/2006")), colWidth, lineHeight, primaryR, primaryG, primaryB, textR, textG, textB)
        
        pdf.Ln(10)
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
    pdf.CellFormat(90, 6, tr("Protocolo: ")+encaminhamento.ID, "", 1, "R", false, 0, "")
    
    // Assinatura
    pdf.Ln(15)
    pdf.SetDrawColor(secondaryR, secondaryG, secondaryB)
    pdf.Line(120, pdf.GetY(), 180, pdf.GetY())
    pdf.SetFont("Arial", "I", 10)
    pdf.SetTextColor(secondaryR, secondaryG, secondaryB)
    pdf.CellFormat(0, 5, tr("Dr.(a) ")+tr(encaminhamento.Medico.Name)+tr(" - CRM: ")+tr(encaminhamento.Medico.CRM), "", 1, "R", false, 0, "")
    
    // Contato da unidade
    pdf.Ln(5)
    pdf.SetFont("Arial", "I", 8)
    pdf.SetTextColor(100, 100, 100)
    pdf.CellFormat(0, 4, tr("Unidade Básica de Saúde - contato@saude.gov.br - (XX) XXXX-XXXX"), "", 1, "C", false, 0, "")

    var buf bytes.Buffer
    err = pdf.Output(&buf)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar PDF"})
        return
    }

    c.Header("Content-Type", "application/pdf")
    c.Header("Content-Disposition", "attachment; filename=encaminhamento-"+encaminhamento.ID+".pdf")
    c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

// Função auxiliar para criar linhas de label e valor
func createLabelValueRow(pdf *gofpdf.Fpdf, label, value string, colWidth, lineHeight float64, 
                        labelR, labelG, labelB int, valueR, valueG, valueB int) {
    pdf.SetFont("Arial", "B", 11)
    pdf.SetTextColor(labelR, labelG, labelB)
    pdf.CellFormat(colWidth, lineHeight, label, "", 0, "L", false, 0, "")
    pdf.SetFont("Arial", "", 11)
    pdf.SetTextColor(valueR, valueG, valueB)
    pdf.CellFormat(colWidth, lineHeight, value, "", 1, "L", false, 0, "")
}