package model

type CadastroMedicoRequest struct {
	ID             uint   `json:"id" binding:"required"`
	Name           string `json:"nomecompleto" binding:"required"`
	CRM            string `json:"crm" binding:"required"`
	Especialidade  string `json:"especialidade" binding:"required"`
	CPF            string `json:"cpf" binding:"required"`
	DataNascimento string `json:"datadenascimento" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Telefone       string `json:"telefone" binding:"required"`
}
