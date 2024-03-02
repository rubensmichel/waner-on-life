package httpserver

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rubensmichel/waner-on-life/internal"

	getuser "github.com/rubensmichel/waner-on-life/internal/usecase/user/get"
)

const (
	port = ":3000"
)

type httpServer struct {
	server *fiber.App
}

func New(ft *internal.Factory) *httpServer {
	app := fiber.New(fiber.Config{ReadBufferSize: 8192})

	router := app.Group("/waner-on-life/v1")

	router.Get("/users", adaptRoute(getuser.NewController(ft)))

	app.Use(routeNotFound())

	return &httpServer{
		server: app,
	}
}

func (srv *httpServer) Listen() error {
	fmt.Print(context.Background(), "Starting HTTP Server on port "+port)
	err := srv.server.Listen(port)
	if err != nil {
		fmt.Print(err)
	}
	return nil
}

func (srv *httpServer) Shutdown() error {
	fmt.Print(context.Background(), "Stopping HTTP Server")
	return srv.server.Shutdown()
}
