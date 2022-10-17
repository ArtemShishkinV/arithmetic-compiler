package semantic

import (
	models2 "arithmetic-compiler/internal/lexical/models"
	"errors"
	"fmt"
)

type semanticAnalyzer struct {
	tokens []models2.Token
}

func NewSemanticAnalyzer(tokens []models2.Token) *semanticAnalyzer {
	return &semanticAnalyzer{tokens: tokens}
}

func (s *semanticAnalyzer) Analyze() ([]models2.Token, error) {
	return s.tokens, s.checkDivisionByZero()
}

func (s *semanticAnalyzer) checkDivisionByZero() error {
	for i, token := range s.tokens {
		if token.Lexeme.Type == models2.OpDiv &&
			(len(s.tokens)-1 != i) && s.tokens[i+1].Value == "0" {
			return errors.New(fmt.Sprintf("semantic error! division by zero on %d position", i+1))
		}
	}
	return nil
}
