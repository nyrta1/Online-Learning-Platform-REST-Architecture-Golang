package config

type Session struct {
	Secret string `env:"SESSION_SECRET"`
	Name   string `env:"SESSION_NAME"`
	Key    string `env:"SESSION_KEY"`
}
