package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/stretchr/testify/assert"
)

func TestShortenURLHappyPath(t *testing.T) {
	server := fiber.New()

	server.Post("/api/v1", ShortenURL)

	// mock data
	happyPathInput := struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"custom_short"`
		Expiry      time.Duration `json:"expiry"`
	}{
		URL:         "www.google.com",
		CustomShort: "testing3432", // need to come up with a way to not have to modify this everytime for tests
		Expiry:      100,
	}

	reqBody, _ := json.Marshal(happyPathInput)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := server.Test(req, -1)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(resp.StatusCode)
	t.Log(string(body))
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestShortenURLDuplicateCustomURL(t *testing.T) {
	server := fiber.New()

	// Add rate limiting middleware
	server.Use(limiter.New(limiter.Config{
		Max:               1,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	server.Post("/api/v1", ShortenURL)

	// mock data
	rateLimitInput := struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"custom_short"`
		Expiry      time.Duration `json:"expiry"`
	}{
		URL:         "www.google.com",
		CustomShort: "testing",
		Expiry:      100,
	}

	reqBody, _ := json.Marshal(rateLimitInput)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// make the request 10 times

	// 11th time should have rate limit error
	resp, _ := server.Test(req, -1)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(resp.StatusCode)
	t.Log(string(body))
	assert.Equal(t, fiber.StatusConflict, resp.StatusCode)
}

// func TestShortenURLBadData(t *testing.T) {
// 	server := fiber.New()

// 	server.Post("/api/v1", ShortenURL)

// 	// mock data
// 	badInput := struct {
// 		URL         string        `json:"url"`
// 		CustomShort string        `json:"custom_short"`
// 		Expiry      time.Duration `json:"expiry"`
// 	}{
// 		URL:         "",
// 		CustomShort: "null",
// 		Expiry:      100,
// 	}

// 	reqBody, _ := json.Marshal(badInput)
// 	req, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// make the request 10 times

// 	// 11th time should have rate limit error
// 	resp, _ := server.Test(req, -1)

// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	// t.Log(resp.StatusCode)
// 	// t.Log(string(body))
// 	assert.Equal(t, 400, resp.StatusCode)
// }
