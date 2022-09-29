package handlers

import (
	"arithmetic-lexical-analyzer/internal/config"
	"arithmetic-lexical-analyzer/internal/lexical"
	"arithmetic-lexical-analyzer/pkg"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	lexemes, err := l.analysis(expression)
	if err != nil {
		return err
	}
	if err = pkg.WriteFile(l.getOutputTokens(lexemes), l.config.OutTokensFileName); err != nil {
		return err
	}
	if err = pkg.WriteFile(l.getOutputVariables(l.getVariablesFromLexemes(lexemes)),
		l.config.OutSymbolsFileName); err != nil {
		return err
	}
	return nil
}

func NewLexicalAnalyzer(argConfig *config.Config) Handler {
	return &lexicalAnalyzer{
		argConfig,
	}
}

func (l *lexicalAnalyzer) analysis(expression string) ([]lexical.Lexeme, error) {
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

func (l *lexicalAnalyzer) getVariablesFromLexemes(lexemes []lexical.Lexeme) []lexical.Lexeme {
	var varLexemes []lexical.Lexeme
	for _, lexeme := range lexemes {
		if lexeme.Type == lexical.Variable {
			varLexemes = append(varLexemes, lexeme)
		}
	}
	return varLexemes
}

func (l *lexicalAnalyzer) getSymbolsFromExpression(expression string) []string {
	delimiter := " "
	regSpaces, _ := regexp.Compile(`\s+`)
	res := regSpaces.ReplaceAllString(expression, delimiter)
	res = strings.ReplaceAll(res, "(", "("+delimiter)
	res = strings.ReplaceAll(res, ")", delimiter+")")
	return strings.Split(res, delimiter)
}

func (l *lexicalAnalyzer) getOutputTokens(lexemes []lexical.Lexeme) []string {
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
func (l *lexicalAnalyzer) getOutputVariables(lexemes []lexical.Lexeme) []string {
	var result []string
	for i, lexeme := range lexemes {
		result = append(result, fmt.Sprintf("%d - %s", i+1, lexeme.Symbol))
	}
	return result
}
