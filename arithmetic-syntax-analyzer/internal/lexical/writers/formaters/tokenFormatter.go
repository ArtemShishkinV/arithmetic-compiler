package formaters

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"fmt"
)

type tokenFormatter struct{}

func NewTokensFormatter() Formatter {
	return &tokenFormatter{}
}

func (t *tokenFormatter) Form(lexemes []models.Lexeme) []string {
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
