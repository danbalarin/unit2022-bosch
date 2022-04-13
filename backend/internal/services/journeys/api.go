package journeys

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"log"
)

type JourneysRestController struct {
	svc      IJourneyService
	authCtrl *auth.AuthRestController
}

func NewController(svc IJourneyService, authCtrl *auth.AuthRestController) *JourneysRestController {
	return &JourneysRestController{
		svc:      svc,
		authCtrl: authCtrl,
	}
}

type AddItemReq struct {
	ItemID      int `json:"itemId"`
	Count       int `json:"count"`
	WarehouseID int `json:"warehouseId"`
}

func (ctrl JourneysRestController) AddItemToCart(c *fiber.Ctx) error {
	req := AddItemReq{}

	if err := c.BodyParser(&req); err != nil {
		if err == fiber.ErrUnprocessableEntity {
			return fiber.ErrUnprocessableEntity
		}
		return errors.WithStack(err)
	}
	log.Println("REQ:", req)
	if req.ItemID <= 0 || req.Count <= 0 {
		return fiber.ErrBadRequest
	}

	user, err := ctrl.authCtrl.GetUser(c)
	if err != nil {
		return err
	}

	err = ctrl.svc.AddItemsToJourney(AddItemsParams{
		ItemID:      uint(req.ItemID),
		Count:       uint(req.Count),
		User:        user,
		WarehouseID: uint(req.WarehouseID),
	})
	if err != nil {
		return errors.WithStack(err)
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
