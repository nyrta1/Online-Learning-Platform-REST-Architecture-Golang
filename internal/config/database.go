package config

type Database struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	Sslmode  string `env:"POSTGRES_SSLMODE"`
	Name     string `env:"POSTGRES_NAME"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}
