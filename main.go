package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abefong54/shorten-url/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	fmt.Println("starting server")

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	app.Use(analytics.Analytics(os.Getenv("ANALYTICS_KEY"))) // Add analytics middleware

	// Add rate limiting middleware
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(logger.New())
	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Use a default port if not set
	}

	fmt.Println("DOMAIN: ", os.Getenv("DOMAIN"))
	fmt.Println("APP_PORT: ", os.Getenv("APP_PORT"))
	fmt.Println("API_QUOTA: ", os.Getenv("API_QUOTA"))
	fmt.Println("ANALYTICS_KEY: ", os.Getenv("ANALYTICS_KEY"))
	fmt.Println("LOCAL_DOMAIN: ", os.Getenv("LOCAL_DOMAIN"))
	fmt.Println("APP_PORT: ", os.Getenv("APP_PORT"))
	fmt.Println("DB_ADDRESS: ", os.Getenv("DB_ADDRESS"))
	fmt.Println("DB_PASS: ", os.Getenv("DB_PASS"))

	app.Listen(port)
}
