package handlers

import "arithmetic-lexical-analyzer/internal/config"

type Handler interface {
	Start() error
}

func NewHandler(config *config.Config) Handler {
	return NewLexicalAnalyzer(config)
}
