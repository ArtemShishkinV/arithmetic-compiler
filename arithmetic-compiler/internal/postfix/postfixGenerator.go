package postfix

import (
	"arithmetic-compiler/internal/code/models"
	models2 "arithmetic-compiler/internal/lexical/models"
	models3 "arithmetic-compiler/internal/syntax/models"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

type postfixGenerator struct {
	actions []models.ThreeAddressCode
	vars    []models2.Token
}

func NewPostfixGenerator(actions []models.ThreeAddressCode, vars []models2.Token) *postfixGenerator {
	return &postfixGenerator{
		actions,
		vars,
	}
}

func (g *postfixGenerator) Generate() ([]string, []models.TableDtoCode) {
	return g.getPostfixExpression(), g.getVarsTable(g.vars)
}

func (g *postfixGenerator) getPostfixExpression() []string {
	var operators []string
	var operands []string

	for _, item := range g.actions {
		operators = append(operators, fmt.Sprintf("<%s>", item.GetOperator().Value))
	}

	tokenOperands := g.filteredOperands()
	fmt.Println(tokenOperands)
	for i := len(tokenOperands) - 1; i >= 0; i-- {
		operands = append(operands, tokenOperands[i].Value)
	}

	return g.concatOperandsAndOperators(operands, operators)
}

func (g *postfixGenerator) filteredOperands() []models2.Token {
	var operands []models2.Token
	for i := len(g.actions) - 1; i >= 0; i-- {
		for _, j := range g.actions[i].GetOperands() {
			if j.Lexeme.Type != models2.Result && !slices.Contains(operands, j) {
				operands = append(operands, j)
			}
		}
	}
	return operands
}

func (g *postfixGenerator) concatOperandsAndOperators(operands []string, operators []string) []string {
	var result []string
	result = append(result, operands...)
	result = append(result, operators...)

	return []string{strings.Join(result, " ")}
}

func (g *postfixGenerator) getVarsTable(vars []models2.Token) []models.TableDtoCode {
	var table []models.TableDtoCode
	for _, item := range vars {
		table = append(table,
			models.NewTableDtoCode(item.Value, item.Lexeme.Symbol,
				models3.GetTypeResult(item.Lexeme)))
	}
	return table
}
