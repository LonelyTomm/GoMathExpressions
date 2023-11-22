package main

import "fmt"

func main() {
	var tokens = Lex([]rune("1 + 23-sin(56) + (34-8)"))

	for _, token := range tokens {
		fmt.Println(token.Value)
	}
}
