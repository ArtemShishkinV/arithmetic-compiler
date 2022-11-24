package models

import (
	"errors"
	"fmt"
)

type convertCode struct {
	operation CodeOperationType
	result    string
	operand   string
}

func NewConvertCode(
	operationType CodeOperationType,
	result string,
	operands []string) (ThreeAddressCode, error) {
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
	return fmt.Sprintf("%s %s %s", string(c.operation), c.result, c.operand)
}

func (c *convertCode) GetOperator() string {
	return string(c.operation)
}

func (c *convertCode) GetOperands() []string {
	return []string{c.operand}
}
