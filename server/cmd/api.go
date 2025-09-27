package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	scrape "github.com/sgainso/Cooklens/internal"
)

// AddURLWithRank adds a URL and its success rank to the database if the URL doesn't already exist
// Returns true if the URL was added, false if it already exists or if there was an error
func AddURLWithRank(url string, successRank bool) (bool, error) {
	// Check if URL already exists in the database
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE url = ?)", url).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if URL exists: %v", err)
	}

	// If URL already exists, return false
	if exists {
		return false, nil
	}
	var successRankVal int = 0
	if successRank {
		successRankVal = 1
	} else {
		successRankVal = 0
	}

	// URL doesn't exist, so add it to the database
	_, err = db.Exec("INSERT INTO urls (url, success_rank) VALUES (?, ?)", url, successRankVal)
	if err != nil {
		return false, fmt.Errorf("error adding URL to database: %v", err)
	}

	return true, nil
}

var db *sql.DB

func main() {

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = "jas"              // Replace with your actual MySQL username
	cfg.Passwd = "strongpassword" // Replace with your actual MySQL password
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "Cooklens"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error creating DB: ", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging DB: ", pingErr)
	}
	fmt.Println("Connected!")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, world!"})
	})
	// https://www.delish.com/cooking/recipe-ideas/a46330/skillet-sicilian-chicken-recipe/
	// https://www.foodnetwork.com/recipes/ground-turkey-enchilada-stir-fry-with-couscous-3416321
	app.Get("/scrape", func(c *fiber.Ctx) error {
		PassedInLink := c.Query("url")
		ings, recipe, succes := scrape.Scrape(PassedInLink)
		if succes == false {
			fmt.Println("Scraping Process Failed")
		}
		AddURLWithRank(PassedInLink, succes)
		fmt.Println("Scrape Ran!")
		return c.JSON(fiber.Map{
			"ingredients": ings,
			"recipe":      recipe,
		})
	})

	app.Listen(":3500")
}
