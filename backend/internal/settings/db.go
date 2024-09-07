package settings

import "fmt"

type Postgres struct {
	Host     string `env:"HOST"`
	Name     string `env:"NAME"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

func (p Postgres) Connection() string {
	return fmt.Sprintf(
		"host=%s dbname=%s port=%d user=%s password=%s sslmode=require",
		p.Host, p.Name, p.Port, p.User, p.Password,
	)
}
