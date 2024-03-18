package tests

import (
	"Web3AuctionApi/business"
	"Web3AuctionApi/handlers"
	"bytes"
	"encoding/json"
	"fmt"
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

func TestSuccessfulBid(t *testing.T) {

	app := fiber.New()

	conn := new(business.EthConnectionMock)

	handler := handlers.EthereumHandler{
		Connection: conn,
	}

	bidAmount := int64(1200)

	app.Post("auction/bid", handler.Bid)

	data, _ := json.Marshal(map[string]int64{
		"amount": bidAmount,
	})

	r := bytes.NewReader(data)

	// http.Request
	req := httptest.NewRequest("POST", "http://127.0.0.1:3000/auction/bid", r)
	req.Header.Set("Content-Type", "application/json")

	conn.On("Bid", bidAmount).Return(nil)

	// http.Response
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status code to be %d but got %d", fiber.StatusOK, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if string(body) != `{"message":"Bid successful"}` {
		t.Fatalf("Expected the response to be %s but got %s", `{"message":"Bid successful"}`, string(body))
	}

	//passing
	//higher bid already made
	////unauthenticated

}

func TestHigherBidAlreadySubmitted(t *testing.T) {

	app := fiber.New()

	conn := new(business.EthConnectionMock)

	handler := handlers.EthereumHandler{
		Connection: conn,
	}

	bidAmount := int64(1200)

	app.Post("auction/bid", handler.Bid)

	data, _ := json.Marshal(map[string]int64{
		"amount": bidAmount,
	})

	r := bytes.NewReader(data)

	// http.Request
	req := httptest.NewRequest("POST", "http://127.0.0.1:3000/auction/bid", r)
	req.Header.Set("Content-Type", "application/json")

	conn.On("Bid", bidAmount).Return(business.HigherBidAlreadySubmitted{})

	// http.Response
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("Expected status code to be %d but got %d", fiber.StatusBadRequest, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if string(body) != fmt.Sprintf(`{"extras":null,"message":"%s"}`, business.HigherBidAlreadySubmitted{}.Error()) {
		t.Fatalf("Expected the response to be %s but got %s", fmt.Sprintf(`{"extras":null,"message":"%s"}`, business.HigherBidAlreadySubmitted{}.Error()), string(body))
	}
}
