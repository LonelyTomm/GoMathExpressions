package main

import "testing"

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

	for _, testCase := range testCases {
		result := isParenthesis(testCase.input)
		if result != testCase.result {
			t.Fatalf("isParentesis for %v expected to return %v got %v instead", testCase.input, testCase.result, result)
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

	for _, testCase := range testCases {
		result := skipSpaces(testCase.source, testCase.position)
		if result != testCase.result {
			t.Fatalf("skipSpaces for (%s, %v) expected to return %v got %v instead", string(testCase.source), testCase.position, testCase.result, result)
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

	for _, testCase := range testCases {
		result := isOneCharacterOperator(testCase.input)
		if result != testCase.result {
			t.Fatalf("isOneCharacterOperator for %s expected to return %v got %v instead", string(testCase.input), testCase.result, result)
		}
	}
}

func TestReadNumber(t *testing.T) {
	// TODO: implement test
}

func TestReadParenthesis(t *testing.T) {
	// TODO: implement test
}

func TestReadOperator(t *testing.T) {
	// TODO: implement test
}
