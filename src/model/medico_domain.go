package model

type Medico struct {
	BaseUser
	CRM             string `json:"crm"`
	Especialidade   string `json:"especialidade"`
	Telefone        string `json:"telefone"`
	DataNascimento  string `json:"data_nascimento"`
	EnderecoConsulta string `json:"endereco_consulta"`
}

func NewMedico(email, password, name string) User {
	return &Medico{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "medico",
		},
	}
}