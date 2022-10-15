package readers

import (
	"arithmetic-compiler/internal/lexical/models"
	"bytes"
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

		if char != expression[pos] {
			return models.Variable, buffer.String()
		}
	}

	return models.Unknown, ""
}

func checkFirstSymbolInVariableName(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}
