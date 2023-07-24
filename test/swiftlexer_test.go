package test

import (
	"testing"

	"OLC2/chore/parser"

	"github.com/antlr4-go/antlr/v4"
)

type testCases struct {
	input          string
	expectedTokens []int
}

func TestSwiftLexerKeywordDataTypes(t *testing.T) {
	testCases := []testCases{
		{
			input: "Int Double String Bool",
			expectedTokens: []int{
				parser.SwiftLexerKw_INT,
				parser.SwiftLexerKw_DOUBLE,
				parser.SwiftLexerKw_STRING,
				parser.SwiftLexerKw_BOOL,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerLiteralDataTypes(t *testing.T) {
	testCases := []testCases{
		{
			input: `
			42               // Integer literal
			3.14159          // Floating-point literal
			"Hello, world!"  // String literal
			true             // Boolean literal
			`,
			expectedTokens: []int{
				parser.SwiftLexerINT,
				parser.SwiftLexerDOUBLE,
				parser.SwiftLexerSTRING,
				parser.SwiftLexerBOOL,
			},
		},
	}

	TraverseCases(t, testCases)

}

func TestSwiftLexerArithmeticOperators(t *testing.T) {
	testCases := []testCases{
		{
			input: `+-*/%`,
			expectedTokens: []int{
				parser.SwiftLexerOp_PLUS,
				parser.SwiftLexerOp_MINUS,
				parser.SwiftLexerOp_MUL,
				parser.SwiftLexerOp_DIV,
				parser.SwiftLexerOp_MOD,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerAssignmentOperators(t *testing.T) {
	testCases := []testCases{
		{
			input: `-> = += -= *= /=`,
			expectedTokens: []int{
				parser.SwiftLexerOp_ARROW,
				parser.SwiftLexerOp_ASSIGN,
				parser.SwiftLexerOp_PLUS_ASSIGN,
				parser.SwiftLexerOp_MINUS_ASSIGN,
				parser.SwiftLexerOp_MUL_ASSIGN,
				parser.SwiftLexerOp_DIV_ASSIGN,
			},
		},
	}

	TraverseCases(t, testCases)

}

func TestSwiftLexerComparisonOperators(t *testing.T) {
	testCases := []testCases{
		{
			input: `== != < > <= >=`,
			expectedTokens: []int{
				parser.SwiftLexerOp_EQ,
				parser.SwiftLexerOp_NEQ,
				parser.SwiftLexerOp_LT,
				parser.SwiftLexerOp_GT,
				parser.SwiftLexerOp_LE,
				parser.SwiftLexerOp_GE,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerLogicalOperators(t *testing.T) {
	testCases := []testCases{
		{
			input: `&& || !`,
			expectedTokens: []int{
				parser.SwiftLexerOp_AND,
				parser.SwiftLexerOp_OR,
				parser.SwiftLexerOp_NOT,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerPunctuation(t *testing.T) {
	testCases := []testCases{
		{
			input: `; , . :`,
			expectedTokens: []int{
				parser.SwiftLexerSEMICOLON,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerDOT,
				parser.SwiftLexerCOLON,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerDelimeters(t *testing.T) {
	testCases := []testCases{
		{
			input: `( ) { } [ ]`,
			expectedTokens: []int{
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerRPAREN,
				parser.SwiftLexerLBRACE,
				parser.SwiftLexerRBRACE,
				parser.SwiftLexerLBRACKET,
				parser.SwiftLexerRBRACKET,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestControlFlowKeywords(t *testing.T) {

	testCases := []testCases{
		{
			input: `
			if else switch case default for while break continue return do
			`,
			expectedTokens: []int{
				parser.SwiftLexerKw_IF,
				parser.SwiftLexerKw_ELSE,
				parser.SwiftLexerKw_SWITCH,
				parser.SwiftLexerKw_CASE,
				parser.SwiftLexerKw_DEFAULT,
				parser.SwiftLexerKw_FOR,
				parser.SwiftLexerKw_WHILE,
				parser.SwiftLexerKw_BREAK,
				parser.SwiftLexerKw_CONTINUE,
				parser.SwiftLexerKw_RETURN,
				parser.SwiftLexerKw_DO,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerKeywordsDeclarations(t *testing.T) {
	testCases := []testCases{
		{
			input: `
			var let func struct
			`,
			expectedTokens: []int{
				parser.SwiftLexerKw_VAR,
				parser.SwiftLexerKw_LET,
				parser.SwiftLexerKw_FUNC,
				parser.SwiftLexerKw_STRUCT,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TestSwiftLexerCompoundBlocks(t *testing.T) {
	testCases := []testCases{
		{
			input: `
			struct FixedLengthRange {
				var firstValue: Int
				let length: Int
			}
			var rangeOfThreeItems = FixedLengthRange(firstValue: 0, length: 3)
			// the range represents integer values 0, 1, and 2
			rangeOfThreeItems.firstValue = 6
			// the range now represents integer values 6, 7, and 8
			`,
			expectedTokens: []int{
				parser.SwiftLexerKw_STRUCT,
				parser.SwiftLexerID,
				parser.SwiftLexerLBRACE,

				parser.SwiftLexerKw_VAR,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerKw_INT,

				parser.SwiftLexerKw_LET,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerKw_INT,

				parser.SwiftLexerRBRACE,

				parser.SwiftLexerKw_VAR,
				parser.SwiftLexerID,
				parser.SwiftLexerOp_ASSIGN,
				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerINT,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerINT,
				parser.SwiftLexerRPAREN,

				parser.SwiftLexerID,
				parser.SwiftLexerDOT,
				parser.SwiftLexerID,
				parser.SwiftLexerOp_ASSIGN,
				parser.SwiftLexerINT,
			},
		},
		{
			input: `func greet(person: String, alreadyGreeted: Bool) -> String {
				if alreadyGreeted {
					return greetAgain(person: person)
				} else {
					return greet(person: person)
				}
			}
			print(greet(person: "Tim", alreadyGreeted: true))`,
			expectedTokens: []int{
				parser.SwiftLexerKw_FUNC,
				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerKw_STRING,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerKw_BOOL,
				parser.SwiftLexerRPAREN,
				parser.SwiftLexerOp_ARROW,
				parser.SwiftLexerKw_STRING,
				parser.SwiftLexerLBRACE,

				parser.SwiftLexerKw_IF,
				parser.SwiftLexerID,
				parser.SwiftLexerLBRACE,

				parser.SwiftLexerKw_RETURN,
				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerID,
				parser.SwiftLexerRPAREN,

				parser.SwiftLexerRBRACE,
				parser.SwiftLexerKw_ELSE,
				parser.SwiftLexerLBRACE,

				parser.SwiftLexerKw_RETURN,
				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerID,
				parser.SwiftLexerRPAREN,

				parser.SwiftLexerRBRACE,

				parser.SwiftLexerRBRACE,

				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerSTRING,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerBOOL,
				parser.SwiftLexerRPAREN,
				parser.SwiftLexerRPAREN,
			},
		},
		{
			input: `
			func arithmeticMean(numbers: Double) -> Double {
				var total: Double = 0
				for number in numbers {
					total += number
				}
				return total / Double(numbers.count)
			}
			arithmeticMean(1, 2, 3, 4, 5)
			// returns 3.0, which is the arithmetic mean of these five numbers
			arithmeticMean(3, 8.25, 18.75)
			`,
			expectedTokens: []int{
				parser.SwiftLexerKw_FUNC,
				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerKw_DOUBLE,
				parser.SwiftLexerRPAREN,
				parser.SwiftLexerOp_ARROW,
				parser.SwiftLexerKw_DOUBLE,
				parser.SwiftLexerLBRACE,

				parser.SwiftLexerKw_VAR,
				parser.SwiftLexerID,
				parser.SwiftLexerCOLON,
				parser.SwiftLexerKw_DOUBLE,
				parser.SwiftLexerOp_ASSIGN,
				parser.SwiftLexerINT,

				parser.SwiftLexerKw_FOR,
				parser.SwiftLexerID,
				parser.SwiftLexerKw_IN,
				parser.SwiftLexerID,
				parser.SwiftLexerLBRACE,

				parser.SwiftLexerID,
				parser.SwiftLexerOp_PLUS_ASSIGN,
				parser.SwiftLexerID,

				parser.SwiftLexerRBRACE,

				parser.SwiftLexerKw_RETURN,
				parser.SwiftLexerID,
				parser.SwiftLexerOp_DIV,
				parser.SwiftLexerKw_DOUBLE,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerID,
				parser.SwiftLexerDOT,
				parser.SwiftLexerID,
				parser.SwiftLexerRPAREN,

				parser.SwiftLexerRBRACE,

				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerINT,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerINT,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerINT,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerINT,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerINT,
				parser.SwiftLexerRPAREN,

				parser.SwiftLexerID,
				parser.SwiftLexerLPAREN,
				parser.SwiftLexerINT,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerDOUBLE,
				parser.SwiftLexerCOMMA,
				parser.SwiftLexerDOUBLE,
				parser.SwiftLexerRPAREN,
			},
		},
	}

	TraverseCases(t, testCases)
}

func TraverseCases(t *testing.T, testCases []testCases) {

	for _, tc := range testCases {
		input := antlr.NewInputStream(tc.input)
		lexer := parser.NewSwiftLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		tokenNames := lexer.GetSymbolicNames()

		for _, expected := range tc.expectedTokens {
			token := stream.LT(1)
			if token.GetTokenType() == antlr.TokenEOF {
				t.Errorf("Esperado token %s pero se alcanzó el fin de archivo", tokenNames[expected])
				break
			}
			if token.GetTokenType() != expected {
				t.Errorf("Esperado token %s pero se obtuvo %s en la línea %d, columna %d", tokenNames[expected], tokenNames[token.GetTokenType()], token.GetLine(), token.GetColumn())
			}
			stream.Consume()
		}
	}
}
