package request

type PasswordRecoveryRequest struct {
    Identifier string `json:"identifier" binding:"required"`
}

type VerifyCodeRequest struct {
    Identifier string `json:"identifier" binding:"required"`
    Code       string `json:"code" binding:"required,min=6,max=6"`
}

type ResetPasswordRequest struct {
    Identifier   string `json:"identifier" binding:"required"`
    Code         string `json:"code" binding:"required,min=6,max=6"`
    NewPassword  string `json:"new_password" binding:"required,min=8"`
    ConfirmPassword string `json:"confirm_password" binding:"required,min=8"`
}