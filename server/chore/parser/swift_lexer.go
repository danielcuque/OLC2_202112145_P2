// Code generated from SwiftLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SwiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftlexerLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
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
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 70, 461, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 1, 0, 4, 0, 143, 8, 0, 11, 0, 12, 0, 144, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 153, 8, 1, 10, 1, 12, 1, 156, 9,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 164, 8, 2, 10, 2, 12, 2, 167,
		9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 4, 30, 327, 8, 30, 11,
		30, 12, 30, 328, 1, 31, 5, 31, 332, 8, 31, 10, 31, 12, 31, 335, 9, 31,
		1, 31, 1, 31, 4, 31, 339, 8, 31, 11, 31, 12, 31, 340, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 352, 8, 32, 1, 33,
		1, 33, 1, 33, 1, 33, 5, 33, 358, 8, 33, 10, 33, 12, 33, 361, 9, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 369, 8, 34, 1, 34, 1, 34,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 5, 36, 379, 8, 36, 10, 36, 12,
		36, 382, 9, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56,
		1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1,
		61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66,
		1, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 165, 0, 70, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30,
		61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39,
		79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48,
		97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113,
		57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64, 129,
		65, 131, 66, 133, 67, 135, 68, 137, 69, 139, 70, 1, 0, 8, 4, 0, 9, 10,
		13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 5, 0, 34, 34,
		92, 92, 110, 110, 114, 114, 116, 116, 3, 0, 10, 10, 13, 13, 34, 34, 3,
		0, 10, 10, 13, 13, 39, 39, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 471, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1,
		0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0,
		105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0,
		0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119,
		1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0,
		0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1,
		0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0, 0, 0, 1,
		142, 1, 0, 0, 0, 3, 148, 1, 0, 0, 0, 5, 159, 1, 0, 0, 0, 7, 173, 1, 0,
		0, 0, 9, 177, 1, 0, 0, 0, 11, 181, 1, 0, 0, 0, 13, 186, 1, 0, 0, 0, 15,
		193, 1, 0, 0, 0, 17, 202, 1, 0, 0, 0, 19, 204, 1, 0, 0, 0, 21, 207, 1,
		0, 0, 0, 23, 212, 1, 0, 0, 0, 25, 219, 1, 0, 0, 0, 27, 224, 1, 0, 0, 0,
		29, 232, 1, 0, 0, 0, 31, 236, 1, 0, 0, 0, 33, 242, 1, 0, 0, 0, 35, 248,
		1, 0, 0, 0, 37, 257, 1, 0, 0, 0, 39, 264, 1, 0, 0, 0, 41, 267, 1, 0, 0,
		0, 43, 274, 1, 0, 0, 0, 45, 277, 1, 0, 0, 0, 47, 283, 1, 0, 0, 0, 49, 287,
		1, 0, 0, 0, 51, 293, 1, 0, 0, 0, 53, 298, 1, 0, 0, 0, 55, 305, 1, 0, 0,
		0, 57, 315, 1, 0, 0, 0, 59, 319, 1, 0, 0, 0, 61, 326, 1, 0, 0, 0, 63, 333,
		1, 0, 0, 0, 65, 351, 1, 0, 0, 0, 67, 353, 1, 0, 0, 0, 69, 364, 1, 0, 0,
		0, 71, 372, 1, 0, 0, 0, 73, 376, 1, 0, 0, 0, 75, 383, 1, 0, 0, 0, 77, 386,
		1, 0, 0, 0, 79, 389, 1, 0, 0, 0, 81, 392, 1, 0, 0, 0, 83, 394, 1, 0, 0,
		0, 85, 396, 1, 0, 0, 0, 87, 399, 1, 0, 0, 0, 89, 402, 1, 0, 0, 0, 91, 404,
		1, 0, 0, 0, 93, 407, 1, 0, 0, 0, 95, 410, 1, 0, 0, 0, 97, 413, 1, 0, 0,
		0, 99, 416, 1, 0, 0, 0, 101, 419, 1, 0, 0, 0, 103, 421, 1, 0, 0, 0, 105,
		423, 1, 0, 0, 0, 107, 425, 1, 0, 0, 0, 109, 427, 1, 0, 0, 0, 111, 429,
		1, 0, 0, 0, 113, 432, 1, 0, 0, 0, 115, 435, 1, 0, 0, 0, 117, 437, 1, 0,
		0, 0, 119, 439, 1, 0, 0, 0, 121, 441, 1, 0, 0, 0, 123, 443, 1, 0, 0, 0,
		125, 445, 1, 0, 0, 0, 127, 447, 1, 0, 0, 0, 129, 449, 1, 0, 0, 0, 131,
		451, 1, 0, 0, 0, 133, 453, 1, 0, 0, 0, 135, 455, 1, 0, 0, 0, 137, 457,
		1, 0, 0, 0, 139, 459, 1, 0, 0, 0, 141, 143, 7, 0, 0, 0, 142, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0,
		145, 146, 1, 0, 0, 0, 146, 147, 6, 0, 0, 0, 147, 2, 1, 0, 0, 0, 148, 149,
		5, 47, 0, 0, 149, 150, 5, 47, 0, 0, 150, 154, 1, 0, 0, 0, 151, 153, 8,
		1, 0, 0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0,
		0, 154, 155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157,
		158, 6, 1, 0, 0, 158, 4, 1, 0, 0, 0, 159, 160, 5, 47, 0, 0, 160, 161, 5,
		42, 0, 0, 161, 165, 1, 0, 0, 0, 162, 164, 9, 0, 0, 0, 163, 162, 1, 0, 0,
		0, 164, 167, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166,
		168, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 5, 42, 0, 0, 169, 170,
		5, 47, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 6, 2, 0, 0, 172, 6, 1, 0,
		0, 0, 173, 174, 5, 108, 0, 0, 174, 175, 5, 101, 0, 0, 175, 176, 5, 116,
		0, 0, 176, 8, 1, 0, 0, 0, 177, 178, 5, 118, 0, 0, 178, 179, 5, 97, 0, 0,
		179, 180, 5, 114, 0, 0, 180, 10, 1, 0, 0, 0, 181, 182, 5, 102, 0, 0, 182,
		183, 5, 117, 0, 0, 183, 184, 5, 110, 0, 0, 184, 185, 5, 99, 0, 0, 185,
		12, 1, 0, 0, 0, 186, 187, 5, 115, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189,
		5, 114, 0, 0, 189, 190, 5, 117, 0, 0, 190, 191, 5, 99, 0, 0, 191, 192,
		5, 116, 0, 0, 192, 14, 1, 0, 0, 0, 193, 194, 5, 109, 0, 0, 194, 195, 5,
		117, 0, 0, 195, 196, 5, 116, 0, 0, 196, 197, 5, 97, 0, 0, 197, 198, 5,
		116, 0, 0, 198, 199, 5, 105, 0, 0, 199, 200, 5, 110, 0, 0, 200, 201, 5,
		103, 0, 0, 201, 16, 1, 0, 0, 0, 202, 203, 5, 38, 0, 0, 203, 18, 1, 0, 0,
		0, 204, 205, 5, 105, 0, 0, 205, 206, 5, 102, 0, 0, 206, 20, 1, 0, 0, 0,
		207, 208, 5, 101, 0, 0, 208, 209, 5, 108, 0, 0, 209, 210, 5, 115, 0, 0,
		210, 211, 5, 101, 0, 0, 211, 22, 1, 0, 0, 0, 212, 213, 5, 115, 0, 0, 213,
		214, 5, 119, 0, 0, 214, 215, 5, 105, 0, 0, 215, 216, 5, 116, 0, 0, 216,
		217, 5, 99, 0, 0, 217, 218, 5, 104, 0, 0, 218, 24, 1, 0, 0, 0, 219, 220,
		5, 99, 0, 0, 220, 221, 5, 97, 0, 0, 221, 222, 5, 115, 0, 0, 222, 223, 5,
		101, 0, 0, 223, 26, 1, 0, 0, 0, 224, 225, 5, 100, 0, 0, 225, 226, 5, 101,
		0, 0, 226, 227, 5, 102, 0, 0, 227, 228, 5, 97, 0, 0, 228, 229, 5, 117,
		0, 0, 229, 230, 5, 108, 0, 0, 230, 231, 5, 116, 0, 0, 231, 28, 1, 0, 0,
		0, 232, 233, 5, 102, 0, 0, 233, 234, 5, 111, 0, 0, 234, 235, 5, 114, 0,
		0, 235, 30, 1, 0, 0, 0, 236, 237, 5, 119, 0, 0, 237, 238, 5, 104, 0, 0,
		238, 239, 5, 105, 0, 0, 239, 240, 5, 108, 0, 0, 240, 241, 5, 101, 0, 0,
		241, 32, 1, 0, 0, 0, 242, 243, 5, 98, 0, 0, 243, 244, 5, 114, 0, 0, 244,
		245, 5, 101, 0, 0, 245, 246, 5, 97, 0, 0, 246, 247, 5, 107, 0, 0, 247,
		34, 1, 0, 0, 0, 248, 249, 5, 99, 0, 0, 249, 250, 5, 111, 0, 0, 250, 251,
		5, 110, 0, 0, 251, 252, 5, 116, 0, 0, 252, 253, 5, 105, 0, 0, 253, 254,
		5, 110, 0, 0, 254, 255, 5, 117, 0, 0, 255, 256, 5, 101, 0, 0, 256, 36,
		1, 0, 0, 0, 257, 258, 5, 114, 0, 0, 258, 259, 5, 101, 0, 0, 259, 260, 5,
		116, 0, 0, 260, 261, 5, 117, 0, 0, 261, 262, 5, 114, 0, 0, 262, 263, 5,
		110, 0, 0, 263, 38, 1, 0, 0, 0, 264, 265, 5, 100, 0, 0, 265, 266, 5, 111,
		0, 0, 266, 40, 1, 0, 0, 0, 267, 268, 5, 114, 0, 0, 268, 269, 5, 101, 0,
		0, 269, 270, 5, 112, 0, 0, 270, 271, 5, 101, 0, 0, 271, 272, 5, 97, 0,
		0, 272, 273, 5, 116, 0, 0, 273, 42, 1, 0, 0, 0, 274, 275, 5, 105, 0, 0,
		275, 276, 5, 110, 0, 0, 276, 44, 1, 0, 0, 0, 277, 278, 5, 103, 0, 0, 278,
		279, 5, 117, 0, 0, 279, 280, 5, 97, 0, 0, 280, 281, 5, 114, 0, 0, 281,
		282, 5, 100, 0, 0, 282, 46, 1, 0, 0, 0, 283, 284, 5, 73, 0, 0, 284, 285,
		5, 110, 0, 0, 285, 286, 5, 116, 0, 0, 286, 48, 1, 0, 0, 0, 287, 288, 5,
		70, 0, 0, 288, 289, 5, 108, 0, 0, 289, 290, 5, 111, 0, 0, 290, 291, 5,
		97, 0, 0, 291, 292, 5, 116, 0, 0, 292, 50, 1, 0, 0, 0, 293, 294, 5, 66,
		0, 0, 294, 295, 5, 111, 0, 0, 295, 296, 5, 111, 0, 0, 296, 297, 5, 108,
		0, 0, 297, 52, 1, 0, 0, 0, 298, 299, 5, 83, 0, 0, 299, 300, 5, 116, 0,
		0, 300, 301, 5, 114, 0, 0, 301, 302, 5, 105, 0, 0, 302, 303, 5, 110, 0,
		0, 303, 304, 5, 103, 0, 0, 304, 54, 1, 0, 0, 0, 305, 306, 5, 67, 0, 0,
		306, 307, 5, 104, 0, 0, 307, 308, 5, 97, 0, 0, 308, 309, 5, 114, 0, 0,
		309, 310, 5, 97, 0, 0, 310, 311, 5, 99, 0, 0, 311, 312, 5, 116, 0, 0, 312,
		313, 5, 101, 0, 0, 313, 314, 5, 114, 0, 0, 314, 56, 1, 0, 0, 0, 315, 316,
		5, 46, 0, 0, 316, 317, 5, 46, 0, 0, 317, 318, 5, 46, 0, 0, 318, 58, 1,
		0, 0, 0, 319, 320, 5, 105, 0, 0, 320, 321, 5, 110, 0, 0, 321, 322, 5, 111,
		0, 0, 322, 323, 5, 117, 0, 0, 323, 324, 5, 116, 0, 0, 324, 60, 1, 0, 0,
		0, 325, 327, 7, 2, 0, 0, 326, 325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328,
		326, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 62, 1, 0, 0, 0, 330, 332, 7,
		2, 0, 0, 331, 330, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0,
		0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336,
		338, 5, 46, 0, 0, 337, 339, 7, 2, 0, 0, 338, 337, 1, 0, 0, 0, 339, 340,
		1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 64, 1, 0,
		0, 0, 342, 343, 5, 116, 0, 0, 343, 344, 5, 114, 0, 0, 344, 345, 5, 117,
		0, 0, 345, 352, 5, 101, 0, 0, 346, 347, 5, 102, 0, 0, 347, 348, 5, 97,
		0, 0, 348, 349, 5, 108, 0, 0, 349, 350, 5, 115, 0, 0, 350, 352, 5, 101,
		0, 0, 351, 342, 1, 0, 0, 0, 351, 346, 1, 0, 0, 0, 352, 66, 1, 0, 0, 0,
		353, 359, 5, 34, 0, 0, 354, 355, 5, 92, 0, 0, 355, 358, 7, 3, 0, 0, 356,
		358, 8, 4, 0, 0, 357, 354, 1, 0, 0, 0, 357, 356, 1, 0, 0, 0, 358, 361,
		1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 362, 1, 0,
		0, 0, 361, 359, 1, 0, 0, 0, 362, 363, 5, 34, 0, 0, 363, 68, 1, 0, 0, 0,
		364, 368, 5, 39, 0, 0, 365, 366, 5, 92, 0, 0, 366, 369, 7, 3, 0, 0, 367,
		369, 8, 5, 0, 0, 368, 365, 1, 0, 0, 0, 368, 367, 1, 0, 0, 0, 369, 370,
		1, 0, 0, 0, 370, 371, 5, 39, 0, 0, 371, 70, 1, 0, 0, 0, 372, 373, 5, 110,
		0, 0, 373, 374, 5, 105, 0, 0, 374, 375, 5, 108, 0, 0, 375, 72, 1, 0, 0,
		0, 376, 380, 7, 6, 0, 0, 377, 379, 7, 7, 0, 0, 378, 377, 1, 0, 0, 0, 379,
		382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 74, 1,
		0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 384, 5, 45, 0, 0, 384, 385, 5, 62,
		0, 0, 385, 76, 1, 0, 0, 0, 386, 387, 5, 61, 0, 0, 387, 388, 5, 61, 0, 0,
		388, 78, 1, 0, 0, 0, 389, 390, 5, 33, 0, 0, 390, 391, 5, 61, 0, 0, 391,
		80, 1, 0, 0, 0, 392, 393, 5, 60, 0, 0, 393, 82, 1, 0, 0, 0, 394, 395, 5,
		62, 0, 0, 395, 84, 1, 0, 0, 0, 396, 397, 5, 60, 0, 0, 397, 398, 5, 61,
		0, 0, 398, 86, 1, 0, 0, 0, 399, 400, 5, 62, 0, 0, 400, 401, 5, 61, 0, 0,
		401, 88, 1, 0, 0, 0, 402, 403, 5, 61, 0, 0, 403, 90, 1, 0, 0, 0, 404, 405,
		5, 42, 0, 0, 405, 406, 5, 61, 0, 0, 406, 92, 1, 0, 0, 0, 407, 408, 5, 47,
		0, 0, 408, 409, 5, 61, 0, 0, 409, 94, 1, 0, 0, 0, 410, 411, 5, 43, 0, 0,
		411, 412, 5, 61, 0, 0, 412, 96, 1, 0, 0, 0, 413, 414, 5, 45, 0, 0, 414,
		415, 5, 61, 0, 0, 415, 98, 1, 0, 0, 0, 416, 417, 5, 37, 0, 0, 417, 418,
		5, 61, 0, 0, 418, 100, 1, 0, 0, 0, 419, 420, 5, 42, 0, 0, 420, 102, 1,
		0, 0, 0, 421, 422, 5, 47, 0, 0, 422, 104, 1, 0, 0, 0, 423, 424, 5, 43,
		0, 0, 424, 106, 1, 0, 0, 0, 425, 426, 5, 45, 0, 0, 426, 108, 1, 0, 0, 0,
		427, 428, 5, 37, 0, 0, 428, 110, 1, 0, 0, 0, 429, 430, 5, 38, 0, 0, 430,
		431, 5, 38, 0, 0, 431, 112, 1, 0, 0, 0, 432, 433, 5, 124, 0, 0, 433, 434,
		5, 124, 0, 0, 434, 114, 1, 0, 0, 0, 435, 436, 5, 33, 0, 0, 436, 116, 1,
		0, 0, 0, 437, 438, 5, 63, 0, 0, 438, 118, 1, 0, 0, 0, 439, 440, 5, 40,
		0, 0, 440, 120, 1, 0, 0, 0, 441, 442, 5, 41, 0, 0, 442, 122, 1, 0, 0, 0,
		443, 444, 5, 123, 0, 0, 444, 124, 1, 0, 0, 0, 445, 446, 5, 125, 0, 0, 446,
		126, 1, 0, 0, 0, 447, 448, 5, 91, 0, 0, 448, 128, 1, 0, 0, 0, 449, 450,
		5, 93, 0, 0, 450, 130, 1, 0, 0, 0, 451, 452, 5, 92, 0, 0, 452, 132, 1,
		0, 0, 0, 453, 454, 5, 44, 0, 0, 454, 134, 1, 0, 0, 0, 455, 456, 5, 59,
		0, 0, 456, 136, 1, 0, 0, 0, 457, 458, 5, 58, 0, 0, 458, 138, 1, 0, 0, 0,
		459, 460, 5, 46, 0, 0, 460, 140, 1, 0, 0, 0, 12, 0, 144, 154, 165, 328,
		333, 340, 351, 357, 359, 368, 380, 1, 6, 0, 0,
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

// SwiftLexerInit initializes any static state used to implement SwiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.once.Do(swiftlexerLexerInit)
}

// NewSwiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftLexer(input antlr.CharStream) *SwiftLexer {
	SwiftLexerInit()
	l := new(SwiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftLexer tokens.
const (
	SwiftLexerWHITESPACE      = 1
	SwiftLexerCOMMENT         = 2
	SwiftLexerBLOCK_COMMENT   = 3
	SwiftLexerKw_LET          = 4
	SwiftLexerKw_VAR          = 5
	SwiftLexerKw_FUNC         = 6
	SwiftLexerKw_STRUCT       = 7
	SwiftLexerKw_MUTATING     = 8
	SwiftLexerKw_AMPER        = 9
	SwiftLexerKw_IF           = 10
	SwiftLexerKw_ELSE         = 11
	SwiftLexerKw_SWITCH       = 12
	SwiftLexerKw_CASE         = 13
	SwiftLexerKw_DEFAULT      = 14
	SwiftLexerKw_FOR          = 15
	SwiftLexerKw_WHILE        = 16
	SwiftLexerKw_BREAK        = 17
	SwiftLexerKw_CONTINUE     = 18
	SwiftLexerKw_RETURN       = 19
	SwiftLexerKw_DO           = 20
	SwiftLexerKw_REPEAT       = 21
	SwiftLexerKw_IN           = 22
	SwiftLexerKw_GUARD        = 23
	SwiftLexerKw_INT          = 24
	SwiftLexerKw_FLOAT        = 25
	SwiftLexerKw_BOOL         = 26
	SwiftLexerKw_STRING       = 27
	SwiftLexerKw_CHAR         = 28
	SwiftLexerKw_RANGE        = 29
	SwiftLexerKw_INOUT        = 30
	SwiftLexerINT             = 31
	SwiftLexerFLOAT           = 32
	SwiftLexerBOOL            = 33
	SwiftLexerSTRING          = 34
	SwiftLexerCHAR            = 35
	SwiftLexerNIL             = 36
	SwiftLexerID              = 37
	SwiftLexerOp_ARROW        = 38
	SwiftLexerOp_EQ           = 39
	SwiftLexerOp_NEQ          = 40
	SwiftLexerOp_LT           = 41
	SwiftLexerOp_GT           = 42
	SwiftLexerOp_LE           = 43
	SwiftLexerOp_GE           = 44
	SwiftLexerOp_ASSIGN       = 45
	SwiftLexerOp_MUL_ASSIGN   = 46
	SwiftLexerOp_DIV_ASSIGN   = 47
	SwiftLexerOp_PLUS_ASSIGN  = 48
	SwiftLexerOp_MINUS_ASSIGN = 49
	SwiftLexerOp_MOD_ASSIGN   = 50
	SwiftLexerOp_MUL          = 51
	SwiftLexerOp_DIV          = 52
	SwiftLexerOp_PLUS         = 53
	SwiftLexerOp_MINUS        = 54
	SwiftLexerOp_MOD          = 55
	SwiftLexerOp_AND          = 56
	SwiftLexerOp_OR           = 57
	SwiftLexerOp_NOT          = 58
	SwiftLexerOp_TERNARY      = 59
	SwiftLexerLPAREN          = 60
	SwiftLexerRPAREN          = 61
	SwiftLexerLBRACE          = 62
	SwiftLexerRBRACE          = 63
	SwiftLexerLBRACKET        = 64
	SwiftLexerRBRACKET        = 65
	SwiftLexerBACKSLASH       = 66
	SwiftLexerCOMMA           = 67
	SwiftLexerSEMICOLON       = 68
	SwiftLexerCOLON           = 69
	SwiftLexerDOT             = 70
)
