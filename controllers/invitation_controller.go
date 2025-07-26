package controllers

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
	"be-undangan-digital/services"
	"be-undangan-digital/validations"
	"fmt"
	"net/http"
	"os"

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

	invitations, err := services.GetInvitations(IdUser)
	if err != nil {
		return lib.RespondError(c, http.StatusInternalServerError, "Query error")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"invitations": invitations,
	})
}

func GenerateLink(c *fiber.Ctx) error {
	req := requests.GenerateLinkRequest{
		IdInvitation: c.Params("id_invitation"),
	}

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateGenerateLinkRequest(req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	invitation, err := services.GetInvitation(req.IdInvitation)
	if err != nil {
		return lib.RespondError(c, http.StatusNotFound, err.Error())
	}

	baseURL := os.Getenv("FRONT_END")
	if baseURL == "" {
		return lib.RespondError(c, http.StatusInternalServerError, "Env front end belum diset")
	}

	link := fmt.Sprintf("%s/invitation/%s", baseURL, invitation.IdInvitation)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"link": link,
	})
}
