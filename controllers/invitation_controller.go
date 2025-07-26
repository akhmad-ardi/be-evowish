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
	LocalsIdUser := c.Locals("id_user")
	IdUser, ok := LocalsIdUser.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized: ID user tidak valid",
		})
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

	if err := c.BodyParser(req); err != nil {
		println(err.Error())
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	validation_errors := validations.ValidateCreateInvitationRequest(*req)
	if validation_errors != nil {
		return lib.RespondValidationError(c, validation_errors)
	}

	file, err := c.FormFile("background_image")
	if err != nil {
		println("Tidak ada gambar")
		req.BackgroundImage = "" // jika tidak ada file
	} else {
		filename, err := lib.UploadImageFile(file, "uploads")
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
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Undangan berhasil didapat",
	})
}
