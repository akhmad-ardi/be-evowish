package routes

import (
	"be-undangan-digital/controllers"
	"be-undangan-digital/middleware"

	"github.com/gofiber/fiber/v2"
)

func InvitationRoutes(app *fiber.App) {
	routes := app.Group("/api/invitation")

	routes.Post("/create", middleware.JWTProtected(), controllers.CreateInvitation)
	routes.Get("/", middleware.JWTProtected(), controllers.GetInvitations)

}
