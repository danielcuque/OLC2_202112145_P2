package test

import (
	"OLC2/chore/parser"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseSwiftVisitor
}

type testCasesParser struct {
	input    string
	expected string
}

func TestParserArithmeticOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input:    "(1+2)*3+(3+2)",
			expected: "14",
		},
	}

	TraverseParserCases(t, testCases)
}

func TraverseParserCases(t *testing.T, testCases []testCasesParser) {
	for _, testCase := range testCases {
		input := testCase.input
		expected := testCase.expected

		inputStream := antlr.NewInputStream(input)
		lexer := parser.NewSwiftLexer(inputStream)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewSwiftParser(tokens)

		tree := p.Expr()
		visitor := Visitor{}
		result := visitor.Visit(tree)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	}
}
