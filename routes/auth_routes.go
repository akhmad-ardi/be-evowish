package routes

import (
	"be-undangan-digital/controllers"
	"be-undangan-digital/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	routes := app.Group("/api/auth")

	routes.Post("/register", controllers.Register)
	routes.Post("/login", controllers.Login)

	routes.Get("/", middleware.JWTProtected(), controllers.CheckAuth)
}
