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
		"", "", "", "", "'let'", "'var'", "'func'", "'struct'", "'mutating'",
		"'if'", "'else'", "'switch'", "'case'", "'default'", "'for'", "'while'",
		"'break'", "'continue'", "'return'", "'do'", "'repeat'", "'in'", "'guard'",
		"'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'...'", "'inout'",
		"", "", "", "", "", "'nil'", "", "'->'", "'=='", "'!='", "'<'", "'>'",
		"'<='", "'>='", "'='", "'*='", "'/='", "'+='", "'-='", "'%='", "'*'",
		"'/'", "'+'", "'-'", "'%'", "'&&'", "'||'", "'!'", "'?'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'\\'", "','", "';'", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_MUTATING", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE",
		"Kw_DEFAULT", "Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN",
		"Kw_DO", "Kw_REPEAT", "Kw_IN", "Kw_GUARD", "Kw_INT", "Kw_FLOAT", "Kw_BOOL",
		"Kw_STRING", "Kw_CHAR", "Kw_RANGE", "Kw_INOUT", "INT", "FLOAT", "BOOL",
		"STRING", "CHAR", "NIL", "ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT",
		"Op_GT", "Op_LE", "Op_GE", "Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN",
		"Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN", "Op_MOD_ASSIGN", "Op_MUL", "Op_DIV",
		"Op_PLUS", "Op_MINUS", "Op_MOD", "Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY",
		"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH",
		"COMMA", "SEMICOLON", "COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"idChain", "program", "block", "statement", "variableType", "variableDeclaration",
		"functionDeclarationStatement", "functionParameters", "functionParameter",
		"functionReturnType", "functionCall", "functionCallArguments", "variableAssignment",
		"ifStatement", "ifTail", "elseStatement", "whileStatement", "switchStatement",
		"switchCase", "switchDefault", "forStatement", "guardStatement", "vectorDeclaration",
		"vectorDefinition", "vectorValues", "vectorAccess", "vectorAssignment",
		"matrixDeclaration", "matrixType", "matrixDefinition", "matrixValues",
		"matrixRepeatingDefinition", "matrixAccess", "matrixAssignment", "structDeclaration",
		"structBody", "structProperty", "expr", "controlFlowStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 478, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 5, 0, 82, 8, 0, 10, 0, 12,
		0, 85, 9, 0, 1, 1, 1, 1, 1, 1, 1, 2, 5, 2, 91, 8, 2, 10, 2, 12, 2, 94,
		9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 111, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 122, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 129, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 137, 8, 5, 3,
		5, 139, 8, 5, 1, 6, 3, 6, 142, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 148,
		8, 6, 1, 6, 1, 6, 3, 6, 152, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 5, 7, 161, 8, 7, 10, 7, 12, 7, 164, 9, 7, 1, 8, 3, 8, 167, 8, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 172, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 3, 10, 181, 8, 10, 1, 10, 1, 10, 3, 10, 185, 8, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 5, 11, 192, 8, 11, 10, 11, 12, 11, 195, 9, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 204, 8, 11, 10, 11,
		12, 11, 207, 9, 11, 3, 11, 209, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 5, 13, 218, 8, 13, 10, 13, 12, 13, 221, 9, 13, 1, 13, 3,
		13, 224, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 5, 17, 247, 8, 17, 10, 17, 12, 17, 250, 9, 17, 1, 17,
		3, 17, 253, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 262, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 268, 8, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 3, 23, 296, 8, 23, 1, 23, 1, 23, 3, 23, 300, 8,
		23, 1, 24, 1, 24, 1, 24, 5, 24, 305, 8, 24, 10, 24, 12, 24, 308, 9, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 326, 8, 27, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 336, 8, 28, 1, 29, 1,
		29, 3, 29, 340, 8, 29, 1, 29, 1, 29, 3, 29, 344, 8, 29, 1, 30, 1, 30, 1,
		30, 5, 30, 349, 8, 30, 10, 30, 12, 30, 352, 9, 30, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 376,
		8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 386,
		8, 32, 10, 32, 12, 32, 389, 9, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 5, 35, 402, 8, 35, 10, 35, 12, 35,
		405, 9, 35, 1, 36, 1, 36, 3, 36, 409, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 430, 8, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 465,
		8, 37, 10, 37, 12, 37, 468, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 474,
		8, 38, 3, 38, 476, 8, 38, 1, 38, 0, 1, 74, 39, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 0, 8, 1, 0, 23, 27,
		1, 0, 4, 5, 2, 0, 44, 44, 47, 48, 1, 0, 50, 51, 1, 0, 52, 53, 2, 0, 41,
		41, 43, 43, 2, 0, 40, 40, 42, 42, 1, 0, 38, 39, 513, 0, 78, 1, 0, 0, 0,
		2, 86, 1, 0, 0, 0, 4, 92, 1, 0, 0, 0, 6, 110, 1, 0, 0, 0, 8, 112, 1, 0,
		0, 0, 10, 138, 1, 0, 0, 0, 12, 141, 1, 0, 0, 0, 14, 157, 1, 0, 0, 0, 16,
		166, 1, 0, 0, 0, 18, 175, 1, 0, 0, 0, 20, 180, 1, 0, 0, 0, 22, 208, 1,
		0, 0, 0, 24, 210, 1, 0, 0, 0, 26, 214, 1, 0, 0, 0, 28, 225, 1, 0, 0, 0,
		30, 231, 1, 0, 0, 0, 32, 236, 1, 0, 0, 0, 34, 242, 1, 0, 0, 0, 36, 256,
		1, 0, 0, 0, 38, 263, 1, 0, 0, 0, 40, 269, 1, 0, 0, 0, 42, 277, 1, 0, 0,
		0, 44, 284, 1, 0, 0, 0, 46, 299, 1, 0, 0, 0, 48, 301, 1, 0, 0, 0, 50, 309,
		1, 0, 0, 0, 52, 314, 1, 0, 0, 0, 54, 318, 1, 0, 0, 0, 56, 335, 1, 0, 0,
		0, 58, 343, 1, 0, 0, 0, 60, 345, 1, 0, 0, 0, 62, 375, 1, 0, 0, 0, 64, 377,
		1, 0, 0, 0, 66, 390, 1, 0, 0, 0, 68, 394, 1, 0, 0, 0, 70, 403, 1, 0, 0,
		0, 72, 408, 1, 0, 0, 0, 74, 429, 1, 0, 0, 0, 76, 475, 1, 0, 0, 0, 78, 83,
		5, 36, 0, 0, 79, 80, 5, 69, 0, 0, 80, 82, 5, 36, 0, 0, 81, 79, 1, 0, 0,
		0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 1, 1,
		0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 3, 4, 2, 0, 87, 88, 5, 0, 0, 1, 88,
		3, 1, 0, 0, 0, 89, 91, 3, 6, 3, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 5, 1, 0, 0, 0, 94, 92, 1,
		0, 0, 0, 95, 111, 3, 24, 12, 0, 96, 111, 3, 10, 5, 0, 97, 111, 3, 26, 13,
		0, 98, 111, 3, 32, 16, 0, 99, 111, 3, 34, 17, 0, 100, 111, 3, 40, 20, 0,
		101, 111, 3, 42, 21, 0, 102, 111, 3, 76, 38, 0, 103, 111, 3, 12, 6, 0,
		104, 111, 3, 20, 10, 0, 105, 111, 3, 44, 22, 0, 106, 111, 3, 52, 26, 0,
		107, 111, 3, 54, 27, 0, 108, 111, 3, 66, 33, 0, 109, 111, 3, 68, 34, 0,
		110, 95, 1, 0, 0, 0, 110, 96, 1, 0, 0, 0, 110, 97, 1, 0, 0, 0, 110, 98,
		1, 0, 0, 0, 110, 99, 1, 0, 0, 0, 110, 100, 1, 0, 0, 0, 110, 101, 1, 0,
		0, 0, 110, 102, 1, 0, 0, 0, 110, 103, 1, 0, 0, 0, 110, 104, 1, 0, 0, 0,
		110, 105, 1, 0, 0, 0, 110, 106, 1, 0, 0, 0, 110, 107, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 7, 1, 0, 0, 0, 112, 113, 7,
		0, 0, 0, 113, 9, 1, 0, 0, 0, 114, 115, 7, 1, 0, 0, 115, 116, 5, 36, 0,
		0, 116, 117, 5, 68, 0, 0, 117, 118, 3, 8, 4, 0, 118, 119, 5, 44, 0, 0,
		119, 121, 3, 74, 37, 0, 120, 122, 5, 67, 0, 0, 121, 120, 1, 0, 0, 0, 121,
		122, 1, 0, 0, 0, 122, 139, 1, 0, 0, 0, 123, 124, 7, 1, 0, 0, 124, 125,
		5, 36, 0, 0, 125, 126, 5, 44, 0, 0, 126, 128, 3, 74, 37, 0, 127, 129, 5,
		67, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 139, 1, 0, 0,
		0, 130, 131, 7, 1, 0, 0, 131, 132, 5, 36, 0, 0, 132, 133, 5, 68, 0, 0,
		133, 134, 3, 8, 4, 0, 134, 136, 5, 58, 0, 0, 135, 137, 5, 67, 0, 0, 136,
		135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 114,
		1, 0, 0, 0, 138, 123, 1, 0, 0, 0, 138, 130, 1, 0, 0, 0, 139, 11, 1, 0,
		0, 0, 140, 142, 5, 8, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0,
		142, 143, 1, 0, 0, 0, 143, 144, 5, 6, 0, 0, 144, 145, 5, 36, 0, 0, 145,
		147, 5, 59, 0, 0, 146, 148, 3, 14, 7, 0, 147, 146, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 5, 60, 0, 0, 150, 152, 3, 18,
		9, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 154, 5, 61, 0, 0, 154, 155, 3, 4, 2, 0, 155, 156, 5, 62, 0, 0, 156,
		13, 1, 0, 0, 0, 157, 162, 3, 16, 8, 0, 158, 159, 5, 66, 0, 0, 159, 161,
		3, 16, 8, 0, 160, 158, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 15, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0,
		165, 167, 5, 36, 0, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 169, 5, 36, 0, 0, 169, 171, 5, 68, 0, 0, 170, 172,
		5, 29, 0, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0,
		0, 0, 173, 174, 3, 8, 4, 0, 174, 17, 1, 0, 0, 0, 175, 176, 5, 37, 0, 0,
		176, 177, 3, 8, 4, 0, 177, 19, 1, 0, 0, 0, 178, 181, 3, 0, 0, 0, 179, 181,
		3, 8, 4, 0, 180, 178, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 184, 5, 59, 0, 0, 183, 185, 3, 22, 11, 0, 184, 183, 1, 0, 0,
		0, 184, 185, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 5, 60, 0, 0, 187,
		21, 1, 0, 0, 0, 188, 193, 3, 74, 37, 0, 189, 190, 5, 66, 0, 0, 190, 192,
		3, 74, 37, 0, 191, 189, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1,
		0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 209, 1, 0, 0, 0, 195, 193, 1, 0, 0,
		0, 196, 197, 5, 36, 0, 0, 197, 198, 5, 68, 0, 0, 198, 205, 3, 74, 37, 0,
		199, 200, 5, 66, 0, 0, 200, 201, 5, 36, 0, 0, 201, 202, 5, 68, 0, 0, 202,
		204, 3, 74, 37, 0, 203, 199, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203,
		1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0,
		0, 0, 208, 188, 1, 0, 0, 0, 208, 196, 1, 0, 0, 0, 209, 23, 1, 0, 0, 0,
		210, 211, 5, 36, 0, 0, 211, 212, 7, 2, 0, 0, 212, 213, 3, 74, 37, 0, 213,
		25, 1, 0, 0, 0, 214, 219, 3, 28, 14, 0, 215, 216, 5, 10, 0, 0, 216, 218,
		3, 28, 14, 0, 217, 215, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1,
		0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0,
		0, 222, 224, 3, 30, 15, 0, 223, 222, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0,
		224, 27, 1, 0, 0, 0, 225, 226, 5, 9, 0, 0, 226, 227, 3, 74, 37, 0, 227,
		228, 5, 61, 0, 0, 228, 229, 3, 4, 2, 0, 229, 230, 5, 62, 0, 0, 230, 29,
		1, 0, 0, 0, 231, 232, 5, 10, 0, 0, 232, 233, 5, 61, 0, 0, 233, 234, 3,
		4, 2, 0, 234, 235, 5, 62, 0, 0, 235, 31, 1, 0, 0, 0, 236, 237, 5, 15, 0,
		0, 237, 238, 3, 74, 37, 0, 238, 239, 5, 61, 0, 0, 239, 240, 3, 4, 2, 0,
		240, 241, 5, 62, 0, 0, 241, 33, 1, 0, 0, 0, 242, 243, 5, 11, 0, 0, 243,
		244, 3, 74, 37, 0, 244, 248, 5, 61, 0, 0, 245, 247, 3, 36, 18, 0, 246,
		245, 1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 253, 3, 38,
		19, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0,
		254, 255, 5, 62, 0, 0, 255, 35, 1, 0, 0, 0, 256, 257, 5, 12, 0, 0, 257,
		258, 3, 74, 37, 0, 258, 259, 5, 68, 0, 0, 259, 261, 3, 4, 2, 0, 260, 262,
		5, 16, 0, 0, 261, 260, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 37, 1, 0,
		0, 0, 263, 264, 5, 13, 0, 0, 264, 265, 5, 68, 0, 0, 265, 267, 3, 4, 2,
		0, 266, 268, 5, 16, 0, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268,
		39, 1, 0, 0, 0, 269, 270, 5, 14, 0, 0, 270, 271, 5, 36, 0, 0, 271, 272,
		5, 21, 0, 0, 272, 273, 3, 74, 37, 0, 273, 274, 5, 61, 0, 0, 274, 275, 3,
		4, 2, 0, 275, 276, 5, 62, 0, 0, 276, 41, 1, 0, 0, 0, 277, 278, 5, 22, 0,
		0, 278, 279, 3, 74, 37, 0, 279, 280, 5, 10, 0, 0, 280, 281, 5, 61, 0, 0,
		281, 282, 3, 4, 2, 0, 282, 283, 5, 62, 0, 0, 283, 43, 1, 0, 0, 0, 284,
		285, 7, 1, 0, 0, 285, 286, 5, 36, 0, 0, 286, 287, 5, 68, 0, 0, 287, 288,
		5, 63, 0, 0, 288, 289, 3, 8, 4, 0, 289, 290, 5, 64, 0, 0, 290, 291, 5,
		44, 0, 0, 291, 292, 3, 46, 23, 0, 292, 45, 1, 0, 0, 0, 293, 295, 5, 63,
		0, 0, 294, 296, 3, 48, 24, 0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0,
		0, 296, 297, 1, 0, 0, 0, 297, 300, 5, 64, 0, 0, 298, 300, 3, 74, 37, 0,
		299, 293, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 47, 1, 0, 0, 0, 301, 306,
		3, 74, 37, 0, 302, 303, 5, 66, 0, 0, 303, 305, 3, 74, 37, 0, 304, 302,
		1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0,
		0, 0, 307, 49, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 310, 3, 0, 0, 0,
		310, 311, 5, 63, 0, 0, 311, 312, 3, 74, 37, 0, 312, 313, 5, 64, 0, 0, 313,
		51, 1, 0, 0, 0, 314, 315, 3, 50, 25, 0, 315, 316, 7, 2, 0, 0, 316, 317,
		3, 74, 37, 0, 317, 53, 1, 0, 0, 0, 318, 319, 7, 1, 0, 0, 319, 320, 3, 0,
		0, 0, 320, 321, 5, 68, 0, 0, 321, 322, 3, 56, 28, 0, 322, 325, 5, 44, 0,
		0, 323, 326, 3, 62, 31, 0, 324, 326, 3, 58, 29, 0, 325, 323, 1, 0, 0, 0,
		325, 324, 1, 0, 0, 0, 326, 55, 1, 0, 0, 0, 327, 328, 5, 63, 0, 0, 328,
		329, 3, 56, 28, 0, 329, 330, 5, 64, 0, 0, 330, 336, 1, 0, 0, 0, 331, 332,
		5, 63, 0, 0, 332, 333, 3, 8, 4, 0, 333, 334, 5, 64, 0, 0, 334, 336, 1,
		0, 0, 0, 335, 327, 1, 0, 0, 0, 335, 331, 1, 0, 0, 0, 336, 57, 1, 0, 0,
		0, 337, 339, 5, 63, 0, 0, 338, 340, 3, 60, 30, 0, 339, 338, 1, 0, 0, 0,
		339, 340, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 344, 5, 64, 0, 0, 342,
		344, 3, 74, 37, 0, 343, 337, 1, 0, 0, 0, 343, 342, 1, 0, 0, 0, 344, 59,
		1, 0, 0, 0, 345, 350, 3, 58, 29, 0, 346, 347, 5, 66, 0, 0, 347, 349, 3,
		58, 29, 0, 348, 346, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348, 1, 0,
		0, 0, 350, 351, 1, 0, 0, 0, 351, 61, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0,
		353, 354, 3, 56, 28, 0, 354, 355, 5, 59, 0, 0, 355, 356, 5, 36, 0, 0, 356,
		357, 5, 68, 0, 0, 357, 358, 3, 62, 31, 0, 358, 359, 5, 66, 0, 0, 359, 360,
		5, 36, 0, 0, 360, 361, 5, 68, 0, 0, 361, 362, 3, 74, 37, 0, 362, 363, 5,
		60, 0, 0, 363, 376, 1, 0, 0, 0, 364, 365, 3, 56, 28, 0, 365, 366, 5, 59,
		0, 0, 366, 367, 5, 36, 0, 0, 367, 368, 5, 68, 0, 0, 368, 369, 3, 74, 37,
		0, 369, 370, 5, 66, 0, 0, 370, 371, 5, 36, 0, 0, 371, 372, 5, 68, 0, 0,
		372, 373, 3, 74, 37, 0, 373, 374, 5, 60, 0, 0, 374, 376, 1, 0, 0, 0, 375,
		353, 1, 0, 0, 0, 375, 364, 1, 0, 0, 0, 376, 63, 1, 0, 0, 0, 377, 378, 3,
		0, 0, 0, 378, 379, 5, 63, 0, 0, 379, 380, 3, 74, 37, 0, 380, 387, 5, 64,
		0, 0, 381, 382, 5, 63, 0, 0, 382, 383, 3, 74, 37, 0, 383, 384, 5, 64, 0,
		0, 384, 386, 1, 0, 0, 0, 385, 381, 1, 0, 0, 0, 386, 389, 1, 0, 0, 0, 387,
		385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 65, 1, 0, 0, 0, 389, 387, 1,
		0, 0, 0, 390, 391, 3, 64, 32, 0, 391, 392, 7, 2, 0, 0, 392, 393, 3, 74,
		37, 0, 393, 67, 1, 0, 0, 0, 394, 395, 5, 7, 0, 0, 395, 396, 5, 36, 0, 0,
		396, 397, 5, 61, 0, 0, 397, 398, 3, 70, 35, 0, 398, 399, 5, 62, 0, 0, 399,
		69, 1, 0, 0, 0, 400, 402, 3, 72, 36, 0, 401, 400, 1, 0, 0, 0, 402, 405,
		1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 71, 1, 0,
		0, 0, 405, 403, 1, 0, 0, 0, 406, 409, 3, 10, 5, 0, 407, 409, 3, 12, 6,
		0, 408, 406, 1, 0, 0, 0, 408, 407, 1, 0, 0, 0, 409, 73, 1, 0, 0, 0, 410,
		411, 6, 37, -1, 0, 411, 430, 3, 20, 10, 0, 412, 430, 3, 50, 25, 0, 413,
		430, 3, 64, 32, 0, 414, 415, 5, 53, 0, 0, 415, 430, 3, 74, 37, 20, 416,
		417, 5, 57, 0, 0, 417, 430, 3, 74, 37, 19, 418, 419, 5, 59, 0, 0, 419,
		420, 3, 74, 37, 0, 420, 421, 5, 60, 0, 0, 421, 430, 1, 0, 0, 0, 422, 430,
		5, 30, 0, 0, 423, 430, 3, 0, 0, 0, 424, 430, 5, 31, 0, 0, 425, 430, 5,
		33, 0, 0, 426, 430, 5, 35, 0, 0, 427, 430, 5, 34, 0, 0, 428, 430, 5, 32,
		0, 0, 429, 410, 1, 0, 0, 0, 429, 412, 1, 0, 0, 0, 429, 413, 1, 0, 0, 0,
		429, 414, 1, 0, 0, 0, 429, 416, 1, 0, 0, 0, 429, 418, 1, 0, 0, 0, 429,
		422, 1, 0, 0, 0, 429, 423, 1, 0, 0, 0, 429, 424, 1, 0, 0, 0, 429, 425,
		1, 0, 0, 0, 429, 426, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 428, 1, 0,
		0, 0, 430, 466, 1, 0, 0, 0, 431, 432, 10, 18, 0, 0, 432, 433, 7, 3, 0,
		0, 433, 465, 3, 74, 37, 19, 434, 435, 10, 17, 0, 0, 435, 436, 7, 4, 0,
		0, 436, 465, 3, 74, 37, 18, 437, 438, 10, 16, 0, 0, 438, 439, 5, 54, 0,
		0, 439, 465, 3, 74, 37, 17, 440, 441, 10, 15, 0, 0, 441, 442, 7, 5, 0,
		0, 442, 465, 3, 74, 37, 16, 443, 444, 10, 14, 0, 0, 444, 445, 7, 6, 0,
		0, 445, 465, 3, 74, 37, 15, 446, 447, 10, 13, 0, 0, 447, 448, 7, 7, 0,
		0, 448, 465, 3, 74, 37, 14, 449, 450, 10, 12, 0, 0, 450, 451, 5, 58, 0,
		0, 451, 452, 3, 74, 37, 0, 452, 453, 5, 68, 0, 0, 453, 454, 3, 74, 37,
		13, 454, 465, 1, 0, 0, 0, 455, 456, 10, 11, 0, 0, 456, 457, 5, 55, 0, 0,
		457, 465, 3, 74, 37, 12, 458, 459, 10, 10, 0, 0, 459, 460, 5, 56, 0, 0,
		460, 465, 3, 74, 37, 11, 461, 462, 10, 9, 0, 0, 462, 463, 5, 28, 0, 0,
		463, 465, 3, 74, 37, 10, 464, 431, 1, 0, 0, 0, 464, 434, 1, 0, 0, 0, 464,
		437, 1, 0, 0, 0, 464, 440, 1, 0, 0, 0, 464, 443, 1, 0, 0, 0, 464, 446,
		1, 0, 0, 0, 464, 449, 1, 0, 0, 0, 464, 455, 1, 0, 0, 0, 464, 458, 1, 0,
		0, 0, 464, 461, 1, 0, 0, 0, 465, 468, 1, 0, 0, 0, 466, 464, 1, 0, 0, 0,
		466, 467, 1, 0, 0, 0, 467, 75, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 469, 476,
		5, 16, 0, 0, 470, 476, 5, 17, 0, 0, 471, 473, 5, 18, 0, 0, 472, 474, 3,
		74, 37, 0, 473, 472, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 476, 1, 0,
		0, 0, 475, 469, 1, 0, 0, 0, 475, 470, 1, 0, 0, 0, 475, 471, 1, 0, 0, 0,
		476, 77, 1, 0, 0, 0, 41, 83, 92, 110, 121, 128, 136, 138, 141, 147, 151,
		162, 166, 171, 180, 184, 193, 205, 208, 219, 223, 248, 252, 261, 267, 295,
		299, 306, 325, 335, 339, 343, 350, 375, 387, 403, 408, 429, 464, 466, 473,
		475,
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
	SwiftParserKw_MUTATING     = 8
	SwiftParserKw_IF           = 9
	SwiftParserKw_ELSE         = 10
	SwiftParserKw_SWITCH       = 11
	SwiftParserKw_CASE         = 12
	SwiftParserKw_DEFAULT      = 13
	SwiftParserKw_FOR          = 14
	SwiftParserKw_WHILE        = 15
	SwiftParserKw_BREAK        = 16
	SwiftParserKw_CONTINUE     = 17
	SwiftParserKw_RETURN       = 18
	SwiftParserKw_DO           = 19
	SwiftParserKw_REPEAT       = 20
	SwiftParserKw_IN           = 21
	SwiftParserKw_GUARD        = 22
	SwiftParserKw_INT          = 23
	SwiftParserKw_FLOAT        = 24
	SwiftParserKw_BOOL         = 25
	SwiftParserKw_STRING       = 26
	SwiftParserKw_CHAR         = 27
	SwiftParserKw_RANGE        = 28
	SwiftParserKw_INOUT        = 29
	SwiftParserINT             = 30
	SwiftParserFLOAT           = 31
	SwiftParserBOOL            = 32
	SwiftParserSTRING          = 33
	SwiftParserCHAR            = 34
	SwiftParserNIL             = 35
	SwiftParserID              = 36
	SwiftParserOp_ARROW        = 37
	SwiftParserOp_EQ           = 38
	SwiftParserOp_NEQ          = 39
	SwiftParserOp_LT           = 40
	SwiftParserOp_GT           = 41
	SwiftParserOp_LE           = 42
	SwiftParserOp_GE           = 43
	SwiftParserOp_ASSIGN       = 44
	SwiftParserOp_MUL_ASSIGN   = 45
	SwiftParserOp_DIV_ASSIGN   = 46
	SwiftParserOp_PLUS_ASSIGN  = 47
	SwiftParserOp_MINUS_ASSIGN = 48
	SwiftParserOp_MOD_ASSIGN   = 49
	SwiftParserOp_MUL          = 50
	SwiftParserOp_DIV          = 51
	SwiftParserOp_PLUS         = 52
	SwiftParserOp_MINUS        = 53
	SwiftParserOp_MOD          = 54
	SwiftParserOp_AND          = 55
	SwiftParserOp_OR           = 56
	SwiftParserOp_NOT          = 57
	SwiftParserOp_TERNARY      = 58
	SwiftParserLPAREN          = 59
	SwiftParserRPAREN          = 60
	SwiftParserLBRACE          = 61
	SwiftParserRBRACE          = 62
	SwiftParserLBRACKET        = 63
	SwiftParserRBRACKET        = 64
	SwiftParserBACKSLASH       = 65
	SwiftParserCOMMA           = 66
	SwiftParserSEMICOLON       = 67
	SwiftParserCOLON           = 68
	SwiftParserDOT             = 69
)

// SwiftParser rules.
const (
	SwiftParserRULE_idChain                      = 0
	SwiftParserRULE_program                      = 1
	SwiftParserRULE_block                        = 2
	SwiftParserRULE_statement                    = 3
	SwiftParserRULE_variableType                 = 4
	SwiftParserRULE_variableDeclaration          = 5
	SwiftParserRULE_functionDeclarationStatement = 6
	SwiftParserRULE_functionParameters           = 7
	SwiftParserRULE_functionParameter            = 8
	SwiftParserRULE_functionReturnType           = 9
	SwiftParserRULE_functionCall                 = 10
	SwiftParserRULE_functionCallArguments        = 11
	SwiftParserRULE_variableAssignment           = 12
	SwiftParserRULE_ifStatement                  = 13
	SwiftParserRULE_ifTail                       = 14
	SwiftParserRULE_elseStatement                = 15
	SwiftParserRULE_whileStatement               = 16
	SwiftParserRULE_switchStatement              = 17
	SwiftParserRULE_switchCase                   = 18
	SwiftParserRULE_switchDefault                = 19
	SwiftParserRULE_forStatement                 = 20
	SwiftParserRULE_guardStatement               = 21
	SwiftParserRULE_vectorDeclaration            = 22
	SwiftParserRULE_vectorDefinition             = 23
	SwiftParserRULE_vectorValues                 = 24
	SwiftParserRULE_vectorAccess                 = 25
	SwiftParserRULE_vectorAssignment             = 26
	SwiftParserRULE_matrixDeclaration            = 27
	SwiftParserRULE_matrixType                   = 28
	SwiftParserRULE_matrixDefinition             = 29
	SwiftParserRULE_matrixValues                 = 30
	SwiftParserRULE_matrixRepeatingDefinition    = 31
	SwiftParserRULE_matrixAccess                 = 32
	SwiftParserRULE_matrixAssignment             = 33
	SwiftParserRULE_structDeclaration            = 34
	SwiftParserRULE_structBody                   = 35
	SwiftParserRULE_structProperty               = 36
	SwiftParserRULE_expr                         = 37
	SwiftParserRULE_controlFlowStatement         = 38
)

// IIdChainContext is an interface to support dynamic dispatch.
type IIdChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdChainContext differentiates from other interfaces.
	IsIdChainContext()
}

type IdChainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdChainContext() *IdChainContext {
	var p = new(IdChainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_idChain
	return p
}

func InitEmptyIdChainContext(p *IdChainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_idChain
}

func (*IdChainContext) IsIdChainContext() {}

func NewIdChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdChainContext {
	var p = new(IdChainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_idChain

	return p
}

func (s *IdChainContext) GetParser() antlr.Parser { return s.parser }

func (s *IdChainContext) CopyAll(ctx *IdChainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IDChainContext struct {
	IdChainContext
}

func NewIDChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDChainContext {
	var p = new(IDChainContext)

	InitEmptyIdChainContext(&p.IdChainContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdChainContext))

	return p
}

func (s *IDChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDChainContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserID)
}

func (s *IDChainContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserID, i)
}

func (s *IDChainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserDOT)
}

func (s *IDChainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserDOT, i)
}

func (s *IDChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIDChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) IdChain() (localctx IIdChainContext) {
	localctx = NewIdChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftParserRULE_idChain)
	var _alt int

	localctx = NewIDChainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
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
				p.SetState(79)
				p.Match(SwiftParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(80)
				p.Match(SwiftParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(85)
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
	p.EnterRule(localctx, 2, SwiftParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Block()
	}
	{
		p.SetState(87)
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
	p.EnterRule(localctx, 4, SwiftParserRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(89)
				p.Statement()
			}

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
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
	VectorDeclaration() IVectorDeclarationContext
	VectorAssignment() IVectorAssignmentContext
	MatrixDeclaration() IMatrixDeclarationContext
	MatrixAssignment() IMatrixAssignmentContext
	StructDeclaration() IStructDeclarationContext

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

func (s *StatementContext) VectorDeclaration() IVectorDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorDeclarationContext)
}

func (s *StatementContext) VectorAssignment() IVectorAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorAssignmentContext)
}

func (s *StatementContext) MatrixDeclaration() IMatrixDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixDeclarationContext)
}

func (s *StatementContext) MatrixAssignment() IMatrixAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixAssignmentContext)
}

func (s *StatementContext) StructDeclaration() IStructDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
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
	p.EnterRule(localctx, 6, SwiftParserRULE_statement)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.VariableAssignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.VariableDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.WhileStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.SwitchStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.ForStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.GuardStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(102)
			p.ControlFlowStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(103)
			p.FunctionDeclarationStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(104)
			p.FunctionCall()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(105)
			p.VectorDeclaration()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(106)
			p.VectorAssignment()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(107)
			p.MatrixDeclaration()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(108)
			p.MatrixAssignment()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(109)
			p.StructDeclaration()
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
	p.EnterRule(localctx, 8, SwiftParserRULE_variableType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260046848) != 0) {
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
	p.EnterRule(localctx, 10, SwiftParserRULE_variableDeclaration)
	var _la int

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)

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
			p.SetState(115)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.VariableType()
		}
		{
			p.SetState(118)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.expr(0)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(120)
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
			p.SetState(123)

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
			p.SetState(124)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.expr(0)
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(127)
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
			p.SetState(130)

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
			p.SetState(131)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.VariableType()
		}
		{
			p.SetState(134)
			p.Match(SwiftParserOp_TERNARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(135)
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
	Kw_MUTATING() antlr.TerminalNode
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

func (s *FunctionDeclarationStatementContext) Kw_MUTATING() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_MUTATING, 0)
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
	p.EnterRule(localctx, 12, SwiftParserRULE_functionDeclarationStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_MUTATING {
		{
			p.SetState(140)
			p.Match(SwiftParserKw_MUTATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(143)
		p.Match(SwiftParserKw_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(SwiftParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserID {
		{
			p.SetState(146)
			p.FunctionParameters()
		}

	}
	{
		p.SetState(149)
		p.Match(SwiftParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserOp_ARROW {
		{
			p.SetState(150)
			p.FunctionReturnType()
		}

	}
	{
		p.SetState(153)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Block()
	}
	{
		p.SetState(155)
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
	p.EnterRule(localctx, 14, SwiftParserRULE_functionParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.FunctionParameter()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserCOMMA {
		{
			p.SetState(158)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.FunctionParameter()
		}

		p.SetState(164)
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
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode
	VariableType() IVariableTypeContext
	Kw_INOUT() antlr.TerminalNode

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

func (s *FunctionParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserID)
}

func (s *FunctionParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserID, i)
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

func (s *FunctionParameterContext) Kw_INOUT() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_INOUT, 0)
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
	p.EnterRule(localctx, 16, SwiftParserRULE_functionParameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(165)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(168)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_INOUT {
		{
			p.SetState(170)
			p.Match(SwiftParserKw_INOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(173)
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
	p.EnterRule(localctx, 18, SwiftParserRULE_functionReturnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(SwiftParserOp_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
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
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	IdChain() IIdChainContext
	VariableType() IVariableTypeContext
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

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *FunctionCallContext) IdChain() IIdChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdChainContext)
}

func (s *FunctionCallContext) VariableType() IVariableTypeContext {
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
	p.EnterRule(localctx, 20, SwiftParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserID:
		{
			p.SetState(178)
			p.IdChain()
		}

	case SwiftParserKw_INT, SwiftParserKw_FLOAT, SwiftParserKw_BOOL, SwiftParserKw_STRING, SwiftParserKw_CHAR:
		{
			p.SetState(179)
			p.VariableType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(182)
		p.Match(SwiftParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&729583276259278848) != 0 {
		{
			p.SetState(183)
			p.FunctionCallArguments()
		}

	}
	{
		p.SetState(186)
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
	p.EnterRule(localctx, 22, SwiftParserRULE_functionCallArguments)
	var _la int

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArgumentsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.expr(0)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftParserCOMMA {
			{
				p.SetState(189)
				p.Match(SwiftParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.expr(0)
			}

			p.SetState(195)
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
			p.SetState(196)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.expr(0)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftParserCOMMA {
			{
				p.SetState(199)
				p.Match(SwiftParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Match(SwiftParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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
				p.expr(0)
			}

			p.SetState(207)
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
	p.EnterRule(localctx, 24, SwiftParserRULE_variableAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&439804651110400) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(212)
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
	p.EnterRule(localctx, 26, SwiftParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.IfTail()
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(215)
				p.Match(SwiftParserKw_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(216)
				p.IfTail()
			}

		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_ELSE {
		{
			p.SetState(222)
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
	p.EnterRule(localctx, 28, SwiftParserRULE_ifTail)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(SwiftParserKw_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.expr(0)
	}
	{
		p.SetState(227)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Block()
	}
	{
		p.SetState(229)
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
	p.EnterRule(localctx, 30, SwiftParserRULE_elseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Block()
	}
	{
		p.SetState(234)
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
	p.EnterRule(localctx, 32, SwiftParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(SwiftParserKw_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.expr(0)
	}
	{
		p.SetState(238)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Block()
	}
	{
		p.SetState(240)
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
	p.EnterRule(localctx, 34, SwiftParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(SwiftParserKw_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}
	{
		p.SetState(244)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserKw_CASE {
		{
			p.SetState(245)
			p.SwitchCase()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_DEFAULT {
		{
			p.SetState(251)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(254)
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
	p.EnterRule(localctx, 36, SwiftParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(SwiftParserKw_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.expr(0)
	}
	{
		p.SetState(258)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Block()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(260)
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
	p.EnterRule(localctx, 38, SwiftParserRULE_switchDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(SwiftParserKw_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Block()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(266)
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
	p.EnterRule(localctx, 40, SwiftParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(SwiftParserKw_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(SwiftParserKw_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.expr(0)
	}
	{
		p.SetState(273)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Block()
	}
	{
		p.SetState(275)
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
	p.EnterRule(localctx, 42, SwiftParserRULE_guardStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(SwiftParserKw_GUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.expr(0)
	}
	{
		p.SetState(279)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Block()
	}
	{
		p.SetState(282)
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

// IVectorDeclarationContext is an interface to support dynamic dispatch.
type IVectorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVarType returns the varType token.
	GetVarType() antlr.Token

	// SetVarType sets the varType token.
	SetVarType(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	LBRACKET() antlr.TerminalNode
	VariableType() IVariableTypeContext
	RBRACKET() antlr.TerminalNode
	Op_ASSIGN() antlr.TerminalNode
	VectorDefinition() IVectorDefinitionContext
	Kw_LET() antlr.TerminalNode
	Kw_VAR() antlr.TerminalNode

	// IsVectorDeclarationContext differentiates from other interfaces.
	IsVectorDeclarationContext()
}

type VectorDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	varType antlr.Token
}

func NewEmptyVectorDeclarationContext() *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorDeclaration
	return p
}

func InitEmptyVectorDeclarationContext(p *VectorDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorDeclaration
}

func (*VectorDeclarationContext) IsVectorDeclarationContext() {}

func NewVectorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_vectorDeclaration

	return p
}

func (s *VectorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorDeclarationContext) GetVarType() antlr.Token { return s.varType }

func (s *VectorDeclarationContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *VectorDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *VectorDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *VectorDeclarationContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, 0)
}

func (s *VectorDeclarationContext) VariableType() IVariableTypeContext {
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

func (s *VectorDeclarationContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, 0)
}

func (s *VectorDeclarationContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *VectorDeclarationContext) VectorDefinition() IVectorDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorDefinitionContext)
}

func (s *VectorDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_LET, 0)
}

func (s *VectorDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_VAR, 0)
}

func (s *VectorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VectorDeclaration() (localctx IVectorDeclarationContext) {
	localctx = NewVectorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftParserRULE_vectorDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VectorDeclarationContext).varType = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftParserKw_LET || _la == SwiftParserKw_VAR) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VectorDeclarationContext).varType = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(285)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(SwiftParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.VariableType()
	}
	{
		p.SetState(289)
		p.Match(SwiftParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(SwiftParserOp_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.VectorDefinition()
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

// IVectorDefinitionContext is an interface to support dynamic dispatch.
type IVectorDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVectorDefinitionContext differentiates from other interfaces.
	IsVectorDefinitionContext()
}

type VectorDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorDefinitionContext() *VectorDefinitionContext {
	var p = new(VectorDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorDefinition
	return p
}

func InitEmptyVectorDefinitionContext(p *VectorDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorDefinition
}

func (*VectorDefinitionContext) IsVectorDefinitionContext() {}

func NewVectorDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorDefinitionContext {
	var p = new(VectorDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_vectorDefinition

	return p
}

func (s *VectorDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorDefinitionContext) CopyAll(ctx *VectorDefinitionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VectorDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorListValueContext struct {
	VectorDefinitionContext
}

func NewVectorListValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorListValueContext {
	var p = new(VectorListValueContext)

	InitEmptyVectorDefinitionContext(&p.VectorDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorDefinitionContext))

	return p
}

func (s *VectorListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorListValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, 0)
}

func (s *VectorListValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, 0)
}

func (s *VectorListValueContext) VectorValues() IVectorValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorValuesContext)
}

func (s *VectorListValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorListValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorSingleValueContext struct {
	VectorDefinitionContext
}

func NewVectorSingleValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorSingleValueContext {
	var p = new(VectorSingleValueContext)

	InitEmptyVectorDefinitionContext(&p.VectorDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorDefinitionContext))

	return p
}

func (s *VectorSingleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorSingleValueContext) Expr() IExprContext {
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

func (s *VectorSingleValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorSingleValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VectorDefinition() (localctx IVectorDefinitionContext) {
	localctx = NewVectorDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftParserRULE_vectorDefinition)
	var _la int

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserLBRACKET:
		localctx = NewVectorListValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(SwiftParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&729583276259278848) != 0 {
			{
				p.SetState(294)
				p.VectorValues()
			}

		}
		{
			p.SetState(297)
			p.Match(SwiftParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserKw_INT, SwiftParserKw_FLOAT, SwiftParserKw_BOOL, SwiftParserKw_STRING, SwiftParserKw_CHAR, SwiftParserINT, SwiftParserFLOAT, SwiftParserBOOL, SwiftParserSTRING, SwiftParserCHAR, SwiftParserNIL, SwiftParserID, SwiftParserOp_MINUS, SwiftParserOp_NOT, SwiftParserLPAREN:
		localctx = NewVectorSingleValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.expr(0)
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

// IVectorValuesContext is an interface to support dynamic dispatch.
type IVectorValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVectorValuesContext differentiates from other interfaces.
	IsVectorValuesContext()
}

type VectorValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorValuesContext() *VectorValuesContext {
	var p = new(VectorValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorValues
	return p
}

func InitEmptyVectorValuesContext(p *VectorValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorValues
}

func (*VectorValuesContext) IsVectorValuesContext() {}

func NewVectorValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorValuesContext {
	var p = new(VectorValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_vectorValues

	return p
}

func (s *VectorValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorValuesContext) AllExpr() []IExprContext {
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

func (s *VectorValuesContext) Expr(i int) IExprContext {
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

func (s *VectorValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOMMA)
}

func (s *VectorValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, i)
}

func (s *VectorValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VectorValues() (localctx IVectorValuesContext) {
	localctx = NewVectorValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftParserRULE_vectorValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.expr(0)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserCOMMA {
		{
			p.SetState(302)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.expr(0)
		}

		p.SetState(308)
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

// IVectorAccessContext is an interface to support dynamic dispatch.
type IVectorAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdChain() IIdChainContext
	LBRACKET() antlr.TerminalNode
	Expr() IExprContext
	RBRACKET() antlr.TerminalNode

	// IsVectorAccessContext differentiates from other interfaces.
	IsVectorAccessContext()
}

type VectorAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorAccessContext() *VectorAccessContext {
	var p = new(VectorAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorAccess
	return p
}

func InitEmptyVectorAccessContext(p *VectorAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorAccess
}

func (*VectorAccessContext) IsVectorAccessContext() {}

func NewVectorAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorAccessContext {
	var p = new(VectorAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_vectorAccess

	return p
}

func (s *VectorAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorAccessContext) IdChain() IIdChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdChainContext)
}

func (s *VectorAccessContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, 0)
}

func (s *VectorAccessContext) Expr() IExprContext {
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

func (s *VectorAccessContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, 0)
}

func (s *VectorAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VectorAccess() (localctx IVectorAccessContext) {
	localctx = NewVectorAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftParserRULE_vectorAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.IdChain()
	}
	{
		p.SetState(310)
		p.Match(SwiftParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.expr(0)
	}
	{
		p.SetState(312)
		p.Match(SwiftParserRBRACKET)
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

// IVectorAssignmentContext is an interface to support dynamic dispatch.
type IVectorAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	VectorAccess() IVectorAccessContext
	Expr() IExprContext
	Op_ASSIGN() antlr.TerminalNode
	Op_PLUS_ASSIGN() antlr.TerminalNode
	Op_MINUS_ASSIGN() antlr.TerminalNode

	// IsVectorAssignmentContext differentiates from other interfaces.
	IsVectorAssignmentContext()
}

type VectorAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyVectorAssignmentContext() *VectorAssignmentContext {
	var p = new(VectorAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorAssignment
	return p
}

func InitEmptyVectorAssignmentContext(p *VectorAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_vectorAssignment
}

func (*VectorAssignmentContext) IsVectorAssignmentContext() {}

func NewVectorAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorAssignmentContext {
	var p = new(VectorAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_vectorAssignment

	return p
}

func (s *VectorAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorAssignmentContext) GetOp() antlr.Token { return s.op }

func (s *VectorAssignmentContext) SetOp(v antlr.Token) { s.op = v }

func (s *VectorAssignmentContext) VectorAccess() IVectorAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorAccessContext)
}

func (s *VectorAssignmentContext) Expr() IExprContext {
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

func (s *VectorAssignmentContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *VectorAssignmentContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_PLUS_ASSIGN, 0)
}

func (s *VectorAssignmentContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MINUS_ASSIGN, 0)
}

func (s *VectorAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) VectorAssignment() (localctx IVectorAssignmentContext) {
	localctx = NewVectorAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftParserRULE_vectorAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.VectorAccess()
	}
	{
		p.SetState(315)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VectorAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&439804651110400) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VectorAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(316)
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

// IMatrixDeclarationContext is an interface to support dynamic dispatch.
type IMatrixDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVarType returns the varType token.
	GetVarType() antlr.Token

	// SetVarType sets the varType token.
	SetVarType(antlr.Token)

	// Getter signatures
	IdChain() IIdChainContext
	COLON() antlr.TerminalNode
	MatrixType() IMatrixTypeContext
	Op_ASSIGN() antlr.TerminalNode
	Kw_LET() antlr.TerminalNode
	Kw_VAR() antlr.TerminalNode
	MatrixRepeatingDefinition() IMatrixRepeatingDefinitionContext
	MatrixDefinition() IMatrixDefinitionContext

	// IsMatrixDeclarationContext differentiates from other interfaces.
	IsMatrixDeclarationContext()
}

type MatrixDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	varType antlr.Token
}

func NewEmptyMatrixDeclarationContext() *MatrixDeclarationContext {
	var p = new(MatrixDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixDeclaration
	return p
}

func InitEmptyMatrixDeclarationContext(p *MatrixDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixDeclaration
}

func (*MatrixDeclarationContext) IsMatrixDeclarationContext() {}

func NewMatrixDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixDeclarationContext {
	var p = new(MatrixDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixDeclaration

	return p
}

func (s *MatrixDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixDeclarationContext) GetVarType() antlr.Token { return s.varType }

func (s *MatrixDeclarationContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *MatrixDeclarationContext) IdChain() IIdChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdChainContext)
}

func (s *MatrixDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *MatrixDeclarationContext) MatrixType() IMatrixTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixTypeContext)
}

func (s *MatrixDeclarationContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *MatrixDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_LET, 0)
}

func (s *MatrixDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_VAR, 0)
}

func (s *MatrixDeclarationContext) MatrixRepeatingDefinition() IMatrixRepeatingDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixRepeatingDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixRepeatingDefinitionContext)
}

func (s *MatrixDeclarationContext) MatrixDefinition() IMatrixDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixDefinitionContext)
}

func (s *MatrixDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixDeclaration() (localctx IMatrixDeclarationContext) {
	localctx = NewMatrixDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftParserRULE_matrixDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MatrixDeclarationContext).varType = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftParserKw_LET || _la == SwiftParserKw_VAR) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MatrixDeclarationContext).varType = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(319)
		p.IdChain()
	}
	{
		p.SetState(320)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.MatrixType()
	}
	{
		p.SetState(322)
		p.Match(SwiftParserOp_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(323)
			p.MatrixRepeatingDefinition()
		}

	case 2:
		{
			p.SetState(324)
			p.MatrixDefinition()
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

// IMatrixTypeContext is an interface to support dynamic dispatch.
type IMatrixTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMatrixTypeContext differentiates from other interfaces.
	IsMatrixTypeContext()
}

type MatrixTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixTypeContext() *MatrixTypeContext {
	var p = new(MatrixTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixType
	return p
}

func InitEmptyMatrixTypeContext(p *MatrixTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixType
}

func (*MatrixTypeContext) IsMatrixTypeContext() {}

func NewMatrixTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixTypeContext {
	var p = new(MatrixTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixType

	return p
}

func (s *MatrixTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixTypeContext) CopyAll(ctx *MatrixTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MatrixTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MatrixTypeSingleContext struct {
	MatrixTypeContext
}

func NewMatrixTypeSingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixTypeSingleContext {
	var p = new(MatrixTypeSingleContext)

	InitEmptyMatrixTypeContext(&p.MatrixTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatrixTypeContext))

	return p
}

func (s *MatrixTypeSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixTypeSingleContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, 0)
}

func (s *MatrixTypeSingleContext) VariableType() IVariableTypeContext {
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

func (s *MatrixTypeSingleContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, 0)
}

func (s *MatrixTypeSingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixTypeSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrixTypeNestedContext struct {
	MatrixTypeContext
}

func NewMatrixTypeNestedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixTypeNestedContext {
	var p = new(MatrixTypeNestedContext)

	InitEmptyMatrixTypeContext(&p.MatrixTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatrixTypeContext))

	return p
}

func (s *MatrixTypeNestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixTypeNestedContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, 0)
}

func (s *MatrixTypeNestedContext) MatrixType() IMatrixTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixTypeContext)
}

func (s *MatrixTypeNestedContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, 0)
}

func (s *MatrixTypeNestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixTypeNested(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixType() (localctx IMatrixTypeContext) {
	localctx = NewMatrixTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftParserRULE_matrixType)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMatrixTypeNestedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(SwiftParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.MatrixType()
		}
		{
			p.SetState(329)
			p.Match(SwiftParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMatrixTypeSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(SwiftParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.VariableType()
		}
		{
			p.SetState(333)
			p.Match(SwiftParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

// IMatrixDefinitionContext is an interface to support dynamic dispatch.
type IMatrixDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	MatrixValues() IMatrixValuesContext
	Expr() IExprContext

	// IsMatrixDefinitionContext differentiates from other interfaces.
	IsMatrixDefinitionContext()
}

type MatrixDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixDefinitionContext() *MatrixDefinitionContext {
	var p = new(MatrixDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixDefinition
	return p
}

func InitEmptyMatrixDefinitionContext(p *MatrixDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixDefinition
}

func (*MatrixDefinitionContext) IsMatrixDefinitionContext() {}

func NewMatrixDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixDefinitionContext {
	var p = new(MatrixDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixDefinition

	return p
}

func (s *MatrixDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixDefinitionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, 0)
}

func (s *MatrixDefinitionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, 0)
}

func (s *MatrixDefinitionContext) MatrixValues() IMatrixValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixValuesContext)
}

func (s *MatrixDefinitionContext) Expr() IExprContext {
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

func (s *MatrixDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixDefinition() (localctx IMatrixDefinitionContext) {
	localctx = NewMatrixDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftParserRULE_matrixDefinition)
	var _la int

	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(337)
			p.Match(SwiftParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-8493788760595496960) != 0 {
			{
				p.SetState(338)
				p.MatrixValues()
			}

		}
		{
			p.SetState(341)
			p.Match(SwiftParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserKw_INT, SwiftParserKw_FLOAT, SwiftParserKw_BOOL, SwiftParserKw_STRING, SwiftParserKw_CHAR, SwiftParserINT, SwiftParserFLOAT, SwiftParserBOOL, SwiftParserSTRING, SwiftParserCHAR, SwiftParserNIL, SwiftParserID, SwiftParserOp_MINUS, SwiftParserOp_NOT, SwiftParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.expr(0)
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

// IMatrixValuesContext is an interface to support dynamic dispatch.
type IMatrixValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMatrixDefinition() []IMatrixDefinitionContext
	MatrixDefinition(i int) IMatrixDefinitionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMatrixValuesContext differentiates from other interfaces.
	IsMatrixValuesContext()
}

type MatrixValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixValuesContext() *MatrixValuesContext {
	var p = new(MatrixValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixValues
	return p
}

func InitEmptyMatrixValuesContext(p *MatrixValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixValues
}

func (*MatrixValuesContext) IsMatrixValuesContext() {}

func NewMatrixValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixValuesContext {
	var p = new(MatrixValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixValues

	return p
}

func (s *MatrixValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixValuesContext) AllMatrixDefinition() []IMatrixDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatrixDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IMatrixDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatrixDefinitionContext); ok {
			tst[i] = t.(IMatrixDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *MatrixValuesContext) MatrixDefinition(i int) IMatrixDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixDefinitionContext); ok {
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

	return t.(IMatrixDefinitionContext)
}

func (s *MatrixValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOMMA)
}

func (s *MatrixValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, i)
}

func (s *MatrixValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixValues() (localctx IMatrixValuesContext) {
	localctx = NewMatrixValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftParserRULE_matrixValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.MatrixDefinition()
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserCOMMA {
		{
			p.SetState(346)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.MatrixDefinition()
		}

		p.SetState(352)
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

// IMatrixRepeatingDefinitionContext is an interface to support dynamic dispatch.
type IMatrixRepeatingDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMatrixRepeatingDefinitionContext differentiates from other interfaces.
	IsMatrixRepeatingDefinitionContext()
}

type MatrixRepeatingDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixRepeatingDefinitionContext() *MatrixRepeatingDefinitionContext {
	var p = new(MatrixRepeatingDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixRepeatingDefinition
	return p
}

func InitEmptyMatrixRepeatingDefinitionContext(p *MatrixRepeatingDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixRepeatingDefinition
}

func (*MatrixRepeatingDefinitionContext) IsMatrixRepeatingDefinitionContext() {}

func NewMatrixRepeatingDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixRepeatingDefinitionContext {
	var p = new(MatrixRepeatingDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixRepeatingDefinition

	return p
}

func (s *MatrixRepeatingDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixRepeatingDefinitionContext) CopyAll(ctx *MatrixRepeatingDefinitionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MatrixRepeatingDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixRepeatingDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MatrixRepeatingDefinitionSingleContext struct {
	MatrixRepeatingDefinitionContext
}

func NewMatrixRepeatingDefinitionSingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixRepeatingDefinitionSingleContext {
	var p = new(MatrixRepeatingDefinitionSingleContext)

	InitEmptyMatrixRepeatingDefinitionContext(&p.MatrixRepeatingDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatrixRepeatingDefinitionContext))

	return p
}

func (s *MatrixRepeatingDefinitionSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixRepeatingDefinitionSingleContext) MatrixType() IMatrixTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixTypeContext)
}

func (s *MatrixRepeatingDefinitionSingleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *MatrixRepeatingDefinitionSingleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserID)
}

func (s *MatrixRepeatingDefinitionSingleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserID, i)
}

func (s *MatrixRepeatingDefinitionSingleContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOLON)
}

func (s *MatrixRepeatingDefinitionSingleContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, i)
}

func (s *MatrixRepeatingDefinitionSingleContext) AllExpr() []IExprContext {
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

func (s *MatrixRepeatingDefinitionSingleContext) Expr(i int) IExprContext {
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

func (s *MatrixRepeatingDefinitionSingleContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, 0)
}

func (s *MatrixRepeatingDefinitionSingleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *MatrixRepeatingDefinitionSingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixRepeatingDefinitionSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrixRepeatingDefinitionNestedContext struct {
	MatrixRepeatingDefinitionContext
}

func NewMatrixRepeatingDefinitionNestedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixRepeatingDefinitionNestedContext {
	var p = new(MatrixRepeatingDefinitionNestedContext)

	InitEmptyMatrixRepeatingDefinitionContext(&p.MatrixRepeatingDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatrixRepeatingDefinitionContext))

	return p
}

func (s *MatrixRepeatingDefinitionNestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixRepeatingDefinitionNestedContext) MatrixType() IMatrixTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixTypeContext)
}

func (s *MatrixRepeatingDefinitionNestedContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *MatrixRepeatingDefinitionNestedContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserID)
}

func (s *MatrixRepeatingDefinitionNestedContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserID, i)
}

func (s *MatrixRepeatingDefinitionNestedContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserCOLON)
}

func (s *MatrixRepeatingDefinitionNestedContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, i)
}

func (s *MatrixRepeatingDefinitionNestedContext) MatrixRepeatingDefinition() IMatrixRepeatingDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixRepeatingDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixRepeatingDefinitionContext)
}

func (s *MatrixRepeatingDefinitionNestedContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOMMA, 0)
}

func (s *MatrixRepeatingDefinitionNestedContext) Expr() IExprContext {
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

func (s *MatrixRepeatingDefinitionNestedContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *MatrixRepeatingDefinitionNestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixRepeatingDefinitionNested(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixRepeatingDefinition() (localctx IMatrixRepeatingDefinitionContext) {
	localctx = NewMatrixRepeatingDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftParserRULE_matrixRepeatingDefinition)
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMatrixRepeatingDefinitionNestedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.MatrixType()
		}
		{
			p.SetState(354)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.MatrixRepeatingDefinition()
		}
		{
			p.SetState(358)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.expr(0)
		}
		{
			p.SetState(362)
			p.Match(SwiftParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMatrixRepeatingDefinitionSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(364)
			p.MatrixType()
		}
		{
			p.SetState(365)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.expr(0)
		}
		{
			p.SetState(369)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.expr(0)
		}
		{
			p.SetState(373)
			p.Match(SwiftParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

// IMatrixAccessContext is an interface to support dynamic dispatch.
type IMatrixAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdChain() IIdChainContext
	AllLBRACKET() []antlr.TerminalNode
	LBRACKET(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRBRACKET() []antlr.TerminalNode
	RBRACKET(i int) antlr.TerminalNode

	// IsMatrixAccessContext differentiates from other interfaces.
	IsMatrixAccessContext()
}

type MatrixAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixAccessContext() *MatrixAccessContext {
	var p = new(MatrixAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixAccess
	return p
}

func InitEmptyMatrixAccessContext(p *MatrixAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixAccess
}

func (*MatrixAccessContext) IsMatrixAccessContext() {}

func NewMatrixAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixAccessContext {
	var p = new(MatrixAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixAccess

	return p
}

func (s *MatrixAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixAccessContext) IdChain() IIdChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdChainContext)
}

func (s *MatrixAccessContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserLBRACKET)
}

func (s *MatrixAccessContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACKET, i)
}

func (s *MatrixAccessContext) AllExpr() []IExprContext {
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

func (s *MatrixAccessContext) Expr(i int) IExprContext {
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

func (s *MatrixAccessContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(SwiftParserRBRACKET)
}

func (s *MatrixAccessContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACKET, i)
}

func (s *MatrixAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixAccess() (localctx IMatrixAccessContext) {
	localctx = NewMatrixAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SwiftParserRULE_matrixAccess)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.IdChain()
	}
	{
		p.SetState(378)
		p.Match(SwiftParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.expr(0)
	}
	{
		p.SetState(380)
		p.Match(SwiftParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.Match(SwiftParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(382)
				p.expr(0)
			}
			{
				p.SetState(383)
				p.Match(SwiftParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
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

// IMatrixAssignmentContext is an interface to support dynamic dispatch.
type IMatrixAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	MatrixAccess() IMatrixAccessContext
	Expr() IExprContext
	Op_ASSIGN() antlr.TerminalNode
	Op_PLUS_ASSIGN() antlr.TerminalNode
	Op_MINUS_ASSIGN() antlr.TerminalNode

	// IsMatrixAssignmentContext differentiates from other interfaces.
	IsMatrixAssignmentContext()
}

type MatrixAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyMatrixAssignmentContext() *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixAssignment
	return p
}

func InitEmptyMatrixAssignmentContext(p *MatrixAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_matrixAssignment
}

func (*MatrixAssignmentContext) IsMatrixAssignmentContext() {}

func NewMatrixAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_matrixAssignment

	return p
}

func (s *MatrixAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixAssignmentContext) GetOp() antlr.Token { return s.op }

func (s *MatrixAssignmentContext) SetOp(v antlr.Token) { s.op = v }

func (s *MatrixAssignmentContext) MatrixAccess() IMatrixAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixAccessContext)
}

func (s *MatrixAssignmentContext) Expr() IExprContext {
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

func (s *MatrixAssignmentContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *MatrixAssignmentContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_PLUS_ASSIGN, 0)
}

func (s *MatrixAssignmentContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MINUS_ASSIGN, 0)
}

func (s *MatrixAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) MatrixAssignment() (localctx IMatrixAssignmentContext) {
	localctx = NewMatrixAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SwiftParserRULE_matrixAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.MatrixAccess()
	}
	{
		p.SetState(391)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MatrixAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&439804651110400) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MatrixAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(392)
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

// IStructDeclarationContext is an interface to support dynamic dispatch.
type IStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	StructBody() IStructBodyContext
	RBRACE() antlr.TerminalNode

	// IsStructDeclarationContext differentiates from other interfaces.
	IsStructDeclarationContext()
}

type StructDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationContext() *StructDeclarationContext {
	var p = new(StructDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_structDeclaration
	return p
}

func InitEmptyStructDeclarationContext(p *StructDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_structDeclaration
}

func (*StructDeclarationContext) IsStructDeclarationContext() {}

func NewStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_structDeclaration

	return p
}

func (s *StructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationContext) Kw_STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_STRUCT, 0)
}

func (s *StructDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *StructDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserLBRACE, 0)
}

func (s *StructDeclarationContext) StructBody() IStructBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
}

func (s *StructDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftParserRBRACE, 0)
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) StructDeclaration() (localctx IStructDeclarationContext) {
	localctx = NewStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftParserRULE_structDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.Match(SwiftParserKw_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.StructBody()
	}
	{
		p.SetState(398)
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

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStructProperty() []IStructPropertyContext
	StructProperty(i int) IStructPropertyContext

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_structBody
	return p
}

func InitEmptyStructBodyContext(p *StructBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_structBody
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) AllStructProperty() []IStructPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructPropertyContext); ok {
			len++
		}
	}

	tst := make([]IStructPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructPropertyContext); ok {
			tst[i] = t.(IStructPropertyContext)
			i++
		}
	}

	return tst
}

func (s *StructBodyContext) StructProperty(i int) IStructPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructPropertyContext); ok {
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

	return t.(IStructPropertyContext)
}

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitStructBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SwiftParserRULE_structBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&368) != 0 {
		{
			p.SetState(400)
			p.StructProperty()
		}

		p.SetState(405)
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

// IStructPropertyContext is an interface to support dynamic dispatch.
type IStructPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	FunctionDeclarationStatement() IFunctionDeclarationStatementContext

	// IsStructPropertyContext differentiates from other interfaces.
	IsStructPropertyContext()
}

type StructPropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructPropertyContext() *StructPropertyContext {
	var p = new(StructPropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_structProperty
	return p
}

func InitEmptyStructPropertyContext(p *StructPropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_structProperty
}

func (*StructPropertyContext) IsStructPropertyContext() {}

func NewStructPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructPropertyContext {
	var p = new(StructPropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_structProperty

	return p
}

func (s *StructPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructPropertyContext) VariableDeclaration() IVariableDeclarationContext {
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

func (s *StructPropertyContext) FunctionDeclarationStatement() IFunctionDeclarationStatementContext {
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

func (s *StructPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitStructProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) StructProperty() (localctx IStructPropertyContext) {
	localctx = NewStructPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SwiftParserRULE_structProperty)
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserKw_LET, SwiftParserKw_VAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(406)
			p.VariableDeclaration()
		}

	case SwiftParserKw_FUNC, SwiftParserKw_MUTATING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(407)
			p.FunctionDeclarationStatement()
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

func (s *IdExprContext) IdChain() IIdChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdChainContext)
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

type VectorAccessExprContext struct {
	ExprContext
}

func NewVectorAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorAccessExprContext {
	var p = new(VectorAccessExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAccessExprContext) VectorAccess() IVectorAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorAccessContext)
}

func (s *VectorAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorAccessExpr(s)

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

type MatrixAccessExprContext struct {
	ExprContext
}

func NewMatrixAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrixAccessExprContext {
	var p = new(MatrixAccessExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MatrixAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixAccessExprContext) MatrixAccess() IMatrixAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixAccessContext)
}

func (s *MatrixAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixAccessExpr(s)

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
	_startState := 74
	p.EnterRecursionRule(localctx, 74, SwiftParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(411)
			p.FunctionCall()
		}

	case 2:
		localctx = NewVectorAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(412)
			p.VectorAccess()
		}

	case 3:
		localctx = NewMatrixAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(413)
			p.MatrixAccess()
		}

	case 4:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.Match(SwiftParserOp_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.expr(20)
		}

	case 5:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(416)
			p.Match(SwiftParserOp_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)

			var _x = p.expr(19)

			localctx.(*NotExprContext).right = _x
		}

	case 6:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(418)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.expr(0)
		}
		{
			p.SetState(420)
			p.Match(SwiftParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(422)
			p.Match(SwiftParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(423)
			p.IdChain()
		}

	case 9:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(424)
			p.Match(SwiftParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(425)
			p.Match(SwiftParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(426)
			p.Match(SwiftParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(427)
			p.Match(SwiftParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(428)
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
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(464)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(432)

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
					p.SetState(433)

					var _x = p.expr(19)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(435)

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
					p.SetState(436)

					var _x = p.expr(18)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 3:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(438)

					var _m = p.Match(SwiftParserOp_MOD)

					localctx.(*ArithmeticExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(439)

					var _x = p.expr(17)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 4:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(440)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(441)

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
					p.SetState(442)

					var _x = p.expr(16)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 5:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(443)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(444)

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
					p.SetState(445)

					var _x = p.expr(15)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 6:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(447)

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
					p.SetState(448)

					var _x = p.expr(14)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(450)
					p.Match(SwiftParserOp_TERNARY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(451)

					var _x = p.expr(0)

					localctx.(*TernaryExprContext).cTrue = _x
				}
				{
					p.SetState(452)
					p.Match(SwiftParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(453)

					var _x = p.expr(13)

					localctx.(*TernaryExprContext).cFalse = _x
				}

			case 8:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(455)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(456)

					var _m = p.Match(SwiftParserOp_AND)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(457)

					var _x = p.expr(12)

					localctx.(*LogicalExprContext).right = _x
				}

			case 9:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(458)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(459)

					var _m = p.Match(SwiftParserOp_OR)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(460)

					var _x = p.expr(11)

					localctx.(*LogicalExprContext).right = _x
				}

			case 10:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RangeExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(461)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(462)
					p.Match(SwiftParserKw_RANGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(463)

					var _x = p.expr(10)

					localctx.(*RangeExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 76, SwiftParserRULE_controlFlowStatement)
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserKw_BREAK:
		localctx = NewControlBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(469)
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
			p.SetState(470)
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
			p.SetState(471)
			p.Match(SwiftParserKw_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(473)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(472)
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
	case 37:
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
