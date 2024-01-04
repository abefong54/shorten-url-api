package routes

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/abefong54/shorten-url/database"
	"github.com/abefong54/shorten-url/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Format structure for request and response
type ShortenURLRequest struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}
type ShortenURLResponse struct {
	URL             string        `json:"url`
	CustomShort     string        `json:"custom_short`
	Expiry          time.Duration `json:"expiry`
	XRateRemaining  int           `json:"rate_limit`
	XRateLimitReset time.Duration `json:"rate_limit_reset`
}

// function that takes in a URL
// and generates a shorter string version
func ShortenURL(c *fiber.Ctx) error {

	body := new(ShortenURLRequest)
	if err := c.BodyParser(&body); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not parse JSON"})
	}

	// implement rate limiting,

	// check the IP of client
	// if the user already has used this service, (we have the ip) then we decrement the rate we allow
	r2 := database.CreateClient(1)
	defer r2.Close() // closes the client after the end of the call stack

	// find the ip address in our db using c.IP
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		// if you didn't find a value, we've never used the service before, so set the limit to 30 minutes
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		// we found the user, decrement user rate count
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			// exceeded rate limit
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":            "Rate limit exceeded. Please try again later!",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
	}

	// check if input is a real URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// check for domain errors (is url accessible publically)
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Domain error!"})
	}

	// enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	// CREATE SHORT URL ID
	var urlID string
	if body.CustomShort == "" {
		// create a random id using UID package
		urlID = uuid.New().String()[:6]
	} else {
		// use custom value if provided by user
		urlID = body.CustomShort
	}

	// validate custom url is unique
	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, urlID).Result()
	if val != "" {
		// the value already exists
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Custom short URL already in use.",
		})
	}

	// SET EXPIRY IF NOT PROVIDED BY USER
	fmt.Println("expiry sent:")
	fmt.Println(body.Expiry)
	if body.Expiry == 0 {
		fmt.Println("expiry not set")
		body.Expiry = 24
	}

	// SAVE VALUES TO DB
	err = r.Set(database.Ctx, urlID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	rate, _ := strconv.Atoi(os.Getenv("API_QUOTA"))
	resp := ShortenURLResponse{
		URL:             body.URL,
		CustomShort:     "",
		Expiry:          body.Expiry,
		XRateRemaining:  rate,
		XRateLimitReset: 30,
	}
	// DECREMENT IP usage each time we call this function
	r2.Decr(database.Ctx, c.IP())

	//GET remaining time for this user
	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	// reset the TTL (time to live) since we creating a new one
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	// build the shortened URL value
	resp.CustomShort = os.Getenv("DOMAIN") + "/" + urlID
	return c.Status(fiber.StatusOK).JSON(resp)

}
