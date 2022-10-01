package models

type LexemeSyntax struct {
	Pos    int
	Lexeme Lexeme
}

func NewLexemeSyntax(pos int, lexeme Lexeme) *LexemeSyntax {
	return &LexemeSyntax{
		Pos:    pos,
		Lexeme: lexeme,
	}
}
