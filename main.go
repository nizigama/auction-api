package main

import (
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

	err = db.AutoMigrate(&models.User{})
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
