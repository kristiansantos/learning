package initializer

type Mongo struct {
	User     string `env:"mongo_user"`
	Pass     string `env:"mongo_pass"`
	Host     string `env:"mongo_host"`
	Args     string `env:"mongo_args"`
	Database string `env:"mongo_database"`
}
