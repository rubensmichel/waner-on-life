package get

import (
	"context"

	factory "github.com/rubensmichel/waner-on-life/internal"
	"github.com/rubensmichel/waner-on-life/internal/infra/transport/http"
	"github.com/rubensmichel/waner-on-life/internal/infra/validators"
)

type Controller struct {
	usc *UseCase
}

func NewController(fc *factory.Factory) http.Controller {
	return &Controller{
		usc: NewUseCase(
			fc.DBUser,
			validators.NewInput(),
		),
	}
}

func (ctrl *Controller) Handler(ctx context.Context, req http.Request) http.Response {
	users, err := ctrl.usc.Get(ctx)
	if err != nil {
		return http.HandlerError(ctx, err)
	}
	return http.Ok(users)
}
