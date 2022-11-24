package handlers

import (
	"arithmetic-compiler/internal/code"
	models2 "arithmetic-compiler/internal/code/models"
	"fmt"
)

type codeGeneratorHandler struct {
	semHandler semanticHandler
	generator  code.CodeGenerator
}

func NewGeneratorCodeHandler(handler semanticHandler) Handler {
	return &codeGeneratorHandler{semHandler: handler}
}

func (g *codeGeneratorHandler) Start(expression string) ([][]string, error) {
	fmt.Println("#generator-code")
	_, node, err := g.semHandler.GetSemanticTree(expression)
	if err != nil {
		return nil, err
	}

	generator := code.NewCodeGenerator(g.semHandler.analyzer.GetVars())
	generator.GetThreeAddressCode(node)

	return [][]string{g.prepareToOutCode(generator.Codes), g.prepareToOutTables(generator.Tables)}, nil
}

func (g *codeGeneratorHandler) prepareToOutCode(code []models2.ThreeAddressCode) []string {
	var codes []string
	for _, item := range code {
		codes = append(codes, item.ToString())
	}
	return codes
}

func (g *codeGeneratorHandler) prepareToOutTables(table []models2.TableDtoCode) []string {
	var codes []string
	for _, item := range table {
		codes = append(codes, item.ToString())
	}
	return codes
}
