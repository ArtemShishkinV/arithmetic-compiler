package handlers

import (
	"arithmetic-compiler/internal/code"
	"arithmetic-compiler/internal/postfix"
	"fmt"
)

type postfixGeneratorHandler struct {
	codeGenerator codeGeneratorHandler
}

func (g *postfixGeneratorHandler) Start(expression string) ([][]string, error) {
	fmt.Println("#postfix-expression-generator")
	_, semanticNode, err := g.codeGenerator.semHandler.GetSemanticTree(expression)
	if err != nil {
		return nil, err
	}
	vars := g.codeGenerator.semHandler.analyzer.GetVars()

	g.codeGenerator.generator = *code.NewCodeGenerator(vars)
	g.codeGenerator.generator.GetThreeAddressCode(semanticNode)
	actions := g.codeGenerator.generator.Codes

	generator := postfix.NewPostfixGenerator(actions, vars)
	postfixExpr, table := generator.Generate()

	return [][]string{postfixExpr, g.codeGenerator.prepareToOutTables(table)}, nil
}
