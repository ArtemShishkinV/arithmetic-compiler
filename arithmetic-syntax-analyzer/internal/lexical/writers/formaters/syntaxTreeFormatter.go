package formaters

import "arithmetic-syntax-analyzer/internal/lexical/models"

type syntaxTreeFormatter struct{}

func NewSyntaxTreeFormatter() Formatter {
	return &syntaxTreeFormatter{}
}

func (s *syntaxTreeFormatter) Form(lexemes []models.Lexeme) []string {
	//TODO implement me
	panic("implement me")
}
