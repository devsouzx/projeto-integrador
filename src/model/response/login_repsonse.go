package response

type LoginResponse struct {
	Message  string       `json:"message"`
	Token    string       `json:"token"`
	User     UserResponse `json:"user"`
	Redirect string       `json:"redirect"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
