package settings

type Postgres struct {
	Host     string
	Name     string
	Port     int
	User     string
	Password string
}

// TODO: implement the connection string
func (p Postgres) Connection() string {
	return ""
}
