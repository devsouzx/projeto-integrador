package request

type LoginRequest struct {
	Role string `form:"role" json:"role"`
	Identificador string `form:"identificador" json:"identificador"`
	Password string `form:"password" json:"password"`
}
