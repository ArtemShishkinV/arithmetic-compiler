package handlers

import (
	"arithmetic-compiler/internal/lexical"
	"arithmetic-compiler/internal/semantic"
	semantic2 "arithmetic-compiler/internal/semantic/writers"
	"arithmetic-compiler/internal/syntax"
	"fmt"
)

type semanticHandler struct{}

func (h *semanticHandler) Start(expression string) ([][]string, error) {
	fmt.Println("#analysis-semantic")
	handler := syntaxHandler{}

	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, err
	}
	tokens := lexical.NewTokenBuilder().GetTokens(handler.prepareLexemesToSyntaxAnalyze(lexemes))
	result, err := syntax.NewSyntaxAnalyzer(tokens).Analyze()
	if err != nil {
		return nil, err
	}
	_, err = semantic.NewSemanticAnalyzer(tokens).Analyze()
	if err != nil {
		return nil, err
	}

	semanticTree := semantic2.NewTreeBuilder(result).Build()
	return [][]string{{semanticTree.Print()}}, nil
}
