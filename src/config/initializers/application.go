package initializers

import "time"

type Application struct {
	ReadTimeout  time.Duration `env:"app_readTimeout"`
	WriteTimeout time.Duration `env:"app_writeTimeout"`
}
