package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

type Handler struct{}

func NewAuthHandler(db *gorm.DB) (*AuthHandler, error) {

	return &AuthHandler{
		db: db,
	}, nil
}

func (h *Handler) Hello(c *fiber.Ctx) error {

	return c.JSON(map[string]interface{}{
		"message": "Hello, World!",
	})
}

func errorResponse(c *fiber.Ctx, status int, message string, extra interface{}) error {

	return c.Status(status).JSON(fiber.Map{"message": message, "extras": extra})
}

func successResponse(c *fiber.Ctx, data interface{}) error {

	return c.Status(fiber.StatusOK).JSON(data)
}

func validateRequest(data interface{}) []string {

	validate := validator.New()

	var validationErrors []string

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {

			validationErrors = append(validationErrors, fmt.Sprintf("%s: %s", err.Field(), err.Error()))
		}
	}

	return validationErrors
}
