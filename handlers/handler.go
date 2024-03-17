package handlers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type AuthHandler struct {
	db            *gorm.DB
	tokenDuration int64
	authSecret    string
}

type Handler struct{}

func NewAuthHandler(db *gorm.DB) (*AuthHandler, error) {

	defaultDurationInSeconds := 7200

	authSecret, found := os.LookupEnv("AUTH_SECRET")
	if !found {
		return nil, errors.New("Authentication secret is needed to sign tokens")
	}

	durationInSeconds, found := os.LookupEnv("AUTH_TOKEN_DURATION")
	if !found {
		durationInSeconds = strconv.Itoa(defaultDurationInSeconds)
	}

	v, err := strconv.Atoi(durationInSeconds)
	if err != nil {
		durationInSeconds = strconv.Itoa(defaultDurationInSeconds)
	}

	return &AuthHandler{
		db:            db,
		tokenDuration: int64(v),
		authSecret:    authSecret,
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
