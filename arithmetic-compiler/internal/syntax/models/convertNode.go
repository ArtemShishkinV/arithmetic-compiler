package models

import "arithmetic-compiler/internal/lexical/models"

type ConvertNode struct {
	convertNode Node
	opNode      Node
	NodeResult  NodeTypeResult
}

func NewConvertNode(node Node) Node {
	binaryNode := ConvertNode{
		convertNode: getConvertNode(),
		opNode:      node,
	}
	binaryNode.NodeResult = Float
	return binaryNode
}

func (c ConvertNode) ToStringNode() string {
	return c.convertNode.GetToken().Value + " (" + c.opNode.ToStringNode() + ")"
}

func (c ConvertNode) GetToken() models.Token {
	return c.convertNode.GetToken()
}

func (c ConvertNode) GetNodeResult() NodeTypeResult {
	return c.NodeResult
}

func (c ConvertNode) GetOperandNode() Node {
	return c.opNode
}

func (c ConvertNode) GetConvertNode() Node {
	return c.convertNode
}

func getConvertNode() Node {
	lexeme, _ := models.NewLexeme(models.Int2Float, "Int2Float")
	token := models.NewToken(*lexeme, "Int2Float")
	return NewOperandNode(*token)
}
