package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/GenerateNU/cooked/backend/internal/errs"
	"github.com/GenerateNU/cooked/backend/internal/settings"
	"github.com/GenerateNU/cooked/backend/internal/types"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func New(settings settings.Postgres) *DB {
	return &DB{db: sqlx.MustConnect("postgres", settings.Connection())}
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) CreateRecipe(ctx context.Context, recipe types.Recipe) (types.Recipe, error) {
	if _, err := db.db.ExecContext(
		ctx,
		`INSERT INTO recipes (id, name, cook_duration, instructions, image_url, meal) VALUES ($1, $2, $3, $4, $5, $6)`,
		recipe.ID, recipe.Name, recipe.Cook, recipe.Instructions, recipe.ImageURL, recipe.Meal,
	); err != nil {
		if db.isUniqueViolation(err) {
			return types.Recipe{}, errs.Conflict("recipe", "name", recipe.Name)
		}
		return types.Recipe{}, err
	}
	return recipe, nil
}

func (db *DB) CreateRecipes(ctx context.Context, recipes []types.Recipe) ([]types.Recipe, error) {
	var values strings.Builder
	for i, recipe := range recipes {
		values.WriteString("('")
		values.WriteString(recipe.ID.String())
		values.WriteString("', '")
		values.WriteString(recipe.Name)
		values.WriteString("', '")
		values.WriteString(fmt.Sprintf("%d", recipe.Cook.Into()))
		values.WriteString("', '")
		values.WriteString(recipe.Instructions)
		values.WriteString("', '")
		url := recipe.ImageURL.Into()
		values.WriteString(url.String())
		values.WriteString("', '")
		values.WriteString(string(recipe.Meal))
		values.WriteString("')")
		if i != len(recipes)-1 {
			values.WriteString(", ")
		}
	}

	if _, err := db.db.ExecContext(
		ctx,
		`INSERT INTO recipes (id, name, cook_duration, instructions, image_url, meal) VALUES`+values.String()+` ON CONFLICT DO NOTHING`,
	); err != nil {
		return []types.Recipe{}, err
	}
	return recipes, nil
}

func (db *DB) GetRecipe(ctx context.Context, id uuid.UUID) (types.Recipe, error) {
	var recipe types.Recipe
	if err := db.db.GetContext(ctx, &recipe, `SELECT * FROM recipes WHERE id = $1`, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.Recipe{}, errs.NotFound("recipe", "id", id)
		}
		return types.Recipe{}, err
	}

	return recipe, nil
}

func (db *DB) GetRecipes(ctx context.Context, page int, limit int) ([]types.Recipe, error) {
	var recipes []types.Recipe
	if err := db.db.SelectContext(ctx, &recipes, `SELECT * FROM recipes LIMIT $1 OFFSET $2`, limit, (page-1)*limit); err != nil {
		return nil, err
	}

	return recipes, nil
}

func (db *DB) isUniqueViolation(err error) bool {
	pgErr, isPGError := err.(*pq.Error)
	return isPGError && pgErr.Code == "23505"
}
