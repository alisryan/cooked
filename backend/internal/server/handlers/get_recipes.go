package handlers

import "github.com/gofiber/fiber/v2"

func (s *Service) GetRecipes(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = c.QueryInt("limit", 10)
	)

	recipes, err := s.store.GetRecipes(c.Context(), page, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(recipes)
}
