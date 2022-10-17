package semantic

import (
	"arithmetic-compiler/internal/syntax/models"
	"fmt"
)

type semanticAnalyzer struct {
	tree models.Node
}

func NewSemanticAnalyzer(tree models.Node) *semanticAnalyzer {
	return &semanticAnalyzer{tree: tree}
}

func (s *semanticAnalyzer) Analyze() (models.Node, error) {
	return s.tree, s.checkDivisionByZero()
}

func (s *semanticAnalyzer) checkDivisionByZero() error {
	tree := s.aroundTree()
	for _, node := range tree {
		fmt.Println(node.ToStringNode())
	}
	return nil
}

func (s *semanticAnalyzer) aroundTree() []models.Node {
	tree := s.tree
	root, ok := tree.(models.BinaryNode)
	if !ok {
		return []models.Node{root}
	}
	return s.getNodes(root)
}

func (s *semanticAnalyzer) getNodes(node models.BinaryNode) []models.Node {
	var nodes []models.Node
	nodes = append(nodes, node)
	for {
		lNode, lOk := node.LeftNode.(models.BinaryNode)
		rNode, rOk := node.RightNode.(models.BinaryNode)
		if !lOk && !rOk {
			if node.LeftNode != nil && node.RightNode != nil {
				nodes = append(nodes, node.LeftNode)
				nodes = append(nodes, node.RightNode)
			}
			break
		} else if lOk && rOk {
			nodes = append(nodes, s.getNodes(lNode)...)
			nodes = append(nodes, s.getNodes(rNode)...)
			break
		} else if lOk {
			nodes = append(nodes, node.RightNode)
			nodes = append(nodes, s.getNodes(lNode)...)
			node = rNode
		} else if rOk {
			nodes = append(nodes, node.LeftNode)
			nodes = append(nodes, s.getNodes(rNode)...)
			node = lNode
		}
	}
	return nodes
}

//func (s *semanticAnalyzer) checkDivisionByZero() error {
//	for i, token := range s.tokens {
//		if token.Lexeme.Type == models2.OpDiv &&
//			(len(s.tokens)-1 != i) {
//			return errors.New(fmt.Sprintf("semantic error! division by zero on %d position", i+1))
//		}
//	}
//	return nil
//}
