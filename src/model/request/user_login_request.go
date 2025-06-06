package request

type UserLogin struct {
	Identificador string `form:"identificador" json:"identificador"`
	Password string `form:"password" json:"password"`
}
