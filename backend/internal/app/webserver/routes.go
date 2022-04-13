package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"github.com/unit2022-bosch/teapot/backend/internal/services/items"
	"github.com/unit2022-bosch/teapot/backend/internal/services/journeys"
	"github.com/unit2022-bosch/teapot/backend/internal/services/warehouse"
	"log"
)

type Router struct {
	auth       *auth.AuthRestController
	items      *items.ItemsRestController
	journeys   *journeys.JourneysRestController
	warehouses *warehouse.WarehouseRestController
}

func NewRouter(
	auth *auth.AuthRestController,
	items *items.ItemsRestController,
	journeys *journeys.JourneysRestController,
	warehouses *warehouse.WarehouseRestController,
) *Router {
	return &Router{
		auth,
		items,
		journeys,
		warehouses,
	}
}

func (r Router) Setup(app *fiber.App) {
	log.Println("Setting up routes")
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello, World!"))
	})

	api.Post("/login", r.auth.Login)
	api.Get("/profile", r.auth.IsUser, r.auth.GetUserProfile)
	api.Get("/items", r.auth.IsUser, r.items.GetItems)
	api.Get("/warehouses", r.auth.IsUser, r.warehouses.GetWarehouses)
	api.Get("/routes", r.auth.IsUser, r.warehouses.GetRoutes)
	api.Post("/order", r.auth.IsUser, r.journeys.AddItemToCart)
	api.Get("/journey", r.auth.IsUser, r.journeys.GetTimeOfJourneys)

	api.Get("/journeys", r.auth.IsAdmin, r.journeys.GetAllJourneys)
}
