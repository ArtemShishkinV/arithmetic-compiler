package syntax

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"errors"
	"fmt"
)

/*------------------------------------------------------------------
 * PARSER RULES
 *------------------------------------------------------------------*/

//    expr : plusminus* EOF ;
//
//    plusminus: multdiv ( ( '+' | '-' ) multdiv )* ;
//
//    multdiv : factor ( ( '*' | '/' ) factor )* ;
//
//    factor : (INT | FLOAT | VARIABLE) | '(' expr ')' ;

type syntaxAnalyzer struct {
	pos     int
	lexemes []models.Lexeme
}

func NewSyntaxAnalyzer(lexemes []models.Lexeme) *syntaxAnalyzer {
	return &syntaxAnalyzer{
		pos:     0,
		lexemes: lexemes,
	}
}

func (s *syntaxAnalyzer) next() models.Lexeme {
	if len(s.lexemes) == s.pos {
		newLexeme, _ := models.NewLexeme(models.Unknown, "")
		return *newLexeme
	}
	lexeme := s.lexemes[s.pos]
	s.pos++
	return lexeme
}
func (s *syntaxAnalyzer) back() {
	s.pos--
}

func (s *syntaxAnalyzer) Analyze() (string, error) {
	return s.plusMinus()
}

func (s *syntaxAnalyzer) plusMinus() (string, error) {
	result, err := s.multDiv()
	var st string
	if err != nil {
		return result, err
	}
	for {
		lexeme := s.next()
		switch lexeme.Type {
		case models.OpPlus, models.OpMinus:
			st, err = s.multDiv()
			result = result + string(lexeme.Type) + st
			if err != nil {
				return result, err
			}
			return result, nil
		case models.RightBracket, models.Unknown:
			s.back()
			return result + lexeme.Symbol, nil
		default:
			return result + lexeme.Symbol, errors.New(fmt.Sprintf("syntax error! unexpected token <%s> at %d position", lexeme.Symbol, s.pos))
		}
	}
}

func (s *syntaxAnalyzer) multDiv() (string, error) {
	var st string
	result, err := s.factor()
	if err != nil {
		return result, err
	}
	for {
		lexeme := s.next()
		switch lexeme.Type {
		case models.OpMul, models.OpDiv:
			st, err = s.factor()
			result = result + string(lexeme.Type) + st
			if err != nil {
				return result, err
			}
			return result, nil
		case models.RightBracket, models.OpPlus, models.OpMinus, models.Unknown:
			s.back()
			return result, nil
		default:
			return result, errors.New(fmt.Sprintf("syntax error! unexpected token <%s> at %d position", lexeme.Symbol, s.pos))
		}
	}
}

func (s *syntaxAnalyzer) factor() (string, error) {
	lexeme := s.next()
	oldPos := s.pos
	switch lexeme.Type {
	case models.IntNumber:
		return lexeme.Symbol, nil
	case models.FloatNumber:
		return lexeme.Symbol, nil
	case models.Variable:
		return lexeme.Symbol, nil
	case models.LeftBracket:
		symbol, err := s.plusMinus()
		if err != nil {
			return lexeme.Symbol + symbol, err
		}
		nextLexeme := s.next()
		if nextLexeme.Type != models.RightBracket {
			return nextLexeme.Symbol, errors.New(fmt.Sprintf(
				"syntax error! missing closing bracket on token <%s> at %d position", lexeme.Symbol, oldPos))
		}
		return lexeme.Symbol + symbol + nextLexeme.Symbol, nil
	default:
		return lexeme.Symbol, errors.New(fmt.Sprintf(
			"syntax error! token <%s> has no operand at %d position", lexeme.Symbol, oldPos))
	}
}
