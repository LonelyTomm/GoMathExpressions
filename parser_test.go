package main

import "testing"

func TestParse(t *testing.T) {
	testCases := []struct {
		tokenPeeker  *tokenPeeker
		expectedTree *node
	}{
		{
			tokenPeeker: newTokenPeeker(
				Lex([]rune("(1*3)+(8/4)")),
				0,
			),
			expectedTree: &node{
				kind:     listNode,
				operator: "+",
				list: []*node{
					{
						kind:     listNode,
						operator: "*",
						list: []*node{
							{
								kind:  literalNode,
								value: "1",
							},
							{
								kind:  literalNode,
								value: "3",
							},
						},
					},
					{
						kind:     listNode,
						operator: "/",
						list: []*node{
							{
								kind:  literalNode,
								value: "8",
							},
							{
								kind:  literalNode,
								value: "4",
							},
						},
					},
				},
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		gotTree := parse(testCase.tokenPeeker)
		gotTreeString := gotTree.printNode()
		expectedTreeString := testCase.expectedTree.printNode()
		if gotTreeString != expectedTreeString {
			t.Fatalf("TCIndex %v: expected to get tree %v, got %v instead", testCaseIdx, expectedTreeString, gotTreeString)
		}
	}
}
