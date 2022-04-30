package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err.Error())
	}
	return c.JSON(data)
}
