package models

import "arithmetic-compiler/internal/lexical/models"

type OperandNode struct {
	Token      models.Token
	NodeResult NodeTypeResult
}

func NewOperandNode(token models.Token) Node {
	operandNode := OperandNode{Token: token}
	operandNode.NodeResult = operandNode.getTypeResult()
	return operandNode
}

func (o OperandNode) ToStringNode() string {
	return o.Token.Value
}

func (o OperandNode) GetToken() models.Token {
	return o.Token
}

func (o OperandNode) GetNodeResult() NodeTypeResult {
	return o.NodeResult
}

func (o OperandNode) getTypeResult() NodeTypeResult {
	if models.IsFloatType(o.Token.Lexeme.Type) {
		return Float
	}
	return Integer
}
