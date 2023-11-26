package main

type nodeKind int

const (
	listNode nodeKind = iota
	literalNode
)

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
		resultString = resultString + "(" + nd.list[0].printNode() + ")"
		resultString = resultString + nd.operator
		resultString = resultString + "(" + nd.list[1].printNode() + ")"
	}

	return resultString
}

func parse(tokenPeeker *tokenPeeker) *node {
	var currentToken = tokenPeeker.peek()
	tokenPeeker.next()
	var ltNode *node
	if currentToken.Kind == Number {
		ltNode = &node{kind: literalNode, value: currentToken.Value}
	}

	if currentToken.Kind == Parenthesis && currentToken.Value == "(" {
		ltNode = parse(tokenPeeker)
	}

	if ltNode == nil {
		panic("Wasn't supposed to get here. Probably error in syntax of expression")
	}

	return parseOperatorPrecedence(ltNode, tokenPeeker, 0)
}

func parseOperatorPrecedence(lhs *node, tp *tokenPeeker, minPrecedence int) *node {
	for {
		var currentToken = tp.peek()

		if currentToken == nil {
			break
		}
		var precedence = getPrecedence(currentToken.Value)
		if currentToken.Kind != Operator || precedence < minPrecedence {
			break
		}

		var operator = currentToken.Value
		tp.next()

		currentToken = tp.peek()

		var rhs *node
		if currentToken.Kind == Number {
			rhs = &node{kind: literalNode, value: currentToken.Value}
		} else if currentToken.Kind == Parenthesis && currentToken.Value == "(" {
			tp.next()
			rhs = parse(tp)
		}

		tp.next()

		for {
			currentToken = tp.peek()

			if currentToken == nil {
				break
			}

			if currentToken.Kind == Parenthesis && currentToken.Value == ")" {
				tp.next()
				return &node{kind: listNode, operator: operator, list: []*node{lhs, rhs}}
			}

			var nextOperationPrecedence = getPrecedence(currentToken.Value)
			if nextOperationPrecedence <= precedence {
				break
			}

			rhs = parseOperatorPrecedence(rhs, tp, nextOperationPrecedence)
		}

		var resultNode = &node{kind: listNode, operator: operator, list: []*node{lhs, rhs}}
		lhs = resultNode
	}

	return lhs
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
