package model

type Paciente struct {
	BaseUser
	CPF 		   string `json:"cpf"`
}

func NewPaciente(email, password, name, cpf string) *Paciente {
	return &Paciente{
		BaseUser: BaseUser{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     "paciente",
		},
		CPF: cpf,
	}
}

func (p *Paciente) GetCPF() string {
	return p.CPF
}