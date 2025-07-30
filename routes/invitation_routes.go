package routes

import (
	"be-undangan-digital/controllers"
	"be-undangan-digital/middleware"

	"github.com/gofiber/fiber/v2"
)

func InvitationRoutes(app *fiber.App) {
	routes := app.Group("/api/invitation")

	routes.Post("/guest_view", controllers.GuestView)
	routes.Get("/:id_invitation", controllers.GetInvitation)

	routes.Post("/create", middleware.JWTProtected(), controllers.CreateInvitation)
	routes.Delete("/delete/:id_invitation", middleware.JWTProtected(), controllers.DeleteInvitation)

	routes.Post("/generate_link", middleware.JWTProtected(), controllers.GenerateLink)
	routes.Post("/share_social_media", middleware.JWTProtected(), controllers.ShareSocialMedia)

	routes.Get("/", middleware.JWTProtected(), controllers.GetInvitations)
}
