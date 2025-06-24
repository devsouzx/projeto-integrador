package request

type CadastroRequest struct {
	NomeCompleto      string `json:"nomecompleto" binding:"required"`
	Cpf               string `json:"cpf" binding:"required"`
	Cartaosus         string `json:"cartaosus" binding:"required"`
	NomeCompletoDaMae string `json:"nomecompletodamae" binding:"required"`
	DataDeNascimento  string `json:"datadenascimento" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Telefone          string `json:"telefone" binding:"required"`
	Senha             string `json:"senha" binding:"required"`
	Endereco          EnderecoRequest `json:"endereco"`
}

type EnderecoRequest struct {
    Cep         string `json:"cep"`
    Logradouro  string `json:"logradouro"`
    Numero      string `json:"numero"`
    Complemento string `json:"complemento"`
    Bairro      string `json:"bairro"`
    Cidade      string `json:"cidade"`
    Uf          string `json:"uf"`
}

type CadastroResponse struct {
	ID                string `json:"id" binding:"required"`
	NomeCompleto      string `json:"nomecompleto" binding:"required"`
	Cpf               string `json:"cpf" binding:"required"`
	Cartaosus         string `json:"cartaosus" binding:"required"`
	NomeCompletoDaMae string `json:"nomecompletodamae" binding:"required"`
	DataDeNascimento  string `json:"datadenascimento" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Telefone          string `json:"telefone" binding:"required"`
	Senha             string `json:"senha" binding:"required"`
}

//cadastro de usuário
