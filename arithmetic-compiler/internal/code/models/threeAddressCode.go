package models

import (
	models2 "arithmetic-compiler/internal/lexical/models"
	"fmt"
)

type ThreeAddressCode interface {
	ToString() string
	GetResult() string
}

func NewThreeAddressCode(
	lexemeType models2.LexemeType,
	order int,
	operands []string) (ThreeAddressCode, error) {
	operation, _ := GetCodeOperation(lexemeType)
	result := fmt.Sprintf("<id,%d>", order)
	if operation == Int2Float {
		return NewConvertCode(operation, result, operands)
	}
	return NewOperatorCode(operation, result, operands)
}
