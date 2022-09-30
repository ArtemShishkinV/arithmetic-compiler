package readers

import (
	"arithmetic-lexical-analyzer/internal/lexical/models"
)

type lexemeReader interface {
	read(string, int) (models.LexemeType, string)
}
