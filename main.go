package main

import (
	"be-undangan-digital/config"
	"be-undangan-digital/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat .env file")
	}

	config.ConnectDatabase()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Static("/public", "./public")

	// Main route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API for Undangan Digital",
		})
	})

	routes.AuthRoutes(app)
	routes.InvitationRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
