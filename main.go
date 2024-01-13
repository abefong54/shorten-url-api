package main

import (
	"fmt"
	"os"

	"github.com/abefong54/shorten-url/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	analytics "github.com/tom-draper/api-analytics/analytics/go/fiber"
)

// list of all available API routes
func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

// APP variable is an instance of the Fiber library
// we will pass it around
func main() {

	fmt.Println("STARTING SERVER")

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	app.Use(analytics.Analytics(os.Getenv("ANALYTICS_KEY")))                             // Add analytics middleware
	app.Get("/metrics", monitor.New(monitor.Config{Title: "URL Shortner Metrics Page"})) // Initialize default middleware settings to /metrics

	app.Use(logger.New())
	setupRoutes(app)

	local := os.Getenv("LOCAL")
	port := ""
	if local == "true" {
		port = "8080" // Use a default port if not set
	} else {
		port = os.Getenv("PORT")
	}
	fmt.Println("STARTED SERVER")

	app.Listen(":" + port)

}
