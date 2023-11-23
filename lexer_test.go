package main

import "testing"

func TestLex(t *testing.T) {
	testCases := []struct {
		source         []rune
		expectedTokens []Token
	}{
		{
			source: []rune("1 -sin(56)"),
			expectedTokens: []Token{
				{Kind: Number, Value: "1", StartPosition: 0},
				{Kind: Operator, Value: "-", StartPosition: 2},
				{Kind: Operator, Value: "sin", StartPosition: 3},
				{Kind: Parenthesis, Value: "(", StartPosition: 6},
				{Kind: Number, Value: "56", StartPosition: 7},
				{Kind: Parenthesis, Value: ")", StartPosition: 9},
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		tokens := Lex(testCase.source)

		if len(tokens) != len(testCase.expectedTokens) {
			t.Fatalf("TCIndex %v: Expected to get %v tokens, got %v instead", testCaseIndex, len(testCase.expectedTokens), len(tokens))
		}

		for tokenIndex, token := range tokens {
			expectedToken := testCase.expectedTokens[tokenIndex]
			if !compareTokens(&token, &expectedToken) {
				t.Fatalf("TCIndex %v: Expected to get %+v token, got %+v token instead", testCaseIndex, expectedToken, token)
			}
		}
	}
}

func TestIsParenthesis(t *testing.T) {
	testCases := []struct {
		input  rune
		result bool
	}{
		{
			input:  '(',
			result: true,
		},
		{
			input:  ')',
			result: true,
		},
		{
			input:  '[',
			result: true,
		},
		{
			input:  ']',
			result: true,
		},
		{
			input:  ' ',
			result: false,
		},
		{
			input:  'a',
			result: false,
		},
		{
			input:  '9',
			result: false,
		},
	}

	for testCaseIndex, testCase := range testCases {
		result := isParenthesis(testCase.input)
		if result != testCase.result {
			t.Fatalf("TCIndex %v: isParentesis for %v expected to return %v got %v instead", testCaseIndex, testCase.input, testCase.result, result)
		}
	}
}

func TestSkipSpaces(t *testing.T) {
	testCases := []struct {
		source   []rune
		position int
		result   int
	}{
		{
			source:   []rune("    1"),
			position: 0,
			result:   4,
		},
	}

	for testCaseIndex, testCase := range testCases {
		result := skipSpaces(testCase.source, testCase.position)
		if result != testCase.result {
			t.Fatalf("TCIndex %v: skipSpaces for (%s, %v) expected to return %v got %v instead", testCaseIndex, string(testCase.source), testCase.position, testCase.result, result)
		}
	}
}

func TestIsOneCharacterOperator(t *testing.T) {
	testCases := []struct {
		input  rune
		result bool
	}{
		{
			input:  '+',
			result: true,
		},
		{
			input:  '-',
			result: true,
		},
		{
			input:  '/',
			result: true,
		},
		{
			input:  '*',
			result: true,
		},
		{
			input:  ' ',
			result: false,
		},
		{
			input:  'a',
			result: false,
		},
		{
			input:  '0',
			result: false,
		},
	}

	for testCaseIndex, testCase := range testCases {
		result := isOneCharacterOperator(testCase.input)
		if result != testCase.result {
			t.Fatalf("TCIndex %v: isOneCharacterOperator for %s expected to return %v got %v instead", testCaseIndex, string(testCase.input), testCase.result, result)
		}
	}
}

func TestReadNumber(t *testing.T) {
	testCases := []struct {
		source           []rune
		position         int
		expectedToken    *Token
		expectedPosition int
	}{
		{
			source:           []rune("- 23.5  )"),
			position:         2,
			expectedToken:    &Token{Kind: Number, Value: "23.5", StartPosition: 2},
			expectedPosition: 6,
		},
		{
			source:           []rune("- 23.53.6888  )"),
			position:         2,
			expectedToken:    &Token{Kind: Number, Value: "23.53.6888", StartPosition: 2},
			expectedPosition: 12,
		},
		{
			source:           []rune("- + )"),
			position:         0,
			expectedToken:    nil,
			expectedPosition: 0,
		},
	}

	for testCaseIndex, testCase := range testCases {
		gotToken, gotPosisition := readNumber(testCase.source, testCase.position)

		if !compareTokens(gotToken, testCase.expectedToken) {
			t.Fatalf("TCIndex %v: Expected token to be %+v got %+v instead", testCaseIndex, testCase.expectedToken, gotToken)
		}

		if testCase.expectedPosition != gotPosisition {
			t.Fatalf("TCIndex %v: Expected position to be %v got %v instead", testCaseIndex, testCase.expectedPosition, gotPosisition)
		}
	}
}

func TestReadParenthesis(t *testing.T) {
	testCases := []struct {
		source           []rune
		position         int
		expectedToken    *Token
		expectedPosition int
	}{
		{
			source:           []rune("((()))"),
			position:         2,
			expectedToken:    &Token{Kind: Parenthesis, Value: "(", StartPosition: 2},
			expectedPosition: 3,
		},
		{
			source:           []rune("("),
			position:         0,
			expectedToken:    &Token{Kind: Parenthesis, Value: "(", StartPosition: 0},
			expectedPosition: 1,
		},
		{
			source:           []rune("- + )"),
			position:         0,
			expectedToken:    nil,
			expectedPosition: 0,
		},
	}

	for testCaseIndex, testCase := range testCases {
		gotToken, gotPosisition := readParenthesis(testCase.source, testCase.position)

		if !compareTokens(gotToken, testCase.expectedToken) {
			t.Fatalf("TCIndex %v: Expected token to be %+v got %+v instead", testCaseIndex, testCase.expectedToken, gotToken)
		}

		if testCase.expectedPosition != gotPosisition {
			t.Fatalf("TCIndex %v: Expected position to be %v got %v instead", testCaseIndex, testCase.expectedPosition, gotPosisition)
		}
	}
}

func TestReadOperator(t *testing.T) {
	testCases := []struct {
		source           []rune
		position         int
		expectedToken    *Token
		expectedPosition int
	}{
		{
			source:           []rune("-2345"),
			position:         0,
			expectedToken:    &Token{Kind: Operator, Value: "-", StartPosition: 0},
			expectedPosition: 1,
		},
		{
			source:           []rune("- sinx(256)  )"),
			position:         2,
			expectedToken:    &Token{Kind: Operator, Value: "sinx", StartPosition: 2},
			expectedPosition: 6,
		},
		{
			source:           []rune(")2 + )"),
			position:         0,
			expectedToken:    nil,
			expectedPosition: 0,
		},
	}

	for testCaseIndex, testCase := range testCases {
		gotToken, gotPosisition := readOperator(testCase.source, testCase.position)

		if !compareTokens(gotToken, testCase.expectedToken) {
			t.Fatalf("TCIndex %v: Expected token to be %+v got %+v instead", testCaseIndex, testCase.expectedToken, gotToken)
		}

		if testCase.expectedPosition != gotPosisition {
			t.Fatalf("TCIndex %v: Expected position to be %v got %v instead", testCaseIndex, testCase.expectedPosition, gotPosisition)
		}
	}
}

func compareTokens(tokenOne *Token, tokenTwo *Token) bool {
	if tokenOne == nil || tokenTwo == nil {
		return tokenOne == tokenTwo
	}

	if tokenOne.Kind == tokenTwo.Kind &&
		tokenOne.Value == tokenTwo.Value &&
		tokenOne.StartPosition == tokenTwo.StartPosition {
		return true
	}

	return false
}
