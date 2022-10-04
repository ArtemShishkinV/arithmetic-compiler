package handlers

import (
	"arithmetic-syntax-analyzer/internal/lexical"
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"arithmetic-syntax-analyzer/internal/syntax"
	"arithmetic-syntax-analyzer/internal/syntax/writers"
	"fmt"
)

type syntaxHandler struct{}

func (h *syntaxHandler) Start(expression string) ([][]string, error) {
	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, err
	}
	endLexeme, _ := models.NewLexeme(models.Unknown, "")
	lexemes = append(lexemes, *endLexeme)
	result, err := syntax.NewSyntaxAnalyzer(lexemes).Analyze()
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(writers.NewTreeBuilder(result).Build())
	return nil, nil
}
