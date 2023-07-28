package test

import (
	"fmt"
	"testing"

	I "OLC2/chore/interfaces"
)

type testCasesParser struct {
	input    string
	expected map[string]I.Value
	desc     string
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
			desc: "Test arithmetic operators",
		},
		{
			// Test precedence
			input: `a = 5 + 5 * 5
					b = 5 * 5 + 5
					c = 5 * 5 / 5
					d = 5 / 5 * 5
					e = 5 + 5 / 5
					f = 5 / 5 + 5
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: 30},
				"b": {ParseValue: 30},
				"c": {ParseValue: 5},
				"d": {ParseValue: 5},
				"e": {ParseValue: 6},
				"f": {ParseValue: 6},
			},
			desc: "Test arithmetic operators precedence",
		},
		{
			// Test precedence with parenthesis
			input: `a = (5 + 5) * 5
					b = 5 * (5 + 5)
					c = (5 * 5) / 5
					d = 5 / (5 * 5)
					e = (5 + 5) / 5
					f = 5 / (5 + 5)
					g = (a * b) / (c + d)
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: 50},
				"b": {ParseValue: 50},
				"c": {ParseValue: 5},
				"d": {ParseValue: 0},
				"e": {ParseValue: 2},
				"f": {ParseValue: 0},
				"g": {ParseValue: 500},
			},
			desc: "Test arithmetic operators precedence with parenthesis",
		},
		{
			// Test with unary operators
			input: `a = -5 + 5
					b = 5 + -5
					c = -5 - 5
					d = 5 - (-5)
					e = -5 * 5
					f = 5 * -5
					g = -5 / 5
					h = 5 / -5
					i = -5 % 5
					j = 5 % -5
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: 0},
				"b": {ParseValue: 0},
				"c": {ParseValue: -10},
				"d": {ParseValue: 10},
				"e": {ParseValue: -25},
				"f": {ParseValue: -25},
				"g": {ParseValue: -1},
				"h": {ParseValue: -1},
				"i": {ParseValue: 0},
				"j": {ParseValue: 0},
			},
			desc: "Test arithmetic operators with unary operators",
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
			desc: "Test relational operators",
		},
		{
			// Test precedence
			input: `a = 5 > 5 + 5
					b = 5 < 5 + 5
					c = 5 >= 5 + 5
					d = 5 <= 5 + 5
					e = 5 == 5 + 5
					f = 5 != 5 + 5
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: false},
				"b": {ParseValue: true},
				"c": {ParseValue: false},
				"d": {ParseValue: true},
				"e": {ParseValue: false},
				"f": {ParseValue: true},
			},
			desc: "Test relational operators precedence",
		},
		{
			// Test precedence with parenthesis
			input: `a = 5 > (5 + 5)
					b = 5 < (5 + 5)
					c = 5 >= (5 + 5)
					d = 5 <= (5 + 5)
					e = 5 == (5 + 5)
					f = 5 != (5 + 5)
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: false},
				"b": {ParseValue: true},
				"c": {ParseValue: false},
				"d": {ParseValue: true},
				"e": {ParseValue: false},
				"f": {ParseValue: true},
			},
			desc: "Test relational operators precedence with parenthesis",
		},
		{
			// Test with decimal numbers
			input: `a = 5.5 > 5.5
					b = 5.5 < 5.5
					c = 5.5 >= 5.5
					d = 5.5 <= 5.5
					e = 5.5 == 5.5
					f = 5.5 != 5.5
			`,
			expected: map[string]I.Value{
				"a": {ParseValue: false},
				"b": {ParseValue: false},
				"c": {ParseValue: true},
				"d": {ParseValue: true},
				"e": {ParseValue: true},
				"f": {ParseValue: false},
			},
			desc: "Test relational operators with decimal numbers",
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
			desc: "Test logical operators",
		},
	}

	TraverseParserCases(t, testCases)
}

func TraverseParserCases(t *testing.T, testCases []testCasesParser) {
	for _, testCase := range testCases {
		expected := testCase.expected

		got := I.NewEvaluator(testCase.input)

		if fmt.Sprint(got) != fmt.Sprint(expected) {
			t.Errorf("\nExpected: %v\nGot: %v \nAt: %s", expected, got, testCase.desc)
		}

	}
}
