package models

import (
	"arithmetic-compiler/internal/syntax/models"
	"fmt"
)

type TableDtoCode struct {
	codeId string
	name   string
	result models.NodeTypeResult
}

func NewTableDtoCode(codeId string, name string, result models.NodeTypeResult) TableDtoCode {
	return TableDtoCode{
		codeId: codeId,
		name:   name,
		result: result,
	}
}

func (c *TableDtoCode) ToString() string {
	return fmt.Sprintf("%s - %s, %s", c.codeId, c.name, c.result)
}
