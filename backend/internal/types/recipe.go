package types

import (
	"github.com/google/uuid"
)

type Recipe struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Cook         Duration  `json:"cook_duration" db:"cook_duration"`
	Instructions string    `json:"instructions"`
	ImageURL     URL       `json:"image_url" db:"image_url"`
	Meal         Meal      `json:"meal"`
}

type Meal string

const (
	MealBreakfast Meal = "breakfast"
	MealLunch     Meal = "lunch"
	MealDinner    Meal = "dinner"
	MealSnack     Meal = "snack"
)

var Meals = map[Meal]struct{}{
	MealBreakfast: {},
	MealLunch:     {},
	MealDinner:    {},
	MealSnack:     {},
}
