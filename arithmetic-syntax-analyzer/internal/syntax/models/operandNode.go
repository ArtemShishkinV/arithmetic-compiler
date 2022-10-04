package models

import "arithmetic-syntax-analyzer/internal/lexical/models"

type OperandNode struct {
	Lexeme models.Lexeme
}

func NewOperandNode(lexeme models.Lexeme) Node {
	return OperandNode{Lexeme: lexeme}
}

func (o OperandNode) ToStringNode() string {
	return o.Lexeme.Symbol
}
