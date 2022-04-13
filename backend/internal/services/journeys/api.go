package journeys

import "github.com/gofiber/fiber/v2"

type JourneysRestController struct {
	svc IJourneyService
}

func NewController(svc IJourneyService) *JourneysRestController {
	return &JourneysRestController{
		svc: svc,
	}
}

func (ctrl JourneysRestController) GetSomething(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
