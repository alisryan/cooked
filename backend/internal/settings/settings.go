package settings

import (
	"github.com/caarlos0/env/v11"
)

type Settings struct {
	Application `envPrefix:"APP_"`
	Postgres    `envPrefix:"DB_"`
}

func Load() (Settings, error) {
	return env.ParseAs[Settings]()
}
