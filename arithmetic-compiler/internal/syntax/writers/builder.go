package writers

import (
	"arithmetic-compiler/internal/syntax/models"
	"github.com/disiqueira/gotree"
)

type TreeBuilder struct {
	node models.Node
}

func NewTreeBuilder(node models.Node) *TreeBuilder {
	return &TreeBuilder{node: node}
}

func (t *TreeBuilder) Build() gotree.Tree {
	node := t.node
	root, ok := node.(models.BinaryNode)
	if !ok {
		return gotree.New(node.ToStringNode())
	}
	return t.buildTreeByRoot(root)
}

func (t *TreeBuilder) buildTreeByRoot(node models.BinaryNode) gotree.Tree {
	tree := gotree.New(node.Operator.Value)
	for {
		lNode, lOk := node.LeftNode.(models.BinaryNode)
		rNode, rOk := node.RightNode.(models.BinaryNode)
		if !lOk && !rOk {
			if node.LeftNode != nil && node.RightNode != nil {
				tree.Add(node.LeftNode.ToStringNode())
				tree.Add(node.RightNode.ToStringNode())
			}
			break
		} else if lOk && rOk {
			lNode = t.checkTypesOperand(lNode)
			rNode = t.checkTypesOperand(rNode)
			tree.AddTree(t.buildTreeByRoot(lNode))
			tree.AddTree(t.buildTreeByRoot(rNode))
			break
		} else if lOk {
			lNode = t.checkTypesOperand(lNode)
			tree.Add(node.RightNode.ToStringNode())
			tree.AddTree(t.buildTreeByRoot(lNode))
			node = rNode
		} else if rOk {
			rNode = t.checkTypesOperand(rNode)
			tree.Add(node.LeftNode.ToStringNode())
			tree.AddTree(t.buildTreeByRoot(rNode))
			node = lNode
		}
	}
	return tree
}

func (t *TreeBuilder) checkTypesOperand(node models.BinaryNode) models.BinaryNode {
	token := node.Operator
	if node.LeftNode.GetNodeResult() == models.Float && node.RightNode.GetNodeResult() == models.Integer ||
		node.LeftNode.GetNodeResult() == models.Integer && node.RightNode.GetNodeResult() == models.Float {
		token.Value = "(Int2Float) " + node.Operator.Value
	}

	newNode := models.NewBinaryNode(token, node.LeftNode, node.RightNode)

	//lNodeType := node.LeftNode.GetToken().Lexeme.Type
	//rNodeType := node.RightNode.GetToken().Lexeme.Type

	//if models2.IsIntType(lNodeType) && models2.IsFloatType(rNodeType) {
	//	token := node.LeftNode.GetToken()
	//	token.Value = "(Int2Float) " + node.LeftNode.GetToken().Lexeme.Symbol
	//	newNode := models.NewBinaryNode(node.Operator, models.NewOperandNode(token), node.RightNode)
	//	return newNode.(models.BinaryNode)
	//}
	//
	//if models2.IsIntType(rNodeType) && (models2.IsFloatType(lNodeType)) {
	//	token := node.RightNode.GetToken()
	//	token.Value = "(Int2Float) " + node.RightNode.GetToken().Lexeme.Symbol
	//	newNode := models.NewBinaryNode(node.Operator, node.LeftNode, models.NewOperandNode(token))
	//	return newNode.(models.BinaryNode)
	//}
	return newNode.(models.BinaryNode)
}
