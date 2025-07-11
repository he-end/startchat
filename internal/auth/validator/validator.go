package authvalidator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/hend41234/startchat/internal/dto"
	"github.com/hend41234/startchat/internal/logger"
	"go.uber.org/zap"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// ========================= field validator register =========================
	// [*] registery otp purpose
	Validate.RegisterValidation("otp_purpose", OtpPurposeValidation)

	// struct validator register
	Validate.RegisterStructValidation(ValidatorOTPRequest, dto.ReqVerifyOTPModel{})
	Validate.RegisterStructValidation(ValidatorRegister, dto.ReqRegisterModel{})
}

type validationError struct {
	Namespace       string `json:"namespace"` // can differ when a custom TagNameFunc is registered or
	Field           string `json:"field"`     // by passing alt name to ReportError like below
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

func ValidationError(err error, rctx context.Context) []map[string]string {
	var validateErrs validator.ValidationErrors
	ctx := logger.FromContext(rctx)
	resErr := []map[string]string{}
	if errors.As(err, &validateErrs) {
		for _, err := range validateErrs {
			e := validationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}
			var errField = map[string]string{}
			errField["field"] = e.StructField
			errField["tag"] = e.ActualTag
			resErr = append(resErr, errField)
			// indent, err := json.MarshalIndent(e, "", "  ")
			// if err != nil {
			// 	fmt.Println(err)
			// 	panic(err)
			// }
			erByte, _ := json.Marshal(e)
			ctx.Debug("validation error", zap.ByteString(e.Field, erByte))
		}
	}
	return resErr
}
