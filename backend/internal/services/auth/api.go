package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
)

const AuthTokenKey = "Authorization"

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
	user, err := ctrl.svc.GetUserFromToken(token)
	if err != nil {
		if errors.Is(err, ErrInvalidToken) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		return errors.WithStack(err)
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}
	if !user.Role.IsUser() {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

func (ctrl AuthRestController) IsAdmin(c *fiber.Ctx) error {
	token := c.Get(AuthTokenKey, "")
	user, err := ctrl.svc.GetUserFromToken(token)
	if err != nil {
		if errors.Is(err, ErrInvalidToken) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}
		return errors.WithStack(err)
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}
	if !user.Role.IsAdmin() {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ctrl AuthRestController) GetUserProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"email": "",
	})
}

func (ctrl AuthRestController) Login(c *fiber.Ctx) error {
	req := LoginReq{}
	if err := c.BodyParser(&req); err != nil {
		if err == fiber.ErrUnprocessableEntity {
			return fiber.ErrUnprocessableEntity
		}
		return errors.WithStack(err)
	}
	if req.Email == "" || req.Password == "" {
		return fiber.ErrBadRequest
	}
	loginRes, err := ctrl.svc.Login(req.Email, req.Password)
	if err != nil {
		if err == ErrInvalidCredentials {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid email or password",
			})
		}
		return errors.WithStack(err)
	}
	return c.JSON(fiber.Map{
		"token": loginRes.Token,
		"role":  loginRes.Role,
	})
}

func (ctrl AuthRestController) GetUser(c *fiber.Ctx) (*entity.User, error) {
	token := c.Get(AuthTokenKey, "")
	user, err := ctrl.svc.GetUserFromToken(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}
