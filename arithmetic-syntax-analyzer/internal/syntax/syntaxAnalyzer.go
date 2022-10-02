package syntax

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
)

type syntaxAnalyzer struct {
	pos     int
	lexemes []models.Lexeme
}

func newSyntaxAnalyzer(lexemes []models.Lexeme) *syntaxAnalyzer {
	return &syntaxAnalyzer{
		pos:     0,
		lexemes: lexemes,
	}
}
