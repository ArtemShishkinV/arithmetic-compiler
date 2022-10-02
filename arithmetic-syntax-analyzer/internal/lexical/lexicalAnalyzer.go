package lexical

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type lexicalAnalyzer struct{}

func NewLexicalAnalyzer() *lexicalAnalyzer {
	return &lexicalAnalyzer{}
}

func (l *lexicalAnalyzer) Analyze(expression string) ([]models.Lexeme, error) {
	fmt.Println("#analysis-lexical")
	var lexemes []models.Lexeme
	var lexemeBuffer LexemeBuffer
	expression = l.getExpressionWithoutSpaces(expression)

	i := 0
	lexemePos := 1

	for i < len(expression) {
		lexeme, err := lexemeBuffer.ReadLexeme(expression, i)
		if err != nil {
			return nil, errors.New(err.Error() + " in " + strconv.Itoa(lexemePos) + " position")
		}
		i += len(lexeme.Symbol)
		lexemes = append(lexemes, *lexeme)
		lexemePos++
	}

	return lexemes, nil
}

func (l *lexicalAnalyzer) getExpressionWithoutSpaces(expression string) string {
	regSpaces, _ := regexp.Compile(`\s+`)
	return regSpaces.ReplaceAllString(expression, "")
}
