package models

type Token struct {
	Lexeme Lexeme
	Value  string
}

func NewToken(lexeme Lexeme, token string) *Token {
	return &Token{
		Lexeme: lexeme,
		Value:  token,
	}
}
