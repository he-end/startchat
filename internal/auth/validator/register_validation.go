package authvalidator

import (
	"github.com/go-playground/validator/v10"
	authpassword "github.com/hend41234/startchat/internal/auth/passwords"
	"github.com/hend41234/startchat/internal/dto"
	"github.com/hend41234/startchat/internal/internalutils"
)

func ValidatorRegister(sl validator.StructLevel) {
	reg := sl.Current().Interface().(dto.ReqRegisterModel)

	// validation email
	if !internalutils.EmailDetetor(reg.Email) {
		sl.ReportError(reg.Email, "Email", "email", "email", "")
	}

	// validation password and confirm password
	if !authpassword.IsValidPassword(reg.Password) || !authpassword.IsValidPassword(reg.ConfirmPassword) {
		sl.ReportError(reg.Password, "password", "Password", "passwordorconfigpasswrd", "")
		sl.ReportError(reg.ConfirmPassword, "confirmpassword", "ConfirmPassword", "passwordorconfigpasswrd", "")
	}
}
