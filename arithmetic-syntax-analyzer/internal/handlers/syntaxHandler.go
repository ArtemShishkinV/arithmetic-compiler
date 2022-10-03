package handlers

import (
	"arithmetic-syntax-analyzer/internal/lexical"
	"arithmetic-syntax-analyzer/internal/syntax"
	"fmt"
)

type syntaxHandler struct{}

func (h *syntaxHandler) Start(expression string) ([][]string, error) {
	lexemes, err := lexical.NewLexicalAnalyzer().Analyze(expression)
	if err != nil {
		return nil, err
	}
	result, err := syntax.NewSyntaxAnalyzer(lexemes).Analyze()
	//fmt.Println(err.Error())
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(result)
	return nil, nil
}
