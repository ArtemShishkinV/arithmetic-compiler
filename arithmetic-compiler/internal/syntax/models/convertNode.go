package models

import "arithmetic-compiler/internal/lexical/models"

type ConvertNode struct {
	rootNode   Node
	NodeResult NodeTypeResult
}

func NewConvertNode() Node {
	binaryNode := ConvertNode{
		rootNode: getConvertNode(),
	}
	binaryNode.NodeResult = Float
	return binaryNode
}

func (c ConvertNode) ToStringNode() string {
	return c.rootNode.GetToken().Value
}

func (c ConvertNode) GetToken() models.Token {
	return c.rootNode.GetToken()
}

func (c ConvertNode) GetNodeResult() NodeTypeResult {
	return c.NodeResult
}

func getConvertNode() Node {
	lexeme, _ := models.NewLexeme(models.Int2Float, "Int2Float")
	token := models.NewToken(*lexeme, "Int2Float")
	return NewOperandNode(*token)
}
