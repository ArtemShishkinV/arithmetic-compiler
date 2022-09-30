package models

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
