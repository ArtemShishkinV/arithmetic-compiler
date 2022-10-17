package models

import "arithmetic-compiler/internal/lexical/models"

type Node interface {
	ToStringNode() string
	GetToken() models.Token
	GetNodeResult() NodeTypeResult
}
