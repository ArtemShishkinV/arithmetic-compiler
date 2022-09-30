package models

import (
	"errors"
	"fmt"
	"regexp"
)

var regValidFloat, _ = regexp.Compile(`^[+-]?\d+(\.\d+)?$`)

type Lexeme struct {
	Symbol string
	Type   LexemeType
}

func NewLexeme(lexemeType LexemeType, symbol string) (*Lexeme, error) {
	lexeme := &Lexeme{
		Symbol: symbol,
		Type:   lexemeType,
	}
	if err := lexeme.isValidLexeme(); err != nil {
		return nil, err
	}
	return lexeme, nil
}

func (l Lexeme) isValidLexeme() error {
	if l.Type == FloatNumber && !regValidFloat.MatchString(l.Symbol) {
		return errors.New(fmt.Sprintf("lexical error: invalid float constant '%s'", l.Symbol))
	}
	return nil
}
