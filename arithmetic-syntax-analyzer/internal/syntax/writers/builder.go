package writers

import (
	"arithmetic-syntax-analyzer/internal/syntax/models"
	"bytes"
)

type TreeBuilder struct {
	node models.Node
}

func NewTreeBuilder(node models.Node) *TreeBuilder {
	return &TreeBuilder{node: node}
}

func (t *TreeBuilder) Build() string {
	var buffer bytes.Buffer
	var lNode, rNode models.BinaryNode
	var lOk, rOk bool
	node := t.node
	root, ok := node.(models.BinaryNode)
	if !ok {
		return ""
	}
	for {
		buffer.WriteString(root.Operator.Symbol + "\n")
		buffer.WriteString(root.LeftNode.ToStringNode() + "\t" + root.RightNode.ToStringNode() + "\n")
		lNode, lOk = root.LeftNode.(models.BinaryNode)
		rNode, rOk = root.RightNode.(models.BinaryNode)
		if !lOk && !rOk {
			return buffer.String()
		} else if lOk {
			root = lNode
		} else {
			root = rNode
		}
	}
	return buffer.String()
}
