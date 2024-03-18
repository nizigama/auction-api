package handlers

import (
	"Web3AuctionApi/business"
	"Web3AuctionApi/models"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"log"
	"os"
	"slices"
	"strconv"
)

type AuthHandler struct {
	db             *gorm.DB
	tokenDuration  int64
	authSecret     string
	AuthMiddleware fiber.Handler
}

type EthereumHandler struct {
	instanceUrl string
	Connection  business.Connection
}

func NewEthereumHandler() (*EthereumHandler, error) {

	instanceUrl, found := os.LookupEnv("INSTANCE_URL")
	if !found {
		return nil, errors.New("Instance url is needed to connect to an ethereum node")
	}

	conn, err := business.NewBlockchainConnection()
	if err != nil {
		return nil, err
	}

	return &EthereumHandler{
		instanceUrl: instanceUrl,
		Connection:  conn,
	}, nil
}

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

	authMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(authSecret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return errorResponse(ctx, fiber.StatusForbidden, err.Error(), nil)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {

			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			username := claims["username"].(string)

			var invalidTokens []models.InvalidToken

			err := db.Model(models.InvalidToken{}).Where("username = ?", username).Find(&invalidTokens).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println(err)
				return errorResponse(ctx, fiber.StatusInternalServerError, "Server error", nil)
			}

			if invalidTokens == nil {
				return ctx.Next()
			}

			tokens := []string{}

			for _, v := range invalidTokens {
				tokens = append(tokens, v.Token)
			}

			idx := slices.Index(tokens, user.Raw)
			if idx != -1 {
				return errorResponse(ctx, fiber.StatusForbidden, "Invalidated token", nil)
			}

			return ctx.Next()

		},
	})

	return &AuthHandler{
		db:             db,
		tokenDuration:  int64(v),
		authSecret:     authSecret,
		AuthMiddleware: authMiddleware,
	}, nil
}

func Hello(c *fiber.Ctx) error {

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
