package test

import (
	"fmt"
	"testing"

	I "OLC2/chore/interfaces"
)

type expectedValue struct {
	values map[string]I.Value
	errors []error
}

type testCasesParser struct {
	input    string
	expected expectedValue
	desc     string
}

func TestParserArithmeticOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `let a = 5 + 5
					let b = 5 - 5
					let c = 5 * 5
					let d = 5 / 5
					let e = 5 % 5
					let f = e + 5
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 10},
					"b": {ParseValue: 0},
					"c": {ParseValue: 25},
					"d": {ParseValue: 1},
					"e": {ParseValue: 0},
					"f": {ParseValue: 5},
				},
				[]error{},
			},
			desc: "Test arithmetic operators",
		},
		{
			// Test precedence
			input: `let a = 5 + 5 * 5
					let b = 5 * 5 + 5
					let c = 5 * 5 / 5
					let d = 5 / 5 * 5
					let e = 5 + 5 / 5
					let f = 5 / 5 + 5
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 30},
					"b": {ParseValue: 30},
					"c": {ParseValue: 5},
					"d": {ParseValue: 5},
					"e": {ParseValue: 6},
					"f": {ParseValue: 6},
				},
				[]error{},
			},
			desc: "Test arithmetic operators precedence",
		},
		{
			// Test precedence with parenthesis
			input: `let a = (5 + 5) * 5
					let b = 5 * (5 + 5)
					let c = (5 * 5) / 5
					let d = 5 / (5 * 5)
					let e = (5 + 5) / 5
					let f = 5 / (5 + 5)
					let g = (a * b) / (c + d)
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 50},
					"b": {ParseValue: 50},
					"c": {ParseValue: 5},
					"d": {ParseValue: 0},
					"e": {ParseValue: 2},
					"f": {ParseValue: 0},
					"g": {ParseValue: 500},
				},
				[]error{},
			},
			desc: "Test arithmetic operators precedence with parenthesis",
		},
		{
			// Test with unary operators
			input: `let a = -5 + 5
					let b = 5 + -5
					let c = -5 - 5
					let d = 5 - (-5)
					let e = -5 * 5
					let f = 5 * -5
					let g = -5 / 5
					let h = 5 / -5
					let i = -5 % 5
					let j = 5 % -5
			`,
			expected: expectedValue{
				map[string]I.Value{
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
				[]error{},
			},
			desc: "Test arithmetic operators with unary operators",
		},
		{
			// With errors
			input: `let a = 5 + false
					let b = 5 - false
					let c = 5 * false
					let d = 5 / false
					let e = 5 % false
					`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: nil},
					"b": {ParseValue: nil},
					"c": {ParseValue: nil},
					"d": {ParseValue: nil},
					"e": {ParseValue: nil},
				},
				[]error{
					fmt.Errorf("Error: No se puede sumar int64 con bool"),
					fmt.Errorf("Error: No se puede restar int64 con bool"),
					fmt.Errorf("Error: No se puede multiplicar int64 con bool"),
					fmt.Errorf("Error: No se puede dividir int64 con bool"),
					fmt.Errorf("Error: No se puede modular int64 con bool"),
				},
			},
			desc: "Test arithmetic operators with errors",
		},
	}

	TraverseParserCases(t, testCases)
}

func TestParserRelationalOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `let a = 5 > 5
					let b = 5 < 5
					let c = 5 >= 5
					let d = 5 <= 5
					let e = 5 == 5
					let f = 5 != 5
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: false},
					"b": {ParseValue: false},
					"c": {ParseValue: true},
					"d": {ParseValue: true},
					"e": {ParseValue: true},
					"f": {ParseValue: false},
				},
				[]error{},
			},
			desc: "Test relational operators",
		},
		{
			// Test precedence
			input: `let a = 5 > 5 + 5
					let b = 5 < 5 + 5
					let c = 5 >= 5 + 5
					let d = 5 <= 5 + 5
					let e = 5 == 5 + 5
					let f = 5 != 5 + 5
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: false},
					"b": {ParseValue: true},
					"c": {ParseValue: false},
					"d": {ParseValue: true},
					"e": {ParseValue: false},
					"f": {ParseValue: true},
				},
				[]error{},
			},
			desc: "Test relational operators precedence",
		},
		{
			// Test precedence with parenthesis
			input: `let a = 5 > (5 + 5)
					let b = 5 < (5 + 5)
					let c = 5 >= (5 + 5)
					let d = 5 <= (5 + 5)
					let e = 5 == (5 + 5)
					let f = 5 != (5 + 5)
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: false},
					"b": {ParseValue: true},
					"c": {ParseValue: false},
					"d": {ParseValue: true},
					"e": {ParseValue: false},
					"f": {ParseValue: true},
				},
				[]error{},
			},
			desc: "Test relational operators precedence with parenthesis",
		},
		{
			// Test with decimal numbers
			input: `let a = 5.5 > 5.5
					let b = 5.5 < 5.5
					let c = 5.5 >= 5.5
					let d = 5.5 <= 5.5
					let e = 5.5 == 5.5
					let f = 5.5 != 5.5
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: false},
					"b": {ParseValue: false},
					"c": {ParseValue: true},
					"d": {ParseValue: true},
					"e": {ParseValue: true},
					"f": {ParseValue: false},
				},
				[]error{},
			},
			desc: "Test relational operators with decimal numbers",
		},
	}

	TraverseParserCases(t, testCases)
}

func TestParserLogicalOperators(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `let a = true && true
					let b = true && false
					let c = true || false
					let d = false || false
					let e = !true
					let f = !false
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: true},
					"b": {ParseValue: false},
					"c": {ParseValue: true},
					"d": {ParseValue: false},
					"e": {ParseValue: false},
					"f": {ParseValue: true},
				},
				[]error{},
			},
			desc: "Test logical operators",
		},
		{
			// Test precedence
			input: `let a = true && true || false
					let b = true && false || false
					let c = true || false && false
					let d = false || false && false
					let e = !true && false
					let f = !false || true
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: true},
					"b": {ParseValue: false},
					"c": {ParseValue: true},
					"d": {ParseValue: false},
					"e": {ParseValue: false},
					"f": {ParseValue: true},
				},
				[]error{},
			},
			desc: "Test logical operators precedence",
		},
		{
			// Test precedence with parenthesis
			input: `let a = true && (true || false)
					let b = true && (false || false)
					let c = (true || false) && false
					let d = (false || false) && false
					let e = !(true && false)
					let f = !(false || true)
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: true},
					"b": {ParseValue: false},
					"c": {ParseValue: false},
					"d": {ParseValue: false},
					"e": {ParseValue: true},
					"f": {ParseValue: false},
				},
				[]error{},
			},
			desc: "Test logical operators precedence with parenthesis",
		},
		{
			// Test for ternary operator
			input: `let a = true ? 1 : 0
					let b = false ? 1 : 0
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 1},
					"b": {ParseValue: 0},
				},
				[]error{},
			},
			desc: "Test ternary operator",
		},
		{
			// Test for ternary operator precedence
			input: `let a = true ? 1 + 1 : 0
					let b = false ? 1 + 1 : 0
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 2},
					"b": {ParseValue: 0},
				},
				[]error{},
			},
			desc: "Test ternary operator precedence",
		},
		{
			// Test for ternary operator precedence with parenthesis
			input: `let a = (true && true) ? 1 + 1 : 0
					let b = (false || false) ? 1 + 1 : 0
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 2},
					"b": {ParseValue: 0},
				},
				[]error{},
			},
			desc: "Test ternary operator precedence with parenthesis",
		},
	}

	TraverseParserCases(t, testCases)
}

func TestParserVariableDeclaration(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `let a = 5
					let b = 10
					let c = 15
			`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 5},
					"b": {ParseValue: 10},
					"c": {ParseValue: 15},
				},
				[]error{},
			},
			desc: "Test variable declaration",
		},
		{
			input: `let a = 5; let b = 10; let c = 15`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 5},
					"b": {ParseValue: 10},
					"c": {ParseValue: 15},
				},
				[]error{},
			},
			desc: "Test variable declaration with semicolon",
		},
	}

	TraverseParserCases(t, testCases)
}

func TestParserRangeExpression(t *testing.T) {
	testCases := []testCasesParser{
		{
			input: `let a = 1...5
					let b = 4...5
					let c = 6...10
					`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: []int{1, 2, 3, 4, 5}},
					"b": {ParseValue: []int{4, 5}},
					"c": {ParseValue: []int{6, 7, 8, 9, 10}},
				},
				[]error{},
			},
			desc: "Test range expression",
		},
		{
			input: `let a = 12
					let b = 14
					let c = a...b
					`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: 12},
					"b": {ParseValue: 14},
					"c": {ParseValue: []int{12, 13, 14}},
				},
				[]error{},
			},
			desc: "Test range expression",
		},
		{
			input: `let a = false
					let b = 14
					let c = a...b
					`,
			expected: expectedValue{
				map[string]I.Value{
					"a": {ParseValue: false},
					"b": {ParseValue: 14},
					"c": {ParseValue: nil},
				},
				[]error{
					fmt.Errorf("Left and right values must be integers"),
				},
			},
			desc: "Test range expression with error value",
		},
	}

	TraverseParserCases(t, testCases)
}

func TraverseParserCases(t *testing.T, testCases []testCasesParser) {
	for _, testCase := range testCases {
		expected := testCase.expected

		got := I.NewEvaluator(testCase.input)

		if fmt.Sprint(got.Memory) != fmt.Sprint(expected.values) || fmt.Sprint(got.Errors) != fmt.Sprint(expected.errors) {
			t.Errorf("\nExpected: \n%v\n%v \n\nGot:\n%v\n\n%v \n\nAt: %s", expected.values, expected.errors, got.Memory, got.Errors, testCase.desc)
		}

	}
}
