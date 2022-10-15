package readers

import "arithmetic-compiler/internal/lexical/models"

type readerOperand struct{}

func (r *readerOperand) read(expression string, pos int) (models.LexemeType, string) {
	char := string(expression[pos])
	switch char {
	case "/":
		return models.OpDiv, char
	case "*":
		return models.OpMul, char
	case "+":
		return models.OpPlus, char
	case "-":
		return models.OpMinus, char
	case "(":
		return models.LeftBracket, char
	case ")":
		return models.RightBracket, char
	}
	return models.Unknown, ""
}
