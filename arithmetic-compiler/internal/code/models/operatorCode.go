package models

import (
	"errors"
	"fmt"
)

type operatorCode struct {
	operation CodeOperationType
	result    string
	first     string
	second    string
}

func NewOperatorCode(
	operationType CodeOperationType,
	result string,
	operands []string) (ThreeAddressCode, error) {
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
	return fmt.Sprintf("%s %s %s %s", string(o.operation), o.result, o.first, o.second)
}

func (o *operatorCode) GetOperator() string {
	return string(o.operation)
}

func (o *operatorCode) GetOperands() []string {
	return []string{o.first, o.second}
}
