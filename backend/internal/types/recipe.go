package types

import (
	"github.com/google/uuid"
)

type Recipe struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Cook         Duration  `json:"cook_duration"`
	Instructions string    `json:"instructions"`
	ImageURL     string    `json:"image_url"`
	Meal         Meal      `json:"meal"`
}

type Meal string

const (
	MealBreakfast Meal = "breakfast"
	MealLunch     Meal = "lunch"
	MealDinner    Meal = "dinner"
	MealSnack     Meal = "snack"
)
