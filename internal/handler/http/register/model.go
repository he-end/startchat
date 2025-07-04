package register

type ReqRegisterModel struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ResRegisterModel struct {
	TokenRegister string `json:"token_register"`
}
