package handlers

import (
	"arithmetic-compiler/internal/lexical"
	"arithmetic-compiler/internal/semantic"
	"arithmetic-compiler/internal/syntax"
	"arithmetic-compiler/internal/syntax/writers"
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
	syntaxTree, err := syntax.NewSyntaxAnalyzer(tokens).Analyze()
	if err != nil {
		return nil, err
	}

	semanticTree, err := semantic.NewSemanticAnalyzer(syntaxTree).Analyze()
	if err != nil {
		return nil, err
	}
	printTree := writers.NewTreeBuilder(semanticTree).Build()
	return [][]string{{printTree.Print()}}, nil
}
