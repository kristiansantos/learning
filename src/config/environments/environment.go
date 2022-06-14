package environments

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/kristiansantos/learning/src/config/initializers"
)

var Instance *initializers.Initializer = nil

func ReadEnvironments(environment, version string) (initializers.Initializer, error) {
	var initializer initializers.Initializer

	if err := cleanenv.ReadEnv(&initializer); err != nil {
		return initializers.Initializer{}, fmt.Errorf(`error reading env: %w`, err)
	}

	initializer.Environment = environment
	initializer.Version = version

	Instance = &initializer

	return initializer, nil
}
