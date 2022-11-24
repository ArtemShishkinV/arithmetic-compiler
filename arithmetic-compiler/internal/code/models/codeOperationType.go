package models

import (
	"arithmetic-compiler/internal/lexical/models"
	"errors"
)

type CodeOperationType string

const (
	Add       CodeOperationType = "add"
	Sub       CodeOperationType = "sub"
	Mul       CodeOperationType = "mul"
	Div       CodeOperationType = "div"
	Int2Float CodeOperationType = "i2f"
	Unknown   CodeOperationType = "неизвестный код операции"
)

func GetCodeOperation(token models.Token) (CodeOperationType, error) {
	lexemeType := token.Lexeme.Type
	switch lexemeType {
	case models.OpPlus:
		return Add, nil
	case models.OpMinus:
		return Sub, nil
	case models.OpMul:
		return Mul, nil
	case models.OpDiv:
		return Div, nil
	case models.Int2Float:
		return Int2Float, nil
	}
	return Unknown, errors.New(string(Unknown))
}
