package main

import "strconv"

func calculate(startNode *node) float64 {
	if startNode.kind == literalNode {
		value, err := strconv.ParseFloat(startNode.value, 64)
		if err != nil {
			panic(err)
		}

		return value
	}

	if startNode.kind == listNode {
		firstValue := calculate(startNode.list[0])
		secondValue := calculate(startNode.list[1])

		switch startNode.operator {
		case "+":
			return firstValue + secondValue
		case "-":
			return firstValue - secondValue
		case "*":
			return firstValue * secondValue
		case "/":
			return firstValue / secondValue
		}
	}

	panic("Couldn't calculate expression")
}
