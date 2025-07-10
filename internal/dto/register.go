package dto

type ReqRegisterModel struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ResRegisterModel struct {
	TokenRegister string `json:"token_register"`
}
