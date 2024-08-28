package settings

type DB interface {
	Connection() string
}

type Postgres struct{}

// TODO: implement the connection string
func (p Postgres) Connection() string {
	return ""
}
