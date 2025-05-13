package handler

import (
	"github.com/decrypt6969/auth-service/internal/model"
	"github.com/decrypt6969/auth-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{userService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req model.User

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, email and password are required",
		})
	}

	if err := h.userService.Register(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":         req.ID,
		"name":       req.Name,
		"email":      req.Email,
		"created_at": req.CreatedAt,
	})
}
