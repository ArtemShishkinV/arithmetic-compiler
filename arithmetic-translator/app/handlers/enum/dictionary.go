package enum

import (
	"errors"
	"strconv"
)

const textError = "incorrect data"

var NumbersExpression = map[int]string{
	1: "один",
	2: "два",
	3: "три",
	4: "четыре",
	5: "пять",
	6: "шесть",
	7: "семь",
	8: "восемь",
	9: "девять",
}

var SymbolsExpression = map[string]string{
	"+": "плюс",
	"-": "минус",
	"*": "умножить на",
	"/": "делить на",
}

func GetValueFromDictionary(symbol string) (string, error) {
	number, err := strconv.Atoi(symbol)
	if err != nil {
		if val, ok := SymbolsExpression[symbol]; ok {
			return val, nil
		}
		return "", errors.New(textError)
	}
	if val, ok := NumbersExpression[number]; ok {
		return val, nil
	}
	return "", errors.New(textError)
}
