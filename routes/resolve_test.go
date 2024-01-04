package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type CreateShortURLResponse struct {
	URL             string        `json:"url`
	CustomShort     string        `json:"custom_short`
	Expiry          time.Duration `json:"expiry`
	XRateRemaining  int           `json:"rate_limit`
	XRateLimitReset time.Duration `json:"rate_limit_reset`
}

var app *fiber.App

func TestMain(m *testing.M) {
	// Load environment variables from the file
	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}

	// Set up your Fiber app or any other setup tasks
	app = fiber.New()

	// Run the tests
	code := m.Run()

	// Perform teardown tasks, if necessary

	// Exit
	os.Exit(code)
}

func TestResolveURLWithoutCustomShort(t *testing.T) {
	server := fiber.New()

	// need to create this in the db first
	server.Post("/api/v1", ShortenURL)

	// CREATE A MOCK ENTRY WITHOUT A CUSTOM SHORT
	createURLMockData := struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"custom_short"`
		Expiry      time.Duration `json:"expiry"`
	}{
		URL:    "www.google.com",
		Expiry: 100,
	}

	var shortURLResponse CreateShortURLResponse

	reqBody, _ := json.Marshal(createURLMockData)
	newUrlReq, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(reqBody))
	newUrlReq.Header.Set("Content-Type", "application/json")
	newUrlresp, _ := server.Test(newUrlReq, -1)
	newUrlBody, _ := ioutil.ReadAll(newUrlresp.Body)
	err := json.Unmarshal(newUrlBody, &shortURLResponse)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if newUrlresp.StatusCode != 200 {
		fmt.Println("Error setting up URL:", err)
		return
	}

	// if newUrlresp.StatusCode == 200 {
	urlId := shortURLResponse.CustomShort
	// }

	// NOW TEST RESOLVING THE NEW URL WE CREATED
	server.Get("/:url", ResolveURL)

	url := os.Getenv("DOMAIN") + urlId

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := server.Test(req, -1)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(resp.StatusCode)
	t.Log(string(body))

	assert.Equal(t, 301, resp.StatusCode)

}

func TestResolveURLExpiredURL(t *testing.T) {

	server := fiber.New()
	// need to create an entry in the db first
	server.Post("/api/v1", ShortenURL)

	// CREATE A MOCK ENTRY WITH VERY SHORT EXPIRY TIME
	createURLMockData := struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"custom_short"`
		Expiry      time.Duration `json:"expiry"`
	}{
		URL:         "www.google.com",
		CustomShort: "testDomain2",
		Expiry:      1 / 3600,
	}

	reqBody, _ := json.Marshal(createURLMockData)
	newUrlReq, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(reqBody))
	newUrlReq.Header.Set("Content-Type", "application/json")
	server.Test(newUrlReq, -1)

	// NOW TEST RESOLVING THE NEW URL WE CREATED
	server.Get("/:url", ResolveURL)

	urlId := "/testDomain2"
	url := os.Getenv("DOMAIN") + urlId
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := server.Test(req, -1)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(resp.StatusCode)
	t.Log(string(body))
	// SHOULD B 404 NOT FOUND
	assert.Equal(t, 404, resp.StatusCode)

}
func TestResolveURLWithCustomShort(t *testing.T) {

	server := fiber.New()
	// need to create an entry in the db first
	server.Post("/api/v1", ShortenURL)

	// CREATE A MOCK ENTRY WITHOUT A CUSTOM SHORT
	createURLMockData := struct {
		URL         string        `json:"url"`
		CustomShort string        `json:"custom_short"`
		Expiry      time.Duration `json:"expiry"`
	}{
		URL:         "www.google.com",
		CustomShort: "testDomain2",
		Expiry:      100,
	}

	reqBody, _ := json.Marshal(createURLMockData)
	newUrlReq, _ := http.NewRequest(http.MethodPost, "/api/v1", bytes.NewBuffer(reqBody))
	newUrlReq.Header.Set("Content-Type", "application/json")
	server.Test(newUrlReq, -1)

	// NOW TEST RESOLVING THE NEW URL WE CREATED
	server.Get("/:url", ResolveURL)

	urlId := "/testDomain2"
	url := os.Getenv("DOMAIN") + urlId
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := server.Test(req, -1)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Log(resp.StatusCode)
	t.Log(string(body))

	assert.Equal(t, 301, resp.StatusCode)

}
