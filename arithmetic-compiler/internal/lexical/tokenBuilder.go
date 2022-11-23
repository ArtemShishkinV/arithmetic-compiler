package lexical

import (
	"arithmetic-compiler/internal/lexical/models"
	"golang.org/x/exp/slices"
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
	var variables []models.Lexeme
	for _, lexeme := range lexemes {
		if lexeme.Type != models.Variable &&
			lexeme.Type != models.FloatVariable {
			result = append(result, lexeme.Symbol)
		} else {
			if !slices.Contains(variables, lexeme) {
				variables = append(variables, lexeme)
			}
			result = append(result, "id, "+strconv.Itoa(len(variables)))
		}
	}
	return result
}
