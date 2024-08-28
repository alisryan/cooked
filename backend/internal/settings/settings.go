package settings

import "log/slog"

type Settings struct {
	Application
	DB
}

// TOOD: implement loading the environment variables
func Load(path string) (Settings, error) {
	slog.Info("loading settings", "path", path)

	return Settings{}, nil
}
