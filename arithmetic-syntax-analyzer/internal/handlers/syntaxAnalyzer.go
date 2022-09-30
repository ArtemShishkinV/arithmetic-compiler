package handlers

import "fmt"

type syntaxAnalyzer struct{}

func NewSyntaxAnalyzer() *syntaxAnalyzer {
	return &syntaxAnalyzer{}
}

func (s *syntaxAnalyzer) Start(expression string) (string, error) {
	fmt.Println("#analysis-syntax")
	return "", nil
}
