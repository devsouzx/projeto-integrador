package request

type LoginRequest struct {
	Identifier string `form:"identificador" json:"identificador"`
	Password string `form:"password" json:"password"`
	Role string `form:"role" json:"role"`
}
