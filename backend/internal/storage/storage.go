package storage

import (
	"context"

	"github.com/GenerateNU/cooked/backend/internal/types"
	"github.com/google/uuid"
)

type Storage interface {
	Ping() error
	Reciper
}

type Reciper interface {
	CreateRecipe(ctx context.Context, recipe types.Recipe) (types.Recipe, error)
	CreateRecipes(ctx context.Context, recipes []types.Recipe) ([]types.Recipe, error)
	GetRecipe(ctx context.Context, id uuid.UUID) (types.Recipe, error)
	GetRecipes(ctx context.Context, page int, limit int) ([]types.Recipe, error)
}
