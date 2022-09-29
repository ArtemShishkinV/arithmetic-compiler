package lexical

import (
	"errors"
	"fmt"
	"regexp"
)

var regValidVariable, _ = regexp.Compile(`^[a-zA-Z_][a-zA-Z\d]*?$`)
var regValidFloat, _ = regexp.Compile(`^[+-]?\d+(\.\d+)?$`)

type Lexeme struct {
	Symbol string
	Type   LexemeType
}

func NewLexeme(symbol string) (*Lexeme, error) {
	lexemeType := getLexemeType(symbol)
	if lexemeType == Unknown {
		return nil, errors.New("lexical error: unknown lexeme")
	}
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
	if l.Type == Variable && !regValidVariable.MatchString(l.Symbol) {
		return errors.New(fmt.Sprintf("lexical error: invalid variable name '%s'", l.Symbol))
	}
	if l.Type == FloatNumber && !regValidFloat.MatchString(l.Symbol) {
		return errors.New(fmt.Sprintf("lexical error: invalid float constant '%s'", l.Symbol))
	}
	return nil
}
