package models

import "arithmetic-syntax-analyzer/internal/lexical/models"

type OperandNode struct {
	Token models.Token
}

func NewOperandNode(token models.Token) Node {
	return OperandNode{Token: token}
}

func (o OperandNode) ToStringNode() string {
	return o.Token.Value
}
