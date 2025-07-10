package controllers

import (
	"be-evowish/lib"
	"be-evowish/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data lib.RegisterRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := lib.Validate.Struct(data); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		return c.Status(400).JSON(fiber.Map{"validation_error": errors})
	}

	_, err := services.RegisterUserService(data.Name, data.Email, data.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Registration successful",
	})
}
