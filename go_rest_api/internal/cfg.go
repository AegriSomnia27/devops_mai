package server

type AppConfig struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port uint16 `envconfig:"PORT" default:"1337"`
}
