package models

type LexemeType string

const (
	OpPlus        LexemeType = "операция сложения"
	OpMinus       LexemeType = "операция вычитания"
	OpMul         LexemeType = "операция умножения"
	OpDiv         LexemeType = "операция деления"
	LeftBracket   LexemeType = "открывающая скобка"
	RightBracket  LexemeType = "закрывающая скобка"
	Variable      LexemeType = "идентификатор с именем целого типа"
	FloatVariable LexemeType = "идентификатор с именем вещественного типа"
	IntNumber     LexemeType = "константа целого типа"
	FloatNumber   LexemeType = "константа вещественного типа"
	Unknown       LexemeType = "неизвестный символ"
)

func IsIntType(lexemeType LexemeType) bool {
	return lexemeType == IntNumber || lexemeType == Variable
}

func IsFloatType(lexemeType LexemeType) bool {
	return lexemeType == FloatVariable || lexemeType == FloatNumber
}
