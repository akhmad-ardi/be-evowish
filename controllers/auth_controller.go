package controllers

import (
	"be-undangan-digital/lib"
	"be-undangan-digital/services"
	"be-undangan-digital/validations"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data lib.RegisterRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message_error": "invalid request"})
	}

	validation_errors := validations.ValidateRegisterRequest(data)
	if validation_errors != nil {
		return c.Status(400).JSON(fiber.Map{"validation_errors": validation_errors})
	}

	_, err := services.CreateUserService(data.Name, data.Email, data.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message_error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Akun berhasil dibuat",
	})
}

func Login(c *fiber.Ctx) error {
	var data lib.LoginRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message_error": "invalid request"})
	}

	validation_errors := validations.ValidateLoginRequest(data)
	if validation_errors != nil {
		return c.Status(400).JSON(fiber.Map{"validation_errors": validation_errors})
	}

	user, err := services.GetUserByField("email", data.Email)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message_error": "Pengguna tidak ditemukan"})
	}

	check_password := lib.CheckPasswordHash(data.Password, user.Password)

	if !check_password {
		return c.Status(400).JSON(fiber.Map{"message_error": "Email atau kata sandi salah"})
	}

	token, err := lib.GenerateJWT(user.IdUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message_error": "Something error"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Login berhasil", "token": token})
}
