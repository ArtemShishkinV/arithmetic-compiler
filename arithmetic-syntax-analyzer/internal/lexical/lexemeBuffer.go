package lexical

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"arithmetic-syntax-analyzer/internal/lexical/readers"
	"errors"
	"fmt"
)

type LexemeBuffer struct{}

func (l *LexemeBuffer) ReadLexeme(expression string, startPos int) (*models.Lexeme, error) {
	lexemeReader := readers.NewReader()
	lexemeType, symbol := lexemeReader.Read(expression, startPos)
	if lexemeType != models.Unknown {
		return models.NewLexeme(lexemeType, symbol)
	}
	fmt.Println(symbol)
	return nil, errors.New("lexical error: unknown lexeme")
}
