package config

type RedisSessionConfig struct {
	ConnectionTimeoutSeconds string `env:"REDIS_CONNECTION_TIMEOUT_SECONDS"`
	NetworkType              string `env:"REDIS_NETWORK_TYPE"`
	Host                     string `env:"REDIS_HOST"`
	Port                     string `env:"REDIS_PORT"`
	Password                 string `env:"REDIS_PASSWORD"`
	DB                       int    `env:"REDIS_DB"`
	Session                  Session
}
