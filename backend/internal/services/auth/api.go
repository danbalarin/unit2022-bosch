package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

const AuthTokenKey = "X-Auth-Token"

type AuthRestController struct {
	svc IAuthService
}

func NewAuthRestController(svc IAuthService) *AuthRestController {
	return &AuthRestController{
		svc: svc,
	}
}

func (ctrl AuthRestController) IsUser(c *fiber.Ctx) error {
	token := c.Get(AuthTokenKey, "")
	userRole, err := ctrl.svc.GetUserRoleFromToken(token)
	if err != nil {
		return errors.WithStack(err)
	}
	if userRole == nil || !userRole.IsUser() {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

func (ctrl AuthRestController) IsAdmin(c *fiber.Ctx) error {
	token := c.Get(AuthTokenKey, "")
	userRole, err := ctrl.svc.GetUserRoleFromToken(token)
	if err != nil {
		return errors.WithStack(err)
	}
	if userRole == nil || !userRole.IsAdmin() {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ctrl AuthRestController) Login(c *fiber.Ctx) error {
	req := LoginReq{}
	if err := c.BodyParser(&req); err != nil {
		return errors.WithStack(err)
	}
	if req.Email == "" || req.Password == "" {
		return fiber.ErrBadRequest
	}
	token, err := ctrl.svc.Login(req.Email, req.Password)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
