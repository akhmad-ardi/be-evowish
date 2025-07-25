package tests

import (
	"be-undangan-digital/config"
	"be-undangan-digital/models"
	"be-undangan-digital/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SetupApp() (*fiber.App, error) {
	// Load ENV khusus test jika perlu
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	config.ConnectDatabase()

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Hello World",
		})
	})

	routes.AuthRoutes(app)

	return app, nil
}

func DeleteDataUser(email *string) error {
	var user *models.User

	result := config.DB.Where("email = ?", *email).Delete(&user)
	return result.Error
}
