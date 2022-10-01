package handlers

import (
	"arithmetic-syntax-analyzer/internal/config"
)

type Handler interface {
	Start(expression string) ([][]string, error)
	//OutResult(files []string) error
}

func NewHandler(config2 config.Config) Handler {
	if config2.Mode == config.Lexical {
		return newLexicalAnalyzer()
	}
	if config2.Mode == config.Syntax {
		return newSyntaxAnalyzer()
	}
	return nil
}
