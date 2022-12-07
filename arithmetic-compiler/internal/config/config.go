package config

import (
	"errors"
	"strings"
)

const textError = "invalid count of arguments"

type Config struct {
	Mode     Mode
	Files    []string
	Optimize bool
}

func NewConfig(args []string) (*Config, error) {
	if len(args) <= 2 {
		return nil, errors.New(textError)
	}

	mode, err := GetMode(args[0])
	if err != nil {
		return nil, err
	}

	optimize := false
	startFileIndex := 1

	if strings.EqualFold(args[1], "opt") {
		optimize = true
		startFileIndex = 2
	}

	if optimize && mode != Generator1 && mode != Generator2 {
		return nil, errors.New("invalid second param")
	}

	cfg := &Config{
		Mode:     mode,
		Files:    args[startFileIndex:],
		Optimize: optimize,
	}

	if err := cfg.checkValid(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) checkValid() error {
	if c.isValidMode() {
		return errors.New(textError)
	}
	return nil
}

func (c *Config) isValidMode() bool {
	if c.Mode == Lexical || c.Mode == Generator1 || c.Mode == Generator2 {
		return len(c.Files) != 3
	}
	if c.Mode == Syntax || c.Mode == Semantic {
		return len(c.Files) != 2
	}
	return false
}
