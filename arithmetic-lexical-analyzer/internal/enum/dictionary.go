package enum

import (
	"errors"
	"regexp"
)

var regVariableName, _ = regexp.Compile("[a-zA-Z_]\\w*")

var SymbolsExpression = map[string]string{
	"+": "операция сложения",
	"-": "операция вычитания",
	"*": "операция умножения",
	"/": "операция деления",
}

func GetToken(symbol string) (string, error) {
	if regVariableName.MatchString(symbol) {
		return "id", nil
	}
	if val, ok := SymbolsExpression[symbol]; ok {
		return val, nil
	}
	return "", errors.New("invalid symbol")
}
