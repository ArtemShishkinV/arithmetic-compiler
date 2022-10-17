package semantic

import (
	models2 "arithmetic-compiler/internal/lexical/models"
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
	root = t.checkTypesOperand(root)
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
	_, lOk := node.LeftNode.(models.BinaryNode)
	_, rOk := node.RightNode.(models.BinaryNode)

	if lOk && rOk {
		if t.checkTypesNode(node.LeftNode, node.RightNode) ||
			t.checkTypesNode(node.RightNode, node.LeftNode) {
			node.Operator = t.getConvertNode(models.NewOperandNode(node.Operator)).GetToken()
		}
		return node
	}

	if t.checkTypesNode(node.LeftNode, node.RightNode) {
		node.RightNode = t.getConvertNode(node.RightNode)
	} else if t.checkTypesNode(node.RightNode, node.LeftNode) {
		node.LeftNode = t.getConvertNode(node.LeftNode)
	}

	return node
}

func (t TreeBuilder) checkTypesNode(node1 models.Node, node2 models.Node) bool {
	return node1.GetNodeResult() == models.Float && node2.GetNodeResult() == models.Integer
}

func (t TreeBuilder) getConvertNode(src models.Node) models.Node {
	token := src.GetToken()
	token.Value = string(models2.Int2Float) + token.Value
	return models.NewOperandNode(token)
}
