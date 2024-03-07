package get

import (
	"context"

	"github.com/rubensmichel/waner-on-life/internal/infra/db"
	"github.com/rubensmichel/waner-on-life/internal/infra/validators"
)

type UseCase struct {
	repoUser       db.Users
	inputValidator validators.InputValidator
}

func NewUseCase(repoUser db.Users, inputValidator validators.InputValidator) *UseCase {
	return &UseCase{
		repoUser:       repoUser,
		inputValidator: inputValidator,
	}
}

func (usc *UseCase) Get(ctx context.Context) ([]Output, error) {
	output := []Output{}

	users, err := usc.repoUser.Find(ctx)
	if err != nil {
		return output, err
	}

	for _, v := range users {
		output = append(output, Output{
			ID: v.ID,
		})
	}

	return output, nil
}
