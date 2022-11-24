package code

import (
	models2 "arithmetic-compiler/internal/code/models"
	models3 "arithmetic-compiler/internal/lexical/models"
	"arithmetic-compiler/internal/syntax/models"
	"strconv"
)

type codeGenerator struct {
	offsetId int
	Codes    []models2.ThreeAddressCode
	Tables   []models2.TableDtoCode
}

func NewCodeGenerator(vars []models3.Token) *codeGenerator {
	generator := &codeGenerator{
		offsetId: len(vars) + 1,
	}
	generator.addVarsInTable(vars)

	return generator
}

func (g *codeGenerator) GetThreeAddressCode(parNode models.Node) models2.ThreeAddressCode {
	var clNode, crNode models.ConvertNode
	var code models2.ThreeAddressCode
	node, ok := parNode.(models.BinaryNode)
	if !ok {
		return g.getCodeByNodeType(parNode)
	}

	for {
		lNode, lOk := node.LeftNode.(models.BinaryNode)
		rNode, rOk := node.RightNode.(models.BinaryNode)
		if !lOk && !rOk {
			if node.LeftNode != nil && node.RightNode != nil {
				crNode, rOk = node.RightNode.(models.ConvertNode)
				clNode, lOk = node.LeftNode.(models.ConvertNode)
				if rOk {
					code = g.createCodesWithConvert(node, crNode)
				} else if lOk {
					code = g.createCodesWithConvert(node, clNode)
				} else {
					token1 := node.LeftNode.GetToken()
					token2 := node.RightNode.GetToken()
					code, _ = models2.NewThreeAddressCode(
						node.Operator.Lexeme.Type,
						len(g.Codes)+g.offsetId, []string{token1.Value, token2.Value})
				}
				g.addCode(code, node)
			}
			break
		} else if lOk && rOk {
			g.GetThreeAddressCode(lNode)
			g.GetThreeAddressCode(rNode)
			break
		} else if lOk {
			code = g.GetThreeAddressCode(lNode)
			secondOperand := node.RightNode.GetToken().Value
			crNode, rOk = node.RightNode.(models.ConvertNode)
			if rOk {
				g.createCodesWithConvert(node, crNode)
				secondOperand = g.Codes[len(g.Codes)-1].GetResult()
			}
			code, _ = models2.NewThreeAddressCode(node.Operator.Lexeme.Type,
				len(g.Codes)+g.offsetId, []string{code.GetResult(), secondOperand})
			g.addCode(code, node)
			node = rNode
		} else if rOk {
			code = g.GetThreeAddressCode(rNode)
			secondOperand := node.LeftNode.GetToken().Value
			crNode, rOk = node.LeftNode.(models.ConvertNode)
			if rOk {
				g.createCodesWithConvert(node, crNode)
				secondOperand = g.Codes[len(g.Codes)-1].GetResult()
			}
			code, _ = models2.NewThreeAddressCode(node.Operator.Lexeme.Type,
				len(g.Codes)+g.offsetId, []string{secondOperand, code.GetResult()})
			g.addCode(code, node)
			node = lNode
		}
	}
	return code
}

func (g *codeGenerator) createCodesWithConvert(node models.BinaryNode, convertNode models.ConvertNode) models2.ThreeAddressCode {
	var tempCode, code models2.ThreeAddressCode
	tempCode = g.GetThreeAddressCode(convertNode)
	g.addCode(tempCode, convertNode)

	if _, ok := node.LeftNode.(models.OperandNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator.Lexeme.Type,
			len(g.Codes)+g.offsetId, []string{node.LeftNode.GetToken().Value, tempCode.GetResult()})
	} else if _, ok := node.RightNode.(models.OperandNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator.Lexeme.Type,
			len(g.Codes)+g.offsetId, []string{tempCode.GetResult(), node.RightNode.GetToken().Value})
	} else if _, ok := node.LeftNode.(models.ConvertNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator.Lexeme.Type,
			len(g.Codes)+g.offsetId, []string{tempCode.GetResult(), g.Codes[len(g.Codes)-2].GetResult()})
	} else if _, ok := node.RightNode.(models.ConvertNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator.Lexeme.Type,
			len(g.Codes)+g.offsetId, []string{g.Codes[len(g.Codes)-2].GetResult(), tempCode.GetResult()})
	}

	return code
}

func (g *codeGenerator) getCodeByNodeType(parNode models.Node) models2.ThreeAddressCode {
	var code models2.ThreeAddressCode

	node, ok := parNode.(models.ConvertNode)
	if ok {
		tempNode, ok := node.GetOperandNode().(models.BinaryNode)
		if ok {
			code = g.GetThreeAddressCode(tempNode)
			code, _ = models2.NewThreeAddressCode(parNode.GetToken().Lexeme.Type,
				len(g.Codes)+g.offsetId, []string{code.GetResult()})
		} else {
			code, _ = models2.NewThreeAddressCode(parNode.GetToken().Lexeme.Type,
				len(g.Codes)+g.offsetId, []string{node.GetOperandNode().GetToken().Value})
		}
	} else {
		code, _ = models2.NewThreeAddressCode(parNode.GetToken().Lexeme.Type,
			len(g.Codes)+g.offsetId, []string{parNode.GetToken().Value})
	}

	return code
}

func (g *codeGenerator) addCode(code models2.ThreeAddressCode, node models.Node) {
	g.Tables = append(g.Tables,
		models2.NewTableDtoCode(code.GetResult(), "T"+strconv.Itoa(len(g.Codes)+1), node.GetNodeResult()))
	g.Codes = append(g.Codes, code)
}

func (g *codeGenerator) addVarsInTable(vars []models3.Token) {
	for _, item := range vars {
		g.Tables = append(g.Tables,
			models2.NewTableDtoCode(item.Value, item.Lexeme.Symbol,
				models.GetTypeResult(item.Lexeme)))
	}
}
