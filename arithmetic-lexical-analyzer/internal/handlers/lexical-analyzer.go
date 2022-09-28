package handlers

import (
	"arithmetic-lexical-analyzer/internal/config"
	"arithmetic-lexical-analyzer/pkg"
	"errors"
	"fmt"
	"regexp"
)

type lexicalAnalyzer struct {
	config *config.Config
}

func (l *lexicalAnalyzer) Start() error {
	fmt.Println("#analysis")
	expressions, err := pkg.ReadFileLines(l.config.SrcFileName)
	if err != nil {
		return err
	}
	if len(expressions) != 1 {
		return errors.New("source file must contain only one string")
	}
	expression := expressions[0]
	_, err = l.analysis(expression)
	if err != nil {
		return err
	}
	return nil
}

func NewLexicalAnalyzer(argConfig *config.Config) Handler {
	return &lexicalAnalyzer{
		argConfig,
	}
}

func (l *lexicalAnalyzer) analysis(expression string) (map[string]string, error) {
	l.getSymbolsFromExpression(expression)
	return nil, nil
}

func (l *lexicalAnalyzer) getSymbolsFromExpression(expression string) string {
	regSpaces, _ := regexp.Compile(`\s+`)
	expressionWithoutSpaces := regSpaces.ReplaceAllString(expression, " ")
	fmt.Println(expressionWithoutSpaces)
	return expressionWithoutSpaces
}
