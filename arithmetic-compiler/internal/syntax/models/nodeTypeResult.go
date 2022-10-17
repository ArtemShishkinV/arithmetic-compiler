package models

import "arithmetic-compiler/internal/lexical/models"

type NodeTypeResult models.LexemeType

const (
	Float   NodeTypeResult = "вещественный"
	Integer NodeTypeResult = "целый"
)

func GetTypeResultOperand(node OperandNode) NodeTypeResult {
	if models.IsFloatType(node.Token.Lexeme.Type) {
		return Float
	}
	return Integer
}

//func GetTypeResultBinary(node BinaryNode) NodeTypeResult {
//	if GetTypeResultOperand(node.LeftNode) == Float || GetTypeResultOperand(node.RightNode) == Float {
//		return Float
//	}
//}
