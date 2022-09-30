package config

import "errors"

type Config struct {
	SrcFileName        string
	OutTokensFileName  string
	OutSymbolsFileName string
}

func NewConfig(args []string) (*Config, error) {
	if len(args) != 3 {
		return nil, errors.New("invalid count of arguments")
	}
	return &Config{
		SrcFileName:        args[0],
		OutTokensFileName:  args[1],
		OutSymbolsFileName: args[2],
	}, nil
}
