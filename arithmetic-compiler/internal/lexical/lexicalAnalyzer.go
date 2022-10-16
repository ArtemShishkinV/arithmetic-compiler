package lexical

import (
	"arithmetic-compiler/internal/lexical/models"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type lexicalAnalyzer struct{}

func NewLexicalAnalyzer() *lexicalAnalyzer {
	return &lexicalAnalyzer{}
}

func (l *lexicalAnalyzer) Analyze(expression string) ([]models.Lexeme, error) {
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
		if lexeme.Type == models.Variable || lexeme.Type == models.FloatVariable {
			lexeme = l.sliceTypeVarSymbols(*lexeme)
		}
		lexemes = append(lexemes, *lexeme)
		lexemePos++
	}

	return lexemes, nil
}

func (l *lexicalAnalyzer) getExpressionWithoutSpaces(expression string) string {
	regSpaces, _ := regexp.Compile(`\s+`)
	return regSpaces.ReplaceAllString(expression, "")
}

func (l *lexicalAnalyzer) sliceTypeVarSymbols(lexeme models.Lexeme) *models.Lexeme {
	if strings.Contains(lexeme.Symbol, "[") {
		lexeme.Symbol = lexeme.Symbol[:len(lexeme.Symbol)-3]
		fmt.Println(lexeme.Symbol)
		//i -= 3
	}
	return &lexeme
}
