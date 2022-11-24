package code

import (
	models2 "arithmetic-compiler/internal/code/models"
	models3 "arithmetic-compiler/internal/lexical/models"
	"arithmetic-compiler/internal/syntax/models"
	"strconv"
)

type CodeGenerator struct {
	offsetId int
	Codes    []models2.ThreeAddressCode
	Tables   []models2.TableDtoCode
}

func NewCodeGenerator(vars []models3.Token) *CodeGenerator {
	generator := &CodeGenerator{
		offsetId: len(vars) + 1,
	}
	generator.addVarsInTable(vars)

	return generator
}

func (g *CodeGenerator) GetThreeAddressCode(parNode models.Node) models2.ThreeAddressCode {
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
						node.Operator,
						len(g.Codes)+g.offsetId, []models3.Token{token1, token2})
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
			secondOperand := node.RightNode.GetToken()
			crNode, rOk = node.RightNode.(models.ConvertNode)
			if rOk {
				g.createCodesWithConvert(node, crNode)
				secondOperand = g.getTokenForResult(g.Codes[len(g.Codes)-1].GetResult())
			}
			code, _ = models2.NewThreeAddressCode(node.Operator,
				len(g.Codes)+g.offsetId, []models3.Token{g.getTokenForResult(code.GetResult()), secondOperand})
			g.addCode(code, node)
			node = rNode
		} else if rOk {
			code = g.GetThreeAddressCode(rNode)
			secondOperand := node.LeftNode.GetToken()
			crNode, rOk = node.LeftNode.(models.ConvertNode)
			if rOk {
				g.createCodesWithConvert(node, crNode)
				secondOperand = g.getTokenForResult(g.Codes[len(g.Codes)-1].GetResult())
			}
			code, _ = models2.NewThreeAddressCode(node.Operator,
				len(g.Codes)+g.offsetId, []models3.Token{secondOperand, g.getTokenForResult(code.GetResult())})
			g.addCode(code, node)
			node = lNode
		}
	}
	return code
}

func (g *CodeGenerator) createCodesWithConvert(node models.BinaryNode, convertNode models.ConvertNode) models2.ThreeAddressCode {
	var tempCode, code models2.ThreeAddressCode
	tempCode = g.GetThreeAddressCode(convertNode)
	g.addCode(tempCode, convertNode)
	tempToken := g.getTokenForResult(tempCode.GetResult())
	if _, ok := node.LeftNode.(models.OperandNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator,
			len(g.Codes)+g.offsetId, []models3.Token{node.LeftNode.GetToken(), tempToken})
	} else if _, ok := node.RightNode.(models.OperandNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator,
			len(g.Codes)+g.offsetId, []models3.Token{tempToken, node.RightNode.GetToken()})
	} else if _, ok := node.LeftNode.(models.ConvertNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator,
			len(g.Codes)+g.offsetId,
			[]models3.Token{
				tempToken,
				g.getTokenForResult(g.Codes[len(g.Codes)-2].GetResult()),
			})
	} else if _, ok := node.RightNode.(models.ConvertNode); ok {
		code, _ = models2.NewThreeAddressCode(
			node.Operator,
			len(g.Codes)+g.offsetId,
			[]models3.Token{
				g.getTokenForResult(g.Codes[len(g.Codes)-2].GetResult()),
				tempToken,
			})
	}

	return code
}

func (g *CodeGenerator) getCodeByNodeType(parNode models.Node) models2.ThreeAddressCode {
	var code models2.ThreeAddressCode

	node, ok := parNode.(models.ConvertNode)
	parNodeToken := parNode.GetToken()
	if ok {
		tempNode, ok := node.GetOperandNode().(models.BinaryNode)
		if ok {
			code = g.GetThreeAddressCode(tempNode)
			code, _ = models2.NewThreeAddressCode(parNodeToken,
				len(g.Codes)+g.offsetId, []models3.Token{g.getTokenForResult(code.GetResult())})
		} else {
			code, _ = models2.NewThreeAddressCode(parNodeToken,
				len(g.Codes)+g.offsetId, []models3.Token{node.GetOperandNode().GetToken()})
		}
	} else {
		code, _ = models2.NewThreeAddressCode(parNodeToken,
			len(g.Codes)+g.offsetId, []models3.Token{parNode.GetToken()})
	}

	return code
}

func (g *CodeGenerator) addCode(code models2.ThreeAddressCode, node models.Node) {
	g.Tables = append(g.Tables,
		models2.NewTableDtoCode(code.GetResult(), "T"+strconv.Itoa(len(g.Codes)+1), node.GetNodeResult()))
	g.Codes = append(g.Codes, code)
}

func (g *CodeGenerator) addVarsInTable(vars []models3.Token) {
	for _, item := range vars {
		g.Tables = append(g.Tables,
			models2.NewTableDtoCode(item.Value, item.Lexeme.Symbol,
				models.GetTypeResult(item.Lexeme)))
	}
}

func (g *CodeGenerator) getTokenForResult(result string) models3.Token {
	lexeme, _ := models3.NewLexeme(models3.Result, result)
	return *models3.NewToken(*lexeme, result)
}
