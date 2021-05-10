package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type ToyIsuumoConfig struct {
	// LogLevel is INFO, DEBUG, or ERROR
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	// Port is the port number toy-isuumo will listen on
	Port int16 `envconfig:"SERVER_PORT" default:"8081"`

	// DBConfig is
	DBConfig DBConfig
}

func ReadFromEnv() (*ToyIsuumoConfig, error) {
	var cfg ToyIsuumoConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, errors.Wrap(err, "envconfig failed to read environment variables")
	}

	return &cfg, nil
}
