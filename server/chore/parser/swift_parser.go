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
		"'Bool'", "'String'", "'Character'", "'...'", "'inout'", "", "", "",
		"", "", "'nil'", "", "'->'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='",
		"'='", "'*='", "'/='", "'+='", "'-='", "'%='", "'*'", "'/'", "'+'",
		"'-'", "'%'", "'&&'", "'||'", "'!'", "'?'", "'('", "')'", "'{'", "'}'",
		"'['", "']'", "'\\'", "','", "';'", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_GUARD", "Kw_INT", "Kw_FLOAT", "Kw_BOOL", "Kw_STRING",
		"Kw_CHAR", "Kw_RANGE", "Kw_INOUT", "INT", "FLOAT", "BOOL", "STRING",
		"CHAR", "NIL", "ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT",
		"Op_LE", "Op_GE", "Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN",
		"Op_MINUS_ASSIGN", "Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS",
		"Op_MOD", "Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"idChain", "program", "block", "statement", "variableType", "variableDeclaration",
		"functionDeclarationStatement", "functionParameters", "functionParameter",
		"functionReturnType", "functionCall", "functionCallArguments", "variableAssignment",
		"ifStatement", "ifTail", "elseStatement", "whileStatement", "switchStatement",
		"switchCase", "switchDefault", "forStatement", "guardStatement", "vectorDeclaration",
		"vectorDefinition", "vectorValues", "vectorAccess", "vectorAssignment",
		"expr", "controlFlowStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 359, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 5, 0, 62, 8, 0, 10,
		0, 12, 0, 65, 9, 0, 1, 1, 1, 1, 1, 1, 1, 2, 5, 2, 71, 8, 2, 10, 2, 12,
		2, 74, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 88, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 99, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 106, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 114, 8, 5, 3, 5, 116, 8, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 3, 6, 122, 8, 6, 1, 6, 1, 6, 3, 6, 126, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 135, 8, 7, 10, 7, 12, 7, 138,
		9, 7, 1, 8, 3, 8, 141, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 146, 8, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 155, 8, 10, 1, 10, 1, 10, 3,
		10, 159, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 166, 8, 11, 10,
		11, 12, 11, 169, 9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 178, 8, 11, 10, 11, 12, 11, 181, 9, 11, 3, 11, 183, 8, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 192, 8, 13, 10, 13, 12,
		13, 195, 9, 13, 1, 13, 3, 13, 198, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 221, 8, 17, 10, 17, 12,
		17, 224, 9, 17, 1, 17, 3, 17, 227, 8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 236, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19,
		242, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 270, 8, 23, 1,
		23, 1, 23, 3, 23, 274, 8, 23, 1, 24, 1, 24, 1, 24, 5, 24, 279, 8, 24, 10,
		24, 12, 24, 282, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27,
		311, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 5, 27, 346, 8, 27, 10, 27, 12, 27, 349, 9, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 355, 8, 28, 3, 28, 357, 8, 28, 1, 28,
		0, 1, 54, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 0, 8, 1, 0, 22, 26,
		1, 0, 4, 5, 2, 0, 43, 43, 46, 47, 1, 0, 49, 50, 1, 0, 51, 52, 2, 0, 40,
		40, 42, 42, 2, 0, 39, 39, 41, 41, 1, 0, 37, 38, 390, 0, 58, 1, 0, 0, 0,
		2, 66, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 89, 1, 0, 0,
		0, 10, 115, 1, 0, 0, 0, 12, 117, 1, 0, 0, 0, 14, 131, 1, 0, 0, 0, 16, 140,
		1, 0, 0, 0, 18, 149, 1, 0, 0, 0, 20, 154, 1, 0, 0, 0, 22, 182, 1, 0, 0,
		0, 24, 184, 1, 0, 0, 0, 26, 188, 1, 0, 0, 0, 28, 199, 1, 0, 0, 0, 30, 205,
		1, 0, 0, 0, 32, 210, 1, 0, 0, 0, 34, 216, 1, 0, 0, 0, 36, 230, 1, 0, 0,
		0, 38, 237, 1, 0, 0, 0, 40, 243, 1, 0, 0, 0, 42, 251, 1, 0, 0, 0, 44, 258,
		1, 0, 0, 0, 46, 273, 1, 0, 0, 0, 48, 275, 1, 0, 0, 0, 50, 283, 1, 0, 0,
		0, 52, 288, 1, 0, 0, 0, 54, 310, 1, 0, 0, 0, 56, 356, 1, 0, 0, 0, 58, 63,
		5, 35, 0, 0, 59, 60, 5, 68, 0, 0, 60, 62, 5, 35, 0, 0, 61, 59, 1, 0, 0,
		0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 1, 1,
		0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 67, 3, 4, 2, 0, 67, 68, 5, 0, 0, 1, 68,
		3, 1, 0, 0, 0, 69, 71, 3, 6, 3, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0,
		0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 5, 1, 0, 0, 0, 74, 72, 1,
		0, 0, 0, 75, 88, 3, 24, 12, 0, 76, 88, 3, 10, 5, 0, 77, 88, 3, 26, 13,
		0, 78, 88, 3, 32, 16, 0, 79, 88, 3, 34, 17, 0, 80, 88, 3, 40, 20, 0, 81,
		88, 3, 42, 21, 0, 82, 88, 3, 56, 28, 0, 83, 88, 3, 12, 6, 0, 84, 88, 3,
		20, 10, 0, 85, 88, 3, 44, 22, 0, 86, 88, 3, 52, 26, 0, 87, 75, 1, 0, 0,
		0, 87, 76, 1, 0, 0, 0, 87, 77, 1, 0, 0, 0, 87, 78, 1, 0, 0, 0, 87, 79,
		1, 0, 0, 0, 87, 80, 1, 0, 0, 0, 87, 81, 1, 0, 0, 0, 87, 82, 1, 0, 0, 0,
		87, 83, 1, 0, 0, 0, 87, 84, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 86, 1,
		0, 0, 0, 88, 7, 1, 0, 0, 0, 89, 90, 7, 0, 0, 0, 90, 9, 1, 0, 0, 0, 91,
		92, 7, 1, 0, 0, 92, 93, 5, 35, 0, 0, 93, 94, 5, 67, 0, 0, 94, 95, 3, 8,
		4, 0, 95, 96, 5, 43, 0, 0, 96, 98, 3, 54, 27, 0, 97, 99, 5, 66, 0, 0, 98,
		97, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 116, 1, 0, 0, 0, 100, 101, 7, 1,
		0, 0, 101, 102, 5, 35, 0, 0, 102, 103, 5, 43, 0, 0, 103, 105, 3, 54, 27,
		0, 104, 106, 5, 66, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		116, 1, 0, 0, 0, 107, 108, 7, 1, 0, 0, 108, 109, 5, 35, 0, 0, 109, 110,
		5, 67, 0, 0, 110, 111, 3, 8, 4, 0, 111, 113, 5, 57, 0, 0, 112, 114, 5,
		66, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 116, 1, 0, 0,
		0, 115, 91, 1, 0, 0, 0, 115, 100, 1, 0, 0, 0, 115, 107, 1, 0, 0, 0, 116,
		11, 1, 0, 0, 0, 117, 118, 5, 6, 0, 0, 118, 119, 5, 35, 0, 0, 119, 121,
		5, 58, 0, 0, 120, 122, 3, 14, 7, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1,
		0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 5, 59, 0, 0, 124, 126, 3, 18,
		9, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0,
		127, 128, 5, 60, 0, 0, 128, 129, 3, 4, 2, 0, 129, 130, 5, 61, 0, 0, 130,
		13, 1, 0, 0, 0, 131, 136, 3, 16, 8, 0, 132, 133, 5, 65, 0, 0, 133, 135,
		3, 16, 8, 0, 134, 132, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0,
		0, 0, 136, 137, 1, 0, 0, 0, 137, 15, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0,
		139, 141, 5, 35, 0, 0, 140, 139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 143, 5, 35, 0, 0, 143, 145, 5, 67, 0, 0, 144, 146,
		5, 28, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 148, 3, 8, 4, 0, 148, 17, 1, 0, 0, 0, 149, 150, 5, 36, 0, 0,
		150, 151, 3, 8, 4, 0, 151, 19, 1, 0, 0, 0, 152, 155, 3, 0, 0, 0, 153, 155,
		3, 8, 4, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 158, 5, 58, 0, 0, 157, 159, 3, 22, 11, 0, 158, 157, 1, 0, 0,
		0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 5, 59, 0, 0, 161,
		21, 1, 0, 0, 0, 162, 167, 3, 54, 27, 0, 163, 164, 5, 65, 0, 0, 164, 166,
		3, 54, 27, 0, 165, 163, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1,
		0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 183, 1, 0, 0, 0, 169, 167, 1, 0, 0,
		0, 170, 171, 5, 35, 0, 0, 171, 172, 5, 67, 0, 0, 172, 179, 3, 54, 27, 0,
		173, 174, 5, 65, 0, 0, 174, 175, 5, 35, 0, 0, 175, 176, 5, 67, 0, 0, 176,
		178, 3, 54, 27, 0, 177, 173, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0,
		0, 0, 182, 162, 1, 0, 0, 0, 182, 170, 1, 0, 0, 0, 183, 23, 1, 0, 0, 0,
		184, 185, 5, 35, 0, 0, 185, 186, 7, 2, 0, 0, 186, 187, 3, 54, 27, 0, 187,
		25, 1, 0, 0, 0, 188, 193, 3, 28, 14, 0, 189, 190, 5, 9, 0, 0, 190, 192,
		3, 28, 14, 0, 191, 189, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1,
		0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0,
		0, 196, 198, 3, 30, 15, 0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0,
		198, 27, 1, 0, 0, 0, 199, 200, 5, 8, 0, 0, 200, 201, 3, 54, 27, 0, 201,
		202, 5, 60, 0, 0, 202, 203, 3, 4, 2, 0, 203, 204, 5, 61, 0, 0, 204, 29,
		1, 0, 0, 0, 205, 206, 5, 9, 0, 0, 206, 207, 5, 60, 0, 0, 207, 208, 3, 4,
		2, 0, 208, 209, 5, 61, 0, 0, 209, 31, 1, 0, 0, 0, 210, 211, 5, 14, 0, 0,
		211, 212, 3, 54, 27, 0, 212, 213, 5, 60, 0, 0, 213, 214, 3, 4, 2, 0, 214,
		215, 5, 61, 0, 0, 215, 33, 1, 0, 0, 0, 216, 217, 5, 10, 0, 0, 217, 218,
		3, 54, 27, 0, 218, 222, 5, 60, 0, 0, 219, 221, 3, 36, 18, 0, 220, 219,
		1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0,
		0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 227, 3, 38, 19,
		0, 226, 225, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		229, 5, 61, 0, 0, 229, 35, 1, 0, 0, 0, 230, 231, 5, 11, 0, 0, 231, 232,
		3, 54, 27, 0, 232, 233, 5, 67, 0, 0, 233, 235, 3, 4, 2, 0, 234, 236, 5,
		15, 0, 0, 235, 234, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 37, 1, 0, 0,
		0, 237, 238, 5, 12, 0, 0, 238, 239, 5, 67, 0, 0, 239, 241, 3, 4, 2, 0,
		240, 242, 5, 15, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		39, 1, 0, 0, 0, 243, 244, 5, 13, 0, 0, 244, 245, 5, 35, 0, 0, 245, 246,
		5, 20, 0, 0, 246, 247, 3, 54, 27, 0, 247, 248, 5, 60, 0, 0, 248, 249, 3,
		4, 2, 0, 249, 250, 5, 61, 0, 0, 250, 41, 1, 0, 0, 0, 251, 252, 5, 21, 0,
		0, 252, 253, 3, 54, 27, 0, 253, 254, 5, 9, 0, 0, 254, 255, 5, 60, 0, 0,
		255, 256, 3, 4, 2, 0, 256, 257, 5, 61, 0, 0, 257, 43, 1, 0, 0, 0, 258,
		259, 7, 1, 0, 0, 259, 260, 5, 35, 0, 0, 260, 261, 5, 67, 0, 0, 261, 262,
		5, 62, 0, 0, 262, 263, 3, 8, 4, 0, 263, 264, 5, 63, 0, 0, 264, 265, 5,
		43, 0, 0, 265, 266, 3, 46, 23, 0, 266, 45, 1, 0, 0, 0, 267, 269, 5, 62,
		0, 0, 268, 270, 3, 48, 24, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1, 0, 0,
		0, 270, 271, 1, 0, 0, 0, 271, 274, 5, 63, 0, 0, 272, 274, 3, 54, 27, 0,
		273, 267, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 47, 1, 0, 0, 0, 275, 280,
		3, 54, 27, 0, 276, 277, 5, 65, 0, 0, 277, 279, 3, 54, 27, 0, 278, 276,
		1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0,
		0, 0, 281, 49, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 284, 3, 0, 0, 0,
		284, 285, 5, 62, 0, 0, 285, 286, 3, 54, 27, 0, 286, 287, 5, 63, 0, 0, 287,
		51, 1, 0, 0, 0, 288, 289, 3, 50, 25, 0, 289, 290, 7, 2, 0, 0, 290, 291,
		3, 54, 27, 0, 291, 53, 1, 0, 0, 0, 292, 293, 6, 27, -1, 0, 293, 311, 3,
		20, 10, 0, 294, 311, 3, 50, 25, 0, 295, 296, 5, 52, 0, 0, 296, 311, 3,
		54, 27, 20, 297, 298, 5, 56, 0, 0, 298, 311, 3, 54, 27, 19, 299, 300, 5,
		58, 0, 0, 300, 301, 3, 54, 27, 0, 301, 302, 5, 59, 0, 0, 302, 311, 1, 0,
		0, 0, 303, 311, 5, 29, 0, 0, 304, 311, 3, 0, 0, 0, 305, 311, 5, 30, 0,
		0, 306, 311, 5, 32, 0, 0, 307, 311, 5, 34, 0, 0, 308, 311, 5, 33, 0, 0,
		309, 311, 5, 31, 0, 0, 310, 292, 1, 0, 0, 0, 310, 294, 1, 0, 0, 0, 310,
		295, 1, 0, 0, 0, 310, 297, 1, 0, 0, 0, 310, 299, 1, 0, 0, 0, 310, 303,
		1, 0, 0, 0, 310, 304, 1, 0, 0, 0, 310, 305, 1, 0, 0, 0, 310, 306, 1, 0,
		0, 0, 310, 307, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 309, 1, 0, 0, 0,
		311, 347, 1, 0, 0, 0, 312, 313, 10, 18, 0, 0, 313, 314, 7, 3, 0, 0, 314,
		346, 3, 54, 27, 19, 315, 316, 10, 17, 0, 0, 316, 317, 7, 4, 0, 0, 317,
		346, 3, 54, 27, 18, 318, 319, 10, 16, 0, 0, 319, 320, 5, 53, 0, 0, 320,
		346, 3, 54, 27, 17, 321, 322, 10, 15, 0, 0, 322, 323, 7, 5, 0, 0, 323,
		346, 3, 54, 27, 16, 324, 325, 10, 14, 0, 0, 325, 326, 7, 6, 0, 0, 326,
		346, 3, 54, 27, 15, 327, 328, 10, 13, 0, 0, 328, 329, 7, 7, 0, 0, 329,
		346, 3, 54, 27, 14, 330, 331, 10, 12, 0, 0, 331, 332, 5, 57, 0, 0, 332,
		333, 3, 54, 27, 0, 333, 334, 5, 67, 0, 0, 334, 335, 3, 54, 27, 13, 335,
		346, 1, 0, 0, 0, 336, 337, 10, 11, 0, 0, 337, 338, 5, 54, 0, 0, 338, 346,
		3, 54, 27, 12, 339, 340, 10, 10, 0, 0, 340, 341, 5, 55, 0, 0, 341, 346,
		3, 54, 27, 11, 342, 343, 10, 9, 0, 0, 343, 344, 5, 27, 0, 0, 344, 346,
		3, 54, 27, 10, 345, 312, 1, 0, 0, 0, 345, 315, 1, 0, 0, 0, 345, 318, 1,
		0, 0, 0, 345, 321, 1, 0, 0, 0, 345, 324, 1, 0, 0, 0, 345, 327, 1, 0, 0,
		0, 345, 330, 1, 0, 0, 0, 345, 336, 1, 0, 0, 0, 345, 339, 1, 0, 0, 0, 345,
		342, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 347, 348,
		1, 0, 0, 0, 348, 55, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 350, 357, 5, 15,
		0, 0, 351, 357, 5, 16, 0, 0, 352, 354, 5, 17, 0, 0, 353, 355, 3, 54, 27,
		0, 354, 353, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 357, 1, 0, 0, 0, 356,
		350, 1, 0, 0, 0, 356, 351, 1, 0, 0, 0, 356, 352, 1, 0, 0, 0, 357, 57, 1,
		0, 0, 0, 31, 63, 72, 87, 98, 105, 113, 115, 121, 125, 136, 140, 145, 154,
		158, 167, 179, 182, 193, 197, 222, 226, 235, 241, 269, 273, 280, 310, 345,
		347, 354, 356,
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
	SwiftParserKw_INOUT        = 28
	SwiftParserINT             = 29
	SwiftParserFLOAT           = 30
	SwiftParserBOOL            = 31
	SwiftParserSTRING          = 32
	SwiftParserCHAR            = 33
	SwiftParserNIL             = 34
	SwiftParserID              = 35
	SwiftParserOp_ARROW        = 36
	SwiftParserOp_EQ           = 37
	SwiftParserOp_NEQ          = 38
	SwiftParserOp_LT           = 39
	SwiftParserOp_GT           = 40
	SwiftParserOp_LE           = 41
	SwiftParserOp_GE           = 42
	SwiftParserOp_ASSIGN       = 43
	SwiftParserOp_MUL_ASSIGN   = 44
	SwiftParserOp_DIV_ASSIGN   = 45
	SwiftParserOp_PLUS_ASSIGN  = 46
	SwiftParserOp_MINUS_ASSIGN = 47
	SwiftParserOp_MOD_ASSIGN   = 48
	SwiftParserOp_MUL          = 49
	SwiftParserOp_DIV          = 50
	SwiftParserOp_PLUS         = 51
	SwiftParserOp_MINUS        = 52
	SwiftParserOp_MOD          = 53
	SwiftParserOp_AND          = 54
	SwiftParserOp_OR           = 55
	SwiftParserOp_NOT          = 56
	SwiftParserOp_TERNARY      = 57
	SwiftParserLPAREN          = 58
	SwiftParserRPAREN          = 59
	SwiftParserLBRACE          = 60
	SwiftParserRBRACE          = 61
	SwiftParserLBRACKET        = 62
	SwiftParserRBRACKET        = 63
	SwiftParserBACKSLASH       = 64
	SwiftParserCOMMA           = 65
	SwiftParserSEMICOLON       = 66
	SwiftParserCOLON           = 67
	SwiftParserDOT             = 68
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
	SwiftParserRULE_expr                         = 27
	SwiftParserRULE_controlFlowStatement         = 28
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
		p.SetState(58)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
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
				p.SetState(59)
				p.Match(SwiftParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(60)
				p.Match(SwiftParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(65)
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
		p.SetState(66)
		p.Block()
	}
	{
		p.SetState(67)
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
	p.SetState(72)
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
				p.SetState(69)
				p.Statement()
			}

		}
		p.SetState(74)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.VariableAssignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.VariableDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			p.WhileStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(79)
			p.SwitchStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(80)
			p.ForStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(81)
			p.GuardStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(82)
			p.ControlFlowStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(83)
			p.FunctionDeclarationStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(84)
			p.FunctionCall()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(85)
			p.VectorDeclaration()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(86)
			p.VectorAssignment()
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
		p.SetState(89)
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
	p.EnterRule(localctx, 10, SwiftParserRULE_variableDeclaration)
	var _la int

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)

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
			p.SetState(92)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.VariableType()
		}
		{
			p.SetState(95)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.expr(0)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(97)
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
			p.SetState(100)

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
			p.SetState(101)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.expr(0)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(104)
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
			p.SetState(107)

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
			p.SetState(108)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.VariableType()
		}
		{
			p.SetState(111)
			p.Match(SwiftParserOp_TERNARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(112)
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
	p.EnterRule(localctx, 12, SwiftParserRULE_functionDeclarationStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(SwiftParserKw_FUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(SwiftParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserID {
		{
			p.SetState(120)
			p.FunctionParameters()
		}

	}
	{
		p.SetState(123)
		p.Match(SwiftParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserOp_ARROW {
		{
			p.SetState(124)
			p.FunctionReturnType()
		}

	}
	{
		p.SetState(127)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Block()
	}
	{
		p.SetState(129)
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
		p.SetState(131)
		p.FunctionParameter()
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
			p.FunctionParameter()
		}

		p.SetState(138)
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
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(139)
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
		p.SetState(142)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_INOUT {
		{
			p.SetState(144)
			p.Match(SwiftParserKw_INOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(147)
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
		p.SetState(149)
		p.Match(SwiftParserOp_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserID:
		{
			p.SetState(152)
			p.IdChain()
		}

	case SwiftParserKw_INT, SwiftParserKw_FLOAT, SwiftParserKw_BOOL, SwiftParserKw_STRING, SwiftParserKw_CHAR:
		{
			p.SetState(153)
			p.VariableType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(156)
		p.Match(SwiftParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&364791638129639424) != 0 {
		{
			p.SetState(157)
			p.FunctionCallArguments()
		}

	}
	{
		p.SetState(160)
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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArgumentsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.expr(0)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftParserCOMMA {
			{
				p.SetState(163)
				p.Match(SwiftParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(164)
				p.expr(0)
			}

			p.SetState(169)
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
			p.SetState(170)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.expr(0)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SwiftParserCOMMA {
			{
				p.SetState(173)
				p.Match(SwiftParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(174)
				p.Match(SwiftParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(175)
				p.Match(SwiftParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(176)
				p.expr(0)
			}

			p.SetState(181)
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
		p.SetState(184)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&219902325555200) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(186)
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
		p.SetState(188)
		p.IfTail()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(189)
				p.Match(SwiftParserKw_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.IfTail()
			}

		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_ELSE {
		{
			p.SetState(196)
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
		p.SetState(199)
		p.Match(SwiftParserKw_IF)
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
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Block()
	}
	{
		p.SetState(203)
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
		p.SetState(205)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Block()
	}
	{
		p.SetState(208)
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
		p.SetState(210)
		p.Match(SwiftParserKw_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.expr(0)
	}
	{
		p.SetState(212)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Block()
	}
	{
		p.SetState(214)
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
		p.SetState(216)
		p.Match(SwiftParserKw_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.expr(0)
	}
	{
		p.SetState(218)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserKw_CASE {
		{
			p.SetState(219)
			p.SwitchCase()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_DEFAULT {
		{
			p.SetState(225)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(228)
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
		p.SetState(230)
		p.Match(SwiftParserKw_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.expr(0)
	}
	{
		p.SetState(232)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Block()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(234)
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
		p.SetState(237)
		p.Match(SwiftParserKw_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Block()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(240)
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
		p.SetState(243)
		p.Match(SwiftParserKw_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(SwiftParserKw_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.expr(0)
	}
	{
		p.SetState(247)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Block()
	}
	{
		p.SetState(249)
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
		p.SetState(251)
		p.Match(SwiftParserKw_GUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.expr(0)
	}
	{
		p.SetState(253)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
		p.Block()
	}
	{
		p.SetState(256)
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
		p.SetState(258)

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
		p.SetState(259)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(SwiftParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.VariableType()
	}
	{
		p.SetState(263)
		p.Match(SwiftParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.Match(SwiftParserOp_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserLBRACKET:
		localctx = NewVectorListValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(267)
			p.Match(SwiftParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&364791638129639424) != 0 {
			{
				p.SetState(268)
				p.VectorValues()
			}

		}
		{
			p.SetState(271)
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
			p.SetState(272)
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
		p.SetState(275)
		p.expr(0)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserCOMMA {
		{
			p.SetState(276)
			p.Match(SwiftParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.expr(0)
		}

		p.SetState(282)
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
		p.SetState(283)
		p.IdChain()
	}
	{
		p.SetState(284)
		p.Match(SwiftParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.expr(0)
	}
	{
		p.SetState(286)
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
		p.SetState(288)
		p.VectorAccess()
	}
	{
		p.SetState(289)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VectorAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&219902325555200) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VectorAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(290)
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

func (p *SwiftParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, SwiftParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(293)
			p.FunctionCall()
		}

	case 2:
		localctx = NewVectorAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(294)
			p.VectorAccess()
		}

	case 3:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(295)
			p.Match(SwiftParserOp_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.expr(20)
		}

	case 4:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(297)
			p.Match(SwiftParserOp_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)

			var _x = p.expr(19)

			localctx.(*NotExprContext).right = _x
		}

	case 5:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(299)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.expr(0)
		}
		{
			p.SetState(301)
			p.Match(SwiftParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(303)
			p.Match(SwiftParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(304)
			p.IdChain()
		}

	case 8:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(305)
			p.Match(SwiftParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(306)
			p.Match(SwiftParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)
			p.Match(SwiftParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(308)
			p.Match(SwiftParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(309)
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
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(312)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(313)

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
					p.SetState(314)

					var _x = p.expr(19)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(316)

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
					p.SetState(317)

					var _x = p.expr(18)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 3:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(318)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(319)

					var _m = p.Match(SwiftParserOp_MOD)

					localctx.(*ArithmeticExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(320)

					var _x = p.expr(17)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 4:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(321)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(322)

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
					p.SetState(323)

					var _x = p.expr(16)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 5:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(325)

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
					p.SetState(326)

					var _x = p.expr(15)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 6:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(328)

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
					p.SetState(329)

					var _x = p.expr(14)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(330)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(331)
					p.Match(SwiftParserOp_TERNARY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(332)

					var _x = p.expr(0)

					localctx.(*TernaryExprContext).cTrue = _x
				}
				{
					p.SetState(333)
					p.Match(SwiftParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(334)

					var _x = p.expr(13)

					localctx.(*TernaryExprContext).cFalse = _x
				}

			case 8:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(337)

					var _m = p.Match(SwiftParserOp_AND)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(338)

					var _x = p.expr(12)

					localctx.(*LogicalExprContext).right = _x
				}

			case 9:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(340)

					var _m = p.Match(SwiftParserOp_OR)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(341)

					var _x = p.expr(11)

					localctx.(*LogicalExprContext).right = _x
				}

			case 10:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RangeExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(343)
					p.Match(SwiftParserKw_RANGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(344)

					var _x = p.expr(10)

					localctx.(*RangeExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 56, SwiftParserRULE_controlFlowStatement)
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserKw_BREAK:
		localctx = NewControlBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
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
			p.SetState(351)
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
			p.SetState(352)
			p.Match(SwiftParserKw_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(353)
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
	case 27:
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
