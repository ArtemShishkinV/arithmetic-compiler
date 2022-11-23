package handlers

import (
	"arithmetic-compiler/internal/code"
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
	_, semanticNode, err := g.semHandler.GetSemanticTree(expression)
	if err != nil {
		return nil, err
	}
	code.NewCodeGenerator(semanticNode).GetThreeAddressCode()
	return [][]string{{}}, nil
}
