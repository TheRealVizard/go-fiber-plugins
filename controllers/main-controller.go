package controllers

import "github.com/gofiber/fiber/v2"

func ScienceController(c *fiber.Ctx) error {
	return c.SendString("Science Rocks!")
}
