package main

import (
	"fmt"
	"unicode"
)

type TokenKind int

const (
	Number TokenKind = iota
	Parenthesis
	Operator
	Comma
)

type Token struct {
	Kind          TokenKind
	Value         string
	StartPosition int
}

type tokenPeeker struct {
	tokens       []Token
	currentIndex int
}

func (tp *tokenPeeker) peek() *Token {
	if tp.currentIndex > len(tp.tokens)-1 {
		return nil
	}

	return &(tp.tokens[tp.currentIndex])
}

func (tp *tokenPeeker) next() {
	tp.currentIndex++
}

func newTokenPeeker(tokens []Token, currentIndex int) *tokenPeeker {
	return &tokenPeeker{tokens: tokens, currentIndex: currentIndex}
}

func Lex(source []rune) []Token {
	var tokens []Token = []Token{}
	var position int = 0
	for position < len(source) {
		position = skipSpaces(source, position)

		if position >= len(source) {
			break
		}

		var token *Token
		if token, position = readNumber(source, position); token != nil {
			tokens = append(tokens, *token)
			continue
		}

		if token, position = readComma(source, position); token != nil {
			tokens = append(tokens, *token)
			continue
		}

		if token, position = readParenthesis(source, position); token != nil {
			tokens = append(tokens, *token)
			continue
		}

		if token, position = readOperator(source, position); token != nil {
			tokens = append(tokens, *token)
			continue
		}

		panic(fmt.Sprintf("Couldn't recognize character at position %v", position))
	}

	return tokens
}

func skipSpaces(source []rune, position int) int {
	for position < len(source) {
		if unicode.IsSpace(source[position]) {
			position++
			continue
		}

		return position
	}

	return position
}

func readNumber(source []rune, position int) (*Token, int) {
	startPosition := position
	var tokenValue []rune = []rune{}
	for position < len(source) {
		if (source[position] >= '0' && source[position] <= '9') || source[position] == '.' {
			tokenValue = append(tokenValue, source[position])
			position++
			continue
		}

		break
	}

	if len(tokenValue) > 0 {
		return &Token{Kind: Number, Value: string(tokenValue), StartPosition: startPosition}, position
	}

	return nil, position
}

func readParenthesis(source []rune, position int) (*Token, int) {
	if isParenthesis(source[position]) {
		return &Token{Kind: Parenthesis, Value: string([]rune{source[position]}), StartPosition: position}, position + 1
	}

	return nil, position
}

func readComma(source []rune, position int) (*Token, int) {
	if source[position] == ',' {
		return &Token{Kind: Comma, Value: ",", StartPosition: position}, position + 1
	}

	return nil, position
}

func readOperator(source []rune, position int) (*Token, int) {
	startPosition := position
	var tokenValue []rune = []rune{}

	if isOneCharacterOperator(source[position]) {
		return &Token{Kind: Operator, Value: string([]rune{source[position]}), StartPosition: startPosition}, position + 1
	}

	for position < len(source) {
		if !isParenthesis(source[position]) && !unicode.IsSpace(source[position]) {
			tokenValue = append(tokenValue, source[position])
			position++
			continue
		}

		break
	}

	if len(tokenValue) > 0 {
		return &Token{Kind: Operator, Value: string(tokenValue), StartPosition: startPosition}, position
	}

	return nil, position
}

func isParenthesis(value rune) bool {
	switch value {
	case '(', ')', '[', ']':
		return true
	default:
		return false
	}
}

func isOneCharacterOperator(value rune) bool {
	switch value {
	case '-', '+', '/', '*':
		return true
	default:
		return false
	}
}
