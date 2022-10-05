package writers

import (
	"arithmetic-syntax-analyzer/internal/syntax/models"
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
	tree := gotree.New(node.Operator.Symbol)
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
			tree.AddTree(t.buildTreeByRoot(lNode))
			tree.AddTree(t.buildTreeByRoot(rNode))
			break
		} else if lOk {
			tree.Add(node.RightNode.ToStringNode())
			tree.AddTree(t.buildTreeByRoot(lNode))
			node = rNode
		} else if rOk {
			tree.Add(node.LeftNode.ToStringNode())
			tree.AddTree(t.buildTreeByRoot(rNode))
			node = lNode
		}
	}
	return tree
}

//func (t *TreeBuilder) Build() string {
//
//	//var buffer bytes.Buffer
//	//var lNode, rNode models.BinaryNode
//	//var lOk, rOk bool
//	//lPreffix, rPreffix := 0, 0
//	node := t.node
//	root, ok := node.(models.BinaryNode)
//	if !ok {
//		return ""
//	}
//	fmt.Println(root.ToStringNode())
//	fmt.Println(getLevelBranch(root, "left"))
//	fmt.Println(getLevelBranch(root, "right"))
//	//for {
//	//	buffer.WriteString(strings.Repeat("\t", lPreffix+rPreffix+1) + root.Operator.Symbol + "\n")
//	//	lNode, lOk = root.LeftNode.(models.BinaryNode)
//	//	rNode, rOk = root.RightNode.(models.BinaryNode)
//	//	if !lOk && !rOk {
//	//		lPreffix++
//	//		rPreffix++
//	//		buffer.WriteString(strings.Repeat("\t", lPreffix) + root.LeftNode.ToStringNode() + strings.Repeat("\t", rPreffix) + root.RightNode.ToStringNode() + "\n")
//	//		return buffer.String()
//	//	} else if lOk {
//	//		rPreffix++
//	//		buffer.WriteString(root.RightNode.ToStringNode())
//	//		root = lNode
//	//	} else {
//	//		lPreffix++
//	//		buffer.WriteString(root.LeftNode.ToStringNode())
//	//		root = rNode
//	//	}
//	//}
//	return ""
//}
//
//func getLevelBranch(root models.BinaryNode, side string) int {
//	if side == "left" {
//		return getLevelSideBranch(root.LeftNode)
//	}
//	return getLevelSideBranch(root.RightNode)
//}

//func getLevelSideBranch(root models.Node) int {
//	level := 0
//	root, ok := root.(models.BinaryNode)
//	if !ok {
//		return 1
//	}
//	for {
//		lNode, ok := root.(models.BinaryNode)
//		if !ok {
//			break
//		}
//		level++
//		root = lNode.LeftNode
//	}
//	return level
//}
