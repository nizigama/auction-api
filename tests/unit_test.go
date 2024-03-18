package tests

import (
	"Web3AuctionApi/handlers"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	app := fiber.New()

	app.Get("/", handlers.Hello)

	// http.Request
	req := httptest.NewRequest("GET", "http://127.0.0.1:3000", nil)

	// http.Response
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status code to be %d but got %d", fiber.StatusOK, resp.StatusCode)
	}

	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		if string(body) != `{"message":"Hello, World!"}` {
			t.Errorf("Expected the response to be %s but got %s", `{"message":"Hello, World!"}`, string(body))
		}
	}
}
