package formaters

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"fmt"
)

type varsFormatter struct{}

func NewVarsFormatter() Formatter {
	return &varsFormatter{}
}

func (v *varsFormatter) Form(lexemes []models.Lexeme) []string {
	var result []string
	indexVariable := 1
	for _, lexeme := range lexemes {
		if lexeme.Type == models.Variable {
			result = append(result, fmt.Sprintf("%d - %s", indexVariable, lexeme.Symbol))
			indexVariable++
		}
	}
	return result
}
