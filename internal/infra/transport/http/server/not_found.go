package httpserver

import (
	"github.com/gofiber/fiber/v2"
)

func routeNotFound() func(ctx *fiber.Ctx) error {
	return func(fc *fiber.Ctx) error {
		return fc.Status(404).JSON("not found route")
	}
}
