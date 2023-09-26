package validation

import (
	"encoding/json"
	"errors"
	rest_err "github.com/Micheli97/crud-campeonato-golang/config/error"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		e := en.New()
		unt := ut.New(e, e)
		transl, _ := unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Inválid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errosCause := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errosCause = append(errosCause, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fileds are inválid", errosCause)
	}

	return rest_err.NewBadRequestError("Error trying to convert fields")
}
