package main

import (
	"strconv"
)

var operators = map[string]func(args []float64) float64{
	"+": func(args []float64) float64 {
		return args[0] + args[1]
	},
	"-": func(args []float64) float64 {
		if len(args) == 1 {
			return -args[0]
		}
		return args[0] - args[1]
	},
	"/": func(args []float64) float64 {
		return args[0] / args[1]
	},
	"*": func(args []float64) float64 {
		return args[0] * args[1]
	},
	"min": func(args []float64) float64 {
		return args[0]
	},
}

func calculate(startNode *node) float64 {
	if startNode.kind == literalNode {
		value, err := strconv.ParseFloat(startNode.value, 64)
		if err != nil {
			panic(err)
		}

		return value
	}

	if startNode.kind == listNode {
		method := operators[startNode.operator]

		args := []float64{}
		for _, nd := range startNode.list {
			args = append(args, calculate(nd))
		}

		return method(args)
	}

	panic("Couldn't calculate expression")
}
