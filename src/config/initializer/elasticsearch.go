package initializer

type Elasticsearch struct {
	Url  string `env:"additional_information"`
	User string `env:"elasticsearch_user"`
	Pass string `env:"elasticsearch_pass"`
}
