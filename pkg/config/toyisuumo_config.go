package config

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

type ToyIsuumoConfig struct {
	// LogLevel is INFO, DEBUG, or ERROR
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	// Port is the port number toy-isuumo will listen on
	Port int16 `envconfig:"SERVER_PORT" default:"8081"`

	// DBConfig is
	DBConfig DBConfig

	// ChairSearchCondition is
	ChairSearchCondition ChairSearchCondition
}

func ReadFromEnv() (*ToyIsuumoConfig, error) {
	var cfg ToyIsuumoConfig
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, xerrors.Errorf("envconfig failed to read environment variables: %w", err)
	}
	cond, err := getChairSearchCondition()
	if err != nil {
		return nil, xerrors.Errorf("failed to retrive chair search condition: %w", err)
	}
	cfg.ChairSearchCondition = *cond

	return &cfg, nil
}
