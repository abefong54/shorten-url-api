package routes

import (
	"fmt"

	"github.com/abefong54/shorten-url/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// CHECK IF A SHORT URL HAS A CORRESPONDING LONG URL SAVED IN OUR DB
// IF SO, REDIRECT THE USER TO THAT LONG URL
func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ShortURL not found."})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not connect to DB"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	// CHECK EXPIRY

	counter := rInr.Incr(database.Ctx, "counter") // increment the use count of this url in the db
	fmt.Println("counter")
	fmt.Println(counter)

	// counterVal, _ = r.Get(database.Ctx, urlID).Result()
	// if val != "" {
	// 	// the value already exists
	// 	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
	// 		"error": "Custom short URL already in use.",
	// 	})
	// }
	return c.Redirect(value, 301)
}
