package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Expected to get expression as argument")
	}

	expression := args[1]
	var tokens = Lex([]rune(expression))
	var tokenPeeker = newTokenPeeker(tokens, 0)

	var startNode = parse(tokenPeeker, 0)
	fmt.Println(calculate(startNode))
}
