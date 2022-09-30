package handlers

import (
	"arithmetic-lexical-analyzer/internal/lexical"
	"arithmetic-lexical-analyzer/internal/lexical/models"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type LexicalAnalyzer struct{}

func NewLexicalAnalyzer() *LexicalAnalyzer {
	return &LexicalAnalyzer{}
}

func (l *LexicalAnalyzer) Start(expression string) ([]string, []string, error) {
	fmt.Println("#analysis")
	lexemes, err := l.analysis(l.getExpressionWithoutSpaces(expression))
	if err != nil {
		return nil, nil, err
	}
	tokens := l.getOutputTokens(lexemes)
	tableVars := l.getOutputVariables(l.getVariablesFromLexemes(lexemes))
	return tokens, tableVars, nil
}

func (l *LexicalAnalyzer) analysis(expression string) ([]models.Lexeme, error) {
	var lexemes []models.Lexeme
	var lexemeBuffer lexical.LexemeBuffer

	i := 0
	lexemePos := 1

	for i < len(expression) {
		lexeme, err := lexemeBuffer.ReadLexeme(expression, i)
		if err != nil {
			return nil, errors.New(err.Error() + " in " + strconv.Itoa(lexemePos) + " position")
		}
		i += len(lexeme.Symbol)
		lexemes = append(lexemes, *lexeme)
		lexemePos++
	}

	return lexemes, nil
}

func (l *LexicalAnalyzer) getExpressionWithoutSpaces(expression string) string {
	regSpaces, _ := regexp.Compile(`\s+`)
	return regSpaces.ReplaceAllString(expression, "")
}

func (l *LexicalAnalyzer) getVariablesFromLexemes(lexemes []models.Lexeme) []models.Lexeme {
	var varLexemes []models.Lexeme
	for _, lexeme := range lexemes {
		if lexeme.Type == models.Variable {
			varLexemes = append(varLexemes, lexeme)
		}
	}
	return varLexemes
}

func (l *LexicalAnalyzer) getOutputTokens(lexemes []models.Lexeme) []string {
	var result []string
	varNumber := 1
	for _, lexeme := range lexemes {
		if lexeme.Type != models.Variable {
			result = append(result, fmt.Sprintf("<%s> - %s", lexeme.Symbol, lexeme.Type))
		} else {
			result = append(result, fmt.Sprintf("<id, %d> - %s %s", varNumber, lexeme.Type, lexeme.Symbol))
			varNumber++
		}
	}
	return result
}
func (l *LexicalAnalyzer) getOutputVariables(lexemes []models.Lexeme) []string {
	var result []string
	for i, lexeme := range lexemes {
		result = append(result, fmt.Sprintf("%d - %s", i+1, lexeme.Symbol))
	}
	return result
}
