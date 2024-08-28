package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) Hello(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Hello %s!", c.Params("name")))
}
