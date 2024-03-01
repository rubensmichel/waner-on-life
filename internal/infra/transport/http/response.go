package http

import (
	"context"
)

type Response struct {
	Code int
	Body any
}
type ErrorMessage struct {
	Error error `json:"error"`
}
type InternalServerErrorMessage struct {
	Error string `json:"error"`
}

func CreateResponse(code int, body ...any) Response {
	var message any
	if len(body) > 0 {
		message = body[0]
	}
	resp := Response{
		Body: message,
		Code: code,
	}
	return resp
}
func Ok(body any) Response {
	return CreateResponse(200, body)
}
func Created(body any) Response {
	return CreateResponse(201, body)
}
func NotFound(body any) Response {
	return CreateResponse(404, body)
}
func UnprocessableEntity(body interface{}) Response {
	return CreateResponse(422, body)
}
func Conflict(body interface{}) Response {
	return CreateResponse(409, body)
}
func Accepted(body any) Response {
	return CreateResponse(202, body)
}
func NoContent() Response {
	return CreateResponse(204, nil)
}
func InternalServerError(ctx context.Context, err error) Response {
	return CreateResponse(500, &InternalServerErrorMessage{
		Error: "Internal Server Error",
	})
}
func ServiceUnavailable(body any) Response {
	return CreateResponse(503, body)
}
