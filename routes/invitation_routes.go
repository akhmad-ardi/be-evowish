package routes

import (
	"be-undangan-digital/controllers"
	"be-undangan-digital/middleware"

	"github.com/gofiber/fiber/v2"
)

func InvitationRoutes(app *fiber.App) {
	routes := app.Group("/api/invitation")

	routes.Post("/guest_view", controllers.GuestView)

	routes.Get("/templates", middleware.JWTProtected(), controllers.GetTemplates)

	routes.Get("/:id_invitation", controllers.GetInvitation)

	routes.Post("/create_invitation", middleware.JWTProtected(), controllers.CreateInvitation)
	routes.Post("/add_data_invitation/:id_invitation", middleware.JWTProtected(), controllers.AddDataInvitation)
	routes.Post("/add_background_image/:id_invitation", middleware.JWTProtected(), controllers.AddBackgroundImage)
	routes.Delete("/delete_invitation/:id_invitation", middleware.JWTProtected(), controllers.DeleteInvitation)

	routes.Post("/generate_link", middleware.JWTProtected(), controllers.GenerateLink)
	routes.Post("/share_social_media", middleware.JWTProtected(), controllers.ShareSocialMedia)

	routes.Get("/", middleware.JWTProtected(), controllers.GetInvitations)
}
