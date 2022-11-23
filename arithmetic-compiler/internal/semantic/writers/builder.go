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

func (t *TreeBuilder) Build() (gotree.Tree, models.Node) {
	node := t.node
	root, ok := node.(models.BinaryNode)
	if !ok {
		return gotree.New(node.ToStringNode()), t.node
	}
	root = t.checkTypesOperand(root)
	return t.buildTreeByRoot(root)
}

func (t *TreeBuilder) buildTreeByRoot(node models.BinaryNode) (gotree.Tree, models.Node) {
	var clNode, crNode models.ConvertNode
	var newTree gotree.Tree
	tree := gotree.New(node.Operator.Value)
	for {
		lNode, lOk := node.LeftNode.(models.BinaryNode)
		rNode, rOk := node.RightNode.(models.BinaryNode)
		if !lOk && !rOk {
			clNode, lOk = node.LeftNode.(models.ConvertNode)
			crNode, rOk = node.RightNode.(models.ConvertNode)
			if lOk {
				tree.AddTree(t.createConvertTree(clNode))
			} else if node.LeftNode != nil {
				tree.Add(node.LeftNode.ToStringNode())
			}
			if rOk {
				tree.AddTree(t.createConvertTree(crNode))
			} else if node.RightNode != nil {
				tree.Add(node.RightNode.ToStringNode())
			}
			break
		} else if lOk && rOk {
			lNode = t.checkTypesOperand(lNode)
			rNode = t.checkTypesOperand(rNode)
			newTree, _ = t.buildTreeByRoot(lNode)
			tree.AddTree(newTree)
			newTree, _ = t.buildTreeByRoot(rNode)
			tree.AddTree(newTree)
			break
		} else if lOk {
			lNode = t.checkTypesOperand(lNode)
			newTree, node.LeftNode = t.buildTreeByRoot(lNode)
			tree.AddTree(newTree)
			crNode, rOk = node.RightNode.(models.ConvertNode)
			if rOk {
				tree.AddTree(t.createConvertTree(crNode))
			} else {
				tree.Add(node.RightNode.ToStringNode())
			}
			if rNode.NodeResult == "" {
				//node.LeftNode = lNode
				return tree, node
			}
			node = rNode
		} else if rOk {
			rNode = t.checkTypesOperand(rNode)
			clNode, lOk = node.LeftNode.(models.ConvertNode)
			if lOk {
				tree.AddTree(t.createConvertTree(clNode))
			} else {
				tree.Add(node.LeftNode.ToStringNode())
			}
			newTree, node.RightNode = t.buildTreeByRoot(rNode)
			tree.AddTree(newTree)
			if lNode.NodeResult == "" {
				//node.RightNode = rNode
				return tree, node
			}
			node = lNode
		}
	}
	return tree, node
}

func (t *TreeBuilder) checkTypesOperand(node models.BinaryNode) models.BinaryNode {
	resNode := models.Node(node)

	lNode, lOk := node.LeftNode.(models.BinaryNode)
	rNode, rOk := node.RightNode.(models.BinaryNode)

	if lOk && rOk {
		if t.checkTypesNode(node.RightNode, node.LeftNode) {
			resNode = models.NewBinaryNode(node.Operator, models.NewConvertNode(lNode), rNode)
		}
		if t.checkTypesNode(node.LeftNode, node.RightNode) {
			resNode = models.NewBinaryNode(node.Operator, lNode, models.NewConvertNode(rNode))
		}
	} else if t.checkTypesNode(node.LeftNode, node.RightNode) {
		resNode = models.NewBinaryNode(node.Operator, node.LeftNode, models.NewConvertNode(node.RightNode))
	} else if t.checkTypesNode(node.RightNode, node.LeftNode) {
		resNode = models.NewBinaryNode(node.Operator, models.NewConvertNode(node.LeftNode), node.RightNode)
	}

	return resNode.(models.BinaryNode)
}

func (t TreeBuilder) checkTypesNode(node1 models.Node, node2 models.Node) bool {
	return node1.GetNodeResult() == models.Float && node2.GetNodeResult() == models.Integer
}

func (t TreeBuilder) getConvertNode(src models.Node) models.Node {
	token := src.GetToken()
	token.Value = string(models2.Int2Float)
	return models.NewOperandNode(token)
}

func (t TreeBuilder) createConvertTree(node models.ConvertNode) gotree.Tree {
	convertTree := gotree.New(node.GetConvertNode().ToStringNode())
	bNode, ok := node.GetOperandNode().(models.BinaryNode)
	if ok {
		tree, _ := t.buildTreeByRoot(bNode)
		convertTree.AddTree(tree)
	} else {
		convertTree.Add(node.GetOperandNode().ToStringNode())
	}
	return convertTree
}
