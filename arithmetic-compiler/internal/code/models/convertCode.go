package models

import (
	models2 "arithmetic-compiler/internal/lexical/models"
	"errors"
	"fmt"
)

type convertCode struct {
	operation models2.Token
	result    string
	operand   models2.Token
}

func NewConvertCode(
	operationType models2.Token,
	result string,
	operands []models2.Token) (ThreeAddressCode, error) {
	if len(operands) != 1 {
		return nil, errors.New("invalid count operands for convert code")
	}
	return &convertCode{
		operation: operationType,
		result:    result,
		operand:   operands[0],
	}, nil
}

func (c *convertCode) GetResult() string {
	return c.result
}

func (c *convertCode) ToString() string {
	operationType, _ := GetCodeOperation(c.operation)
	return fmt.Sprintf("%s %s %s", string(operationType), c.result, c.operand.Value)
}

func (c *convertCode) GetOperator() models2.Token {
	return c.operation
}

func (c *convertCode) GetOperands() []models2.Token {
	return []models2.Token{c.operand}
}
