package models

import (
	models2 "arithmetic-compiler/internal/lexical/models"
	"fmt"
)

type ThreeAddressCode interface {
	ToString() string
	GetResult() string
	GetOperator() models2.Token
	GetOperands() []models2.Token
}

func NewThreeAddressCode(
	token models2.Token,
	order int,
	operands []models2.Token) (ThreeAddressCode, error) {
	operation, _ := GetCodeOperation(token)
	result := fmt.Sprintf("<id,%d>", order)
	if operation == Int2Float {
		return NewConvertCode(token, result, operands)
	}
	return NewOperatorCode(token, result, operands)
}
