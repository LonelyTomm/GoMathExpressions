package main

import "fmt"

type nodeKind int

const (
	listNode nodeKind = iota
	literalNode
)

const maxPrecedence = 2

type node struct {
	kind     nodeKind
	operator string
	value    string
	list     []*node
}

func (nd *node) printNode() string {
	resultString := ""
	if nd.kind == literalNode {
		resultString = resultString + nd.value
	}

	if nd.kind == listNode {
		switch nd.operator {
		case "+", "-", "*", "/":
			if nd.operator == "-" && len(nd.list) == 1 {
				resultString = resultString + "(" + "-" + nd.list[0].printNode() + ")"
			} else {
				resultString = resultString + "(" + nd.list[0].printNode() + ")"
				resultString = resultString + nd.operator
				resultString = resultString + "(" + nd.list[1].printNode() + ")"
			}
		default:
			resultString = resultString + nd.operator + "("
			for i := 0; i < len(nd.list); i++ {
				resultString = resultString + nd.list[i].printNode()
				if i != len(nd.list) {
					resultString = resultString + ", "
				}
			}

			resultString = resultString + ")"
		}
	}

	return resultString
}

func parse(tokenPeeker *tokenPeeker, precedence int) *node {
	if precedence >= maxPrecedence {
		return parsePrimary(tokenPeeker)
	}

	lhs := parse(tokenPeeker, precedence+1)

	currentToken := tokenPeeker.peek()
	if currentToken != nil {
		if isBinaryInfixOperator(currentToken.Value) && getPrecedence(currentToken.Value) == precedence {
			tokenPeeker.next()
			rhs := parse(tokenPeeker, precedence)

			return &node{
				kind:     listNode,
				operator: currentToken.Value,
				list:     []*node{lhs, rhs},
			}
		}
	}

	return lhs
}

func parsePrimary(tokenPeeker *tokenPeeker) *node {
	currentToken := tokenPeeker.peek()
	tokenPeeker.next()
	if currentToken != nil {
		if currentToken.Kind == Operator && currentToken.Value == "-" {
			nextToken := tokenPeeker.peek()
			operator := currentToken.Value
			var operand *node
			if nextToken.Kind == Parenthesis && nextToken.Value == "(" {
				tokenPeeker.next()
				operand = parse(tokenPeeker, 0)
				currentToken = tokenPeeker.peek()

				if currentToken.Kind != Parenthesis || currentToken.Value != ")" {
					panic(fmt.Sprintf("Expected to get ), after unary operator %s got %s instead", operator, currentToken.Value))
				}
				tokenPeeker.next()
			} else if nextToken.Kind == Number {
				tokenPeeker.next()
				operand = &node{
					kind:  literalNode,
					value: nextToken.Value,
				}
			}

			return &node{
				kind:     listNode,
				operator: operator,
				list:     []*node{operand},
			}
		} else if currentToken.Kind == Parenthesis && currentToken.Value == "(" {
			node := parse(tokenPeeker, 0)
			currentToken = tokenPeeker.peek()
			if currentToken.Kind != Parenthesis || currentToken.Value != ")" {
				panic(fmt.Sprintf("Expected to get (, got %s instead", currentToken.Value))
			}
			tokenPeeker.next()

			return node
		} else if currentToken.Kind == Number {
			return &node{
				kind:  literalNode,
				value: currentToken.Value,
			}
		} else if currentToken.Kind == Operator {
			nextToken := tokenPeeker.peek()
			tokenPeeker.next()
			if nextToken.Kind != Parenthesis || nextToken.Value != "(" {
				panic(fmt.Sprintf("Expected to get ( after operator %s, got %s instead", currentToken.Value, nextToken.Value))
			}

			args := []*node{}
			args = append(args, parse(tokenPeeker, 0))

			nextToken = tokenPeeker.peek()
			for nextToken.Kind == Comma && nextToken.Value == "," {
				tokenPeeker.next()
				args = append(args, parse(tokenPeeker, 0))
				nextToken = tokenPeeker.peek()
			}

			if nextToken.Kind != Parenthesis || nextToken.Value != ")" {
				panic(fmt.Sprintf("Expected to get ) after operator %s, got %s instead", currentToken.Value, nextToken.Value))
			}

			tokenPeeker.next()

			return &node{
				kind:     listNode,
				operator: currentToken.Value,
				list:     args,
			}
		}
	}

	panic("Unexpected end of expression")
}

func getPrecedence(operator string) int {
	switch operator {
	case "+", "-":
		return 0
	case "*", "/":
		return 1
	default:
		return 0
	}
}

func isBinaryInfixOperator(operator string) bool {
	switch operator {
	case "+", "-", "/", "*":
		return true
	default:
		return false
	}
}
