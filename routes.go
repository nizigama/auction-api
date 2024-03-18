package main

import (
	"Web3AuctionApi/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func loadRoutes(app *fiber.App, db *gorm.DB) error {

	authHandler, err := handlers.NewAuthHandler(db)
	if err != nil {
		return err
	}

	ethHandler, err := handlers.NewEthereumHandler()
	if err != nil {
		return err
	}

	app.Use(recover.New())

	app.Get("/", handlers.Hello)
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)
	app.Use(authHandler.AuthMiddleware).Post("/logout", authHandler.Logout)

	auctionRoutes := app.Use(authHandler.AuthMiddleware).Group("auction")

	auctionRoutes.Get("status", ethHandler.GetStatus)
	auctionRoutes.Get("bids", ethHandler.History)
	auctionRoutes.Post("bid", ethHandler.Bid)
	auctionRoutes.Get("statistics", ethHandler.GetStatistics)
	auctionRoutes.Post("deploy", ethHandler.Deploy)

	return nil
}
