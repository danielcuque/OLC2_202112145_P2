package test

import (
	"testing"

	I "OLC2/chore/interfaces"
	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type testCasesParser struct {
	input    string
	expected interface{}
}

func TestParserArithmeticOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input:    "(1+2)*3+(3+2)",
			expected: 14,
		},
	}

	TraverseParserCases(t, testCases)
}

func TraverseParserCases(t *testing.T, testCases []testCasesParser) {
	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected

		lexer := parser.NewSwiftLexer(antlr.NewInputStream(input))
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSwiftParser(stream)

		p.BuildParseTrees = true
		tree := p.Expr()
		eval := I.Visitor{}
		result := eval.Visit(tree)

		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
		}
	}
}
