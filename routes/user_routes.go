package routes

import (
	"be-evowish/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	routes := app.Group("/api/user")

	routes.Post("/register", controllers.Register)
}
