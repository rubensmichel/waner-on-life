package get

import (
	"context"

	entity "github.com/rubensmichel/waner-on-life/internal/domain"
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
	output := []Output{}

	users := []entity.User{{ID: 1}}

	for _, v := range users {
		output = append(output, Output{
			ID: v.ID,
		})
	}

	return output, nil
}
