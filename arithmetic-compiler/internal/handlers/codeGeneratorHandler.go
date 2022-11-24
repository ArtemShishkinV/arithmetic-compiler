package handlers

import (
	"arithmetic-compiler/internal/code"
	"arithmetic-compiler/internal/lexical"
	"arithmetic-compiler/internal/semantic"
	semantic2 "arithmetic-compiler/internal/semantic/writers"
	"arithmetic-compiler/internal/syntax"
	"fmt"
)

type codeGeneratorHandler struct {
	semHandler semanticHandler
}

func NewGeneratorCodeHandler(handler semanticHandler) Handler {
	return &codeGeneratorHandler{semHandler: handler}
}

func (g *codeGeneratorHandler) Start(expression string) ([][]string, error) {
	fmt.Println("#generator-code")

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
	semAnalyzer := semantic.NewSemanticAnalyzer(tokens)

	if _, err = semAnalyzer.Analyze(); err != nil {
		return nil, err
	}

	tree, node := semantic2.NewTreeBuilder(result).Build()

	generator := code.NewCodeGenerator(semAnalyzer.GetVars())
	generator.GetThreeAddressCode(node)

	for _, item := range generator.Codes {
		fmt.Println(item.ToString())
	}
	for _, item := range generator.Tables {
		fmt.Println(item)
	}

	return [][]string{{tree.Print()}}, nil
}
