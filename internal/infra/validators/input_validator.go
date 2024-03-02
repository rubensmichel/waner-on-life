package validators

import "context"

type InputValidator interface {
	Required(field any) InputValidator
	Validate(ctx context.Context, input interface{}) error
}
