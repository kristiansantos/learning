package environment

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kristiansantos/learning/src/config/initializer"
)

var Instance *initializer.Application = nil

func ReadEnvironments(environment, version string) (initializer.Application, error) {
	var app initializer.Application

	if err := cleanenv.ReadEnv(&app); err != nil {
		return initializer.Application{}, fmt.Errorf(`error reading env: %w`, err)
	}

	app.Environment = environment
	app.Version = version

	Instance = &app

	return app, nil
}
