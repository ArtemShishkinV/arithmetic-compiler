package handlers

import (
	"arithmetic-compiler/internal/lexical"
	"arithmetic-compiler/internal/lexical/models"
	"arithmetic-compiler/internal/syntax"
	"arithmetic-compiler/internal/syntax/writers"
)

type syntaxHandler struct {
}

func (h *syntaxHandler) Start(expression string) ([][]string, error) {
	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, err
	}
	tokens := lexical.NewTokenBuilder().GetTokens(h.prepareLexemesToSyntaxAnalyze(lexemes))

	result, err := syntax.NewSyntaxAnalyzer(tokens).Analyze()
	if err != nil {
		return nil, err
	}

	syntaxTree := writers.NewTreeBuilder(result).Build()
	return [][]string{{syntaxTree.Print()}}, nil
}

func (h *syntaxHandler) prepareLexemesToSyntaxAnalyze(lexemes []models.Lexeme) []models.Lexeme {
	endLexeme, _ := models.NewLexeme(models.Unknown, "")
	lexemes = append(lexemes, *endLexeme)
	return lexemes
}
