package handlers

import (
	"arithmetic-compiler/internal/config"
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
	if config2.Mode == config.Semantic {
		return &semanticHandler{}
	}

	if config2.Mode == config.Generator1 {
		return &generatorCodeHandler{}
	}

	if config2.Mode == config.Generator2 {
		return &generatorPostfixHandler{}
	}

	return nil
}
