package request

type CadastroProfissionalRequest struct {
	Name     string `json:"nomecompleto" binding:"required"`
	Password string `form:"password" json:"password"`
	CPF      string `json:"cpf" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Telefone string `json:"telefone" binding:"required"`
	Role     string `json:"role" binding:"required"`

	// Dados para medico
	CRM           string `json:"crm" binding:"omitempty"`
	Especialidade string `json:"especialidade" binding:"omitempty"`

	// Dados para enfermeira
	COREN string `json:"coren" binding:"omitempty"`
}
