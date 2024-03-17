package handlers

import (
	"Web3AuctionApi/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type registerRequest struct {
	Username             string `validate:"required,email" json:"username"`
	Password             string `validate:"required,min=6" json:"password"`
	PasswordConfirmation string `validate:"required,min=6,eqfield=Password" json:"passwordConfirmation"`
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {

	registrationData := registerRequest{}

	err := c.BodyParser(&registrationData)
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	validationErrors := validateRequest(registrationData)
	if validationErrors != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	var existingUsername int64

	err = ah.db.Model(&models.User{}).Where("username = ?", registrationData.Username).Count(&existingUsername).Error
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed creating new user", nil)
	}

	if existingUsername != 0 {
		return errorResponse(c, fiber.StatusForbidden, "Username already in use", nil)
	}

	hashedP, err := bcrypt.GenerateFromPassword([]byte(registrationData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed creating new user", nil)
	}

	user := models.User{Username: registrationData.Username, Password: string(hashedP)}

	err = ah.db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Failed creating new user", nil)
	}

	return successResponse(c, map[string]string{
		"message": "Registration successful",
	})
}
