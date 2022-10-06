package formaters

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"fmt"
)

type tokenFormatter struct {
	tokens []models.Token
}

func NewTokensFormatter(token []models.Token) Formatter {
	return &tokenFormatter{
		tokens: token,
	}
}

func (t *tokenFormatter) Form() []string {
	var result []string
	varNumber := 1
	for _, token := range t.tokens {
		if token.Lexeme.Type != models.Variable {
			result = append(result, fmt.Sprintf("<%s> - %s", token.Value, token.Lexeme.Type))
		} else {
			result = append(result, fmt.Sprintf("<%s> - %s %s", token.Value, token.Lexeme.Type, token.Lexeme.Symbol))
			varNumber++
		}
	}
	return result
}
