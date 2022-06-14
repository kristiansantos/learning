package initializers

import "time"

type Application struct {
	ApiClientKey string        `env:"app_apiClientKey"`
	ApiKey       string        `env:"app_apiKey"`
	ReadTimeout  time.Duration `env:"app_readTimeout"`
	WriteTimeout time.Duration `env:"app_writeTimeout"`
	Jwt          Jwt
}
