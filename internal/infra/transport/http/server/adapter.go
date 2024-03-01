package httpserver

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/google/uuid"
	"github.com/rubensmichel/waner-on-life/internal/infra/telemetry/request"
	"github.com/rubensmichel/waner-on-life/internal/infra/transport/http"
)

const (
	httpRequestTimeout = 30 * time.Second
)

type Route struct {
	RequestTimeout time.Duration
}

type RouteOption func(*Route)

func WithTimeout(duration time.Duration) RouteOption {
	return func(r *Route) {
		r.RequestTimeout = duration
	}
}

func adaptRoute(ctrl http.Controller, opts ...RouteOption) func(c *fiber.Ctx) error {
	routeOptions := &Route{}

	for _, opt := range opts {
		opt(routeOptions)
	}

	return timeout.NewWithContext(
		func(ctx *fiber.Ctx) error {
			response := handleRequest(ctx, ctrl)
			return ctx.Status(response.Code).JSON(response.Body)
		},
		routeOptions.RequestTimeout,
	)
}

func handleRequest(fctx *fiber.Ctx, ctrl http.Controller) http.Response {
	headers := adaptHeaders(fctx.GetReqHeaders())
	ctx := request.NewContext(fctx.UserContext(), request.New(uuid.New().String()))
	body := adaptBody(fctx.Body())

	request := http.NewRequestBuilder().
		DomainName(fctx.Hostname()).
		UrlPath(fctx.Path()).
		Method(fctx.Method()).
		Headers(headers).
		Body(body).
		Params(adaptParams(fctx)).
		Query(adaptQuery(fctx)).
		Build()

	return ctrl.Handler(ctx, *request)
}

func adaptHeaders(headers map[string][]string) map[string]string {
	newHeaders := map[string]string{}
	i := 0
	for k, v := range headers {
		newHeaders[strings.ToLower(k)] = v[i]
		i += 1
	}
	return newHeaders
}

func adaptParams(c *fiber.Ctx) map[string]string {
	values := map[string]string{}
	args := c.Route().Params
	for _, v := range args {
		values[v] = c.Params(v)
	}
	return values
}

func adaptQuery(c *fiber.Ctx) map[string]string {
	values := map[string]string{}
	args := c.Context().QueryArgs()

	args.VisitAll(func(key, value []byte) {
		k := string(key)
		v := string(value)

		values[k] = v
	})

	return values
}

func adaptBody(mutableBytes []byte) []byte {
	buffer := make([]byte, len(mutableBytes))
	copy(buffer, mutableBytes)
	return buffer
}
