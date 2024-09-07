package handlers

import (
	"fmt"

	"github.com/GenerateNU/cooked/backend/internal/errs"
	"github.com/GenerateNU/cooked/backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type createRecipeRequest struct {
	Name         string         `json:"name"`
	Cook         types.Duration `json:"cook_duration"`
	Instructions string         `json:"instructions"`
	ImageURL     types.URL      `json:"image_url"`
	Meal         types.Meal     `json:"meal"`
}

func (s *Service) CreateRecipe(c *fiber.Ctx) error {
	var req createRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return errs.InvalidJSON()
	}

	if validationErrs := req.validate(); len(validationErrs) > 0 {
		return errs.InvalidRequestData(validationErrs)
	}

	recipe, err := s.store.CreateRecipe(c.Context(), req.into())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(recipe)

}

func (r *createRecipeRequest) validate() map[string]string {
	errs := make(map[string]string)

	if r.Name == "" {
		errs["name"] = "name is required"
	}

	if r.Cook == 0 {
		errs["cook_duration"] = "cook_duration is required"
	}

	if r.Instructions == "" {
		errs["instructions"] = "instructions is required"
	}

	if r.Meal == "" {
		errs["meal"] = "meal is required"
	}

	if _, ok := types.Meals[r.Meal]; !ok {
		errs["meal"] = fmt.Sprintf("meal is invalid got: '%s'", r.Meal)
	}

	return errs
}

func (r *createRecipeRequest) into() types.Recipe {
	return types.Recipe{
		ID:           uuid.New(),
		Name:         r.Name,
		Cook:         r.Cook,
		Instructions: r.Instructions,
		ImageURL:     r.ImageURL,
		Meal:         r.Meal,
	}
}
