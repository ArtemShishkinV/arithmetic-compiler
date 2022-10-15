package handlers

import (
	"arithmetic-compiler/internal/lexical"
	"arithmetic-compiler/internal/lexical/writers/formaters"
)

type lexicalHandler struct{}

func (h *lexicalHandler) Start(expression string) ([][]string, error) {
	handler := lexical.NewLexicalAnalyzer()

	lexemes, err := handler.Analyze(expression)
	if err != nil {
		return nil, err
	}
	tokens := lexical.NewTokenBuilder().GetTokens(lexemes)
	tokensOut := formaters.NewTokensFormatter(tokens).Form()
	vars := formaters.NewVarsFormatter(tokens).Form()
	return [][]string{tokensOut, vars}, nil
}
