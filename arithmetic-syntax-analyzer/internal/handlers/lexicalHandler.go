package handlers

import (
	"arithmetic-syntax-analyzer/internal/lexical"
	"arithmetic-syntax-analyzer/internal/lexical/writers/formaters"
)

type lexicalHandler struct{}

func (h *lexicalHandler) Start(expression string) ([][]string, error) {
	handler := lexical.NewLexicalAnalyzer()

	lexemes, err := handler.Analyze(expression)
	if err != nil {
		return nil, err
	}

	tokens := formaters.NewTokensFormatter().Form(lexemes)
	vars := formaters.NewVarsFormatter().Form(lexemes)
	return [][]string{tokens, vars}, nil
}
