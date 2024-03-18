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

type deployRequest struct {
	DurationInSeconds  int64  `validate:"required,numeric,min=1" json:"durationInSeconds"`
	BeneficiaryAddress string `validate:"required,hexadecimal" json:"beneficiaryAddress"`
}

// GetStatus godoc
// @Summary      Get auction status
// @Description  Auction status
// @Tags         auction
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {array}   business.AuctionStatus
// @Router       /auction/status [get]
func (eh *EthereumHandler) GetStatus(c *fiber.Ctx) error {

	status, err := eh.Connection.GetAuctionStatus()
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	return c.JSON(status)
}

// GetStatistics godoc
// @Summary      Get auction statistics
// @Description  Auction statistics
// @Tags         auction
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {array}   business.Stats
// @Router       /auction/statistics [get]
func (eh *EthereumHandler) GetStatistics(c *fiber.Ctx) error {

	stats, err := eh.Connection.Stats()
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	return c.JSON(stats)
}

// History godoc
// @Summary      List bids
// @Description  Bids history
// @Tags         auction
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {array}   []business.Bid
// @Router       /auction/bids [get]
func (eh *EthereumHandler) History(c *fiber.Ctx) error {

	bids, err := eh.Connection.ListAllBids()
	if err != nil {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	return c.JSON(bids)
}

// Bid func bidding
// @Description Bid.
// @Summary Bid
// @Tags auction
// @Accept json
// @Produce json
// @Param request body bidRequest true "Bid"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string "Unauthenticated."
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auction/bid [post]
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

	err = eh.Connection.Bid(bidData.Amount)
	if err != nil && !errors.Is(err, business.HigherBidAlreadySubmitted{}) && !errors.Is(err, business.AuctionEnded{}) {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	if errors.Is(err, business.HigherBidAlreadySubmitted{}) {

		return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if errors.Is(err, business.AuctionEnded{}) {

		return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return successResponse(c, map[string]string{
		"message": "Bid successful",
	})
}

// Deploy func deployment
// @Description Deploy contract.
// @Summary Deploy contract
// @Tags auction
// @Accept json
// @Produce json
// @Param request body deployRequest true "Deployment"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string "Unauthenticated."
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auction/deploy [post]
func (eh *EthereumHandler) Deploy(c *fiber.Ctx) error {

	reqData := deployRequest{}

	err := c.BodyParser(&reqData)
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	validationErrors := validateRequest(reqData)
	if validationErrors != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Validation failed", validationErrors)
	}

	contractAddr, txHash, err := eh.Connection.Deploy(reqData.DurationInSeconds, reqData.BeneficiaryAddress)
	if err != nil && !errors.Is(err, business.HigherBidAlreadySubmitted{}) {
		log.Println(err)
		return errorResponse(c, fiber.StatusInternalServerError, "Server error", nil)
	}

	return successResponse(c, map[string]string{
		"message":         "Deployment successful",
		"contractAddress": contractAddr,
		"transactionHash": txHash,
	})
}
