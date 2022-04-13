package warehouse

import "github.com/gofiber/fiber/v2"

type WarehouseRestController struct {
	svc IWarehouseService
}

func NewController(svc IWarehouseService) *WarehouseRestController {
	return &WarehouseRestController{
		svc: svc,
	}
}

func (ctrl WarehouseRestController) GetWarehouses(c *fiber.Ctx) error {
	warehouses, err := ctrl.svc.GetWarehouses()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"warehouses": warehouses,
	})
}

func (ctrl WarehouseRestController) GetRoutes(c *fiber.Ctx) error {
	routes, err := ctrl.svc.GetRoutes()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"routes": routes,
	})
}
