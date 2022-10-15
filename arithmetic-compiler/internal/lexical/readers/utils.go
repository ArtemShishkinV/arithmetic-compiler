package readers

import (
	"bytes"
)

func ReadNumbers(expression string, pos int) string {
	var buffer bytes.Buffer
	currentPos := pos
	char := expression[pos]

	for char >= '0' && char <= '9' {
		buffer.WriteString(string(char))
		currentPos++
		if len(expression) == currentPos {
			char = '\n'
			break
		}
		char = expression[currentPos]
	}

	return buffer.String()
}
