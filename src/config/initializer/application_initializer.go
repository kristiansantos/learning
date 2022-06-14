package initializer

import "time"

type Application struct {
	Environment   string
	Version       string
	Mongo         Mongo
	Elasticsearch Elasticsearch
	ReadTimeout   time.Duration `env:"app_readTimeout"`
	WriteTimeout  time.Duration `env:"app_writeTimeout"`
}
