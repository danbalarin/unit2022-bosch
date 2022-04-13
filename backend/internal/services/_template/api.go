package template

import "github.com/gofiber/fiber/v2"

type Controller struct {
	svc ITemplateService
}

func NewController(svc ITemplateService) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (ctrl Controller) GetSomething(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
