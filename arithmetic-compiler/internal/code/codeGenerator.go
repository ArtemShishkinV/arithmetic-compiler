package code

import (
	"arithmetic-compiler/internal/syntax/models"
)

type codeGenerator struct {
	node models.Node
}

func NewCodeGenerator(node models.Node) *codeGenerator {
	return &codeGenerator{node: node}
}

func (g *codeGenerator) GetThreeAddressCode() {
}

func GetSymbolsTable() []string {
	return []string{}
}
