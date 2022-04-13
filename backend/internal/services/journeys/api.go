package journeys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type JourneysRestController struct {
	svc IJourneyService
}

func NewController(svc IJourneyService) *JourneysRestController {
	return &JourneysRestController{
		svc: svc,
	}
}

type AddItemReq struct {
	ItemID int `json:"item_id"`
	Count  int `json:"count"`
}

func (ctrl JourneysRestController) AddItemToCart(c *fiber.Ctx) error {
	req := AddItemReq{}

	if err := c.BodyParser(&req); err != nil {
		if err == fiber.ErrUnprocessableEntity {
			return fiber.ErrUnprocessableEntity
		}
		return errors.WithStack(err)
	}
	if req.ItemID > 0 || req.Count > 0 {
		return fiber.ErrBadRequest
	}

	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (ctrl JourneysRestController) GetTimeOfJourneys(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"nextCartDeparture": "2020-01-01T00:00:00Z",
		"cartArrival":       "2020-01-01T00:00:00Z",
	})
}

func (ctrl JourneysRestController) GetAllJourneys(c *fiber.Ctx) error {
	journeys, err := ctrl.svc.GetAllJourneys()
	if err != nil {
		return errors.WithStack(err)
	}

	return c.JSON(fiber.Map{
		"journeys": journeys,
	})
}

func (ctrl JourneysRestController) GetSomething(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
