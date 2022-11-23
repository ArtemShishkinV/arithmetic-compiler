package handlers

import (
	"arithmetic-compiler/internal/lexical"
	"arithmetic-compiler/internal/semantic"
	semantic2 "arithmetic-compiler/internal/semantic/writers"
	"arithmetic-compiler/internal/syntax"
	"arithmetic-compiler/internal/syntax/models"
	"fmt"
	"github.com/disiqueira/gotree"
)

type semanticHandler struct{}

func (h *semanticHandler) Start(expression string) ([][]string, error) {
	fmt.Println("#analysis-semantic")
	tree, _, err := h.GetSemanticTree(expression)
	if err != nil {
		return nil, err
	}
	return [][]string{{tree.Print()}}, nil
}

func (h *semanticHandler) GetSemanticTree(expression string) (gotree.Tree, models.Node, error) {
	handler := syntaxHandler{}

	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, nil, err
	}
	tokens := lexical.NewTokenBuilder().GetTokens(handler.prepareLexemesToSyntaxAnalyze(lexemes))
	result, err := syntax.NewSyntaxAnalyzer(tokens).Analyze()
	if err != nil {
		return nil, result, err
	}
	_, err = semantic.NewSemanticAnalyzer(tokens).Analyze()
	if err != nil {
		return nil, result, err
	}

	tree, node := semantic2.NewTreeBuilder(result).Build()

	return tree, node, nil
}
