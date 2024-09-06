package settings

type Settings struct {
	Application
	Postgres
}

// TOOD: implement loading the environment variables
func Load() (Settings, error) {
	return Settings{}, nil
}
