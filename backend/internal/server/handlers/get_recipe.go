package handlers

import (
	"fmt"

	"github.com/GenerateNU/cooked/backend/internal/errs"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Service) GetRecipe(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return errs.BadRequest(fmt.Errorf("failed to parse id. got: %s", idParam))
	}

	recipe, err := s.store.GetRecipe(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(recipe)
}
