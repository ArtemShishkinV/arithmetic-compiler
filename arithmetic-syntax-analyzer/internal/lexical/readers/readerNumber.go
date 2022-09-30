package readers

import (
	"arithmetic-syntax-analyzer/internal/lexical/models"
	"bytes"
)

type readerNumber struct{}

func (r *readerNumber) read(expression string, pos int) (models.LexemeType, string) {
	var buffer bytes.Buffer
	isFloat := false

	result := ReadNumbers(expression, pos)
	if len(result) == 0 {
		return models.Unknown, ""
	}

	currentPos := pos + len(result)
	buffer.WriteString(result)

	for currentPos != len(expression) && expression[currentPos] == '.' {
		isFloat = true
		result = ReadNumbers(expression, currentPos+1)
		buffer.WriteString(string(expression[currentPos]))
		buffer.WriteString(result)
		currentPos += len(result) + 1
	}

	if isFloat {
		return models.FloatNumber, buffer.String()
	}

	return models.IntNumber, buffer.String()
}
