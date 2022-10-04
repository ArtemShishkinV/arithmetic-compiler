package models

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
)

type BinaryNode struct {
	operator  models.Lexeme
	leftNode  Node
	rightNode Node
}

func NewBinaryNode(operator models.Lexeme, leftNode Node, rightNode Node) Node {
	return &BinaryNode{
		operator:  operator,
		leftNode:  leftNode,
		rightNode: rightNode,
	}
}

func (b BinaryNode) ToStringNode() string {
	return b.leftNode.ToStringNode() + b.operator.Symbol + b.rightNode.ToStringNode()
}
