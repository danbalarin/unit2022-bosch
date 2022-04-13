package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"log"
)

type Router struct {
	auth *auth.AuthRestController
}

func NewRouter(
	auth *auth.AuthRestController,
) *Router {
	return &Router{
		auth,
	}
}

func (r Router) Setup(app *fiber.App) {
	log.Println("Setting up routes")
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello, World!"))
	})

	api.Get("/login", r.auth.Login)
}
