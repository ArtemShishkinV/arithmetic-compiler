package models

import "arithmetic-compiler/internal/lexical/models"

type NodeTypeResult models.LexemeType

const (
	Float   NodeTypeResult = "вещественный"
	Integer NodeTypeResult = "целый"
)

func GetTypeResult(lexeme models.Lexeme) NodeTypeResult {
	if models.IsFloatType(lexeme.Type) {
		return Float
	}
	return Integer
}
