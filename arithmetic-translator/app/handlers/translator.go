package handlers

import (
	"arithmetic-translator/app/handlers/enum"
	"arithmetic-translator/utils"
	"errors"
	"fmt"
	"strings"
)

type Translator struct {
	srcFileName string
	outFileName string
}

func NewTranslator(args []string) (Handler, error) {
	if len(args) == 2 {
		return &Translator{
			srcFileName: args[0],
			outFileName: args[1],
		}, nil
	}
	return nil, errors.New("invalid number of arguments")
}

func (t *Translator) Start() error {
	fmt.Println("#translating....")
	var translatedExpressions []string

	expressions, err := utils.ReadFileLines(t.srcFileName)
	if err != nil {
		return err
	}

	translatedExpressions, err = t.getTranslatedExpressions(expressions)
	if err != nil {
		return err
	}

	if err = utils.WriteFile(translatedExpressions, t.outFileName); err != nil {
		return err
	}

	return nil
}

func (t *Translator) getTranslatedExpressions(expressions []string) ([]string, error) {
	var translatedExpressions []string
	copy(translatedExpressions, expressions)
	//singleSpacePattern := regexp.MustCompile(`\s+`)

	for _, expression := range expressions {
		translatedExpression, err := t.translateExpression(expression)
		if err != nil {
			return nil, err
		}
		translatedExpressions = append(translatedExpressions, translatedExpression)
	}

	return translatedExpressions, nil
}

func (t *Translator) translateExpression(expression string) (string, error) {
	translatedExpression := strings.Split(expression, " ")
	var symbol string
	var err error
	var result []string

	for _, v := range translatedExpression {
		symbol, err = t.getTranslatedSymbol(v)
		if err != nil {
			return "", err
		}
		result = append(result, symbol)
	}

	return strings.Join(result, " "), nil
}

func (t *Translator) getTranslatedSymbol(srcSymbol string) (string, error) {
	symbol := srcSymbol
	var err error

	symbol, err = enum.GetValueFromDictionary(symbol)
	if err != nil {
		return "", err
	}
	return symbol, nil
}
