package controllers

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/requests"
	"be-undangan-digital/services"
	"be-undangan-digital/validations"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var req requests.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	if errs := validations.ValidateRegisterRequest(req); errs != nil {
		return lib.RespondValidationError(c, errs)
	}

	if _, err := services.CreateUserService(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Akun berhasil dibuat",
	})
}

func Login(c *fiber.Ctx) error {
	var req requests.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return lib.RespondError(c, http.StatusBadRequest, "Permintaan tidak valid")
	}

	if errs := validations.ValidateLoginRequest(req); errs != nil {
		return lib.RespondValidationError(c, errs)
	}

	user, err := services.GetUserByField("email", req.Email)
	if err != nil {
		return lib.RespondError(c, http.StatusNotFound, "Pengguna tidak ditemukan")
	}

	if !lib.CheckPasswordHash(req.Password, user.Password) {
		return lib.RespondError(c, http.StatusUnauthorized, "Email atau kata sandi salah")
	}

	token, err := lib.GenerateJWT(user.IdUser)
	if err != nil {
		println(err)
		return lib.RespondError(c, http.StatusInternalServerError, "Terjadi kesalahan saat menghasilkan token")
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
	})
}

func CheckAuth(c *fiber.Ctx) error {
	IdUser, errGetUserID := lib.GetUserIDFromContext(c)
	if errGetUserID != nil {
		return lib.RespondError(c, http.StatusUnauthorized, errGetUserID.Error())
	}

	_, errGetUser := services.GetUserByField("id_user", IdUser)
	if errGetUser != nil {
		return lib.RespondError(c, http.StatusUnauthorized, errGetUser.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"is_auth": true,
	})
}
