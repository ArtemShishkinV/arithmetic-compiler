package readers

import (
	"arithmetic-lexical-analyzer/internal/lexical/models"
)

type ReaderDefault struct {
	readers []lexemeReader
}

func NewReader() *ReaderDefault {
	return &ReaderDefault{
		readers: []lexemeReader{
			&readerOperand{},
			&readerNumber{},
			&readerVariable{},
		},
	}
}

func (r *ReaderDefault) Read(expression string, pos int) (models.LexemeType, string) {
	lexemeType := models.Unknown
	value := ""

	for _, reader := range r.readers {
		lexemeType, value = reader.read(expression, pos)
		if lexemeType != models.Unknown {
			return lexemeType, value
		}
	}

	return models.Unknown, ""
}
