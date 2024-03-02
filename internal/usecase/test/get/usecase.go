package get

import (
	"context"

	"github.com/rubensmichel/waner-on-life/internal/infra/validators"
)

type UseCase struct {
	inputValidator validators.InputValidator
}

func NewUseCase(inputValidator validators.InputValidator) *UseCase {
	return &UseCase{
		inputValidator: inputValidator,
	}
}

func (usc *UseCase) Get(ctx context.Context) ([]Output, error) {
	return []Output{}, nil
}
