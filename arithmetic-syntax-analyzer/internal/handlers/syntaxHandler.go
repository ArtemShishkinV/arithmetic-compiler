package handlers

import (
	"arithmetic-syntax-analyzer/internal/lexical"
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"arithmetic-syntax-analyzer/internal/syntax"
	"arithmetic-syntax-analyzer/internal/syntax/writers"
)

type syntaxHandler struct{}

func (h *syntaxHandler) Start(expression string) ([][]string, error) {
	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, err
	}
	lexemes = h.prepareLexemesToSyntaxAnalyze(lexemes)
	result, err := syntax.NewSyntaxAnalyzer(lexemes).Analyze()
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
