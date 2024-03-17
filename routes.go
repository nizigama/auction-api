package main

import (
	"Web3AuctionApi/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func loadRoutes(app *fiber.App, db *gorm.DB) error {

	handler := new(handlers.Handler)

	authHandler, err := handlers.NewAuthHandler(db)
	if err != nil {
		return err
	}

	app.Get("/", handler.Hello)
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)
	app.Use(authHandler.AuthMiddleware).Post("/logout", authHandler.Logout)

	return nil
}
