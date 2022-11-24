package models

import (
	models2 "arithmetic-compiler/internal/lexical/models"
	"errors"
	"fmt"
)

type operatorCode struct {
	operation models2.Token
	result    string
	first     models2.Token
	second    models2.Token
}

func NewOperatorCode(
	operationType models2.Token,
	result string,
	operands []models2.Token) (ThreeAddressCode, error) {
	if len(operands) != 2 {
		return nil, errors.New("invalid count operands for math operator code")
	}
	return &operatorCode{
		operation: operationType,
		result:    result,
		first:     operands[0],
		second:    operands[1],
	}, nil
}

func (o *operatorCode) GetResult() string {
	return o.result
}

func (o *operatorCode) ToString() string {
	operationType, _ := GetCodeOperation(o.operation)
	return fmt.Sprintf("%s %s %s %s", string(operationType), o.result, o.first.Value, o.second.Value)
}

func (o *operatorCode) GetOperator() models2.Token {
	return o.operation
}

func (o *operatorCode) GetOperands() []models2.Token {
	return []models2.Token{o.first, o.second}
}
