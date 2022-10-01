package handlers

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"arithmetic-syntax-analyzer/pkg"
	"errors"
	"fmt"
)

type syntaxAnalyzer struct{}

func newSyntaxAnalyzer() *syntaxAnalyzer {
	return &syntaxAnalyzer{}
}

func (s *syntaxAnalyzer) Start(expression string) ([][]string, error) {
	fmt.Println("#analysis-syntax")
	lexemes, err := newLexicalAnalyzer().analysis(expression)
	if err != nil {
		return nil, err
	}
	if err := s.checkSyntaxErrors(lexemes); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *syntaxAnalyzer) checkSyntaxErrors(lexemes []models.Lexeme) error {
	filteredLexemes := s.getBranchesAndOtherLexemes(lexemes)
	if err := s.checkBranchesErrors(filteredLexemes[0], filteredLexemes[1]); err != nil {
		return err
	}
	return nil
}

func (s *syntaxAnalyzer) checkBranchesErrors(leftBrackets []models.LexemeSyntax, rightBrackets []models.LexemeSyntax) error {
	var usedRightBrackets []int
	var rightBracketPos int

	if len(leftBrackets) > len(rightBrackets) {
		for i := len(leftBrackets) - 1; i >= 0; i-- {
			rightBracketPos = s.findRightBracket(leftBrackets[i].Pos, rightBrackets, usedRightBrackets)
			if rightBracketPos == -1 {
				return errors.New(fmt.Sprintf("syntax error! not found closing brace for token <%s> at position %d",
					leftBrackets[i].Lexeme.Symbol, leftBrackets[i].Pos))
			}
			usedRightBrackets = append(usedRightBrackets, rightBracketPos)
		}
	}

	return nil
}

func (s *syntaxAnalyzer) findRightBracket(lbPos int, rightBrackets []models.LexemeSyntax, usedBrackets []int) int {
	for _, v := range rightBrackets {
		if lbPos < v.Pos && !pkg.Contains(usedBrackets, v.Pos) {
			return v.Pos
		}
	}
	return -1
}

func (s *syntaxAnalyzer) checkOperandsForOperation(lexemes []models.Lexeme, pos int) {
	//err := nil
	//for err == nil {
	//	if pos == 0 && (lexemes[pos].Type == models.OpPlus || lexemes[pos].Type == models.OpMinus) {
	//		continue
	//	}
	//	err = errors.New()
	//}
}

func (s *syntaxAnalyzer) getBranchesAndOtherLexemes(lexemes []models.Lexeme) [][]models.LexemeSyntax {
	var leftBrackets []models.LexemeSyntax
	var rightBrackets []models.LexemeSyntax
	var otherLexemes []models.LexemeSyntax
	var pos int

	for i, lexeme := range lexemes {
		pos = i + 1
		switch lexeme.Type {
		case models.LeftBracket:
			leftBrackets = append(leftBrackets, *models.NewLexemeSyntax(pos, lexeme))
		case models.RightBracket:
			rightBrackets = append(rightBrackets, *models.NewLexemeSyntax(pos, lexeme))
		default:
			otherLexemes = append(otherLexemes, *models.NewLexemeSyntax(pos, lexeme))
		}
	}

	return [][]models.LexemeSyntax{leftBrackets, rightBrackets, otherLexemes}
}
