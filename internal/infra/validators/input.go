package validators

import (
	"context"
	"reflect"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/rubensmichel/waner-on-life/internal/types/errors"
)

type Input struct {
	validator *validator.Validate
}

func NewInput() *Input {
	input := &Input{
		validator: validator.New(),
	}
	input.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		split := strings.SplitN(fld.Tag.Get("json"), ",", 2)
		if len(split) == 0 {
			return ""
		}
		name := split[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return input
}

func (val *Input) Required(value any) InputValidator {
	return val
}

func (val *Input) Validate(ctx context.Context, input any) error {
	err := val.validator.Struct(input)
	if err == nil {
		return nil
	}

	detailedErr := errors.ErrInvalidInput.WithContext(ctx)
	for _, valErr := range err.(validator.ValidationErrors) {
		detailedErr = val.buildError(detailedErr, valErr)
	}
	return detailedErr
}

func (val *Input) buildError(err errors.Error, fieldErr validator.FieldError) errors.Error {
	fieldPath := val.attribute(fieldErr.Namespace())
	switch fieldErr.Tag() {
	case "required", "required_for":
		return err.WithDetailMessage(fieldPath, errors.DetailRequiredAttributeMissing)
	default:
		return err.WithDetailMessage(fieldPath, errors.DetailInvalidValue)
	}
}

func (val *Input) attribute(namespace string) string {
	split := strings.Split(namespace, ".")
	if len(split) > 1 {
		return strings.Join(split[1:], ".")
	}
	return split[0]
}
