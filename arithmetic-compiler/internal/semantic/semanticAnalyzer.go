package semantic

import (
	models2 "arithmetic-compiler/internal/lexical/models"
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
)

type semanticAnalyzer struct {
	tokens []models2.Token
}

func NewSemanticAnalyzer(tokens []models2.Token) *semanticAnalyzer {
	return &semanticAnalyzer{tokens: tokens}
}

func (s *semanticAnalyzer) Analyze() ([]models2.Token, error) {
	return s.tokens, s.checkErrors()
}

func (s *semanticAnalyzer) GetVars() []models2.Token {
	var vars []models2.Token
	for _, token := range s.tokens {
		if (token.Lexeme.Type == models2.Variable ||
			token.Lexeme.Type == models2.FloatVariable) && !slices.Contains(vars, token) {
			vars = append(vars, token)
		}
	}
	return vars
}

func (s *semanticAnalyzer) checkErrors() error {
	err := s.checkDivisionByZero()
	if err != nil {
		return err
	}
	return s.checkRepeatVariablesType()
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

func (s *semanticAnalyzer) checkRepeatVariablesType() error {
	for _, token := range s.tokens {
		if token.Lexeme.Type == models2.Variable {
			if err := s.checkContainsVariablesInTokens(token); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *semanticAnalyzer) checkContainsVariablesInTokens(token models2.Token) error {
	for i, tkn := range s.tokens {
		if tkn.Lexeme.Symbol == token.Lexeme.Symbol && tkn.Lexeme.Type != token.Lexeme.Type {
			return errors.New(fmt.Sprintf("semantic error! variable on %d position already defined!", i+1))
		}
	}
	return nil
}
