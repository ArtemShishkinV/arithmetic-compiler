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
	var lexType models.LexemeType
	indexVariable := 1
	for _, token := range v.tokens {
		lexType = token.Lexeme.Type
		if lexType == models.Variable || lexType == models.FloatVariable {
			result = append(result, fmt.Sprintf("%d - %s %s", indexVariable,
				token.Lexeme.Symbol, v.getSuffixByType(lexType)))
			indexVariable++
		}
	}
	return result
}

func (v *varsFormatter) getSuffixByType(lexemeType models.LexemeType) string {
	if lexemeType == models.Variable {
		return "[целый]"
	}
	return "[вещественный]"
}
