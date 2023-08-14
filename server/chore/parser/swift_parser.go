// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftParser struct {
	*antlr.BaseParser
}

var SwiftParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftParserInit() {
	staticData := &SwiftParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'let'", "'var'", "'func'", "'struct'", "'if'", "'else'",
		"'switch'", "'case'", "'default'", "'for'", "'while'", "'break'", "'continue'",
		"'return'", "'do'", "'repeat'", "'in'", "'guard'", "'Int'", "'Float'",
		"'Bool'", "'String'", "'Character'", "'...'", "", "", "", "", "", "'nil'",
		"", "'->'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'='", "'*='",
		"'/='", "'+='", "'-='", "'%='", "'*'", "'/'", "'+'", "'-'", "'%'", "'&&'",
		"'||'", "'!'", "'?'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'\\'",
		"','", "';'", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_GUARD", "Kw_INT", "Kw_FLOAT", "Kw_BOOL", "Kw_STRING",
		"Kw_CHAR", "Kw_RANGE", "INT", "FLOAT", "BOOL", "STRING", "CHAR", "NIL",
		"ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE",
		"Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN",
		"Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", "Op_MOD",
		"Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"program", "block", "statement", "variableType", "variableDeclaration",
		"functionDeclarationStatement", "functionParameters", "functionParameter",
		"functionReturnType", "functionCall", "functionCallArguments", "variableAssignment",
		"ifStatement", "ifTail", "elseStatement", "whileStatement", "switchStatement",
		"switchCase", "switchDefault", "forStatement", "guardStatement", "expr",
		"controlFlowStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 293, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 51, 8, 1, 10, 1,
		12, 1, 54, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 66, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 77, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 92, 8, 4, 3, 4, 94, 8, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 100, 8, 5, 1, 5, 1, 5, 3, 5, 104, 8, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 113, 8, 6, 10, 6, 12, 6, 116, 9, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 128,
		8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 135, 8, 10, 10, 10, 12, 10,
		138, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 147,
		8, 10, 10, 10, 12, 10, 150, 9, 10, 3, 10, 152, 8, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 161, 8, 12, 10, 12, 12, 12, 164,
		9, 12, 1, 12, 3, 12, 167, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 190, 8, 16, 10, 16, 12, 16, 193,
		9, 16, 1, 16, 3, 16, 196, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 3, 17, 205, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 211, 8,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 245, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 280, 8, 21, 10, 21, 12, 21,
		283, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 289, 8, 22, 3, 22, 291,
		8, 22, 1, 22, 0, 1, 42, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 8, 1, 0, 22, 26, 1, 0, 4,
		5, 2, 0, 42, 42, 45, 46, 1, 0, 48, 49, 1, 0, 50, 51, 2, 0, 39, 39, 41,
		41, 2, 0, 38, 38, 40, 40, 1, 0, 36, 37, 320, 0, 46, 1, 0, 0, 0, 2, 52,
		1, 0, 0, 0, 4, 65, 1, 0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10,
		95, 1, 0, 0, 0, 12, 109, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 121, 1, 0,
		0, 0, 18, 124, 1, 0, 0, 0, 20, 151, 1, 0, 0, 0, 22, 153, 1, 0, 0, 0, 24,
		157, 1, 0, 0, 0, 26, 168, 1, 0, 0, 0, 28, 174, 1, 0, 0, 0, 30, 179, 1,
		0, 0, 0, 32, 185, 1, 0, 0, 0, 34, 199, 1, 0, 0, 0, 36, 206, 1, 0, 0, 0,
		38, 212, 1, 0, 0, 0, 40, 220, 1, 0, 0, 0, 42, 244, 1, 0, 0, 0, 44, 290,
		1, 0, 0, 0, 46, 47, 3, 2, 1, 0, 47, 48, 5, 0, 0, 1, 48, 1, 1, 0, 0, 0,
		49, 51, 3, 4, 2, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1,
		0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 3, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55,
		66, 3, 22, 11, 0, 56, 66, 3, 8, 4, 0, 57, 66, 3, 24, 12, 0, 58, 66, 3,
		30, 15, 0, 59, 66, 3, 32, 16, 0, 60, 66, 3, 38, 19, 0, 61, 66, 3, 40, 20,
		0, 62, 66, 3, 44, 22, 0, 63, 66, 3, 10, 5, 0, 64, 66, 3, 18, 9, 0, 65,
		55, 1, 0, 0, 0, 65, 56, 1, 0, 0, 0, 65, 57, 1, 0, 0, 0, 65, 58, 1, 0, 0,
		0, 65, 59, 1, 0, 0, 0, 65, 60, 1, 0, 0, 0, 65, 61, 1, 0, 0, 0, 65, 62,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 5, 1, 0, 0, 0,
		67, 68, 7, 0, 0, 0, 68, 7, 1, 0, 0, 0, 69, 70, 7, 1, 0, 0, 70, 71, 5, 34,
		0, 0, 71, 72, 5, 66, 0, 0, 72, 73, 3, 6, 3, 0, 73, 74, 5, 42, 0, 0, 74,
		76, 3, 42, 21, 0, 75, 77, 5, 65, 0, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0,
		0, 0, 77, 94, 1, 0, 0, 0, 78, 79, 7, 1, 0, 0, 79, 80, 5, 34, 0, 0, 80,
		81, 5, 42, 0, 0, 81, 83, 3, 42, 21, 0, 82, 84, 5, 65, 0, 0, 83, 82, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 94, 1, 0, 0, 0, 85, 86, 7, 1, 0, 0, 86,
		87, 5, 34, 0, 0, 87, 88, 5, 66, 0, 0, 88, 89, 3, 6, 3, 0, 89, 91, 5, 56,
		0, 0, 90, 92, 5, 65, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		94, 1, 0, 0, 0, 93, 69, 1, 0, 0, 0, 93, 78, 1, 0, 0, 0, 93, 85, 1, 0, 0,
		0, 94, 9, 1, 0, 0, 0, 95, 96, 5, 6, 0, 0, 96, 97, 5, 34, 0, 0, 97, 99,
		5, 57, 0, 0, 98, 100, 3, 12, 6, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 101, 1, 0, 0, 0, 101, 103, 5, 58, 0, 0, 102, 104, 3, 16, 8, 0,
		103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105,
		106, 5, 59, 0, 0, 106, 107, 3, 2, 1, 0, 107, 108, 5, 60, 0, 0, 108, 11,
		1, 0, 0, 0, 109, 114, 3, 14, 7, 0, 110, 111, 5, 64, 0, 0, 111, 113, 3,
		14, 7, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 13, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117,
		118, 5, 34, 0, 0, 118, 119, 5, 66, 0, 0, 119, 120, 3, 6, 3, 0, 120, 15,
		1, 0, 0, 0, 121, 122, 5, 35, 0, 0, 122, 123, 3, 6, 3, 0, 123, 17, 1, 0,
		0, 0, 124, 125, 5, 34, 0, 0, 125, 127, 5, 57, 0, 0, 126, 128, 3, 20, 10,
		0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		130, 5, 58, 0, 0, 130, 19, 1, 0, 0, 0, 131, 136, 3, 42, 21, 0, 132, 133,
		5, 64, 0, 0, 133, 135, 3, 42, 21, 0, 134, 132, 1, 0, 0, 0, 135, 138, 1,
		0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 152, 1, 0, 0,
		0, 138, 136, 1, 0, 0, 0, 139, 140, 5, 34, 0, 0, 140, 141, 5, 66, 0, 0,
		141, 148, 3, 42, 21, 0, 142, 143, 5, 64, 0, 0, 143, 144, 5, 34, 0, 0, 144,
		145, 5, 66, 0, 0, 145, 147, 3, 42, 21, 0, 146, 142, 1, 0, 0, 0, 147, 150,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 152, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 151, 131, 1, 0, 0, 0, 151, 139, 1, 0, 0, 0,
		152, 21, 1, 0, 0, 0, 153, 154, 5, 34, 0, 0, 154, 155, 7, 2, 0, 0, 155,
		156, 3, 42, 21, 0, 156, 23, 1, 0, 0, 0, 157, 162, 3, 26, 13, 0, 158, 159,
		5, 9, 0, 0, 159, 161, 3, 26, 13, 0, 160, 158, 1, 0, 0, 0, 161, 164, 1,
		0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 166, 1, 0, 0,
		0, 164, 162, 1, 0, 0, 0, 165, 167, 3, 28, 14, 0, 166, 165, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 25, 1, 0, 0, 0, 168, 169, 5, 8, 0, 0, 169, 170,
		3, 42, 21, 0, 170, 171, 5, 59, 0, 0, 171, 172, 3, 2, 1, 0, 172, 173, 5,
		60, 0, 0, 173, 27, 1, 0, 0, 0, 174, 175, 5, 9, 0, 0, 175, 176, 5, 59, 0,
		0, 176, 177, 3, 2, 1, 0, 177, 178, 5, 60, 0, 0, 178, 29, 1, 0, 0, 0, 179,
		180, 5, 14, 0, 0, 180, 181, 3, 42, 21, 0, 181, 182, 5, 59, 0, 0, 182, 183,
		3, 2, 1, 0, 183, 184, 5, 60, 0, 0, 184, 31, 1, 0, 0, 0, 185, 186, 5, 10,
		0, 0, 186, 187, 3, 42, 21, 0, 187, 191, 5, 59, 0, 0, 188, 190, 3, 34, 17,
		0, 189, 188, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191,
		192, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 196,
		3, 36, 18, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1,
		0, 0, 0, 197, 198, 5, 60, 0, 0, 198, 33, 1, 0, 0, 0, 199, 200, 5, 11, 0,
		0, 200, 201, 3, 42, 21, 0, 201, 202, 5, 66, 0, 0, 202, 204, 3, 2, 1, 0,
		203, 205, 5, 15, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		35, 1, 0, 0, 0, 206, 207, 5, 12, 0, 0, 207, 208, 5, 66, 0, 0, 208, 210,
		3, 2, 1, 0, 209, 211, 5, 15, 0, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0,
		0, 0, 211, 37, 1, 0, 0, 0, 212, 213, 5, 13, 0, 0, 213, 214, 5, 34, 0, 0,
		214, 215, 5, 20, 0, 0, 215, 216, 3, 42, 21, 0, 216, 217, 5, 59, 0, 0, 217,
		218, 3, 2, 1, 0, 218, 219, 5, 60, 0, 0, 219, 39, 1, 0, 0, 0, 220, 221,
		5, 21, 0, 0, 221, 222, 3, 42, 21, 0, 222, 223, 5, 9, 0, 0, 223, 224, 5,
		59, 0, 0, 224, 225, 3, 2, 1, 0, 225, 226, 5, 60, 0, 0, 226, 41, 1, 0, 0,
		0, 227, 228, 6, 21, -1, 0, 228, 245, 3, 18, 9, 0, 229, 230, 5, 51, 0, 0,
		230, 245, 3, 42, 21, 20, 231, 232, 5, 55, 0, 0, 232, 245, 3, 42, 21, 19,
		233, 234, 5, 57, 0, 0, 234, 235, 3, 42, 21, 0, 235, 236, 5, 58, 0, 0, 236,
		245, 1, 0, 0, 0, 237, 245, 5, 28, 0, 0, 238, 245, 5, 34, 0, 0, 239, 245,
		5, 29, 0, 0, 240, 245, 5, 31, 0, 0, 241, 245, 5, 33, 0, 0, 242, 245, 5,
		32, 0, 0, 243, 245, 5, 30, 0, 0, 244, 227, 1, 0, 0, 0, 244, 229, 1, 0,
		0, 0, 244, 231, 1, 0, 0, 0, 244, 233, 1, 0, 0, 0, 244, 237, 1, 0, 0, 0,
		244, 238, 1, 0, 0, 0, 244, 239, 1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 244,
		241, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 281,
		1, 0, 0, 0, 246, 247, 10, 18, 0, 0, 247, 248, 7, 3, 0, 0, 248, 280, 3,
		42, 21, 19, 249, 250, 10, 17, 0, 0, 250, 251, 7, 4, 0, 0, 251, 280, 3,
		42, 21, 18, 252, 253, 10, 16, 0, 0, 253, 254, 5, 52, 0, 0, 254, 280, 3,
		42, 21, 17, 255, 256, 10, 15, 0, 0, 256, 257, 7, 5, 0, 0, 257, 280, 3,
		42, 21, 16, 258, 259, 10, 14, 0, 0, 259, 260, 7, 6, 0, 0, 260, 280, 3,
		42, 21, 15, 261, 262, 10, 13, 0, 0, 262, 263, 7, 7, 0, 0, 263, 280, 3,
		42, 21, 14, 264, 265, 10, 12, 0, 0, 265, 266, 5, 56, 0, 0, 266, 267, 3,
		42, 21, 0, 267, 268, 5, 66, 0, 0, 268, 269, 3, 42, 21, 13, 269, 280, 1,
		0, 0, 0, 270, 271, 10, 11, 0, 0, 271, 272, 5, 53, 0, 0, 272, 280, 3, 42,
		21, 12, 273, 274, 10, 10, 0, 0, 274, 275, 5, 54, 0, 0, 275, 280, 3, 42,
		21, 11, 276, 277, 10, 9, 0, 0, 277, 278, 5, 27, 0, 0, 278, 280, 3, 42,
		21, 10, 279, 246, 1, 0, 0, 0, 279, 249, 1, 0, 0, 0, 279, 252, 1, 0, 0,
		0, 279, 255, 1, 0, 0, 0, 279, 258, 1, 0, 0, 0, 279, 261, 1, 0, 0, 0, 279,
		264, 1, 0, 0, 0, 279, 270, 1, 0, 0, 0, 279, 273, 1, 0, 0, 0, 279, 276,
		1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0,
		0, 0, 282, 43, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284, 291, 5, 15, 0, 0,
		285, 291, 5, 16, 0, 0, 286, 288, 5, 17, 0, 0, 287, 289, 3, 42, 21, 0, 288,
		287, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0, 0, 290, 284,
		1, 0, 0, 0, 290, 285, 1, 0, 0, 0, 290, 286, 1, 0, 0, 0, 291, 45, 1, 0,
		0, 0, 24, 52, 65, 76, 83, 91, 93, 99, 103, 114, 127, 136, 148, 151, 162,
		166, 191, 195, 204, 210, 244, 279, 281, 288, 290,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftParserInit initializes any static state used to implement SwiftParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftParserInit() {
	staticData := &SwiftParserStaticData
	staticData.once.Do(swiftParserInit)
}

// NewSwiftParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftParser(input antlr.TokenStream) *SwiftParser {
	SwiftParserInit()
	this := new(SwiftParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Swift.g4"

	return this
}

// SwiftParser tokens.
const (
	SwiftParserEOF             = antlr.TokenEOF
	SwiftParserWHITESPACE      = 1
	SwiftParserCOMMENT         = 2
	SwiftParserBLOCK_COMMENT   = 3
	SwiftParserKw_LET          = 4
	SwiftParserKw_VAR          = 5
	SwiftParserKw_FUNC         = 6
	SwiftParserKw_STRUCT       = 7
	SwiftParserKw_IF           = 8
	SwiftParserKw_ELSE         = 9
	SwiftParserKw_SWITCH       = 10
	SwiftParserKw_CASE         = 11
	SwiftParserKw_DEFAULT      = 12
	SwiftParserKw_FOR          = 13
	SwiftParserKw_WHILE        = 14
	SwiftParserKw_BREAK        = 15
	SwiftParserKw_CONTINUE     = 16
	SwiftParserKw_RETURN       = 17
	SwiftParserKw_DO           = 18
	SwiftParserKw_REPEAT       = 19
	SwiftParserKw_IN           = 20
	SwiftParserKw_GUARD        = 21
	SwiftParserKw_INT          = 22
	SwiftParserKw_FLOAT        = 23
	SwiftParserKw_BOOL         = 24
	SwiftParserKw_STRING       = 25
	SwiftParserKw_CHAR         = 26
	SwiftParserKw_RANGE        = 27
	SwiftParserINT             = 28
	SwiftParserFLOAT           = 29
	SwiftParserBOOL            = 30
	SwiftParserSTRING          = 31
	SwiftParserCHAR            = 32
	SwiftParserNIL             = 33
	SwiftParserID              = 34
	SwiftParserOp_ARROW        = 35
	SwiftParserOp_EQ           = 36
	SwiftParserOp_NEQ          = 37
	SwiftParserOp_LT           = 38
	SwiftParserOp_GT           = 39
	SwiftParserOp_LE           = 40
	SwiftParserOp_GE           = 41
	SwiftParserOp_ASSIGN       = 42
	SwiftParserOp_MUL_ASSIGN   = 43
	SwiftParserOp_DIV_ASSIGN   = 44
	SwiftParserOp_PLUS_ASSIGN  = 45
	SwiftParserOp_MINUS_ASSIGN = 46
	SwiftParserOp_MOD_ASSIGN   = 47
	SwiftParserOp_MUL          = 48
	SwiftParserOp_DIV          = 49
	SwiftParserOp_PLUS         = 50
	SwiftParserOp_MINUS        = 51
	SwiftParserOp_MOD          = 52
	SwiftParserOp_AND          = 53
	SwiftParserOp_OR           = 54
	SwiftParserOp_NOT          = 55
	SwiftParserOp_TERNARY      = 56
	SwiftParserLPAREN          = 57
	SwiftParserRPAREN          = 58
	SwiftParserLBRACE          = 59
	SwiftParserRBRACE          = 60
	SwiftParserLBRACKET        = 61
	SwiftParserRBRACKET        = 62
	SwiftParserBACKSLASH       = 63
	SwiftParserCOMMA           = 64
	SwiftParserSEMICOLON       = 65
	SwiftParserCOLON           = 66
	SwiftParserDOT             = 67
)

// SwiftParser rules.
const (
	SwiftParserRULE_program                      = 0
	SwiftParserRULE_block                        = 1
	SwiftParserRULE_statement                    = 2
	SwiftParserRULE_variableType                 = 3
	SwiftParserRULE_variableDeclaration          = 4
	SwiftParserRULE_functionDeclarationStatement = 5
	SwiftParserRULE_functionParameters           = 6
	SwiftParserRULE_functionParameter            = 7
	SwiftParserRULE_functionReturnType           = 8
	SwiftParserRULE_functionCall                 = 9
	SwiftParserRULE_functionCallArguments        = 10
	SwiftParserRULE_variableAssignment           = 11
	SwiftParserRULE_ifStatement                  = 12
	SwiftParserRULE_ifTail                       = 13
	SwiftParserRULE_elseStatement                = 14
	SwiftParserRULE_whileStatement               = 15
	SwiftParserRULE_switchStatement              = 16
	SwiftParserRULE_switchCase                   = 17
	SwiftParserRULE_switchDefault                = 18
	SwiftParserRULE_forStatement                 = 19
	SwiftParserRULE_guardStatement               = 20
	SwiftParserRULE_expr                         = 21
	SwiftParserRULE_controlFlowStatement         = 22
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Block()
	}
	{
		p.SetState(47)
		p.Match(SwiftParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftParserRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(49)
				p.Statement()
			}

		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableAssignment() IVariableAssignmentContext
	VariableDeclaration() IVariableDeclarationContext
	IfStatement() IIfStatementContext
	WhileStatement() IWhileStatementContext
	SwitchStatement() ISwitchStatementContext
	ForStatement() IForStatementContext
	GuardStatement() IGuardStatementContext
	ControlFlowStatement() IControlFlowStatementContext
	FunctionDeclarationStatement() IFunctionDeclarationStatementContext
	FunctionCall() IFunctionCallContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableAssignment() IVariableAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableAssignmentContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) GuardStatement() IGuardStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardStatementContext)
}

func (s *StatementContext) ControlFlowStatement() IControlFlowStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlFlowStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlFlowStatementContext)
}

func (s *StatementContext) FunctionDeclarationStatement() IFunctionDeclarationStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationStatementContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftParserRULE_statement)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.VariableAssignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.VariableDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.WhileStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.SwitchStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.ForStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.GuardStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.ControlFlowStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(63)
			p.FunctionDeclarationStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(64)
			p.FunctionCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableTypeContext is an interface to support dynamic dispatch.
type IVariableTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_INT() antlr.TerminalNode
	Kw_FLOAT() antlr.TerminalNode
	Kw_BOOL() antlr.TerminalNode
	Kw_STRING() antlr.TerminalNode
	Kw_CHAR() antlr.TerminalNode

	// IsVariableTypeContext differentiates from other interfaces.
	IsVariableTypeContext()
}

type VariableTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableTypeContext() *VariableTypeContext {
	var p = new(VariableTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_variableType
	return p
}

func InitEmptyVariableTypeContext(p *VariableTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_variableType
}

func (*VariableTypeContext) IsVariableTypeContext() {}

func NewVariableTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableTypeContext {
	var p = new(VariableTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_variableType

	return p
}

func (s *VariableTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableTypeContext) Kw_INT() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_INT, 0)
}

func (s *VariableTypeContext) Kw_FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_FLOAT, 0)
}

func (s *VariableTypeContext) Kw_BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_BOOL, 0)
}

func (s *VariableTypeContext) Kw_STRING() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_STRING, 0)
}

func (s *VariableTypeContext) Kw_CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_CHAR, 0)
}

func (s *VariableTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVariableType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VariableType() (localctx IVariableTypeContext) {
	localctx = NewVariableTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftParserRULE_variableType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&130023424) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) CopyAll(ctx *VariableDeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueDeclarationContext struct {
	VariableDeclarationContext
	varType antlr.Token
}

func NewValueDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueDeclarationContext {
	var p = new(ValueDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *ValueDeclarationContext) GetVarType() antlr.Token { return s.varType }

func (s *ValueDeclarationContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *ValueDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *ValueDeclarationContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *ValueDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_LET, 0)
}

func (s *ValueDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_VAR, 0)
}

func (s *ValueDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserSEMICOLON, 0)
}

func (s *ValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitValueDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeValueDeclarationContext struct {
	VariableDeclarationContext
	varType antlr.Token
}

func NewTypeValueDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeValueDeclarationContext {
	var p = new(TypeValueDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *TypeValueDeclarationContext) GetVarType() antlr.Token { return s.varType }

func (s *TypeValueDeclarationContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *TypeValueDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeValueDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *TypeValueDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *TypeValueDeclarationContext) VariableType() IVariableTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *TypeValueDeclarationContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *TypeValueDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeValueDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_LET, 0)
}

func (s *TypeValueDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_VAR, 0)
}

func (s *TypeValueDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserSEMICOLON, 0)
}

func (s *TypeValueDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitTypeValueDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclarationContext struct {
	VariableDeclarationContext
	varType antlr.Token
}

func NewTypeDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	InitEmptyVariableDeclarationContext(&p.VariableDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclarationContext))

	return p
}

func (s *TypeDeclarationContext) GetVarType() antlr.Token { return s.varType }

func (s *TypeDeclarationContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *TypeDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *TypeDeclarationContext) VariableType() IVariableTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *TypeDeclarationContext) Op_TERNARY() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_TERNARY, 0)
}

func (s *TypeDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_LET, 0)
}

func (s *TypeDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_VAR, 0)
}

func (s *TypeDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserSEMICOLON, 0)
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftParserRULE_variableDeclaration)
	var _la int

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TypeValueDeclarationContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftParserKw_LET || _la == SwiftParserKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TypeValueDeclarationContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(70)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.VariableType()
		}
		{
			p.SetState(73)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.expr(0)
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(75)
				p.Match(SwiftParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueDeclarationContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftParserKw_LET || _la == SwiftParserKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueDeclarationContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(79)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(82)
				p.Match(SwiftParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewTypeDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TypeDeclarationContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftParserKw_LET || _la == SwiftParserKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TypeDeclarationContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(86)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.VariableType()
		}
		{
			p.SetState(89)
			p.Match(SwiftParserOp_TERNARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(90)
				p.Match(SwiftParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclarationStatementContext is an interface to support dynamic dispatch.
type IFunctionDeclarationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode
	FunctionParameters() IFunctionParametersContext
	FunctionReturnType() IFunctionReturnTypeContext

	// IsFunctionDeclarationStatementContext differentiates from other interfaces.
	IsFunctionDeclarationStatementContext()
}

type FunctionDeclarationStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationStatementContext() *FunctionDeclarationStatementContext {
	var p = new(FunctionDeclarationStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionDeclarationStatement
	return p
}

func InitEmptyFunctionDeclarationStatementContext(p *FunctionDeclarationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionDeclarationStatement
}

func (*FunctionDeclarationStatementContext) IsFunctionDeclarationStatementContext() {}

func NewFunctionDeclarationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationStatementContext {
	var p = new(FunctionDeclarationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_functionDeclarationStatement

	return p
}

func (s *FunctionDeclarationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationStatementContext) Kw_FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_FUNC, 0)
}

func (s *FunctionDeclarationStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *FunctionDeclarationStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *FunctionDeclarationStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *FunctionDeclarationStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *FunctionDeclarationStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *FunctionDeclarationStatementContext) FunctionParameters() IFunctionParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionDeclarationStatementContext) FunctionReturnType() IFunctionReturnTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionReturnTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnTypeContext)
}

func (s *FunctionDeclarationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionDeclarationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) FunctionDeclarationStatement() (localctx IFunctionDeclarationStatementContext) {
	localctx = NewFunctionDeclarationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftParserRULE_functionDeclarationStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(SwiftParserKw_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(SwiftParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserID {
		{
			p.SetState(98)
			p.FunctionParameters()
		}

	}
	{
		p.SetState(101)
		p.Match(SwiftParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserOp_ARROW {
		{
			p.SetState(102)
			p.FunctionReturnType()
		}

	}
	{
		p.SetState(105)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Block()
	}
	{
		p.SetState(107)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionParametersContext is an interface to support dynamic dispatch.
type IFunctionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunctionParameter() []IFunctionParameterContext
	FunctionParameter(i int) IFunctionParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionParametersContext differentiates from other interfaces.
	IsFunctionParametersContext()
}

type FunctionParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParametersContext() *FunctionParametersContext {
	var p = new(FunctionParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionParameters
	return p
}

func InitEmptyFunctionParametersContext(p *FunctionParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionParameters
}

func (*FunctionParametersContext) IsFunctionParametersContext() {}

func NewFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_functionParameters

	return p
}

func (s *FunctionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParametersContext) AllFunctionParameter() []IFunctionParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionParameterContext); ok {
			len++
		}
	}

	tst := make([]IFunctionParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionParameterContext); ok {
			tst[i] = t.(IFunctionParameterContext)
			i++
		}
	}

	return tst
}

func (s *FunctionParametersContext) FunctionParameter(i int) IFunctionParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *FunctionParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOMMA)
}

func (s *FunctionParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, i)
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftParserRULE_functionParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.FunctionParameter()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserCOMMA {
		{
			p.SetState(110)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.FunctionParameter()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	VariableType() IVariableTypeContext

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionParameter
	return p
}

func InitEmptyFunctionParameterContext(p *FunctionParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionParameter
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *FunctionParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *FunctionParameterContext) VariableType() IVariableTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftParserRULE_functionParameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.VariableType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionReturnTypeContext is an interface to support dynamic dispatch.
type IFunctionReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Op_ARROW() antlr.TerminalNode
	VariableType() IVariableTypeContext

	// IsFunctionReturnTypeContext differentiates from other interfaces.
	IsFunctionReturnTypeContext()
}

type FunctionReturnTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionReturnTypeContext() *FunctionReturnTypeContext {
	var p = new(FunctionReturnTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionReturnType
	return p
}

func InitEmptyFunctionReturnTypeContext(p *FunctionReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionReturnType
}

func (*FunctionReturnTypeContext) IsFunctionReturnTypeContext() {}

func NewFunctionReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnTypeContext {
	var p = new(FunctionReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_functionReturnType

	return p
}

func (s *FunctionReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnTypeContext) Op_ARROW() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ARROW, 0)
}

func (s *FunctionReturnTypeContext) VariableType() IVariableTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FunctionReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) FunctionReturnType() (localctx IFunctionReturnTypeContext) {
	localctx = NewFunctionReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftParserRULE_functionReturnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(SwiftParserOp_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.VariableType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	FunctionCallArguments() IFunctionCallArgumentsContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *FunctionCallContext) FunctionCallArguments() IFunctionCallArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgumentsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(SwiftParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&182395818999808000) != 0 {
		{
			p.SetState(126)
			p.FunctionCallArguments()
		}

	}
	{
		p.SetState(129)
		p.Match(SwiftParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallArgumentsContext is an interface to support dynamic dispatch.
type IFunctionCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionCallArgumentsContext differentiates from other interfaces.
	IsFunctionCallArgumentsContext()
}

type FunctionCallArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallArgumentsContext() *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionCallArguments
	return p
}

func InitEmptyFunctionCallArgumentsContext(p *FunctionCallArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_functionCallArguments
}

func (*FunctionCallArgumentsContext) IsFunctionCallArgumentsContext() {}

func NewFunctionCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_functionCallArguments

	return p
}

func (s *FunctionCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallArgumentsContext) CopyAll(ctx *FunctionCallArgumentsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgumentsContext struct {
	FunctionCallArgumentsContext
}

func NewArgumentsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsContext {
	var p = new(ArgumentsContext)

	InitEmptyFunctionCallArgumentsContext(&p.FunctionCallArgumentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionCallArgumentsContext))

	return p
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, i)
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

type NamedArgumentsContext struct {
	FunctionCallArgumentsContext
}

func NewNamedArgumentsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedArgumentsContext {
	var p = new(NamedArgumentsContext)

	InitEmptyFunctionCallArgumentsContext(&p.FunctionCallArgumentsContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionCallArgumentsContext))

	return p
}

func (s *NamedArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgumentsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserID)
}

func (s *NamedArgumentsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserID, i)
}

func (s *NamedArgumentsContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOLON)
}

func (s *NamedArgumentsContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, i)
}

func (s *NamedArgumentsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NamedArgumentsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NamedArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOMMA)
}

func (s *NamedArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, i)
}

func (s *NamedArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitNamedArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) FunctionCallArguments() (localctx IFunctionCallArgumentsContext) {
	localctx = NewFunctionCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftParserRULE_functionCallArguments)
	var _la int

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArgumentsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.expr(0)
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftParserCOMMA {
			{
				p.SetState(132)
				p.Match(SwiftParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(133)
				p.expr(0)
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewNamedArgumentsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.expr(0)
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftParserCOMMA {
			{
				p.SetState(142)
				p.Match(SwiftParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Match(SwiftParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)
				p.Match(SwiftParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(145)
				p.expr(0)
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableAssignmentContext is an interface to support dynamic dispatch.
type IVariableAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext
	Op_ASSIGN() antlr.TerminalNode
	Op_PLUS_ASSIGN() antlr.TerminalNode
	Op_MINUS_ASSIGN() antlr.TerminalNode

	// IsVariableAssignmentContext differentiates from other interfaces.
	IsVariableAssignmentContext()
}

type VariableAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyVariableAssignmentContext() *VariableAssignmentContext {
	var p = new(VariableAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_variableAssignment
	return p
}

func InitEmptyVariableAssignmentContext(p *VariableAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_variableAssignment
}

func (*VariableAssignmentContext) IsVariableAssignmentContext() {}

func NewVariableAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableAssignmentContext {
	var p = new(VariableAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_variableAssignment

	return p
}

func (s *VariableAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableAssignmentContext) GetOp() antlr.Token { return s.op }

func (s *VariableAssignmentContext) SetOp(v antlr.Token) { s.op = v }

func (s *VariableAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *VariableAssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VariableAssignmentContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *VariableAssignmentContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_PLUS_ASSIGN, 0)
}

func (s *VariableAssignmentContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MINUS_ASSIGN, 0)
}

func (s *VariableAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVariableAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VariableAssignment() (localctx IVariableAssignmentContext) {
	localctx = NewVariableAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftParserRULE_variableAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&109951162777600) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(155)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIfTail() []IIfTailContext
	IfTail(i int) IIfTailContext
	AllKw_ELSE() []antlr.TerminalNode
	Kw_ELSE(i int) antlr.TerminalNode
	ElseStatement() IElseStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) AllIfTail() []IIfTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfTailContext); ok {
			len++
		}
	}

	tst := make([]IIfTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfTailContext); ok {
			tst[i] = t.(IIfTailContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) IfTail(i int) IIfTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfTailContext)
}

func (s *IfStatementContext) AllKw_ELSE() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserKw_ELSE)
}

func (s *IfStatementContext) Kw_ELSE(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_ELSE, i)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.IfTail()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(158)
				p.Match(SwiftParserKw_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(159)
				p.IfTail()
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_ELSE {
		{
			p.SetState(165)
			p.ElseStatement()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfTailContext is an interface to support dynamic dispatch.
type IIfTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_IF() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsIfTailContext differentiates from other interfaces.
	IsIfTailContext()
}

type IfTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfTailContext() *IfTailContext {
	var p = new(IfTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_ifTail
	return p
}

func InitEmptyIfTailContext(p *IfTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_ifTail
}

func (*IfTailContext) IsIfTailContext() {}

func NewIfTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfTailContext {
	var p = new(IfTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_ifTail

	return p
}

func (s *IfTailContext) GetParser() antlr.Parser { return s.parser }

func (s *IfTailContext) Kw_IF() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_IF, 0)
}

func (s *IfTailContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfTailContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *IfTailContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfTailContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *IfTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIfTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) IfTail() (localctx IIfTailContext) {
	localctx = NewIfTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftParserRULE_ifTail)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(SwiftParserKw_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.expr(0)
	}
	{
		p.SetState(170)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Block()
	}
	{
		p.SetState(172)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_ELSE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_elseStatement
	return p
}

func InitEmptyElseStatementContext(p *ElseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_elseStatement
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) Kw_ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_ELSE, 0)
}

func (s *ElseStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *ElseStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftParserRULE_elseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Block()
	}
	{
		p.SetState(177)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_WHILE() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) Kw_WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_WHILE, 0)
}

func (s *WhileStatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(SwiftParserKw_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.expr(0)
	}
	{
		p.SetState(181)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Block()
	}
	{
		p.SetState(183)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	SwitchDefault() ISwitchDefaultContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Kw_SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_SWITCH, 0)
}

func (s *SwitchStatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *SwitchStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *SwitchStatementContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *SwitchStatementContext) SwitchDefault() ISwitchDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchDefaultContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(SwiftParserKw_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.expr(0)
	}
	{
		p.SetState(187)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserKw_CASE {
		{
			p.SetState(188)
			p.SwitchCase()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_DEFAULT {
		{
			p.SetState(194)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(197)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_CASE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	Block() IBlockContext
	Kw_BREAK() antlr.TerminalNode

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) Kw_CASE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_CASE, 0)
}

func (s *SwitchCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *SwitchCaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchCaseContext) Kw_BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_BREAK, 0)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(SwiftParserKw_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.expr(0)
	}
	{
		p.SetState(201)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Block()
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(203)
			p.Match(SwiftParserKw_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchDefaultContext is an interface to support dynamic dispatch.
type ISwitchDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Block() IBlockContext
	Kw_BREAK() antlr.TerminalNode

	// IsSwitchDefaultContext differentiates from other interfaces.
	IsSwitchDefaultContext()
}

type SwitchDefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchDefaultContext() *SwitchDefaultContext {
	var p = new(SwitchDefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_switchDefault
	return p
}

func InitEmptySwitchDefaultContext(p *SwitchDefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_switchDefault
}

func (*SwitchDefaultContext) IsSwitchDefaultContext() {}

func NewSwitchDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDefaultContext {
	var p = new(SwitchDefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_switchDefault

	return p
}

func (s *SwitchDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchDefaultContext) Kw_DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_DEFAULT, 0)
}

func (s *SwitchDefaultContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *SwitchDefaultContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchDefaultContext) Kw_BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_BREAK, 0)
}

func (s *SwitchDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitSwitchDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) SwitchDefault() (localctx ISwitchDefaultContext) {
	localctx = NewSwitchDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftParserRULE_switchDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(SwiftParserKw_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Block()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(209)
			p.Match(SwiftParserKw_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	Kw_IN() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) Kw_FOR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_FOR, 0)
}

func (s *ForStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *ForStatementContext) Kw_IN() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_IN, 0)
}

func (s *ForStatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(SwiftParserKw_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(SwiftParserKw_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.expr(0)
	}
	{
		p.SetState(216)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Block()
	}
	{
		p.SetState(218)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardStatementContext is an interface to support dynamic dispatch.
type IGuardStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_GUARD() antlr.TerminalNode
	Expr() IExprContext
	Kw_ELSE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	Block() IBlockContext
	RBRACE() antlr.TerminalNode

	// IsGuardStatementContext differentiates from other interfaces.
	IsGuardStatementContext()
}

type GuardStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardStatementContext() *GuardStatementContext {
	var p = new(GuardStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_guardStatement
	return p
}

func InitEmptyGuardStatementContext(p *GuardStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_guardStatement
}

func (*GuardStatementContext) IsGuardStatementContext() {}

func NewGuardStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardStatementContext {
	var p = new(GuardStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_guardStatement

	return p
}

func (s *GuardStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardStatementContext) Kw_GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_GUARD, 0)
}

func (s *GuardStatementContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardStatementContext) Kw_ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_ELSE, 0)
}

func (s *GuardStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *GuardStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *GuardStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitGuardStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) GuardStatement() (localctx IGuardStatementContext) {
	localctx = NewGuardStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftParserRULE_guardStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(SwiftParserKw_GUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.expr(0)
	}
	{
		p.SetState(222)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Block()
	}
	{
		p.SetState(225)
		p.Match(SwiftParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftParserBOOL, 0)
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatExprContext struct {
	ExprContext
}

func NewFloatExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatExprContext {
	var p = new(FloatExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftParserFLOAT, 0)
}

func (s *FloatExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFloatExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExprContext struct {
	ExprContext
}

func NewNilExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExprContext {
	var p = new(NilExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftParserNIL, 0)
}

func (s *NilExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitNilExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RangeExprContext struct {
	ExprContext
	left  IExprContext
	right IExprContext
}

func NewRangeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangeExprContext {
	var p = new(RangeExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RangeExprContext) GetLeft() IExprContext { return s.left }

func (s *RangeExprContext) GetRight() IExprContext { return s.right }

func (s *RangeExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *RangeExprContext) SetRight(v IExprContext) { s.right = v }

func (s *RangeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeExprContext) Kw_RANGE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_RANGE, 0)
}

func (s *RangeExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangeExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitRangeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	ExprContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Op_MINUS() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MINUS, 0)
}

func (s *UnaryExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharExprContext struct {
	ExprContext
}

func NewCharExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharExprContext {
	var p = new(CharExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CharExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserCHAR, 0)
}

func (s *CharExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitCharExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExprContext struct {
	ExprContext
}

func NewFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExprContext {
	var p = new(FunctionCallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeticExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewArithmeticExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticExprContext {
	var p = new(ArithmeticExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArithmeticExprContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticExprContext) GetLeft() IExprContext { return s.left }

func (s *ArithmeticExprContext) GetRight() IExprContext { return s.right }

func (s *ArithmeticExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ArithmeticExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ArithmeticExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArithmeticExprContext) Op_MUL() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MUL, 0)
}

func (s *ArithmeticExprContext) Op_DIV() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_DIV, 0)
}

func (s *ArithmeticExprContext) Op_PLUS() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_PLUS, 0)
}

func (s *ArithmeticExprContext) Op_MINUS() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MINUS, 0)
}

func (s *ArithmeticExprContext) Op_MOD() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MOD, 0)
}

func (s *ArithmeticExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitArithmeticExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExprContext struct {
	ExprContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *ParExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrExprContext struct {
	ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftParserSTRING, 0)
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparasionExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewComparasionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparasionExprContext {
	var p = new(ComparasionExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparasionExprContext) GetOp() antlr.Token { return s.op }

func (s *ComparasionExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparasionExprContext) GetLeft() IExprContext { return s.left }

func (s *ComparasionExprContext) GetRight() IExprContext { return s.right }

func (s *ComparasionExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ComparasionExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ComparasionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparasionExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparasionExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComparasionExprContext) Op_GE() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_GE, 0)
}

func (s *ComparasionExprContext) Op_GT() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_GT, 0)
}

func (s *ComparasionExprContext) Op_LE() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_LE, 0)
}

func (s *ComparasionExprContext) Op_LT() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_LT, 0)
}

func (s *ComparasionExprContext) Op_EQ() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_EQ, 0)
}

func (s *ComparasionExprContext) Op_NEQ() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_NEQ, 0)
}

func (s *ComparasionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitComparasionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
	right IExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRight() IExprContext { return s.right }

func (s *NotExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Op_NOT() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_NOT, 0)
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftParserINT, 0)
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewLogicalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExprContext {
	var p = new(LogicalExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalExprContext) GetLeft() IExprContext { return s.left }

func (s *LogicalExprContext) GetRight() IExprContext { return s.right }

func (s *LogicalExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *LogicalExprContext) SetRight(v IExprContext) { s.right = v }

func (s *LogicalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicalExprContext) Op_AND() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_AND, 0)
}

func (s *LogicalExprContext) Op_OR() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_OR, 0)
}

func (s *LogicalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitLogicalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExprContext struct {
	ExprContext
	condition IExprContext
	cTrue     IExprContext
	cFalse    IExprContext
}

func NewTernaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExprContext {
	var p = new(TernaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TernaryExprContext) GetCondition() IExprContext { return s.condition }

func (s *TernaryExprContext) GetCTrue() IExprContext { return s.cTrue }

func (s *TernaryExprContext) GetCFalse() IExprContext { return s.cFalse }

func (s *TernaryExprContext) SetCondition(v IExprContext) { s.condition = v }

func (s *TernaryExprContext) SetCTrue(v IExprContext) { s.cTrue = v }

func (s *TernaryExprContext) SetCFalse(v IExprContext) { s.cFalse = v }

func (s *TernaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExprContext) Op_TERNARY() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_TERNARY, 0)
}

func (s *TernaryExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *TernaryExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *TernaryExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TernaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitTernaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SwiftParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(228)
			p.FunctionCall()
		}

	case 2:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(229)
			p.Match(SwiftParserOp_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.expr(20)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Match(SwiftParserOp_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)

			var _x = p.expr(19)

			localctx.(*NotExprContext).right = _x
		}

	case 4:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(233)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.expr(0)
		}
		{
			p.SetState(235)
			p.Match(SwiftParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(237)
			p.Match(SwiftParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)
			p.Match(SwiftParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Match(SwiftParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(241)
			p.Match(SwiftParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Match(SwiftParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(SwiftParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(279)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(247)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_MUL || _la == SwiftParserOp_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(248)

					var _x = p.expr(19)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(249)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(250)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_PLUS || _la == SwiftParserOp_MINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(251)

					var _x = p.expr(18)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 3:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(253)

					var _m = p.Match(SwiftParserOp_MOD)

					localctx.(*ArithmeticExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(254)

					var _x = p.expr(17)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 4:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(256)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparasionExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_GT || _la == SwiftParserOp_GE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparasionExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(257)

					var _x = p.expr(16)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 5:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(259)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparasionExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_LT || _la == SwiftParserOp_LE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparasionExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(260)

					var _x = p.expr(15)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 6:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(262)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparasionExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_EQ || _la == SwiftParserOp_NEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparasionExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(263)

					var _x = p.expr(14)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(265)
					p.Match(SwiftParserOp_TERNARY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(266)

					var _x = p.expr(0)

					localctx.(*TernaryExprContext).cTrue = _x
				}
				{
					p.SetState(267)
					p.Match(SwiftParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(268)

					var _x = p.expr(13)

					localctx.(*TernaryExprContext).cFalse = _x
				}

			case 8:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(271)

					var _m = p.Match(SwiftParserOp_AND)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(272)

					var _x = p.expr(12)

					localctx.(*LogicalExprContext).right = _x
				}

			case 9:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(273)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(274)

					var _m = p.Match(SwiftParserOp_OR)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(275)

					var _x = p.expr(11)

					localctx.(*LogicalExprContext).right = _x
				}

			case 10:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RangeExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(277)
					p.Match(SwiftParserKw_RANGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(278)

					var _x = p.expr(10)

					localctx.(*RangeExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IControlFlowStatementContext is an interface to support dynamic dispatch.
type IControlFlowStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsControlFlowStatementContext differentiates from other interfaces.
	IsControlFlowStatementContext()
}

type ControlFlowStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlFlowStatementContext() *ControlFlowStatementContext {
	var p = new(ControlFlowStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_controlFlowStatement
	return p
}

func InitEmptyControlFlowStatementContext(p *ControlFlowStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_controlFlowStatement
}

func (*ControlFlowStatementContext) IsControlFlowStatementContext() {}

func NewControlFlowStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlFlowStatementContext {
	var p = new(ControlFlowStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_controlFlowStatement

	return p
}

func (s *ControlFlowStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlFlowStatementContext) CopyAll(ctx *ControlFlowStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ControlFlowStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlFlowStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ControlReturnContext struct {
	ControlFlowStatementContext
}

func NewControlReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ControlReturnContext {
	var p = new(ControlReturnContext)

	InitEmptyControlFlowStatementContext(&p.ControlFlowStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlFlowStatementContext))

	return p
}

func (s *ControlReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlReturnContext) Kw_RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_RETURN, 0)
}

func (s *ControlReturnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ControlReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitControlReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type ControlBreakContext struct {
	ControlFlowStatementContext
}

func NewControlBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ControlBreakContext {
	var p = new(ControlBreakContext)

	InitEmptyControlFlowStatementContext(&p.ControlFlowStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlFlowStatementContext))

	return p
}

func (s *ControlBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlBreakContext) Kw_BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_BREAK, 0)
}

func (s *ControlBreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitControlBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

type ControlContinueContext struct {
	ControlFlowStatementContext
}

func NewControlContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ControlContinueContext {
	var p = new(ControlContinueContext)

	InitEmptyControlFlowStatementContext(&p.ControlFlowStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ControlFlowStatementContext))

	return p
}

func (s *ControlContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlContinueContext) Kw_CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_CONTINUE, 0)
}

func (s *ControlContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitControlContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) ControlFlowStatement() (localctx IControlFlowStatementContext) {
	localctx = NewControlFlowStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftParserRULE_controlFlowStatement)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserKw_BREAK:
		localctx = NewControlBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(SwiftParserKw_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserKw_CONTINUE:
		localctx = NewControlContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(SwiftParserKw_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserKw_RETURN:
		localctx = NewControlReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(286)
			p.Match(SwiftParserKw_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(287)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
