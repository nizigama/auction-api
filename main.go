package main

import (
	_ "Web3AuctionApi/docs"
	"Web3AuctionApi/models"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

const defaultServerPort = "3000"

// @title Auction contract API
// @version 1.0
// @description This is the API for the auction contract
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort, found := os.LookupEnv("APP_PORT")
	if !found {
		appPort = defaultServerPort
	}

	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.InvalidToken{})
	if err != nil {
		log.Fatalln(err)
	}

	err = loadRoutes(app, db)
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Listen(fmt.Sprintf(":%s", appPort))
	if err != nil {
		log.Fatalln(err)
	}
}
