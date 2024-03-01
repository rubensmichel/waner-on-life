package http

import (
	"context"
	"errors"
	"net/http"

	domainErrors "github.com/rubensmichel/waner-on-life/internal/types/errors"
)

func HandlerError(ctx context.Context, err error) Response {
	var converted domainErrors.Error
	if errors.As(err, &converted) {
		converted = converted.WithContext(ctx)
		if converted.Kind == domainErrors.KindInternalError {
			return InternalServerError(ctx, err)
		}
		return CreateResponse(getHTTPStatusCodeFromDomainError(converted), ErrorMessage{Error: converted})
	}
	return InternalServerError(ctx, err)
}

func getHTTPStatusCodeFromDomainError(err domainErrors.Error) int {
	switch err.Kind {
	case domainErrors.KindInvalidInput:
		return http.StatusBadRequest
	case domainErrors.KindBusinessRule:
		return http.StatusConflict
	case domainErrors.KindInternalError:
		return http.StatusInternalServerError
	case domainErrors.KindNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
