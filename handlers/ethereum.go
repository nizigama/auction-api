package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

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
