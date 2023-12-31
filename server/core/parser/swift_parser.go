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

type Swift struct {
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
		"'&'", "'if'", "'else'", "'switch'", "'case'", "'default'", "'for'",
		"'while'", "'break'", "'continue'", "'return'", "'do'", "'repeat'",
		"'in'", "'guard'", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'",
		"'...'", "'inout'", "", "", "", "", "", "'nil'", "", "'->'", "'=='",
		"'!='", "'<'", "'>'", "'<='", "'>='", "'='", "'*='", "'/='", "'+='",
		"'-='", "'%='", "'*'", "'/'", "'+'", "'-'", "'%'", "'&&'", "'||'", "'!'",
		"'?'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'\\'", "','", "';'",
		"':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_MUTATING", "Kw_AMPER", "Kw_IF", "Kw_ELSE", "Kw_SWITCH",
		"Kw_CASE", "Kw_DEFAULT", "Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE",
		"Kw_RETURN", "Kw_DO", "Kw_REPEAT", "Kw_IN", "Kw_GUARD", "Kw_INT", "Kw_FLOAT",
		"Kw_BOOL", "Kw_STRING", "Kw_CHAR", "Kw_RANGE", "Kw_INOUT", "INT", "FLOAT",
		"BOOL", "STRING", "CHAR", "NIL", "ID", "Op_ARROW", "Op_EQ", "Op_NEQ",
		"Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN",
		"Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN", "Op_MOD_ASSIGN", "Op_MUL", "Op_DIV",
		"Op_PLUS", "Op_MINUS", "Op_MOD", "Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY",
		"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH",
		"COMMA", "SEMICOLON", "COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"idChain", "program", "block", "statement", "variableType", "variableDeclaration",
		"functionDeclarationStatement", "functionParameters", "functionParameter",
		"functionParameterCompound", "functionReturnType", "functionCall", "functionCallArguments",
		"variableAssignment", "variableAssignmentObject", "ifStatement", "ifTail",
		"elseStatement", "whileStatement", "switchStatement", "switchCase",
		"switchDefault", "forStatement", "guardStatement", "vectorDeclaration",
		"vectorDefinition", "vectorValues", "objectChain", "vectorAccess", "vectorAssignment",
		"matrixDeclaration", "matrixType", "matrixDefinition", "matrixValues",
		"matrixRepeatingDefinition", "matrixAccess", "matrixAssignment", "structDeclaration",
		"structBody", "structProperty", "expr", "controlFlowStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 540, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 1, 0, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 1, 1, 1, 1, 1,
		1, 2, 5, 2, 97, 8, 2, 10, 2, 12, 2, 100, 9, 2, 1, 3, 1, 3, 3, 3, 104, 8,
		3, 1, 3, 1, 3, 3, 3, 108, 8, 3, 1, 3, 1, 3, 3, 3, 112, 8, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 123, 8, 3, 1, 3, 1, 3,
		3, 3, 127, 8, 3, 1, 3, 1, 3, 3, 3, 131, 8, 3, 1, 3, 1, 3, 3, 3, 135, 8,
		3, 1, 3, 1, 3, 3, 3, 139, 8, 3, 1, 3, 3, 3, 142, 8, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 162, 8, 5, 3, 5, 164, 8, 5, 1, 6, 3, 6, 167, 8,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 173, 8, 6, 1, 6, 1, 6, 3, 6, 177, 8, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 186, 8, 7, 10, 7, 12, 7,
		189, 9, 7, 1, 8, 3, 8, 192, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 197, 8, 8, 1,
		8, 1, 8, 3, 8, 201, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 3, 9, 211, 8, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 218, 8, 11,
		1, 11, 1, 11, 3, 11, 222, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5,
		12, 229, 8, 12, 10, 12, 12, 12, 232, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 5, 12, 241, 8, 12, 10, 12, 12, 12, 244, 9, 12, 3,
		12, 246, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 5, 15, 259, 8, 15, 10, 15, 12, 15, 262, 9, 15, 1,
		15, 3, 15, 265, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 5, 19, 288, 8, 19, 10, 19, 12, 19, 291, 9, 19,
		1, 19, 3, 19, 294, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 303, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 309, 8, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 343, 8, 24, 1, 25, 1, 25, 3, 25, 347, 8, 25, 1, 25, 1, 25, 3, 25,
		351, 8, 25, 1, 26, 1, 26, 1, 26, 5, 26, 356, 8, 26, 10, 26, 12, 26, 359,
		9, 26, 1, 27, 1, 27, 1, 27, 4, 27, 364, 8, 27, 11, 27, 12, 27, 365, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 384, 8, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 394, 8, 31, 1, 32, 1, 32,
		3, 32, 398, 8, 32, 1, 32, 1, 32, 3, 32, 402, 8, 32, 1, 33, 1, 33, 1, 33,
		5, 33, 407, 8, 33, 10, 33, 12, 33, 410, 9, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 434, 8,
		34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 444,
		8, 35, 10, 35, 12, 35, 447, 9, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 5, 38, 460, 8, 38, 10, 38, 12, 38,
		463, 9, 38, 1, 39, 1, 39, 3, 39, 467, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 3, 40, 484, 8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40,
		492, 8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 5, 40, 527, 8, 40, 10, 40, 12, 40, 530, 9, 40,
		1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 536, 8, 41, 3, 41, 538, 8, 41, 1, 41,
		0, 1, 80, 42, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 0, 8, 2, 0, 24, 28, 37, 37, 1, 0, 4, 5,
		2, 0, 45, 45, 48, 49, 1, 0, 51, 52, 1, 0, 53, 54, 2, 0, 42, 42, 44, 44,
		2, 0, 41, 41, 43, 43, 1, 0, 39, 40, 585, 0, 84, 1, 0, 0, 0, 2, 92, 1, 0,
		0, 0, 4, 98, 1, 0, 0, 0, 6, 141, 1, 0, 0, 0, 8, 143, 1, 0, 0, 0, 10, 163,
		1, 0, 0, 0, 12, 166, 1, 0, 0, 0, 14, 182, 1, 0, 0, 0, 16, 191, 1, 0, 0,
		0, 18, 210, 1, 0, 0, 0, 20, 212, 1, 0, 0, 0, 22, 217, 1, 0, 0, 0, 24, 245,
		1, 0, 0, 0, 26, 247, 1, 0, 0, 0, 28, 251, 1, 0, 0, 0, 30, 255, 1, 0, 0,
		0, 32, 266, 1, 0, 0, 0, 34, 272, 1, 0, 0, 0, 36, 277, 1, 0, 0, 0, 38, 283,
		1, 0, 0, 0, 40, 297, 1, 0, 0, 0, 42, 304, 1, 0, 0, 0, 44, 310, 1, 0, 0,
		0, 46, 318, 1, 0, 0, 0, 48, 342, 1, 0, 0, 0, 50, 350, 1, 0, 0, 0, 52, 352,
		1, 0, 0, 0, 54, 360, 1, 0, 0, 0, 56, 367, 1, 0, 0, 0, 58, 372, 1, 0, 0,
		0, 60, 376, 1, 0, 0, 0, 62, 393, 1, 0, 0, 0, 64, 401, 1, 0, 0, 0, 66, 403,
		1, 0, 0, 0, 68, 433, 1, 0, 0, 0, 70, 435, 1, 0, 0, 0, 72, 448, 1, 0, 0,
		0, 74, 452, 1, 0, 0, 0, 76, 461, 1, 0, 0, 0, 78, 466, 1, 0, 0, 0, 80, 491,
		1, 0, 0, 0, 82, 537, 1, 0, 0, 0, 84, 89, 5, 37, 0, 0, 85, 86, 5, 70, 0,
		0, 86, 88, 5, 37, 0, 0, 87, 85, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87,
		1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 1, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0,
		92, 93, 3, 4, 2, 0, 93, 94, 5, 0, 0, 1, 94, 3, 1, 0, 0, 0, 95, 97, 3, 6,
		3, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98,
		99, 1, 0, 0, 0, 99, 5, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 3, 26,
		13, 0, 102, 104, 5, 68, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 142, 1, 0, 0, 0, 105, 107, 3, 10, 5, 0, 106, 108, 5, 68, 0, 0,
		107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 142, 1, 0, 0, 0, 109,
		111, 3, 28, 14, 0, 110, 112, 5, 68, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 142, 1, 0, 0, 0, 113, 142, 3, 30, 15, 0, 114, 142, 3,
		36, 18, 0, 115, 142, 3, 38, 19, 0, 116, 142, 3, 44, 22, 0, 117, 142, 3,
		46, 23, 0, 118, 142, 3, 82, 41, 0, 119, 142, 3, 12, 6, 0, 120, 122, 3,
		22, 11, 0, 121, 123, 5, 68, 0, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0,
		0, 0, 123, 142, 1, 0, 0, 0, 124, 126, 3, 48, 24, 0, 125, 127, 5, 68, 0,
		0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 142, 1, 0, 0, 0, 128,
		130, 3, 58, 29, 0, 129, 131, 5, 68, 0, 0, 130, 129, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 142, 1, 0, 0, 0, 132, 134, 3, 60, 30, 0, 133, 135, 5,
		68, 0, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 142, 1, 0, 0,
		0, 136, 138, 3, 72, 36, 0, 137, 139, 5, 68, 0, 0, 138, 137, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 142, 3, 74, 37, 0, 141,
		101, 1, 0, 0, 0, 141, 105, 1, 0, 0, 0, 141, 109, 1, 0, 0, 0, 141, 113,
		1, 0, 0, 0, 141, 114, 1, 0, 0, 0, 141, 115, 1, 0, 0, 0, 141, 116, 1, 0,
		0, 0, 141, 117, 1, 0, 0, 0, 141, 118, 1, 0, 0, 0, 141, 119, 1, 0, 0, 0,
		141, 120, 1, 0, 0, 0, 141, 124, 1, 0, 0, 0, 141, 128, 1, 0, 0, 0, 141,
		132, 1, 0, 0, 0, 141, 136, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 7, 1,
		0, 0, 0, 143, 144, 7, 0, 0, 0, 144, 9, 1, 0, 0, 0, 145, 146, 7, 1, 0, 0,
		146, 147, 5, 37, 0, 0, 147, 148, 5, 69, 0, 0, 148, 149, 3, 8, 4, 0, 149,
		150, 5, 45, 0, 0, 150, 151, 3, 80, 40, 0, 151, 164, 1, 0, 0, 0, 152, 153,
		7, 1, 0, 0, 153, 154, 5, 37, 0, 0, 154, 155, 5, 45, 0, 0, 155, 164, 3,
		80, 40, 0, 156, 157, 7, 1, 0, 0, 157, 158, 5, 37, 0, 0, 158, 159, 5, 69,
		0, 0, 159, 161, 3, 8, 4, 0, 160, 162, 5, 59, 0, 0, 161, 160, 1, 0, 0, 0,
		161, 162, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 145, 1, 0, 0, 0, 163,
		152, 1, 0, 0, 0, 163, 156, 1, 0, 0, 0, 164, 11, 1, 0, 0, 0, 165, 167, 5,
		8, 0, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0,
		0, 168, 169, 5, 6, 0, 0, 169, 170, 5, 37, 0, 0, 170, 172, 5, 60, 0, 0,
		171, 173, 3, 14, 7, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173,
		174, 1, 0, 0, 0, 174, 176, 5, 61, 0, 0, 175, 177, 3, 20, 10, 0, 176, 175,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 5, 62,
		0, 0, 179, 180, 3, 4, 2, 0, 180, 181, 5, 63, 0, 0, 181, 13, 1, 0, 0, 0,
		182, 187, 3, 16, 8, 0, 183, 184, 5, 67, 0, 0, 184, 186, 3, 16, 8, 0, 185,
		183, 1, 0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188,
		1, 0, 0, 0, 188, 15, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 192, 5, 37,
		0, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0,
		193, 194, 5, 37, 0, 0, 194, 196, 5, 69, 0, 0, 195, 197, 5, 30, 0, 0, 196,
		195, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 201,
		3, 8, 4, 0, 199, 201, 3, 62, 31, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1,
		0, 0, 0, 201, 17, 1, 0, 0, 0, 202, 203, 5, 64, 0, 0, 203, 204, 3, 18, 9,
		0, 204, 205, 5, 65, 0, 0, 205, 211, 1, 0, 0, 0, 206, 207, 5, 64, 0, 0,
		207, 208, 3, 8, 4, 0, 208, 209, 5, 65, 0, 0, 209, 211, 1, 0, 0, 0, 210,
		202, 1, 0, 0, 0, 210, 206, 1, 0, 0, 0, 211, 19, 1, 0, 0, 0, 212, 213, 5,
		38, 0, 0, 213, 214, 3, 8, 4, 0, 214, 21, 1, 0, 0, 0, 215, 218, 3, 0, 0,
		0, 216, 218, 3, 8, 4, 0, 217, 215, 1, 0, 0, 0, 217, 216, 1, 0, 0, 0, 218,
		219, 1, 0, 0, 0, 219, 221, 5, 60, 0, 0, 220, 222, 3, 24, 12, 0, 221, 220,
		1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 5, 61,
		0, 0, 224, 23, 1, 0, 0, 0, 225, 230, 3, 80, 40, 0, 226, 227, 5, 67, 0,
		0, 227, 229, 3, 80, 40, 0, 228, 226, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0,
		230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 246, 1, 0, 0, 0, 232,
		230, 1, 0, 0, 0, 233, 234, 5, 37, 0, 0, 234, 235, 5, 69, 0, 0, 235, 242,
		3, 80, 40, 0, 236, 237, 5, 67, 0, 0, 237, 238, 5, 37, 0, 0, 238, 239, 5,
		69, 0, 0, 239, 241, 3, 80, 40, 0, 240, 236, 1, 0, 0, 0, 241, 244, 1, 0,
		0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0,
		244, 242, 1, 0, 0, 0, 245, 225, 1, 0, 0, 0, 245, 233, 1, 0, 0, 0, 246,
		25, 1, 0, 0, 0, 247, 248, 3, 0, 0, 0, 248, 249, 7, 2, 0, 0, 249, 250, 3,
		80, 40, 0, 250, 27, 1, 0, 0, 0, 251, 252, 3, 54, 27, 0, 252, 253, 7, 2,
		0, 0, 253, 254, 3, 80, 40, 0, 254, 29, 1, 0, 0, 0, 255, 260, 3, 32, 16,
		0, 256, 257, 5, 11, 0, 0, 257, 259, 3, 32, 16, 0, 258, 256, 1, 0, 0, 0,
		259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		264, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 265, 3, 34, 17, 0, 264, 263,
		1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 31, 1, 0, 0, 0, 266, 267, 5, 10,
		0, 0, 267, 268, 3, 80, 40, 0, 268, 269, 5, 62, 0, 0, 269, 270, 3, 4, 2,
		0, 270, 271, 5, 63, 0, 0, 271, 33, 1, 0, 0, 0, 272, 273, 5, 11, 0, 0, 273,
		274, 5, 62, 0, 0, 274, 275, 3, 4, 2, 0, 275, 276, 5, 63, 0, 0, 276, 35,
		1, 0, 0, 0, 277, 278, 5, 16, 0, 0, 278, 279, 3, 80, 40, 0, 279, 280, 5,
		62, 0, 0, 280, 281, 3, 4, 2, 0, 281, 282, 5, 63, 0, 0, 282, 37, 1, 0, 0,
		0, 283, 284, 5, 12, 0, 0, 284, 285, 3, 80, 40, 0, 285, 289, 5, 62, 0, 0,
		286, 288, 3, 40, 20, 0, 287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289,
		287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289,
		1, 0, 0, 0, 292, 294, 3, 42, 21, 0, 293, 292, 1, 0, 0, 0, 293, 294, 1,
		0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 5, 63, 0, 0, 296, 39, 1, 0, 0,
		0, 297, 298, 5, 13, 0, 0, 298, 299, 3, 80, 40, 0, 299, 300, 5, 69, 0, 0,
		300, 302, 3, 4, 2, 0, 301, 303, 5, 17, 0, 0, 302, 301, 1, 0, 0, 0, 302,
		303, 1, 0, 0, 0, 303, 41, 1, 0, 0, 0, 304, 305, 5, 14, 0, 0, 305, 306,
		5, 69, 0, 0, 306, 308, 3, 4, 2, 0, 307, 309, 5, 17, 0, 0, 308, 307, 1,
		0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 43, 1, 0, 0, 0, 310, 311, 5, 15, 0,
		0, 311, 312, 5, 37, 0, 0, 312, 313, 5, 22, 0, 0, 313, 314, 3, 80, 40, 0,
		314, 315, 5, 62, 0, 0, 315, 316, 3, 4, 2, 0, 316, 317, 5, 63, 0, 0, 317,
		45, 1, 0, 0, 0, 318, 319, 5, 23, 0, 0, 319, 320, 3, 80, 40, 0, 320, 321,
		5, 11, 0, 0, 321, 322, 5, 62, 0, 0, 322, 323, 3, 4, 2, 0, 323, 324, 5,
		63, 0, 0, 324, 47, 1, 0, 0, 0, 325, 326, 7, 1, 0, 0, 326, 327, 5, 37, 0,
		0, 327, 328, 5, 69, 0, 0, 328, 329, 5, 64, 0, 0, 329, 330, 3, 8, 4, 0,
		330, 331, 5, 65, 0, 0, 331, 332, 5, 45, 0, 0, 332, 333, 3, 50, 25, 0, 333,
		343, 1, 0, 0, 0, 334, 335, 7, 1, 0, 0, 335, 336, 5, 37, 0, 0, 336, 337,
		5, 45, 0, 0, 337, 338, 5, 64, 0, 0, 338, 339, 5, 37, 0, 0, 339, 340, 5,
		65, 0, 0, 340, 341, 5, 60, 0, 0, 341, 343, 5, 61, 0, 0, 342, 325, 1, 0,
		0, 0, 342, 334, 1, 0, 0, 0, 343, 49, 1, 0, 0, 0, 344, 346, 5, 64, 0, 0,
		345, 347, 3, 52, 26, 0, 346, 345, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347,
		348, 1, 0, 0, 0, 348, 351, 5, 65, 0, 0, 349, 351, 3, 80, 40, 0, 350, 344,
		1, 0, 0, 0, 350, 349, 1, 0, 0, 0, 351, 51, 1, 0, 0, 0, 352, 357, 3, 80,
		40, 0, 353, 354, 5, 67, 0, 0, 354, 356, 3, 80, 40, 0, 355, 353, 1, 0, 0,
		0, 356, 359, 1, 0, 0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358,
		53, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 360, 363, 3, 56, 28, 0, 361, 362,
		5, 70, 0, 0, 362, 364, 5, 37, 0, 0, 363, 361, 1, 0, 0, 0, 364, 365, 1,
		0, 0, 0, 365, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 55, 1, 0, 0,
		0, 367, 368, 3, 0, 0, 0, 368, 369, 5, 64, 0, 0, 369, 370, 3, 80, 40, 0,
		370, 371, 5, 65, 0, 0, 371, 57, 1, 0, 0, 0, 372, 373, 3, 56, 28, 0, 373,
		374, 7, 2, 0, 0, 374, 375, 3, 80, 40, 0, 375, 59, 1, 0, 0, 0, 376, 377,
		7, 1, 0, 0, 377, 378, 3, 0, 0, 0, 378, 379, 5, 69, 0, 0, 379, 380, 3, 62,
		31, 0, 380, 383, 5, 45, 0, 0, 381, 384, 3, 68, 34, 0, 382, 384, 3, 64,
		32, 0, 383, 381, 1, 0, 0, 0, 383, 382, 1, 0, 0, 0, 384, 61, 1, 0, 0, 0,
		385, 386, 5, 64, 0, 0, 386, 387, 3, 62, 31, 0, 387, 388, 5, 65, 0, 0, 388,
		394, 1, 0, 0, 0, 389, 390, 5, 64, 0, 0, 390, 391, 3, 8, 4, 0, 391, 392,
		5, 65, 0, 0, 392, 394, 1, 0, 0, 0, 393, 385, 1, 0, 0, 0, 393, 389, 1, 0,
		0, 0, 394, 63, 1, 0, 0, 0, 395, 397, 5, 64, 0, 0, 396, 398, 3, 66, 33,
		0, 397, 396, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399,
		402, 5, 65, 0, 0, 400, 402, 3, 80, 40, 0, 401, 395, 1, 0, 0, 0, 401, 400,
		1, 0, 0, 0, 402, 65, 1, 0, 0, 0, 403, 408, 3, 64, 32, 0, 404, 405, 5, 67,
		0, 0, 405, 407, 3, 64, 32, 0, 406, 404, 1, 0, 0, 0, 407, 410, 1, 0, 0,
		0, 408, 406, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 67, 1, 0, 0, 0, 410,
		408, 1, 0, 0, 0, 411, 412, 3, 62, 31, 0, 412, 413, 5, 60, 0, 0, 413, 414,
		5, 37, 0, 0, 414, 415, 5, 69, 0, 0, 415, 416, 3, 68, 34, 0, 416, 417, 5,
		67, 0, 0, 417, 418, 5, 37, 0, 0, 418, 419, 5, 69, 0, 0, 419, 420, 3, 80,
		40, 0, 420, 421, 5, 61, 0, 0, 421, 434, 1, 0, 0, 0, 422, 423, 3, 62, 31,
		0, 423, 424, 5, 60, 0, 0, 424, 425, 5, 37, 0, 0, 425, 426, 5, 69, 0, 0,
		426, 427, 3, 80, 40, 0, 427, 428, 5, 67, 0, 0, 428, 429, 5, 37, 0, 0, 429,
		430, 5, 69, 0, 0, 430, 431, 3, 80, 40, 0, 431, 432, 5, 61, 0, 0, 432, 434,
		1, 0, 0, 0, 433, 411, 1, 0, 0, 0, 433, 422, 1, 0, 0, 0, 434, 69, 1, 0,
		0, 0, 435, 436, 3, 0, 0, 0, 436, 437, 5, 64, 0, 0, 437, 438, 3, 80, 40,
		0, 438, 445, 5, 65, 0, 0, 439, 440, 5, 64, 0, 0, 440, 441, 3, 80, 40, 0,
		441, 442, 5, 65, 0, 0, 442, 444, 1, 0, 0, 0, 443, 439, 1, 0, 0, 0, 444,
		447, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 71, 1,
		0, 0, 0, 447, 445, 1, 0, 0, 0, 448, 449, 3, 70, 35, 0, 449, 450, 7, 2,
		0, 0, 450, 451, 3, 80, 40, 0, 451, 73, 1, 0, 0, 0, 452, 453, 5, 7, 0, 0,
		453, 454, 5, 37, 0, 0, 454, 455, 5, 62, 0, 0, 455, 456, 3, 76, 38, 0, 456,
		457, 5, 63, 0, 0, 457, 75, 1, 0, 0, 0, 458, 460, 3, 78, 39, 0, 459, 458,
		1, 0, 0, 0, 460, 463, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 462, 1, 0,
		0, 0, 462, 77, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0, 464, 467, 3, 10, 5, 0,
		465, 467, 3, 12, 6, 0, 466, 464, 1, 0, 0, 0, 466, 465, 1, 0, 0, 0, 467,
		79, 1, 0, 0, 0, 468, 469, 6, 40, -1, 0, 469, 492, 3, 54, 27, 0, 470, 492,
		3, 22, 11, 0, 471, 492, 3, 56, 28, 0, 472, 492, 3, 70, 35, 0, 473, 474,
		5, 54, 0, 0, 474, 492, 3, 80, 40, 20, 475, 476, 5, 58, 0, 0, 476, 492,
		3, 80, 40, 19, 477, 478, 5, 60, 0, 0, 478, 479, 3, 80, 40, 0, 479, 480,
		5, 61, 0, 0, 480, 492, 1, 0, 0, 0, 481, 492, 5, 31, 0, 0, 482, 484, 5,
		9, 0, 0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 485, 1, 0, 0,
		0, 485, 492, 3, 0, 0, 0, 486, 492, 5, 32, 0, 0, 487, 492, 5, 34, 0, 0,
		488, 492, 5, 36, 0, 0, 489, 492, 5, 35, 0, 0, 490, 492, 5, 33, 0, 0, 491,
		468, 1, 0, 0, 0, 491, 470, 1, 0, 0, 0, 491, 471, 1, 0, 0, 0, 491, 472,
		1, 0, 0, 0, 491, 473, 1, 0, 0, 0, 491, 475, 1, 0, 0, 0, 491, 477, 1, 0,
		0, 0, 491, 481, 1, 0, 0, 0, 491, 483, 1, 0, 0, 0, 491, 486, 1, 0, 0, 0,
		491, 487, 1, 0, 0, 0, 491, 488, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0, 491,
		490, 1, 0, 0, 0, 492, 528, 1, 0, 0, 0, 493, 494, 10, 18, 0, 0, 494, 495,
		7, 3, 0, 0, 495, 527, 3, 80, 40, 19, 496, 497, 10, 17, 0, 0, 497, 498,
		7, 4, 0, 0, 498, 527, 3, 80, 40, 18, 499, 500, 10, 16, 0, 0, 500, 501,
		5, 55, 0, 0, 501, 527, 3, 80, 40, 17, 502, 503, 10, 15, 0, 0, 503, 504,
		7, 5, 0, 0, 504, 527, 3, 80, 40, 16, 505, 506, 10, 14, 0, 0, 506, 507,
		7, 6, 0, 0, 507, 527, 3, 80, 40, 15, 508, 509, 10, 13, 0, 0, 509, 510,
		7, 7, 0, 0, 510, 527, 3, 80, 40, 14, 511, 512, 10, 12, 0, 0, 512, 513,
		5, 59, 0, 0, 513, 514, 3, 80, 40, 0, 514, 515, 5, 69, 0, 0, 515, 516, 3,
		80, 40, 13, 516, 527, 1, 0, 0, 0, 517, 518, 10, 11, 0, 0, 518, 519, 5,
		56, 0, 0, 519, 527, 3, 80, 40, 12, 520, 521, 10, 10, 0, 0, 521, 522, 5,
		57, 0, 0, 522, 527, 3, 80, 40, 11, 523, 524, 10, 9, 0, 0, 524, 525, 5,
		29, 0, 0, 525, 527, 3, 80, 40, 10, 526, 493, 1, 0, 0, 0, 526, 496, 1, 0,
		0, 0, 526, 499, 1, 0, 0, 0, 526, 502, 1, 0, 0, 0, 526, 505, 1, 0, 0, 0,
		526, 508, 1, 0, 0, 0, 526, 511, 1, 0, 0, 0, 526, 517, 1, 0, 0, 0, 526,
		520, 1, 0, 0, 0, 526, 523, 1, 0, 0, 0, 527, 530, 1, 0, 0, 0, 528, 526,
		1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 81, 1, 0, 0, 0, 530, 528, 1, 0,
		0, 0, 531, 538, 5, 17, 0, 0, 532, 538, 5, 18, 0, 0, 533, 535, 5, 19, 0,
		0, 534, 536, 3, 80, 40, 0, 535, 534, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0,
		536, 538, 1, 0, 0, 0, 537, 531, 1, 0, 0, 0, 537, 532, 1, 0, 0, 0, 537,
		533, 1, 0, 0, 0, 538, 83, 1, 0, 0, 0, 52, 89, 98, 103, 107, 111, 122, 126,
		130, 134, 138, 141, 161, 163, 166, 172, 176, 187, 191, 196, 200, 210, 217,
		221, 230, 242, 245, 260, 264, 289, 293, 302, 308, 342, 346, 350, 357, 365,
		383, 393, 397, 401, 408, 433, 445, 461, 466, 483, 491, 526, 528, 535, 537,
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

// SwiftInit initializes any static state used to implement Swift. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwift(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftInit() {
	staticData := &SwiftParserStaticData
	staticData.once.Do(swiftParserInit)
}

// NewSwift produces a new parser instance for the optional input antlr.TokenStream.
func NewSwift(input antlr.TokenStream) *Swift {
	SwiftInit()
	this := new(Swift)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Swift.g4"

	return this
}

// Swift tokens.
const (
	SwiftEOF             = antlr.TokenEOF
	SwiftWHITESPACE      = 1
	SwiftCOMMENT         = 2
	SwiftBLOCK_COMMENT   = 3
	SwiftKw_LET          = 4
	SwiftKw_VAR          = 5
	SwiftKw_FUNC         = 6
	SwiftKw_STRUCT       = 7
	SwiftKw_MUTATING     = 8
	SwiftKw_AMPER        = 9
	SwiftKw_IF           = 10
	SwiftKw_ELSE         = 11
	SwiftKw_SWITCH       = 12
	SwiftKw_CASE         = 13
	SwiftKw_DEFAULT      = 14
	SwiftKw_FOR          = 15
	SwiftKw_WHILE        = 16
	SwiftKw_BREAK        = 17
	SwiftKw_CONTINUE     = 18
	SwiftKw_RETURN       = 19
	SwiftKw_DO           = 20
	SwiftKw_REPEAT       = 21
	SwiftKw_IN           = 22
	SwiftKw_GUARD        = 23
	SwiftKw_INT          = 24
	SwiftKw_FLOAT        = 25
	SwiftKw_BOOL         = 26
	SwiftKw_STRING       = 27
	SwiftKw_CHAR         = 28
	SwiftKw_RANGE        = 29
	SwiftKw_INOUT        = 30
	SwiftINT             = 31
	SwiftFLOAT           = 32
	SwiftBOOL            = 33
	SwiftSTRING          = 34
	SwiftCHAR            = 35
	SwiftNIL             = 36
	SwiftID              = 37
	SwiftOp_ARROW        = 38
	SwiftOp_EQ           = 39
	SwiftOp_NEQ          = 40
	SwiftOp_LT           = 41
	SwiftOp_GT           = 42
	SwiftOp_LE           = 43
	SwiftOp_GE           = 44
	SwiftOp_ASSIGN       = 45
	SwiftOp_MUL_ASSIGN   = 46
	SwiftOp_DIV_ASSIGN   = 47
	SwiftOp_PLUS_ASSIGN  = 48
	SwiftOp_MINUS_ASSIGN = 49
	SwiftOp_MOD_ASSIGN   = 50
	SwiftOp_MUL          = 51
	SwiftOp_DIV          = 52
	SwiftOp_PLUS         = 53
	SwiftOp_MINUS        = 54
	SwiftOp_MOD          = 55
	SwiftOp_AND          = 56
	SwiftOp_OR           = 57
	SwiftOp_NOT          = 58
	SwiftOp_TERNARY      = 59
	SwiftLPAREN          = 60
	SwiftRPAREN          = 61
	SwiftLBRACE          = 62
	SwiftRBRACE          = 63
	SwiftLBRACKET        = 64
	SwiftRBRACKET        = 65
	SwiftBACKSLASH       = 66
	SwiftCOMMA           = 67
	SwiftSEMICOLON       = 68
	SwiftCOLON           = 69
	SwiftDOT             = 70
)

// Swift rules.
const (
	SwiftRULE_idChain                      = 0
	SwiftRULE_program                      = 1
	SwiftRULE_block                        = 2
	SwiftRULE_statement                    = 3
	SwiftRULE_variableType                 = 4
	SwiftRULE_variableDeclaration          = 5
	SwiftRULE_functionDeclarationStatement = 6
	SwiftRULE_functionParameters           = 7
	SwiftRULE_functionParameter            = 8
	SwiftRULE_functionParameterCompound    = 9
	SwiftRULE_functionReturnType           = 10
	SwiftRULE_functionCall                 = 11
	SwiftRULE_functionCallArguments        = 12
	SwiftRULE_variableAssignment           = 13
	SwiftRULE_variableAssignmentObject     = 14
	SwiftRULE_ifStatement                  = 15
	SwiftRULE_ifTail                       = 16
	SwiftRULE_elseStatement                = 17
	SwiftRULE_whileStatement               = 18
	SwiftRULE_switchStatement              = 19
	SwiftRULE_switchCase                   = 20
	SwiftRULE_switchDefault                = 21
	SwiftRULE_forStatement                 = 22
	SwiftRULE_guardStatement               = 23
	SwiftRULE_vectorDeclaration            = 24
	SwiftRULE_vectorDefinition             = 25
	SwiftRULE_vectorValues                 = 26
	SwiftRULE_objectChain                  = 27
	SwiftRULE_vectorAccess                 = 28
	SwiftRULE_vectorAssignment             = 29
	SwiftRULE_matrixDeclaration            = 30
	SwiftRULE_matrixType                   = 31
	SwiftRULE_matrixDefinition             = 32
	SwiftRULE_matrixValues                 = 33
	SwiftRULE_matrixRepeatingDefinition    = 34
	SwiftRULE_matrixAccess                 = 35
	SwiftRULE_matrixAssignment             = 36
	SwiftRULE_structDeclaration            = 37
	SwiftRULE_structBody                   = 38
	SwiftRULE_structProperty               = 39
	SwiftRULE_expr                         = 40
	SwiftRULE_controlFlowStatement         = 41
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
	p.RuleIndex = SwiftRULE_idChain
	return p
}

func InitEmptyIdChainContext(p *IdChainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_idChain
}

func (*IdChainContext) IsIdChainContext() {}

func NewIdChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdChainContext {
	var p = new(IdChainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_idChain

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
	return s.GetTokens(SwiftID)
}

func (s *IDChainContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *IDChainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SwiftDOT)
}

func (s *IDChainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SwiftDOT, i)
}

func (s *IDChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIDChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) IdChain() (localctx IIdChainContext) {
	localctx = NewIdChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftRULE_idChain)
	var _alt int

	localctx = NewIDChainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(SwiftID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(89)
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
				p.SetState(85)
				p.Match(SwiftDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(86)
				p.Match(SwiftID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(91)
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
	p.RuleIndex = SwiftRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_program

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
	return s.GetToken(SwiftEOF, 0)
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

func (p *Swift) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Block()
	}
	{
		p.SetState(93)
		p.Match(SwiftEOF)
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
	p.RuleIndex = SwiftRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_block

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

func (p *Swift) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
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
				p.SetState(95)
				p.Statement()
			}

		}
		p.SetState(100)
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
	SEMICOLON() antlr.TerminalNode
	VariableDeclaration() IVariableDeclarationContext
	VariableAssignmentObject() IVariableAssignmentObjectContext
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
	p.RuleIndex = SwiftRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_statement

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

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SwiftSEMICOLON, 0)
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

func (s *StatementContext) VariableAssignmentObject() IVariableAssignmentObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableAssignmentObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableAssignmentObjectContext)
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

func (p *Swift) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftRULE_statement)
	var _la int

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.VariableAssignment()
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(102)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.VariableDeclaration()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(106)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.VariableAssignmentObject()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(110)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.SwitchStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(116)
			p.ForStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(117)
			p.GuardStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(118)
			p.ControlFlowStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(119)
			p.FunctionDeclarationStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(120)
			p.FunctionCall()
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(121)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(124)
			p.VectorDeclaration()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(125)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(128)
			p.VectorAssignment()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(129)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(132)
			p.MatrixDeclaration()
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(133)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(136)
			p.MatrixAssignment()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftSEMICOLON {
			{
				p.SetState(137)
				p.Match(SwiftSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(140)
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
	ID() antlr.TerminalNode

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
	p.RuleIndex = SwiftRULE_variableType
	return p
}

func InitEmptyVariableTypeContext(p *VariableTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_variableType
}

func (*VariableTypeContext) IsVariableTypeContext() {}

func NewVariableTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableTypeContext {
	var p = new(VariableTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_variableType

	return p
}

func (s *VariableTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableTypeContext) Kw_INT() antlr.TerminalNode {
	return s.GetToken(SwiftKw_INT, 0)
}

func (s *VariableTypeContext) Kw_FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftKw_FLOAT, 0)
}

func (s *VariableTypeContext) Kw_BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftKw_BOOL, 0)
}

func (s *VariableTypeContext) Kw_STRING() antlr.TerminalNode {
	return s.GetToken(SwiftKw_STRING, 0)
}

func (s *VariableTypeContext) Kw_CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_CHAR, 0)
}

func (s *VariableTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftID, 0)
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

func (p *Swift) VariableType() (localctx IVariableTypeContext) {
	localctx = NewVariableTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftRULE_variableType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137959047168) != 0) {
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
	p.RuleIndex = SwiftRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_variableDeclaration

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
	return s.GetToken(SwiftID, 0)
}

func (s *ValueDeclarationContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_ASSIGN, 0)
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
	return s.GetToken(SwiftKw_LET, 0)
}

func (s *ValueDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_VAR, 0)
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
	return s.GetToken(SwiftID, 0)
}

func (s *TypeValueDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, 0)
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
	return s.GetToken(SwiftOp_ASSIGN, 0)
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
	return s.GetToken(SwiftKw_LET, 0)
}

func (s *TypeValueDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_VAR, 0)
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
	return s.GetToken(SwiftID, 0)
}

func (s *TypeDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, 0)
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

func (s *TypeDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftKw_LET, 0)
}

func (s *TypeDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_VAR, 0)
}

func (s *TypeDeclarationContext) Op_TERNARY() antlr.TerminalNode {
	return s.GetToken(SwiftOp_TERNARY, 0)
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftRULE_variableDeclaration)
	var _la int

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TypeValueDeclarationContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftKw_LET || _la == SwiftKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TypeValueDeclarationContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.VariableType()
		}
		{
			p.SetState(149)
			p.Match(SwiftOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.expr(0)
		}

	case 2:
		localctx = NewValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueDeclarationContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftKw_LET || _la == SwiftKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueDeclarationContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(153)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(SwiftOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.expr(0)
		}

	case 3:
		localctx = NewTypeDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TypeDeclarationContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftKw_LET || _la == SwiftKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TypeDeclarationContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(157)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.VariableType()
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftOp_TERNARY {
			{
				p.SetState(160)
				p.Match(SwiftOp_TERNARY)
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
	p.RuleIndex = SwiftRULE_functionDeclarationStatement
	return p
}

func InitEmptyFunctionDeclarationStatementContext(p *FunctionDeclarationStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionDeclarationStatement
}

func (*FunctionDeclarationStatementContext) IsFunctionDeclarationStatementContext() {}

func NewFunctionDeclarationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationStatementContext {
	var p = new(FunctionDeclarationStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionDeclarationStatement

	return p
}

func (s *FunctionDeclarationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationStatementContext) Kw_FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftKw_FUNC, 0)
}

func (s *FunctionDeclarationStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftID, 0)
}

func (s *FunctionDeclarationStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftLPAREN, 0)
}

func (s *FunctionDeclarationStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftRPAREN, 0)
}

func (s *FunctionDeclarationStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
}

func (s *FunctionDeclarationStatementContext) Kw_MUTATING() antlr.TerminalNode {
	return s.GetToken(SwiftKw_MUTATING, 0)
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

func (p *Swift) FunctionDeclarationStatement() (localctx IFunctionDeclarationStatementContext) {
	localctx = NewFunctionDeclarationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftRULE_functionDeclarationStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftKw_MUTATING {
		{
			p.SetState(165)
			p.Match(SwiftKw_MUTATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(168)
		p.Match(SwiftKw_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(SwiftID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(SwiftLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftID {
		{
			p.SetState(171)
			p.FunctionParameters()
		}

	}
	{
		p.SetState(174)
		p.Match(SwiftRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftOp_ARROW {
		{
			p.SetState(175)
			p.FunctionReturnType()
		}

	}
	{
		p.SetState(178)
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Block()
	}
	{
		p.SetState(180)
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_functionParameters
	return p
}

func InitEmptyFunctionParametersContext(p *FunctionParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionParameters
}

func (*FunctionParametersContext) IsFunctionParametersContext() {}

func NewFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionParameters

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
	return s.GetTokens(SwiftCOMMA)
}

func (s *FunctionParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOMMA, i)
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

func (p *Swift) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftRULE_functionParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.FunctionParameter()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftCOMMA {
		{
			p.SetState(183)
			p.Match(SwiftCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.FunctionParameter()
		}

		p.SetState(189)
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
	MatrixType() IMatrixTypeContext
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
	p.RuleIndex = SwiftRULE_functionParameter
	return p
}

func InitEmptyFunctionParameterContext(p *FunctionParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionParameter
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftID)
}

func (s *FunctionParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *FunctionParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, 0)
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

func (s *FunctionParameterContext) MatrixType() IMatrixTypeContext {
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

func (s *FunctionParameterContext) Kw_INOUT() antlr.TerminalNode {
	return s.GetToken(SwiftKw_INOUT, 0)
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

func (p *Swift) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftRULE_functionParameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(190)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(193)
		p.Match(SwiftID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(SwiftCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftKw_INOUT {
		{
			p.SetState(195)
			p.Match(SwiftKw_INOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftKw_INT, SwiftKw_FLOAT, SwiftKw_BOOL, SwiftKw_STRING, SwiftKw_CHAR, SwiftID:
		{
			p.SetState(198)
			p.VariableType()
		}

	case SwiftLBRACKET:
		{
			p.SetState(199)
			p.MatrixType()
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

// IFunctionParameterCompoundContext is an interface to support dynamic dispatch.
type IFunctionParameterCompoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionParameterCompoundContext differentiates from other interfaces.
	IsFunctionParameterCompoundContext()
}

type FunctionParameterCompoundContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterCompoundContext() *FunctionParameterCompoundContext {
	var p = new(FunctionParameterCompoundContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionParameterCompound
	return p
}

func InitEmptyFunctionParameterCompoundContext(p *FunctionParameterCompoundContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionParameterCompound
}

func (*FunctionParameterCompoundContext) IsFunctionParameterCompoundContext() {}

func NewFunctionParameterCompoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterCompoundContext {
	var p = new(FunctionParameterCompoundContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionParameterCompound

	return p
}

func (s *FunctionParameterCompoundContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterCompoundContext) CopyAll(ctx *FunctionParameterCompoundContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionParameterCompoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterCompoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionParameterCompoundNestedContext struct {
	FunctionParameterCompoundContext
}

func NewFunctionParameterCompoundNestedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionParameterCompoundNestedContext {
	var p = new(FunctionParameterCompoundNestedContext)

	InitEmptyFunctionParameterCompoundContext(&p.FunctionParameterCompoundContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionParameterCompoundContext))

	return p
}

func (s *FunctionParameterCompoundNestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterCompoundNestedContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACKET, 0)
}

func (s *FunctionParameterCompoundNestedContext) FunctionParameterCompound() IFunctionParameterCompoundContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionParameterCompoundContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterCompoundContext)
}

func (s *FunctionParameterCompoundNestedContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, 0)
}

func (s *FunctionParameterCompoundNestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionParameterCompoundNested(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionParameterCompoundSingleContext struct {
	FunctionParameterCompoundContext
}

func NewFunctionParameterCompoundSingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionParameterCompoundSingleContext {
	var p = new(FunctionParameterCompoundSingleContext)

	InitEmptyFunctionParameterCompoundContext(&p.FunctionParameterCompoundContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionParameterCompoundContext))

	return p
}

func (s *FunctionParameterCompoundSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterCompoundSingleContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACKET, 0)
}

func (s *FunctionParameterCompoundSingleContext) VariableType() IVariableTypeContext {
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

func (s *FunctionParameterCompoundSingleContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, 0)
}

func (s *FunctionParameterCompoundSingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitFunctionParameterCompoundSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) FunctionParameterCompound() (localctx IFunctionParameterCompoundContext) {
	localctx = NewFunctionParameterCompoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftRULE_functionParameterCompound)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionParameterCompoundNestedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.FunctionParameterCompound()
		}
		{
			p.SetState(204)
			p.Match(SwiftRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionParameterCompoundSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.VariableType()
		}
		{
			p.SetState(208)
			p.Match(SwiftRBRACKET)
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
	p.RuleIndex = SwiftRULE_functionReturnType
	return p
}

func InitEmptyFunctionReturnTypeContext(p *FunctionReturnTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionReturnType
}

func (*FunctionReturnTypeContext) IsFunctionReturnTypeContext() {}

func NewFunctionReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnTypeContext {
	var p = new(FunctionReturnTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionReturnType

	return p
}

func (s *FunctionReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnTypeContext) Op_ARROW() antlr.TerminalNode {
	return s.GetToken(SwiftOp_ARROW, 0)
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

func (p *Swift) FunctionReturnType() (localctx IFunctionReturnTypeContext) {
	localctx = NewFunctionReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftRULE_functionReturnType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(SwiftOp_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
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
	p.RuleIndex = SwiftRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftRPAREN, 0)
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

func (p *Swift) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(215)
			p.IdChain()
		}

	case 2:
		{
			p.SetState(216)
			p.VariableType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(219)
		p.Match(SwiftLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1459166552518558208) != 0 {
		{
			p.SetState(220)
			p.FunctionCallArguments()
		}

	}
	{
		p.SetState(223)
		p.Match(SwiftRPAREN)
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
	p.RuleIndex = SwiftRULE_functionCallArguments
	return p
}

func InitEmptyFunctionCallArgumentsContext(p *FunctionCallArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_functionCallArguments
}

func (*FunctionCallArgumentsContext) IsFunctionCallArgumentsContext() {}

func NewFunctionCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_functionCallArguments

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
	return s.GetTokens(SwiftCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOMMA, i)
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
	return s.GetTokens(SwiftID)
}

func (s *NamedArgumentsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *NamedArgumentsContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SwiftCOLON)
}

func (s *NamedArgumentsContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, i)
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
	return s.GetTokens(SwiftCOMMA)
}

func (s *NamedArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOMMA, i)
}

func (s *NamedArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitNamedArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) FunctionCallArguments() (localctx IFunctionCallArgumentsContext) {
	localctx = NewFunctionCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftRULE_functionCallArguments)
	var _la int

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArgumentsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.expr(0)
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftCOMMA {
			{
				p.SetState(226)
				p.Match(SwiftCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(227)
				p.expr(0)
			}

			p.SetState(232)
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
			p.SetState(233)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.expr(0)
		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftCOMMA {
			{
				p.SetState(236)
				p.Match(SwiftCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(237)
				p.Match(SwiftID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(238)
				p.Match(SwiftCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(239)
				p.expr(0)
			}

			p.SetState(244)
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
	IdChain() IIdChainContext
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
	p.RuleIndex = SwiftRULE_variableAssignment
	return p
}

func InitEmptyVariableAssignmentContext(p *VariableAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_variableAssignment
}

func (*VariableAssignmentContext) IsVariableAssignmentContext() {}

func NewVariableAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableAssignmentContext {
	var p = new(VariableAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_variableAssignment

	return p
}

func (s *VariableAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableAssignmentContext) GetOp() antlr.Token { return s.op }

func (s *VariableAssignmentContext) SetOp(v antlr.Token) { s.op = v }

func (s *VariableAssignmentContext) IdChain() IIdChainContext {
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
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *VariableAssignmentContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_PLUS_ASSIGN, 0)
}

func (s *VariableAssignmentContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_MINUS_ASSIGN, 0)
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

func (p *Swift) VariableAssignment() (localctx IVariableAssignmentContext) {
	localctx = NewVariableAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftRULE_variableAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.IdChain()
	}
	{
		p.SetState(248)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&879609302220800) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(249)
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

// IVariableAssignmentObjectContext is an interface to support dynamic dispatch.
type IVariableAssignmentObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ObjectChain() IObjectChainContext
	Expr() IExprContext
	Op_ASSIGN() antlr.TerminalNode
	Op_PLUS_ASSIGN() antlr.TerminalNode
	Op_MINUS_ASSIGN() antlr.TerminalNode

	// IsVariableAssignmentObjectContext differentiates from other interfaces.
	IsVariableAssignmentObjectContext()
}

type VariableAssignmentObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyVariableAssignmentObjectContext() *VariableAssignmentObjectContext {
	var p = new(VariableAssignmentObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_variableAssignmentObject
	return p
}

func InitEmptyVariableAssignmentObjectContext(p *VariableAssignmentObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_variableAssignmentObject
}

func (*VariableAssignmentObjectContext) IsVariableAssignmentObjectContext() {}

func NewVariableAssignmentObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableAssignmentObjectContext {
	var p = new(VariableAssignmentObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_variableAssignmentObject

	return p
}

func (s *VariableAssignmentObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableAssignmentObjectContext) GetOp() antlr.Token { return s.op }

func (s *VariableAssignmentObjectContext) SetOp(v antlr.Token) { s.op = v }

func (s *VariableAssignmentObjectContext) ObjectChain() IObjectChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectChainContext)
}

func (s *VariableAssignmentObjectContext) Expr() IExprContext {
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

func (s *VariableAssignmentObjectContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *VariableAssignmentObjectContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_PLUS_ASSIGN, 0)
}

func (s *VariableAssignmentObjectContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_MINUS_ASSIGN, 0)
}

func (s *VariableAssignmentObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableAssignmentObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableAssignmentObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVariableAssignmentObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) VariableAssignmentObject() (localctx IVariableAssignmentObjectContext) {
	localctx = NewVariableAssignmentObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftRULE_variableAssignmentObject)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.ObjectChain()
	}
	{
		p.SetState(252)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableAssignmentObjectContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&879609302220800) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableAssignmentObjectContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(253)
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
	p.RuleIndex = SwiftRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_ifStatement

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
	return s.GetTokens(SwiftKw_ELSE)
}

func (s *IfStatementContext) Kw_ELSE(i int) antlr.TerminalNode {
	return s.GetToken(SwiftKw_ELSE, i)
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

func (p *Swift) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.IfTail()
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(256)
				p.Match(SwiftKw_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(257)
				p.IfTail()
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftKw_ELSE {
		{
			p.SetState(263)
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
	p.RuleIndex = SwiftRULE_ifTail
	return p
}

func InitEmptyIfTailContext(p *IfTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_ifTail
}

func (*IfTailContext) IsIfTailContext() {}

func NewIfTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfTailContext {
	var p = new(IfTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_ifTail

	return p
}

func (s *IfTailContext) GetParser() antlr.Parser { return s.parser }

func (s *IfTailContext) Kw_IF() antlr.TerminalNode {
	return s.GetToken(SwiftKw_IF, 0)
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
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) IfTail() (localctx IIfTailContext) {
	localctx = NewIfTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftRULE_ifTail)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(SwiftKw_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.expr(0)
	}
	{
		p.SetState(268)
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Block()
	}
	{
		p.SetState(270)
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_elseStatement
	return p
}

func InitEmptyElseStatementContext(p *ElseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_elseStatement
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) Kw_ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftKw_ELSE, 0)
}

func (s *ElseStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftRULE_elseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(SwiftKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(SwiftLBRACE)
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
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) Kw_WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftKw_WHILE, 0)
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
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(SwiftKw_WHILE)
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
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.Block()
	}
	{
		p.SetState(281)
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Kw_SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftKw_SWITCH, 0)
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
	return s.GetToken(SwiftLBRACE, 0)
}

func (s *SwitchStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(SwiftKw_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.expr(0)
	}
	{
		p.SetState(285)
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftKw_CASE {
		{
			p.SetState(286)
			p.SwitchCase()
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftKw_DEFAULT {
		{
			p.SetState(292)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(295)
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) Kw_CASE() antlr.TerminalNode {
	return s.GetToken(SwiftKw_CASE, 0)
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
	return s.GetToken(SwiftCOLON, 0)
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
	return s.GetToken(SwiftKw_BREAK, 0)
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

func (p *Swift) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(SwiftKw_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.expr(0)
	}
	{
		p.SetState(299)
		p.Match(SwiftCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Block()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftKw_BREAK {
		{
			p.SetState(301)
			p.Match(SwiftKw_BREAK)
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
	p.RuleIndex = SwiftRULE_switchDefault
	return p
}

func InitEmptySwitchDefaultContext(p *SwitchDefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_switchDefault
}

func (*SwitchDefaultContext) IsSwitchDefaultContext() {}

func NewSwitchDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchDefaultContext {
	var p = new(SwitchDefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_switchDefault

	return p
}

func (s *SwitchDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchDefaultContext) Kw_DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftKw_DEFAULT, 0)
}

func (s *SwitchDefaultContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, 0)
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
	return s.GetToken(SwiftKw_BREAK, 0)
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

func (p *Swift) SwitchDefault() (localctx ISwitchDefaultContext) {
	localctx = NewSwitchDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftRULE_switchDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(SwiftKw_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(SwiftCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Block()
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftKw_BREAK {
		{
			p.SetState(307)
			p.Match(SwiftKw_BREAK)
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
	p.RuleIndex = SwiftRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) Kw_FOR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_FOR, 0)
}

func (s *ForStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftID, 0)
}

func (s *ForStatementContext) Kw_IN() antlr.TerminalNode {
	return s.GetToken(SwiftKw_IN, 0)
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
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(SwiftKw_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(SwiftID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Match(SwiftKw_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.expr(0)
	}
	{
		p.SetState(314)
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Block()
	}
	{
		p.SetState(316)
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_guardStatement
	return p
}

func InitEmptyGuardStatementContext(p *GuardStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_guardStatement
}

func (*GuardStatementContext) IsGuardStatementContext() {}

func NewGuardStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardStatementContext {
	var p = new(GuardStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_guardStatement

	return p
}

func (s *GuardStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardStatementContext) Kw_GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftKw_GUARD, 0)
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
	return s.GetToken(SwiftKw_ELSE, 0)
}

func (s *GuardStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) GuardStatement() (localctx IGuardStatementContext) {
	localctx = NewGuardStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftRULE_guardStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(SwiftKw_GUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.expr(0)
	}
	{
		p.SetState(320)
		p.Match(SwiftKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Block()
	}
	{
		p.SetState(323)
		p.Match(SwiftRBRACE)
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
	// IsVectorDeclarationContext differentiates from other interfaces.
	IsVectorDeclarationContext()
}

type VectorDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorDeclarationContext() *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_vectorDeclaration
	return p
}

func InitEmptyVectorDeclarationContext(p *VectorDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_vectorDeclaration
}

func (*VectorDeclarationContext) IsVectorDeclarationContext() {}

func NewVectorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_vectorDeclaration

	return p
}

func (s *VectorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorDeclarationContext) CopyAll(ctx *VectorDeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VectorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorTypeValueContext struct {
	VectorDeclarationContext
	varType antlr.Token
}

func NewVectorTypeValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorTypeValueContext {
	var p = new(VectorTypeValueContext)

	InitEmptyVectorDeclarationContext(&p.VectorDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorDeclarationContext))

	return p
}

func (s *VectorTypeValueContext) GetVarType() antlr.Token { return s.varType }

func (s *VectorTypeValueContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *VectorTypeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorTypeValueContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftID, 0)
}

func (s *VectorTypeValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, 0)
}

func (s *VectorTypeValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACKET, 0)
}

func (s *VectorTypeValueContext) VariableType() IVariableTypeContext {
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

func (s *VectorTypeValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, 0)
}

func (s *VectorTypeValueContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *VectorTypeValueContext) VectorDefinition() IVectorDefinitionContext {
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

func (s *VectorTypeValueContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftKw_LET, 0)
}

func (s *VectorTypeValueContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_VAR, 0)
}

func (s *VectorTypeValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorTypeValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorStructValueContext struct {
	VectorDeclarationContext
	varType antlr.Token
}

func NewVectorStructValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorStructValueContext {
	var p = new(VectorStructValueContext)

	InitEmptyVectorDeclarationContext(&p.VectorDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorDeclarationContext))

	return p
}

func (s *VectorStructValueContext) GetVarType() antlr.Token { return s.varType }

func (s *VectorStructValueContext) SetVarType(v antlr.Token) { s.varType = v }

func (s *VectorStructValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorStructValueContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftID)
}

func (s *VectorStructValueContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *VectorStructValueContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *VectorStructValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACKET, 0)
}

func (s *VectorStructValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, 0)
}

func (s *VectorStructValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftLPAREN, 0)
}

func (s *VectorStructValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftRPAREN, 0)
}

func (s *VectorStructValueContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftKw_LET, 0)
}

func (s *VectorStructValueContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_VAR, 0)
}

func (s *VectorStructValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitVectorStructValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) VectorDeclaration() (localctx IVectorDeclarationContext) {
	localctx = NewVectorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftRULE_vectorDeclaration)
	var _la int

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorTypeValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VectorTypeValueContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftKw_LET || _la == SwiftKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VectorTypeValueContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(326)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.VariableType()
		}
		{
			p.SetState(330)
			p.Match(SwiftRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(SwiftOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.VectorDefinition()
		}

	case 2:
		localctx = NewVectorStructValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VectorStructValueContext).varType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftKw_LET || _la == SwiftKw_VAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VectorStructValueContext).varType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(335)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(SwiftOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(SwiftRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(SwiftLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(SwiftRPAREN)
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
	p.RuleIndex = SwiftRULE_vectorDefinition
	return p
}

func InitEmptyVectorDefinitionContext(p *VectorDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_vectorDefinition
}

func (*VectorDefinitionContext) IsVectorDefinitionContext() {}

func NewVectorDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorDefinitionContext {
	var p = new(VectorDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_vectorDefinition

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
	return s.GetToken(SwiftLBRACKET, 0)
}

func (s *VectorListValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, 0)
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

func (p *Swift) VectorDefinition() (localctx IVectorDefinitionContext) {
	localctx = NewVectorDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftRULE_vectorDefinition)
	var _la int

	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftLBRACKET:
		localctx = NewVectorListValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1459166552518558208) != 0 {
			{
				p.SetState(345)
				p.VectorValues()
			}

		}
		{
			p.SetState(348)
			p.Match(SwiftRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftKw_AMPER, SwiftKw_INT, SwiftKw_FLOAT, SwiftKw_BOOL, SwiftKw_STRING, SwiftKw_CHAR, SwiftINT, SwiftFLOAT, SwiftBOOL, SwiftSTRING, SwiftCHAR, SwiftNIL, SwiftID, SwiftOp_MINUS, SwiftOp_NOT, SwiftLPAREN:
		localctx = NewVectorSingleValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
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
	p.RuleIndex = SwiftRULE_vectorValues
	return p
}

func InitEmptyVectorValuesContext(p *VectorValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_vectorValues
}

func (*VectorValuesContext) IsVectorValuesContext() {}

func NewVectorValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorValuesContext {
	var p = new(VectorValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_vectorValues

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
	return s.GetTokens(SwiftCOMMA)
}

func (s *VectorValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOMMA, i)
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

func (p *Swift) VectorValues() (localctx IVectorValuesContext) {
	localctx = NewVectorValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftRULE_vectorValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.expr(0)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftCOMMA {
		{
			p.SetState(353)
			p.Match(SwiftCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.expr(0)
		}

		p.SetState(359)
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

// IObjectChainContext is an interface to support dynamic dispatch.
type IObjectChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VectorAccess() IVectorAccessContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsObjectChainContext differentiates from other interfaces.
	IsObjectChainContext()
}

type ObjectChainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectChainContext() *ObjectChainContext {
	var p = new(ObjectChainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_objectChain
	return p
}

func InitEmptyObjectChainContext(p *ObjectChainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_objectChain
}

func (*ObjectChainContext) IsObjectChainContext() {}

func NewObjectChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectChainContext {
	var p = new(ObjectChainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_objectChain

	return p
}

func (s *ObjectChainContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectChainContext) VectorAccess() IVectorAccessContext {
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

func (s *ObjectChainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SwiftDOT)
}

func (s *ObjectChainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SwiftDOT, i)
}

func (s *ObjectChainContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftID)
}

func (s *ObjectChainContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *ObjectChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitObjectChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) ObjectChain() (localctx IObjectChainContext) {
	localctx = NewObjectChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftRULE_objectChain)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.VectorAccess()
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(361)
				p.Match(SwiftDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(362)
				p.Match(SwiftID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
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
	p.RuleIndex = SwiftRULE_vectorAccess
	return p
}

func InitEmptyVectorAccessContext(p *VectorAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_vectorAccess
}

func (*VectorAccessContext) IsVectorAccessContext() {}

func NewVectorAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorAccessContext {
	var p = new(VectorAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_vectorAccess

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
	return s.GetToken(SwiftLBRACKET, 0)
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
	return s.GetToken(SwiftRBRACKET, 0)
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

func (p *Swift) VectorAccess() (localctx IVectorAccessContext) {
	localctx = NewVectorAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftRULE_vectorAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.IdChain()
	}
	{
		p.SetState(368)
		p.Match(SwiftLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.expr(0)
	}
	{
		p.SetState(370)
		p.Match(SwiftRBRACKET)
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
	p.RuleIndex = SwiftRULE_vectorAssignment
	return p
}

func InitEmptyVectorAssignmentContext(p *VectorAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_vectorAssignment
}

func (*VectorAssignmentContext) IsVectorAssignmentContext() {}

func NewVectorAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorAssignmentContext {
	var p = new(VectorAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_vectorAssignment

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
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *VectorAssignmentContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_PLUS_ASSIGN, 0)
}

func (s *VectorAssignmentContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_MINUS_ASSIGN, 0)
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

func (p *Swift) VectorAssignment() (localctx IVectorAssignmentContext) {
	localctx = NewVectorAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftRULE_vectorAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.VectorAccess()
	}
	{
		p.SetState(373)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VectorAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&879609302220800) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VectorAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(374)
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
	p.RuleIndex = SwiftRULE_matrixDeclaration
	return p
}

func InitEmptyMatrixDeclarationContext(p *MatrixDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixDeclaration
}

func (*MatrixDeclarationContext) IsMatrixDeclarationContext() {}

func NewMatrixDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixDeclarationContext {
	var p = new(MatrixDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixDeclaration

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
	return s.GetToken(SwiftCOLON, 0)
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
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *MatrixDeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftKw_LET, 0)
}

func (s *MatrixDeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftKw_VAR, 0)
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

func (p *Swift) MatrixDeclaration() (localctx IMatrixDeclarationContext) {
	localctx = NewMatrixDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftRULE_matrixDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MatrixDeclarationContext).varType = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftKw_LET || _la == SwiftKw_VAR) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MatrixDeclarationContext).varType = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(377)
		p.IdChain()
	}
	{
		p.SetState(378)
		p.Match(SwiftCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.MatrixType()
	}
	{
		p.SetState(380)
		p.Match(SwiftOp_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(381)
			p.MatrixRepeatingDefinition()
		}

	case 2:
		{
			p.SetState(382)
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
	p.RuleIndex = SwiftRULE_matrixType
	return p
}

func InitEmptyMatrixTypeContext(p *MatrixTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixType
}

func (*MatrixTypeContext) IsMatrixTypeContext() {}

func NewMatrixTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixTypeContext {
	var p = new(MatrixTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixType

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
	return s.GetToken(SwiftLBRACKET, 0)
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
	return s.GetToken(SwiftRBRACKET, 0)
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
	return s.GetToken(SwiftLBRACKET, 0)
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
	return s.GetToken(SwiftRBRACKET, 0)
}

func (s *MatrixTypeNestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixTypeNested(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) MatrixType() (localctx IMatrixTypeContext) {
	localctx = NewMatrixTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftRULE_matrixType)
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMatrixTypeNestedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.MatrixType()
		}
		{
			p.SetState(387)
			p.Match(SwiftRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMatrixTypeSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.VariableType()
		}
		{
			p.SetState(391)
			p.Match(SwiftRBRACKET)
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
	p.RuleIndex = SwiftRULE_matrixDefinition
	return p
}

func InitEmptyMatrixDefinitionContext(p *MatrixDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixDefinition
}

func (*MatrixDefinitionContext) IsMatrixDefinitionContext() {}

func NewMatrixDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixDefinitionContext {
	var p = new(MatrixDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixDefinition

	return p
}

func (s *MatrixDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixDefinitionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACKET, 0)
}

func (s *MatrixDefinitionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, 0)
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

func (p *Swift) MatrixDefinition() (localctx IMatrixDefinitionContext) {
	localctx = NewMatrixDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SwiftRULE_matrixDefinition)
	var _la int

	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Match(SwiftLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-9)) & ^0x3f) == 0 && ((int64(1)<<(_la-9))&38878731691851777) != 0 {
			{
				p.SetState(396)
				p.MatrixValues()
			}

		}
		{
			p.SetState(399)
			p.Match(SwiftRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftKw_AMPER, SwiftKw_INT, SwiftKw_FLOAT, SwiftKw_BOOL, SwiftKw_STRING, SwiftKw_CHAR, SwiftINT, SwiftFLOAT, SwiftBOOL, SwiftSTRING, SwiftCHAR, SwiftNIL, SwiftID, SwiftOp_MINUS, SwiftOp_NOT, SwiftLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
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
	p.RuleIndex = SwiftRULE_matrixValues
	return p
}

func InitEmptyMatrixValuesContext(p *MatrixValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixValues
}

func (*MatrixValuesContext) IsMatrixValuesContext() {}

func NewMatrixValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixValuesContext {
	var p = new(MatrixValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixValues

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
	return s.GetTokens(SwiftCOMMA)
}

func (s *MatrixValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOMMA, i)
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

func (p *Swift) MatrixValues() (localctx IMatrixValuesContext) {
	localctx = NewMatrixValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SwiftRULE_matrixValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.MatrixDefinition()
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftCOMMA {
		{
			p.SetState(404)
			p.Match(SwiftCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.MatrixDefinition()
		}

		p.SetState(410)
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
	p.RuleIndex = SwiftRULE_matrixRepeatingDefinition
	return p
}

func InitEmptyMatrixRepeatingDefinitionContext(p *MatrixRepeatingDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixRepeatingDefinition
}

func (*MatrixRepeatingDefinitionContext) IsMatrixRepeatingDefinitionContext() {}

func NewMatrixRepeatingDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixRepeatingDefinitionContext {
	var p = new(MatrixRepeatingDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixRepeatingDefinition

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
	return s.GetToken(SwiftLPAREN, 0)
}

func (s *MatrixRepeatingDefinitionSingleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftID)
}

func (s *MatrixRepeatingDefinitionSingleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *MatrixRepeatingDefinitionSingleContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SwiftCOLON)
}

func (s *MatrixRepeatingDefinitionSingleContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, i)
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
	return s.GetToken(SwiftCOMMA, 0)
}

func (s *MatrixRepeatingDefinitionSingleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftRPAREN, 0)
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
	return s.GetToken(SwiftLPAREN, 0)
}

func (s *MatrixRepeatingDefinitionNestedContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftID)
}

func (s *MatrixRepeatingDefinitionNestedContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftID, i)
}

func (s *MatrixRepeatingDefinitionNestedContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(SwiftCOLON)
}

func (s *MatrixRepeatingDefinitionNestedContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, i)
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
	return s.GetToken(SwiftCOMMA, 0)
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
	return s.GetToken(SwiftRPAREN, 0)
}

func (s *MatrixRepeatingDefinitionNestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitMatrixRepeatingDefinitionNested(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) MatrixRepeatingDefinition() (localctx IMatrixRepeatingDefinitionContext) {
	localctx = NewMatrixRepeatingDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftRULE_matrixRepeatingDefinition)
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMatrixRepeatingDefinitionNestedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.MatrixType()
		}
		{
			p.SetState(412)
			p.Match(SwiftLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.MatrixRepeatingDefinition()
		}
		{
			p.SetState(416)
			p.Match(SwiftCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(SwiftCOLON)
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
			p.Match(SwiftRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMatrixRepeatingDefinitionSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.MatrixType()
		}
		{
			p.SetState(423)
			p.Match(SwiftLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.expr(0)
		}
		{
			p.SetState(427)
			p.Match(SwiftCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(SwiftID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Match(SwiftCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.expr(0)
		}
		{
			p.SetState(431)
			p.Match(SwiftRPAREN)
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
	p.RuleIndex = SwiftRULE_matrixAccess
	return p
}

func InitEmptyMatrixAccessContext(p *MatrixAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixAccess
}

func (*MatrixAccessContext) IsMatrixAccessContext() {}

func NewMatrixAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixAccessContext {
	var p = new(MatrixAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixAccess

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
	return s.GetTokens(SwiftLBRACKET)
}

func (s *MatrixAccessContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SwiftLBRACKET, i)
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
	return s.GetTokens(SwiftRBRACKET)
}

func (s *MatrixAccessContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(SwiftRBRACKET, i)
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

func (p *Swift) MatrixAccess() (localctx IMatrixAccessContext) {
	localctx = NewMatrixAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SwiftRULE_matrixAccess)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.IdChain()
	}
	{
		p.SetState(436)
		p.Match(SwiftLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(437)
		p.expr(0)
	}
	{
		p.SetState(438)
		p.Match(SwiftRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(439)
				p.Match(SwiftLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(440)
				p.expr(0)
			}
			{
				p.SetState(441)
				p.Match(SwiftRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
	p.RuleIndex = SwiftRULE_matrixAssignment
	return p
}

func InitEmptyMatrixAssignmentContext(p *MatrixAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_matrixAssignment
}

func (*MatrixAssignmentContext) IsMatrixAssignmentContext() {}

func NewMatrixAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixAssignmentContext {
	var p = new(MatrixAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_matrixAssignment

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
	return s.GetToken(SwiftOp_ASSIGN, 0)
}

func (s *MatrixAssignmentContext) Op_PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_PLUS_ASSIGN, 0)
}

func (s *MatrixAssignmentContext) Op_MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftOp_MINUS_ASSIGN, 0)
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

func (p *Swift) MatrixAssignment() (localctx IMatrixAssignmentContext) {
	localctx = NewMatrixAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SwiftRULE_matrixAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.MatrixAccess()
	}
	{
		p.SetState(449)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MatrixAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&879609302220800) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MatrixAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(450)
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
	p.RuleIndex = SwiftRULE_structDeclaration
	return p
}

func InitEmptyStructDeclarationContext(p *StructDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_structDeclaration
}

func (*StructDeclarationContext) IsStructDeclarationContext() {}

func NewStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_structDeclaration

	return p
}

func (s *StructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationContext) Kw_STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftKw_STRUCT, 0)
}

func (s *StructDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftID, 0)
}

func (s *StructDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SwiftLBRACE, 0)
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
	return s.GetToken(SwiftRBRACE, 0)
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

func (p *Swift) StructDeclaration() (localctx IStructDeclarationContext) {
	localctx = NewStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SwiftRULE_structDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(SwiftKw_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.Match(SwiftID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.Match(SwiftLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.StructBody()
	}
	{
		p.SetState(456)
		p.Match(SwiftRBRACE)
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
	p.RuleIndex = SwiftRULE_structBody
	return p
}

func InitEmptyStructBodyContext(p *StructBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_structBody
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_structBody

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

func (p *Swift) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SwiftRULE_structBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&368) != 0 {
		{
			p.SetState(458)
			p.StructProperty()
		}

		p.SetState(463)
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
	p.RuleIndex = SwiftRULE_structProperty
	return p
}

func InitEmptyStructPropertyContext(p *StructPropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_structProperty
}

func (*StructPropertyContext) IsStructPropertyContext() {}

func NewStructPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructPropertyContext {
	var p = new(StructPropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_structProperty

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

func (p *Swift) StructProperty() (localctx IStructPropertyContext) {
	localctx = NewStructPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SwiftRULE_structProperty)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftKw_LET, SwiftKw_VAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.VariableDeclaration()
		}

	case SwiftKw_FUNC, SwiftKw_MUTATING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
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
	p.RuleIndex = SwiftRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_expr

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
	return s.GetToken(SwiftBOOL, 0)
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
	return s.GetToken(SwiftFLOAT, 0)
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
	return s.GetToken(SwiftNIL, 0)
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

func (s *IdExprContext) Kw_AMPER() antlr.TerminalNode {
	return s.GetToken(SwiftKw_AMPER, 0)
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewComparisonExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparisonExprContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonExprContext) GetLeft() IExprContext { return s.left }

func (s *ComparisonExprContext) GetRight() IExprContext { return s.right }

func (s *ComparisonExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ComparisonExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) AllExpr() []IExprContext {
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

func (s *ComparisonExprContext) Expr(i int) IExprContext {
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

func (s *ComparisonExprContext) Op_GE() antlr.TerminalNode {
	return s.GetToken(SwiftOp_GE, 0)
}

func (s *ComparisonExprContext) Op_GT() antlr.TerminalNode {
	return s.GetToken(SwiftOp_GT, 0)
}

func (s *ComparisonExprContext) Op_LE() antlr.TerminalNode {
	return s.GetToken(SwiftOp_LE, 0)
}

func (s *ComparisonExprContext) Op_LT() antlr.TerminalNode {
	return s.GetToken(SwiftOp_LT, 0)
}

func (s *ComparisonExprContext) Op_EQ() antlr.TerminalNode {
	return s.GetToken(SwiftOp_EQ, 0)
}

func (s *ComparisonExprContext) Op_NEQ() antlr.TerminalNode {
	return s.GetToken(SwiftOp_NEQ, 0)
}

func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitComparisonExpr(s)

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
	return s.GetToken(SwiftKw_RANGE, 0)
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
	return s.GetToken(SwiftOp_MINUS, 0)
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
	return s.GetToken(SwiftCHAR, 0)
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
	return s.GetToken(SwiftOp_MUL, 0)
}

func (s *ArithmeticExprContext) Op_DIV() antlr.TerminalNode {
	return s.GetToken(SwiftOp_DIV, 0)
}

func (s *ArithmeticExprContext) Op_PLUS() antlr.TerminalNode {
	return s.GetToken(SwiftOp_PLUS, 0)
}

func (s *ArithmeticExprContext) Op_MINUS() antlr.TerminalNode {
	return s.GetToken(SwiftOp_MINUS, 0)
}

func (s *ArithmeticExprContext) Op_MOD() antlr.TerminalNode {
	return s.GetToken(SwiftOp_MOD, 0)
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
	return s.GetToken(SwiftLPAREN, 0)
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
	return s.GetToken(SwiftRPAREN, 0)
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
	return s.GetToken(SwiftSTRING, 0)
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitStrExpr(s)

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
	return s.GetToken(SwiftOp_NOT, 0)
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
	return s.GetToken(SwiftINT, 0)
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
	return s.GetToken(SwiftOp_AND, 0)
}

func (s *LogicalExprContext) Op_OR() antlr.TerminalNode {
	return s.GetToken(SwiftOp_OR, 0)
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
	return s.GetToken(SwiftOp_TERNARY, 0)
}

func (s *TernaryExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftCOLON, 0)
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

type ObjectChainExprContext struct {
	ExprContext
}

func NewObjectChainExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectChainExprContext {
	var p = new(ObjectChainExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ObjectChainExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectChainExprContext) ObjectChain() IObjectChainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectChainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectChainContext)
}

func (s *ObjectChainExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitObjectChainExpr(s)

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

func (p *Swift) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *Swift) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 80
	p.EnterRecursionRule(localctx, 80, SwiftRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		localctx = NewObjectChainExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(469)
			p.ObjectChain()
		}

	case 2:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(470)
			p.FunctionCall()
		}

	case 3:
		localctx = NewVectorAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(471)
			p.VectorAccess()
		}

	case 4:
		localctx = NewMatrixAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(472)
			p.MatrixAccess()
		}

	case 5:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(473)
			p.Match(SwiftOp_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(474)
			p.expr(20)
		}

	case 6:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(475)
			p.Match(SwiftOp_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)

			var _x = p.expr(19)

			localctx.(*NotExprContext).right = _x
		}

	case 7:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(477)
			p.Match(SwiftLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.expr(0)
		}
		{
			p.SetState(479)
			p.Match(SwiftRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(481)
			p.Match(SwiftINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftKw_AMPER {
			{
				p.SetState(482)
				p.Match(SwiftKw_AMPER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(485)
			p.IdChain()
		}

	case 10:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(486)
			p.Match(SwiftFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(487)
			p.Match(SwiftSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(488)
			p.Match(SwiftNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(489)
			p.Match(SwiftCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(490)
			p.Match(SwiftBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(526)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(493)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(494)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftOp_MUL || _la == SwiftOp_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(495)

					var _x = p.expr(19)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(496)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(497)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftOp_PLUS || _la == SwiftOp_MINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(498)

					var _x = p.expr(18)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 3:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(499)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(500)

					var _m = p.Match(SwiftOp_MOD)

					localctx.(*ArithmeticExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(501)

					var _x = p.expr(17)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 4:
				localctx = NewComparisonExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparisonExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(502)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(503)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftOp_GT || _la == SwiftOp_GE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(504)

					var _x = p.expr(16)

					localctx.(*ComparisonExprContext).right = _x
				}

			case 5:
				localctx = NewComparisonExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparisonExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(505)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(506)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftOp_LT || _la == SwiftOp_LE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(507)

					var _x = p.expr(15)

					localctx.(*ComparisonExprContext).right = _x
				}

			case 6:
				localctx = NewComparisonExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparisonExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(508)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(509)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftOp_EQ || _la == SwiftOp_NEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(510)

					var _x = p.expr(14)

					localctx.(*ComparisonExprContext).right = _x
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(512)
					p.Match(SwiftOp_TERNARY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(513)

					var _x = p.expr(0)

					localctx.(*TernaryExprContext).cTrue = _x
				}
				{
					p.SetState(514)
					p.Match(SwiftCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(515)

					var _x = p.expr(13)

					localctx.(*TernaryExprContext).cFalse = _x
				}

			case 8:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(517)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(518)

					var _m = p.Match(SwiftOp_AND)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(519)

					var _x = p.expr(12)

					localctx.(*LogicalExprContext).right = _x
				}

			case 9:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(520)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(521)

					var _m = p.Match(SwiftOp_OR)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(522)

					var _x = p.expr(11)

					localctx.(*LogicalExprContext).right = _x
				}

			case 10:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RangeExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftRULE_expr)
				p.SetState(523)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(524)
					p.Match(SwiftKw_RANGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(525)

					var _x = p.expr(10)

					localctx.(*RangeExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
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
	p.RuleIndex = SwiftRULE_controlFlowStatement
	return p
}

func InitEmptyControlFlowStatementContext(p *ControlFlowStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftRULE_controlFlowStatement
}

func (*ControlFlowStatementContext) IsControlFlowStatementContext() {}

func NewControlFlowStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlFlowStatementContext {
	var p = new(ControlFlowStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftRULE_controlFlowStatement

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
	return s.GetToken(SwiftKw_RETURN, 0)
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
	return s.GetToken(SwiftKw_BREAK, 0)
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
	return s.GetToken(SwiftKw_CONTINUE, 0)
}

func (s *ControlContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitControlContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Swift) ControlFlowStatement() (localctx IControlFlowStatementContext) {
	localctx = NewControlFlowStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SwiftRULE_controlFlowStatement)
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftKw_BREAK:
		localctx = NewControlBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)
			p.Match(SwiftKw_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftKw_CONTINUE:
		localctx = NewControlContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.Match(SwiftKw_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftKw_RETURN:
		localctx = NewControlReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(533)
			p.Match(SwiftKw_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(534)
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

func (p *Swift) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 40:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Swift) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
