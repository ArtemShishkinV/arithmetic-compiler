package formaters

import (
	"arithmetic-compiler/internal/lexical/models"
	"fmt"
)

type varsFormatter struct {
	tokens []models.Token
}

func NewVarsFormatter(tokens []models.Token) Formatter {
	return &varsFormatter{
		tokens: tokens,
	}
}

func (v *varsFormatter) Form() []string {
	var result []string
	indexVariable := 1
	for _, token := range v.tokens {
		if token.Lexeme.Type == models.Variable {
			result = append(result, fmt.Sprintf("%d - %s", indexVariable, token.Lexeme.Symbol))
			indexVariable++
		}
	}
	return result
}
