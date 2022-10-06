package models

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
)

type BinaryNode struct {
	Operator  models.Token
	LeftNode  Node
	RightNode Node
}

func NewBinaryNode(operator models.Token, leftNode Node, rightNode Node) Node {
	return BinaryNode{
		Operator:  operator,
		LeftNode:  leftNode,
		RightNode: rightNode,
	}
}

func (b BinaryNode) ToStringNode() string {
	return b.LeftNode.ToStringNode() + b.Operator.Value + b.RightNode.ToStringNode()
}
