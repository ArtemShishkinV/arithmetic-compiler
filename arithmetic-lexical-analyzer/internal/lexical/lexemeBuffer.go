package lexical

import (
	"bytes"
	"errors"
)

type LexemeBuffer struct{}

func (l *LexemeBuffer) ReadLexeme(expression string, startPos int) (*Lexeme, error) {
	if lexemeType := getOperand(string(expression[startPos])); lexemeType != Unknown {
		return NewLexeme(lexemeType, string(expression[startPos]))
	}
	lexemeType, symbol := getVariable(expression, startPos)
	if lexemeType != Unknown {
		return NewLexeme(lexemeType, symbol)
	}
	lexemeType, symbol = getFloatNumber(expression, startPos)
	if lexemeType != Unknown {
		return NewLexeme(lexemeType, symbol)
	}
	lexemeType, symbol = getIntNumber(expression, startPos)
	if lexemeType != Unknown {
		return NewLexeme(lexemeType, symbol)
	}
	return nil, errors.New("lexical error: unknown lexeme")
}

func getOperand(symbol string) LexemeType {
	switch symbol {
	case "/":
		return OpDiv
	case "*":
		return OpMul
	case "+":
		return OpPlus
	case "-":
		return OpMinus
	case "(":
		return LeftBracket
	case ")":
		return RightBracket
	}
	return Unknown
}

func getFloatNumber(expression string, startPos int) (LexemeType, string) {
	var buffer bytes.Buffer
	currentPos := startPos
	char := expression[startPos]
	isFloat := false

	for (char >= '0' && char <= '9') || char == '.' {
		if char == '.' {
			isFloat = true
		}
		buffer.WriteString(string(char))
		currentPos++
		if len(expression) == currentPos {
			char = '\n'
			break
		}
		char = expression[currentPos]
	}

	if char != expression[startPos] && isFloat {
		return FloatNumber, buffer.String()
	}

	return Unknown, ""
}

func getIntNumber(expression string, startPos int) (LexemeType, string) {
	var buffer bytes.Buffer
	currentPos := startPos
	char := expression[startPos]

	for char >= '0' && char <= '9' {
		buffer.WriteString(string(char))
		currentPos++
		if len(expression) == currentPos {
			char = '\n'
			break
		}
		char = expression[currentPos]
	}

	if char != expression[startPos] {
		return IntNumber, buffer.String()
	}

	return Unknown, ""
}

func getVariable(expression string, startPos int) (LexemeType, string) {
	if checkFirstSymbolInVariableName(expression[startPos]) {
		buffer := bytes.NewBufferString(string(expression[startPos]))
		currentPos := startPos + 1
		char := expression[currentPos]

		for char >= '0' && char <= '9' || checkFirstSymbolInVariableName(char) {
			buffer.WriteString(string(char))
			currentPos++
			if len(expression) == currentPos {
				char = '\n'
				break
			}
			char = expression[currentPos]
		}

		if char != expression[startPos] {
			return Variable, buffer.String()
		}
	}

	return Unknown, ""
}

func checkFirstSymbolInVariableName(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}
