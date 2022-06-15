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

type Elasticsearch struct {
	Url  string `env:"additional_information"`
	User string `env:"elasticsearch_user"`
	Pass string `env:"elasticsearch_pass"`
}

type Mongo struct {
	User     string `env:"mongo_user"`
	Pass     string `env:"mongo_pass"`
	Host     string `env:"mongo_host"`
	Args     string `env:"mongo_args"`
	Database string `env:"mongo_database"`
}
