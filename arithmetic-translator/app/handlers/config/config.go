package config

import (
	"arithmetic-translator/app/handlers/enum"
	"errors"
)

type Config struct {
	Mode    enum.Mode
	Setting []string
}

func NewConfig(settings []string) (*Config, error) {
	if len(settings) == 0 {
		return nil, errors.New("invalid number of arguments")
	}

	mode, err := enum.GetMode(settings[0])

	if err != nil {
		return nil, err
	}

	return &Config{
		Mode:    mode,
		Setting: settings[1:],
	}, nil
}
