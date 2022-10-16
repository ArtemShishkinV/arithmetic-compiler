package readers

import (
	"arithmetic-compiler/internal/lexical/models"
	"bytes"
	"strings"
)

type readerVariable struct{}

func (r *readerVariable) read(expression string, pos int) (models.LexemeType, string) {
	if checkFirstSymbolInVariableName(expression[pos]) {
		buffer := bytes.NewBufferString(string(expression[pos]))
		currentPos := pos + 1
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

		if char == '[' {
			lexType, s := r.readTypeVariable(expression[currentPos : currentPos+3])
			buffer.WriteString(s)
			return lexType, buffer.String()
		}

		if char != expression[pos] {
			return models.Variable, buffer.String()
		}
	}

	return models.Unknown, ""
}

func checkFirstSymbolInVariableName(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func (r *readerVariable) readTypeVariable(expression string) (models.LexemeType, string) {
	i := 0
	symbol := ""
	lexType := models.Unknown

	buffer := bytes.NewBufferString(string(expression[i]))
	i++

	lexType, symbol = r.getTypeVariable(expression[i])
	if lexType == models.Unknown {
		return lexType, ""
	}
	buffer.WriteString(symbol)
	i++
	if expression[i] == ']' {
		buffer.WriteString(string(expression[i]))
		return lexType, buffer.String()
	}

	return models.Unknown, ""
}

func (r *readerVariable) getTypeVariable(char uint8) (models.LexemeType, string) {
	symbol := strings.ToLower(string(char))
	lexType := models.Unknown

	switch symbol {
	case "f":
		lexType = models.FloatVariable
	case "i":
		lexType = models.Variable
	default:
		symbol = ""
	}

	return lexType, symbol
}
