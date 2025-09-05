package main
import (
	"github.com/gofiber/fiber/v2"
	"github.com/sgainso/Cooklens/internal/scrape"
)

func main() {
    app := fiber.New()
    app.Get("/hello", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello, world!"})
    })
    app.Listen(":8080")

    app.Get("/scrape", func(c *fiber.Ctx) error {
        ings, recipe := scrape.Scrape(c.Query("https://www.recipetineats.com/mexican-corn-salad/"))
        return c.JSON(fiber.Map{
            "ingredients": ings,
            "recipe":      recipe,
        })
    })
}
