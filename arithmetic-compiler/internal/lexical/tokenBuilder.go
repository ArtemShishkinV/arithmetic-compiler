package lexical

import (
	"arithmetic-compiler/internal/lexical/models"
	"strconv"
)

type TokenBuilder struct{}

func NewTokenBuilder() *TokenBuilder {
	return &TokenBuilder{}
}

func (t *TokenBuilder) GetTokens(lexemes []models.Lexeme) []models.Token {
	var tokens []models.Token
	values := t.getTokenValues(lexemes)
	for i, lexeme := range lexemes {
		tokens = append(tokens, *models.NewToken(lexeme, values[i]))
	}
	return tokens
}

func (t *TokenBuilder) getTokenValues(lexemes []models.Lexeme) []string {
	var result []string
	varNumber := 1
	for _, lexeme := range lexemes {
		if lexeme.Type != models.Variable {
			result = append(result, lexeme.Symbol)
		} else {
			result = append(result, "id, "+strconv.Itoa(varNumber))
			varNumber++
		}
	}
	return result
}
