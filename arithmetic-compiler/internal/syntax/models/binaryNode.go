package models

import (
	"arithmetic-compiler/internal/lexical/models"
	"gopkg.in/Knetic/govaluate.v2"
)

type BinaryNode struct {
	Operator   models.Token
	LeftNode   Node
	RightNode  Node
	NodeResult NodeTypeResult
}

func NewBinaryNode(operator models.Token, leftNode Node, rightNode Node) Node {
	binaryNode := BinaryNode{
		Operator:  operator,
		LeftNode:  leftNode,
		RightNode: rightNode,
	}
	binaryNode.NodeResult = binaryNode.getTypeResult()
	return binaryNode
}

func (b BinaryNode) ToStringNode() string {
	return b.LeftNode.ToStringNode() + b.Operator.Value + b.RightNode.ToStringNode()
}

func (b BinaryNode) GetToken() models.Token {
	return b.Operator
}

func (b BinaryNode) GetNodeResult() NodeTypeResult {
	return b.NodeResult
}

func (b BinaryNode) getTypeResult() NodeTypeResult {
	if b.LeftNode.GetNodeResult() == Float || b.RightNode.GetNodeResult() == Float {
		return Float
	}
	return Integer
}

func (b BinaryNode) calculateNode() Node {
	expression, err := govaluate.NewEvaluableExpression(b.LeftNode.GetToken().Value + b.Operator.Value + b.RightNode.GetToken().Value)
	if err != nil {
		return b
	}
	lexType := b.LeftNode.GetNodeResult()
	lexeme := models.NewLexeme()
	token := models.NewToken()

	return NewOperandNode(*token)
}
