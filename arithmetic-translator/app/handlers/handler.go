package handlers

import (
	"arithmetic-translator/app/handlers/config"
	"arithmetic-translator/app/handlers/enum"
)

type Handler interface {
	Start() error
}

func NewHandler(config *config.Config) (Handler, error) {
	switch config.Mode {
	case enum.Translation:
		return NewTranslator(config.Setting)
	case enum.Generation:
		return NewGenerator(config.Setting)
	}
	return nil, nil
}
