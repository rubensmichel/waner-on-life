package request

import "context"

type contextKey struct{}

var requestKey = contextKey{}

type Request struct {
	ID string
}

func New(id string) *Request {
	return &Request{id}
}

func FromContext(ctx context.Context) *Request {
	if ctx == nil {
		return nil
	}
	request, ok := ctx.Value(requestKey).(*Request)
	if !ok {
		return nil
	}
	return request
}

func NewContext(ctx context.Context, r *Request) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, requestKey, r)
}

func RequestID(r *Request) string {
	if r == nil {
		return ""
	}
	return r.ID
}
