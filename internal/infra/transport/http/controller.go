package http

import "context"

type Controller interface {
	Handler(ctx context.Context, request Request) Response
}
