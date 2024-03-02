package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type StructValidator = validator.Validate

func NewStructValidator() *StructValidator {
	validator := validator.New()
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validator
}

type Error struct {
	Field       string `json:"field"`
	StructField string `json:"structField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func (e *Error) Error() string {
	return "Error when validating field '" + e.Field + "' (' " + e.StructField + " ') with value '" + e.Value + "': '" + e.Tag + "'"
}
