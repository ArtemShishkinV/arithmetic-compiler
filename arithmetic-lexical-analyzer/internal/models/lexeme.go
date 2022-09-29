package models

import (
	"errors"
	"regexp"
	"strconv"
)

var regVariableName, _ = regexp.Compile("[a-zA-Z_]\\w*")
var regFloatNumber, _ = regexp.Compile("[+-]?(\\d+).\\d+")
var regIntNumber, _ = regexp.Compile("[+-]?(\\d+)")

type Lexeme struct {
	Symbol          string
	PosInExpression int
	Type            LexemeType
}

//TODO: change method getTypeLexeme and refactoring code

func GetLexeme(symbol string, pos int) (*Lexeme, error) {
	if regVariableName.MatchString(symbol) {
		return newLexeme(symbol, pos, Variable), nil
	}
	if typeLex, ok := getOperand(symbol); ok {
		return newLexeme(symbol, pos, typeLex), nil
	}
	if regFloatNumber.MatchString(symbol) {
		return newLexeme(symbol, pos, FloatNumber), nil
	}
	if regIntNumber.MatchString(symbol) {
		return newLexeme(symbol, pos, IntNumber), nil
	}
	return nil, errors.New("lexical error: invalid symbol in " + strconv.Itoa(pos))
}

func newLexeme(symbol string, pos int, lexemeType LexemeType) *Lexeme {
	return &Lexeme{
		Symbol:          symbol,
		PosInExpression: pos,
		Type:            lexemeType,
	}
}

func getOperand(symbol string) (LexemeType, bool) {
	switch symbol {
	case "/":
		return OpDiv, true
	case "*":
		return OpMul, true
	case "+":
		return OpPlus, true
	case "-":
		return OpMinus, true
	case "(":
		return LeftBracket, true
	case ")":
		return RightBracket, true
	}
	return Unknown, false
}
