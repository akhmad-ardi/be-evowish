package main

import (
	"be-evowish/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat .env file")
	}

	app := fiber.New()
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Hello World",
		})
	})
	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
