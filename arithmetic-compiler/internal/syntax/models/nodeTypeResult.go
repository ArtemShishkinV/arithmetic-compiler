package models

import "arithmetic-compiler/internal/lexical/models"

type NodeTypeResult models.LexemeType

const (
	Float   NodeTypeResult = "вещественный"
	Integer NodeTypeResult = "целый"
)
