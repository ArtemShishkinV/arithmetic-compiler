package models

import models2 "arithmetic-compiler/internal/lexical/models"

type NodeResultType string

const (
	Float   NodeResultType = "вещественное число"
	Integer NodeResultType = "целое чиcло"
)

func getTypeResultBinaryNode(node BinaryNode) NodeResultType {
	if node.LeftNode.GetToken().Lexeme.Type == models2.FloatNumber ||
		node.LeftNode.GetToken().Lexeme.Type == models2.FloatVariable ||
		node.RightNode.GetToken().Lexeme.Type == models2.FloatNumber ||
		node.RightNode.GetToken().Lexeme.Type == models2.FloatVariable {
		return Float
	}
	return Integer
}
