package items

import "github.com/gofiber/fiber/v2"

type ItemsRestController struct {
	svc IItemsService
}

func NewController(svc IItemsService) *ItemsRestController {
	return &ItemsRestController{
		svc: svc,
	}
}

func (ctrl ItemsRestController) GetItems(c *fiber.Ctx) error {
	items, err := ctrl.svc.GetItems()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"items": items,
	})
}
