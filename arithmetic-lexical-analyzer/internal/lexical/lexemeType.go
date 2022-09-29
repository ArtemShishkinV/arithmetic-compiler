package lexical

import (
	"regexp"
)

var regVariableName, _ = regexp.Compile("[a-zA-Z_]\\w*")
var regFloatNumber, _ = regexp.Compile("[+-]?(\\d+).\\d+")
var regIntNumber, _ = regexp.Compile("[+-]?(\\d+)")

type LexemeType string

const (
	OpPlus       LexemeType = "операция сложения"
	OpMinus      LexemeType = "операция вычитания"
	OpMul        LexemeType = "операция умножения"
	OpDiv        LexemeType = "операция деления"
	LeftBracket  LexemeType = "открывающая скобка"
	RightBracket LexemeType = "закрывающая скобка"
	Variable     LexemeType = "идентификатор с именем"
	IntNumber    LexemeType = "константа целого типа"
	FloatNumber  LexemeType = "контсанта вещественного типа"
	Unknown      LexemeType = "неизвестный символ"
)

func getLexemeType(symbol string) LexemeType {
	if regVariableName.MatchString(symbol) {
		return Variable
	}
	if lexemeType := getOperand(symbol); lexemeType != Unknown {
		return lexemeType
	}
	if regFloatNumber.MatchString(symbol) {
		return FloatNumber
	}
	if regIntNumber.MatchString(symbol) {
		return IntNumber
	}
	return Unknown
}

func getOperand(symbol string) LexemeType {
	switch symbol {
	case "/":
		return OpDiv
	case "*":
		return OpMul
	case "+":
		return OpPlus
	case "-":
		return OpMinus
	case "(":
		return LeftBracket
	case ")":
		return RightBracket
	}
	return Unknown
}
