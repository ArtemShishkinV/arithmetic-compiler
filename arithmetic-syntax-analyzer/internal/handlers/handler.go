package handlers

import (
	"arithmetic-syntax-analyzer/internal/config"
)

type Handler interface {
	Start(expression string) ([][]string, error)
}

func NewHandler(config2 config.Config) Handler {
	if config2.Mode == config.Lexical {
		return &lexicalHandler{}
	}
	if config2.Mode == config.Syntax {
		return &syntaxHandler{}
	}
	return nil
}
