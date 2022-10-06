package config

import (
	"errors"
)

const textError = "invalid count of arguments"

type Config struct {
	Mode  Mode
	Files []string
}

func NewConfig(args []string) (*Config, error) {
	if len(args) == 0 {
		return nil, errors.New(textError)
	}

	mode, err := GetMode(args[0])
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Mode:  mode,
		Files: args[1:],
	}

	if err := cfg.checkValid(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) checkValid() error {
	if c.Mode == Lexical && len(c.Files) != 3 ||
		c.Mode == Syntax && len(c.Files) != 2 {
		return errors.New(textError)
	}
	return nil
}
