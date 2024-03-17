package handlers

import (
	"Web3AuctionApi/business"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

type bidRequest struct {
	Amount int64 `validate:"required,numeric,min=1" json:"amount"`
}

func (eh *EthereumHandler) GetStatus(c *fiber.Ctx) error {

	status, err := eh.connection.GetAuctionStatus()
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	return c.JSON(status)
}

func (eh *EthereumHandler) History(c *fiber.Ctx) error {

	bids, err := eh.connection.ListAllBids()
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	return c.JSON(bids)
}

func (eh *EthereumHandler) Bid(c *fiber.Ctx) error {

	bidData := bidRequest{}

	err := c.BodyParser(&bidData)
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	validationErrors := validateRequest(bidData)
	if validationErrors != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	err = eh.connection.Bid(bidData.Amount)
	if err != nil && !errors.Is(err, business.HigherBidAlreadySubmitted{}) {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	if errors.Is(err, business.HigherBidAlreadySubmitted{}) {

		return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return successResponse(c, map[string]string{
		"message": "Bid successful",
	})
}
