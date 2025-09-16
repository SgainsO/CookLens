package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	scrape "github.com/sgainso/Cooklens/internal"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, world!"})
	})

	app.Get("/scrape", func(c *fiber.Ctx) error {
		ings, recipe, succes := scrape.Scrape("https://www.recipetineats.com/mexican-corn-salad/")
		if succes == false {
			fmt.Println("Ditionary Loading Error")
		}
		fmt.Println("Scrape Ran!")
		return c.JSON(fiber.Map{
			"ingredients": ings,
			"recipe":      recipe,
		})
	})
	app.Listen(":8080")
}
