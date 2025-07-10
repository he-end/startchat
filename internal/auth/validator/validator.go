package authvalidator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hend41234/startchat/internal/dto"
	"github.com/hend41234/startchat/internal/logger"
	"go.uber.org/zap"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()

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

func ValidationError(err error, rctx context.Context) {
	var validateErrs validator.ValidationErrors
	ctx := logger.FromContext(rctx)

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

			// indent, err := json.MarshalIndent(e, "", "  ")
			// if err != nil {
			// 	fmt.Println(err)
			// 	panic(err)
			// }
			erByte, _ := json.Marshal(e)
			ctx.Debug("validation error", zap.ByteString(e.Field, erByte))
		}
	}
}
