package test

import (
	"fmt"
	"testing"

	I "OLC2/chore/interfaces"
	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type testCasesParser struct {
	input    string
	expected map[string]I.Value
}

func TestParserArithmeticOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `a = 5 + 5
					b = 5 - 5
					c = 5 * 5
					d = 5 / 5
					e = 5 % 5
					f = e + 5
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: 10},
				"b": {ParseValue: 0},
				"c": {ParseValue: 25},
				"d": {ParseValue: 1},
				"e": {ParseValue: 0},
				"f": {ParseValue: 5},
			},
		},
	}

	TraverseParserCases(t, testCases)
}

func TestParserRelationalOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `a = 5 > 5
					b = 5 < 5
					c = 5 >= 5
					d = 5 <= 5
					e = 5 == 5
					f = 5 != 5
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: false},
				"b": {ParseValue: false},
				"c": {ParseValue: true},
				"d": {ParseValue: true},
				"e": {ParseValue: true},
				"f": {ParseValue: false},
			},
		},
	}

	TraverseParserCases(t, testCases)
}

func TestParserLogicalOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `a = true && true
					b = true && false
					c = true || false
					d = false || false
					e = !true
					f = !false
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: true},
				"b": {ParseValue: false},
				"c": {ParseValue: true},
				"d": {ParseValue: false},
				"e": {ParseValue: false},
				"f": {ParseValue: true},
			},
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
		tree := p.Program()
		eval := I.Visitor{Memory: make(map[string]I.Value)}
		eval.Visit(tree)

		got := eval.Memory

		if fmt.Sprint(got) != fmt.Sprint(expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}

	}
}
