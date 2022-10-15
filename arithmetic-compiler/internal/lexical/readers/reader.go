package readers

import (
	"arithmetic-compiler/internal/lexical/models"
)

type lexemeReader interface {
	read(string, int) (models.LexemeType, string)
}
