package controllers

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
	"be-undangan-digital/services"
	"be-undangan-digital/validations"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

func GetTemplates(c *fiber.Ctx) error {
	_, errGetUserID := lib.GetUserIDFromContext(c)
	if errGetUserID != nil {
		return lib.RespondError(c, http.StatusUnauthorized, errGetUserID.Error())
	}

	templates, errGetTemplates := services.GetTemplatesService()
	if errGetTemplates != nil {
		return lib.RespondError(c, http.StatusInternalServerError, errGetTemplates.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"templates": templates,
	})
}

func CreateInvitation(c *fiber.Ctx) error {
	IdUser, errGetUserID := lib.GetUserIDFromContext(c)
	if errGetUserID != nil {
		return lib.RespondError(c, http.StatusUnauthorized, errGetUserID.Error())
	}

	var req requests.CreateInvitationRequest

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateCreateInvitationRequest(req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	invitation, errCreateInvitation := services.CreateInvitationService(IdUser, &req)
	if errCreateInvitation != nil {
		return lib.RespondError(c, http.StatusBadRequest, errCreateInvitation.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message":       "Undangan berhasil dibuat",
		"id_invitation": invitation.IdInvitation,
	})
}

func AddDataInvitation(c *fiber.Ctx) error {
	_, errGetUserID := lib.GetUserIDFromContext(c)
	if errGetUserID != nil {
		return lib.RespondError(c, http.StatusUnauthorized, errGetUserID.Error())
	}

	var req requests.AddDataInvitation

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateAddDataInvitation(req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	jsonData, _ := json.Marshal(req.DataInvitation)
	updates := map[string]interface{}{
		"data_invitation": datatypes.JSON(jsonData),
	}

	_, errUpdateInvitation := services.UpdateInvitationService(c.Params("id_invitation"), updates)
	if errUpdateInvitation != nil {
		return lib.RespondError(c, http.StatusInternalServerError, errUpdateInvitation.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Undangan siap dibagikan",
	})
}

func AddBackgroundImage(c *fiber.Ctx) error {
	_, errGetUserID := lib.GetUserIDFromContext(c)
	if errGetUserID != nil {
		return lib.RespondError(c, http.StatusUnauthorized, errGetUserID.Error())
	}

	file, errFormFile := c.FormFile("bg_image")
	if errFormFile != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Gagal membaca file")
	}

	req := requests.AddBackgroundImageRequest{
		File: file,
	}

	validation_errors := validations.ValidateAddBackgroundImageRequest(req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	filename, errUploadImageFile := lib.UploadImageFile(req.File, "public")
	if errUploadImageFile != nil {
		return lib.RespondError(c, http.StatusInternalServerError, errUploadImageFile.Error())
	}

	updates := map[string]interface{}{
		"background_image": filename,
	}

	_, errUpdate := services.UpdateInvitationService(c.Params("id_invitation"), updates)
	if errUpdate != nil {
		return lib.RespondError(c, http.StatusInternalServerError, errUpdate.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Latar belakang berhasil ditambahkan",
	})
}

func GetInvitations(c *fiber.Ctx) error {
	IdUser, err := lib.GetUserIDFromContext(c)
	if err != nil {
		return lib.RespondError(c, http.StatusUnauthorized, err.Error())
	}

	invitations, err := services.GetInvitationsService(IdUser)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"invitations": invitations,
	})
}

func GetInvitation(c *fiber.Ctx) error {
	id_invitation := c.Params("id_invitation")

	invitation, err := services.GetInvitationService(id_invitation)
	if err != nil {
		return lib.RespondError(c, http.StatusNotFound, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"invitation": invitation,
	})
}

func DeleteInvitation(c *fiber.Ctx) error {
	id_invitation := c.Params("id_invitation")

	invitation, err := services.GetInvitationService(id_invitation)
	if err != nil {
		return lib.RespondError(c, http.StatusNotFound, err.Error())
	}

	_, err = services.DeleteInvitationService(invitation.IdInvitation)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Undangan berhasil dihapus",
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
		"message": "Undangan link berhasil dibuat",
		"link":    invitation_link.Link,
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

func GuestView(c *fiber.Ctx) error {
	var req requests.GuestViewRequest

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateGuestViewRequest(req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	_, err := services.CreateGuestViewService(req)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": req.UserAgent + " telah melihat undangan",
	})
}
