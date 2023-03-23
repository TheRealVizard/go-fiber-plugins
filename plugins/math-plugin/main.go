package main

import (
	"github.com/gofiber/fiber/v2"
)

func MathController(c *fiber.Ctx) error {
	return c.SendString("Math Rocks too!")
}

func GetFunctions() ([]string, []func(c *fiber.Ctx) error) {
	paths := []string{"/math"}
	functions := []func(c *fiber.Ctx) error{MathController}
	return paths, functions
}
