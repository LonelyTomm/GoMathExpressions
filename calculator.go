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
		minValue := args[0]
		for i := 1; i < len(args); i++ {
			if args[i] < minValue {
				minValue = args[i]
			}
		}

		return minValue
	},
	"max": func(args []float64) float64 {
		maxValue := args[0]
		for i := 1; i < len(args); i++ {
			if args[i] > maxValue {
				maxValue = args[i]
			}
		}

		return maxValue
	},
	"abs": func(args []float64) float64 {
		val := args[0]
		if val < 0 {
			return val * -1
		}

		return val
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
