package controllers

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
	"be-undangan-digital/services"
	"be-undangan-digital/validations"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateInvitation(c *fiber.Ctx) error {
	IdUser, err := lib.GetUserIDFromContext(c)
	if err != nil {
		return err
	}

	req := new(requests.CreateInvitationRequest)
	req.Title = c.FormValue("title")
	req.Date = c.FormValue("date")
	req.Time = c.FormValue("time")
	req.Location = c.FormValue("location")
	req.Description = c.FormValue("description")
	req.IdTemplate = c.FormValue("id_template")
	req.PrimaryColor = c.FormValue("primary_color")
	req.SecondaryColor = c.FormValue("secondary_color")
	req.BackgroundImage = c.FormValue("background_image")

	if err := c.BodyParser(req); err != nil {
		println(err.Error())
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateCreateInvitationRequest(*req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	file, err := c.FormFile("background_image")
	if err == nil {
		filename, err := lib.UploadImageFile(file, "public")
		if err != nil {
			return lib.RespondError(c, fiber.StatusInternalServerError, err.Error())
		}
		req.BackgroundImage = filename
	}

	if _, err := services.CreateInvitationService(IdUser, req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Undangan berhasil dibuat",
	})
}

func GetInvitations(c *fiber.Ctx) error {
	IdUser, err := lib.GetUserIDFromContext(c)
	if err != nil {
		return err
	}

	invitations, err := services.GetInvitationsService(IdUser)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, "Query error")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"invitations": invitations,
	})
}

func GenerateLink(c *fiber.Ctx) error {
	var req requests.GenerateLinkRequest

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateGenerateLinkRequest(req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	invitation, err := services.GetInvitationService(req.IdInvitation)
	if err != nil {
		return lib.RespondError(c, http.StatusNotFound, err.Error())
	}

	link, err := lib.GenerateInvitationLink(invitation.IdInvitation)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	invitation_link, err := services.CreateInvitationLink(invitation.IdInvitation, link)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"link": invitation_link.Link,
	})
}

func ShareSocialMedia(c *fiber.Ctx) error {
	var req requests.ShareSocialMediaRequest

	// Parse request
	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	// Validasi request
	validationErrors := validations.ValidateShareSocialMediaRequest(req)
	if validationErrors != nil {
		return lib.RespondValidationError(c, validationErrors)
	}

	// Cek apakah sudah pernah dibagikan ke platform ini
	if _, err := services.GetSharedSocialService(req.IdInvitation, req.NamePlatform); err == nil {
		return lib.RespondError(c, http.StatusBadRequest, "Sudah pernah share di "+req.NamePlatform)
	}

	// Ambil atau buat link undangan
	invitationLink, err := services.GetInvitationLink(req.IdInvitation)
	if err != nil || invitationLink == nil {
		// Jika belum ada, generate dan simpan link
		link, err := lib.GenerateInvitationLink(req.IdInvitation)
		if err != nil {
			return lib.RespondError(c, http.StatusInternalServerError, err.Error())
		}

		invitationLink, err = services.CreateInvitationLink(req.IdInvitation, link)
		if err != nil {
			return lib.RespondError(c, http.StatusInternalServerError, err.Error())
		}
	}

	// Simpan data share ke platform
	if _, err := services.CreateSharedSocialService(invitationLink.IdInvitationLink, req.NamePlatform); err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	// Berhasil
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"link": invitationLink.Link,
	})
}
