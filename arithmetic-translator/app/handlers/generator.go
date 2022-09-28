package handlers

import (
	"arithmetic-translator/app/handlers/enum"
	"arithmetic-translator/utils"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type generator struct {
	outFileName      string
	countExpression  int
	minCountOperands int
	maxCountOperands int
}

func NewGenerator(args []string) (Handler, error) {
	if numbers, err := isValidArgs(args); err != nil {
		return nil, err
	} else {
		return &generator{
			outFileName:      args[0],
			countExpression:  numbers[0],
			minCountOperands: numbers[1],
			maxCountOperands: numbers[2],
		}, nil
	}
}

func (g *generator) Start() error {
	fmt.Println("#generating....")

	if err := utils.WriteFile(g.getExpressions(), g.outFileName); err != nil {
		return err
	}

	return nil
}

func (g *generator) getExpressions() []string {
	var expressions []string

	for i := 0; i < g.countExpression; i++ {
		expressions = append(expressions, g.generateExpression())
	}

	return expressions
}

func (g *generator) generateExpression() string {
	countOperands := rand.Intn(g.maxCountOperands-g.minCountOperands+1) + g.minCountOperands
	var expression []string

	for i := 0; i < countOperands*2-1; i++ {
		if i%2 == 0 {
			expression = append(expression, fmt.Sprintf("%v", utils.RandMapKey(enum.NumbersExpression)))
		} else {
			expression = append(expression, fmt.Sprintf("%v", utils.RandMapKey(enum.SymbolsExpression)))
		}
	}

	return strings.Join(expression, " ")
}

func isValidArgs(args []string) ([]int, error) {
	if len(args) != 4 {
		return nil, errors.New("invalid number of arguments")
	}
	countExpression, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("third arg must be int")
	}
	minCountOperands, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, errors.New("fourth arg must be int")
	}
	maxCountOperands, err := strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("fifth arg must be int")
	}
	if countExpression <= 0 || minCountOperands <= 0 || maxCountOperands <= 0 {
		return nil, errors.New("all args must be more zero")
	}
	if minCountOperands > maxCountOperands {
		return nil, errors.New("fourth arg must be more fifth")
	}

	return []int{countExpression, minCountOperands, maxCountOperands}, nil
}
