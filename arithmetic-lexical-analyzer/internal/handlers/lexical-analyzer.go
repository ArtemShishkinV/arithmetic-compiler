package handlers

import (
	"arithmetic-lexical-analyzer/internal/lexical"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type LexicalAnalyzer struct{}

func NewLexicalAnalyzer() *LexicalAnalyzer {
	return &LexicalAnalyzer{}
}

func (l *LexicalAnalyzer) Start(expression string) ([]string, []string, error) {
	fmt.Println("#analysis")
	lexemes, err := l.analysis(expression)
	if err != nil {
		return nil, nil, err
	}
	tokens := l.getOutputTokens(lexemes)
	tableVars := l.getOutputVariables(l.getVariablesFromLexemes(lexemes))
	return tokens, tableVars, nil
}

func (l *LexicalAnalyzer) analysis(expression string) ([]lexical.Lexeme, error) {
	symbols := l.getSymbolsFromExpression(expression)
	var lexemes []lexical.Lexeme

	for i, symbol := range symbols {
		lexeme, err := lexical.NewLexeme(symbol)
		if err != nil {
			return nil, errors.New(err.Error() + " in " + strconv.Itoa(i+1) + " position")
		}
		lexemes = append(lexemes, *lexeme)
	}

	return lexemes, nil
}

func (l *LexicalAnalyzer) getVariablesFromLexemes(lexemes []lexical.Lexeme) []lexical.Lexeme {
	var varLexemes []lexical.Lexeme
	for _, lexeme := range lexemes {
		if lexeme.Type == lexical.Variable {
			varLexemes = append(varLexemes, lexeme)
		}
	}
	return varLexemes
}

func (l *LexicalAnalyzer) getSymbolsFromExpression(expression string) []string {
	delimiter := " "
	regSpaces, _ := regexp.Compile(`\s+`)
	res := regSpaces.ReplaceAllString(expression, delimiter)
	res = strings.ReplaceAll(res, "(", "("+delimiter)
	res = strings.ReplaceAll(res, ")", delimiter+")")
	return strings.Split(res, delimiter)
}

func (l *LexicalAnalyzer) getOutputTokens(lexemes []lexical.Lexeme) []string {
	var result []string
	varNumber := 1
	for _, lexeme := range lexemes {
		if lexeme.Type != lexical.Variable {
			result = append(result, fmt.Sprintf("<%s> - %s", lexeme.Symbol, lexeme.Type))
		} else {
			result = append(result, fmt.Sprintf("<id, %d> - %s %s", varNumber, lexeme.Type, lexeme.Symbol))
			varNumber++
		}
	}
	return result
}
func (l *LexicalAnalyzer) getOutputVariables(lexemes []lexical.Lexeme) []string {
	var result []string
	for i, lexeme := range lexemes {
		result = append(result, fmt.Sprintf("%d - %s", i+1, lexeme.Symbol))
	}
	return result
}
