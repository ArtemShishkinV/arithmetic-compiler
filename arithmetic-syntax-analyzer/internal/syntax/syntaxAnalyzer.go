package syntax

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	syntaxModels "arithmetic-syntax-analyzer/internal/syntax/models"
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
	lexeme := s.lexemes[s.pos]
	s.pos++
	return lexeme
}
func (s *syntaxAnalyzer) back() {
	s.pos--
}

func (s *syntaxAnalyzer) Analyze() (syntaxModels.Node, error) {
	return s.plusMinus()
}

func (s *syntaxAnalyzer) plusMinus() (syntaxModels.Node, error) {
	lNode, err := s.multDiv()
	if err != nil {
		return nil, err
	}
	for {
		lexeme := s.next()
		switch lexeme.Type {
		case models.OpPlus, models.OpMinus:
			rNode, err := s.multDiv()
			if err != nil {
				return nil, err
			}
			lNode = syntaxModels.NewBinaryNode(lexeme, lNode, rNode)
		case models.RightBracket, models.Unknown:
			s.back()
			return lNode, nil
		default:
			return nil, errors.New(fmt.Sprintf("syntax error! unexpected token <%s> at %d position", lexeme.Symbol, s.pos))
		}
	}
}

func (s *syntaxAnalyzer) multDiv() (syntaxModels.Node, error) {
	lNode, err := s.factor()
	if err != nil {
		return nil, err
	}
	for {
		lexeme := s.next()
		switch lexeme.Type {
		case models.OpMul, models.OpDiv:
			rNode, err := s.factor()
			if err != nil {
				return nil, err
			}
			lNode = syntaxModels.NewBinaryNode(lexeme, lNode, rNode)
		case models.RightBracket, models.OpPlus, models.OpMinus, models.Unknown:
			s.back()
			return lNode, nil
		default:
			return nil, errors.New(fmt.Sprintf("syntax error! unexpected token <%s> at %d position", lexeme.Symbol, s.pos))
		}
	}
}

func (s *syntaxAnalyzer) factor() (syntaxModels.Node, error) {
	lexeme := s.next()
	oldPos := s.pos
	switch lexeme.Type {
	case models.IntNumber:
		return syntaxModels.NewOperandNode(lexeme), nil
	case models.FloatNumber:
		return syntaxModels.NewOperandNode(lexeme), nil
	case models.Variable:
		return syntaxModels.NewOperandNode(lexeme), nil
	case models.LeftBracket:
		node, err := s.plusMinus()
		if err != nil {
			return nil, err
		}
		nextLexeme := s.next()
		if nextLexeme.Type != models.RightBracket {
			return nil, errors.New(fmt.Sprintf(
				"syntax error! missing closing bracket on token <%s> at %d position", lexeme.Symbol, oldPos))
		}
		return node, nil
	default:
		return nil, errors.New(fmt.Sprintf(
			"syntax error! token <%s> has no operand at %d position", lexeme.Symbol, oldPos))
	}
}
