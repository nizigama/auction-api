package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "Hello, World!",
		})
	})

	err = app.Listen(fmt.Sprintf(":%s", appPort))
	if err != nil {
		log.Fatalln(err)
	}
}
