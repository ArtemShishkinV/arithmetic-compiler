package syntax

import (
	"arithmetic-compiler/internal/lexical/models"
	syntaxModels "arithmetic-compiler/internal/syntax/models"
	"errors"
	"fmt"
)

/*------------------------------------------------------------------
 * PARSER RULES
 *------------------------------------------------------------------*/

//    expr ::= <plusminus> * EOF ;
//
//    plusminus ::= <multdiv> { ( '+' | '-' ) <multdiv> } ;
//
//    multdiv ::= <factor> { ( '*' | '/' ) <factor> } ;
//
//    factor ::= (INT | FLOAT | VARIABLE) | '(' expr ')' ;

type syntaxAnalyzer struct {
	pos    int
	tokens []models.Token
}

func NewSyntaxAnalyzer(tokens []models.Token) *syntaxAnalyzer {
	return &syntaxAnalyzer{
		pos:    0,
		tokens: tokens,
	}
}

func (s *syntaxAnalyzer) next() models.Token {
	token := s.tokens[s.pos]
	s.pos++
	return token
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
		token := s.next()
		switch token.Lexeme.Type {
		case models.OpPlus, models.OpMinus:
			rNode, err := s.multDiv()
			if err != nil {
				return nil, err
			}
			lNode = syntaxModels.NewBinaryNode(token, lNode, rNode)
		case models.RightBracket, models.Unknown:
			s.back()
			return lNode, nil
		default:
			return nil, errors.New(fmt.Sprintf("syntax error! unexpected token <%s> at %d position", token.Value, s.pos))
		}
	}
}

func (s *syntaxAnalyzer) multDiv() (syntaxModels.Node, error) {
	lNode, err := s.factor()
	if err != nil {
		return nil, err
	}
	for {
		token := s.next()
		switch token.Lexeme.Type {
		case models.OpMul, models.OpDiv:
			rNode, err := s.factor()
			if err != nil {
				return nil, err
			}
			lNode = syntaxModels.NewBinaryNode(token, lNode, rNode)
		case models.RightBracket, models.OpPlus, models.OpMinus, models.Unknown:
			s.back()
			return lNode, nil
		default:
			return nil, errors.New(fmt.Sprintf("syntax error! unexpected token <%s> at %d position", token.Value, s.pos))
		}
	}
}

func (s *syntaxAnalyzer) factor() (syntaxModels.Node, error) {
	token := s.next()
	oldPos := s.pos
	switch token.Lexeme.Type {
	case models.IntNumber:
		return syntaxModels.NewOperandNode(token), nil
	case models.FloatNumber:
		return syntaxModels.NewOperandNode(token), nil
	case models.Variable:
		return syntaxModels.NewOperandNode(token), nil
	case models.LeftBracket:
		node, err := s.plusMinus()
		if err != nil {
			return nil, err
		}
		nextToken := s.next()
		if nextToken.Lexeme.Type != models.RightBracket {
			return nil, errors.New(fmt.Sprintf(
				"syntax error! missing closing bracket on token <%s> at %d position", token.Value, oldPos))
		}
		return node, nil
	default:
		return nil, errors.New(fmt.Sprintf(
			"syntax error! token <%s> has no operand at %d position", token.Value, oldPos))
	}
}
