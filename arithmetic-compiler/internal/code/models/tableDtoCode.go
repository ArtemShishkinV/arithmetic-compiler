package models

import "arithmetic-compiler/internal/syntax/models"

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
