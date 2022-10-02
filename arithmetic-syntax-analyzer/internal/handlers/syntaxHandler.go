package handlers

import (
	"arithmetic-syntax-analyzer/internal/lexical"
	"arithmetic-syntax-analyzer/internal/lexical/writers/formaters"
)

type syntaxHandler struct{}

func (h *syntaxHandler) Start(expression string) ([][]string, error) {
	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, err
	}
	tree := formaters.NewSyntaxTreeFormatter().Form(lexemes)
	return [][]string{tree}, nil
}
