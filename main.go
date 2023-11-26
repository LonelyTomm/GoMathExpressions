package main

import "fmt"

func main() {
	var tokens = Lex([]rune("(1 * 3 + ((8 / 2) *(3+4)))"))
	var tokenPeeker = newTokenPeeker(tokens, 0)

	var startNode = parse(tokenPeeker)
	fmt.Printf(startNode.operator)
	return
}
