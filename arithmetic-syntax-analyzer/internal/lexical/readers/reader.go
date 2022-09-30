package readers

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
)

type lexemeReader interface {
	read(string, int) (models.LexemeType, string)
}
