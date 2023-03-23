package main

import (
	"github.com/gofiber/fiber/v2"
)

func HistoryController(c *fiber.Ctx) error {
	return c.SendString("History Rocks too!")
}

func GetFunctions() ([]string, []func(c *fiber.Ctx) error) {
	paths := []string{"/history"}
	functions := []func(c *fiber.Ctx) error{HistoryController}
	return paths, functions
}
