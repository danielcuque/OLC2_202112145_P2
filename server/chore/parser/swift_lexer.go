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
		"", "", "", "", "'let'", "'var'", "'func'", "'struct'", "'if'", "'else'",
		"'switch'", "'case'", "'default'", "'for'", "'while'", "'break'", "'continue'",
		"'return'", "'do'", "'repeat'", "'in'", "'guard'", "'Int'", "'Float'",
		"'Bool'", "'String'", "'Character'", "'nil'", "'...'", "", "", "", "",
		"", "", "'->'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'='",
		"'*='", "'/='", "'+='", "'-='", "'%='", "'*'", "'/'", "'+'", "'-'",
		"'%'", "'&&'", "'||'", "'!'", "'?'", "'('", "')'", "'{'", "'}'", "'['",
		"']'", "'\\'", "','", "';'", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_GUARD", "Kw_INT", "Kw_FLOAT", "Kw_BOOL", "Kw_STRING",
		"Kw_CHAR", "Kw_NIL", "Kw_RANGE", "INT", "FLOAT", "BOOL", "STRING", "CHAR",
		"ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE",
		"Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN",
		"Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", "Op_MOD",
		"Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_GUARD", "Kw_INT", "Kw_FLOAT", "Kw_BOOL", "Kw_STRING",
		"Kw_CHAR", "Kw_NIL", "Kw_RANGE", "INT", "FLOAT", "BOOL", "STRING", "CHAR",
		"ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE",
		"Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN",
		"Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", "Op_MOD",
		"Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 67, 441, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 1, 0, 4, 0,
		137, 8, 0, 11, 0, 12, 0, 138, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		147, 8, 1, 10, 1, 12, 1, 150, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 158, 8, 2, 10, 2, 12, 2, 161, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 4, 28,
		308, 8, 28, 11, 28, 12, 28, 309, 1, 29, 5, 29, 313, 8, 29, 10, 29, 12,
		29, 316, 9, 29, 1, 29, 1, 29, 4, 29, 320, 8, 29, 11, 29, 12, 29, 321, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 333,
		8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 339, 8, 31, 10, 31, 12, 31, 342,
		9, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 350, 8, 32, 10,
		32, 12, 32, 353, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 359, 8, 33,
		10, 33, 12, 33, 362, 9, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1,
		52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57,
		1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1,
		62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 159, 0,
		67, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93,
		47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55,
		111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63,
		127, 64, 129, 65, 131, 66, 133, 67, 1, 0, 8, 4, 0, 9, 10, 13, 13, 32, 32,
		92, 92, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34,
		92, 92, 3, 0, 10, 10, 13, 13, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92,
		92, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		452, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0,
		0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107,
		1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0,
		0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1,
		0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0,
		129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 1, 136, 1, 0,
		0, 0, 3, 142, 1, 0, 0, 0, 5, 153, 1, 0, 0, 0, 7, 167, 1, 0, 0, 0, 9, 171,
		1, 0, 0, 0, 11, 175, 1, 0, 0, 0, 13, 180, 1, 0, 0, 0, 15, 187, 1, 0, 0,
		0, 17, 190, 1, 0, 0, 0, 19, 195, 1, 0, 0, 0, 21, 202, 1, 0, 0, 0, 23, 207,
		1, 0, 0, 0, 25, 215, 1, 0, 0, 0, 27, 219, 1, 0, 0, 0, 29, 225, 1, 0, 0,
		0, 31, 231, 1, 0, 0, 0, 33, 240, 1, 0, 0, 0, 35, 247, 1, 0, 0, 0, 37, 250,
		1, 0, 0, 0, 39, 257, 1, 0, 0, 0, 41, 260, 1, 0, 0, 0, 43, 266, 1, 0, 0,
		0, 45, 270, 1, 0, 0, 0, 47, 276, 1, 0, 0, 0, 49, 281, 1, 0, 0, 0, 51, 288,
		1, 0, 0, 0, 53, 298, 1, 0, 0, 0, 55, 302, 1, 0, 0, 0, 57, 307, 1, 0, 0,
		0, 59, 314, 1, 0, 0, 0, 61, 332, 1, 0, 0, 0, 63, 334, 1, 0, 0, 0, 65, 345,
		1, 0, 0, 0, 67, 356, 1, 0, 0, 0, 69, 363, 1, 0, 0, 0, 71, 366, 1, 0, 0,
		0, 73, 369, 1, 0, 0, 0, 75, 372, 1, 0, 0, 0, 77, 374, 1, 0, 0, 0, 79, 376,
		1, 0, 0, 0, 81, 379, 1, 0, 0, 0, 83, 382, 1, 0, 0, 0, 85, 384, 1, 0, 0,
		0, 87, 387, 1, 0, 0, 0, 89, 390, 1, 0, 0, 0, 91, 393, 1, 0, 0, 0, 93, 396,
		1, 0, 0, 0, 95, 399, 1, 0, 0, 0, 97, 401, 1, 0, 0, 0, 99, 403, 1, 0, 0,
		0, 101, 405, 1, 0, 0, 0, 103, 407, 1, 0, 0, 0, 105, 409, 1, 0, 0, 0, 107,
		412, 1, 0, 0, 0, 109, 415, 1, 0, 0, 0, 111, 417, 1, 0, 0, 0, 113, 419,
		1, 0, 0, 0, 115, 421, 1, 0, 0, 0, 117, 423, 1, 0, 0, 0, 119, 425, 1, 0,
		0, 0, 121, 427, 1, 0, 0, 0, 123, 429, 1, 0, 0, 0, 125, 431, 1, 0, 0, 0,
		127, 433, 1, 0, 0, 0, 129, 435, 1, 0, 0, 0, 131, 437, 1, 0, 0, 0, 133,
		439, 1, 0, 0, 0, 135, 137, 7, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137, 138,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 141, 6, 0, 0, 0, 141, 2, 1, 0, 0, 0, 142, 143, 5, 47, 0, 0,
		143, 144, 5, 47, 0, 0, 144, 148, 1, 0, 0, 0, 145, 147, 8, 1, 0, 0, 146,
		145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 6, 1,
		0, 0, 152, 4, 1, 0, 0, 0, 153, 154, 5, 47, 0, 0, 154, 155, 5, 42, 0, 0,
		155, 159, 1, 0, 0, 0, 156, 158, 9, 0, 0, 0, 157, 156, 1, 0, 0, 0, 158,
		161, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 162,
		1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163, 5, 42, 0, 0, 163, 164, 5, 47,
		0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 6, 2, 0, 0, 166, 6, 1, 0, 0, 0, 167,
		168, 5, 108, 0, 0, 168, 169, 5, 101, 0, 0, 169, 170, 5, 116, 0, 0, 170,
		8, 1, 0, 0, 0, 171, 172, 5, 118, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174,
		5, 114, 0, 0, 174, 10, 1, 0, 0, 0, 175, 176, 5, 102, 0, 0, 176, 177, 5,
		117, 0, 0, 177, 178, 5, 110, 0, 0, 178, 179, 5, 99, 0, 0, 179, 12, 1, 0,
		0, 0, 180, 181, 5, 115, 0, 0, 181, 182, 5, 116, 0, 0, 182, 183, 5, 114,
		0, 0, 183, 184, 5, 117, 0, 0, 184, 185, 5, 99, 0, 0, 185, 186, 5, 116,
		0, 0, 186, 14, 1, 0, 0, 0, 187, 188, 5, 105, 0, 0, 188, 189, 5, 102, 0,
		0, 189, 16, 1, 0, 0, 0, 190, 191, 5, 101, 0, 0, 191, 192, 5, 108, 0, 0,
		192, 193, 5, 115, 0, 0, 193, 194, 5, 101, 0, 0, 194, 18, 1, 0, 0, 0, 195,
		196, 5, 115, 0, 0, 196, 197, 5, 119, 0, 0, 197, 198, 5, 105, 0, 0, 198,
		199, 5, 116, 0, 0, 199, 200, 5, 99, 0, 0, 200, 201, 5, 104, 0, 0, 201,
		20, 1, 0, 0, 0, 202, 203, 5, 99, 0, 0, 203, 204, 5, 97, 0, 0, 204, 205,
		5, 115, 0, 0, 205, 206, 5, 101, 0, 0, 206, 22, 1, 0, 0, 0, 207, 208, 5,
		100, 0, 0, 208, 209, 5, 101, 0, 0, 209, 210, 5, 102, 0, 0, 210, 211, 5,
		97, 0, 0, 211, 212, 5, 117, 0, 0, 212, 213, 5, 108, 0, 0, 213, 214, 5,
		116, 0, 0, 214, 24, 1, 0, 0, 0, 215, 216, 5, 102, 0, 0, 216, 217, 5, 111,
		0, 0, 217, 218, 5, 114, 0, 0, 218, 26, 1, 0, 0, 0, 219, 220, 5, 119, 0,
		0, 220, 221, 5, 104, 0, 0, 221, 222, 5, 105, 0, 0, 222, 223, 5, 108, 0,
		0, 223, 224, 5, 101, 0, 0, 224, 28, 1, 0, 0, 0, 225, 226, 5, 98, 0, 0,
		226, 227, 5, 114, 0, 0, 227, 228, 5, 101, 0, 0, 228, 229, 5, 97, 0, 0,
		229, 230, 5, 107, 0, 0, 230, 30, 1, 0, 0, 0, 231, 232, 5, 99, 0, 0, 232,
		233, 5, 111, 0, 0, 233, 234, 5, 110, 0, 0, 234, 235, 5, 116, 0, 0, 235,
		236, 5, 105, 0, 0, 236, 237, 5, 110, 0, 0, 237, 238, 5, 117, 0, 0, 238,
		239, 5, 101, 0, 0, 239, 32, 1, 0, 0, 0, 240, 241, 5, 114, 0, 0, 241, 242,
		5, 101, 0, 0, 242, 243, 5, 116, 0, 0, 243, 244, 5, 117, 0, 0, 244, 245,
		5, 114, 0, 0, 245, 246, 5, 110, 0, 0, 246, 34, 1, 0, 0, 0, 247, 248, 5,
		100, 0, 0, 248, 249, 5, 111, 0, 0, 249, 36, 1, 0, 0, 0, 250, 251, 5, 114,
		0, 0, 251, 252, 5, 101, 0, 0, 252, 253, 5, 112, 0, 0, 253, 254, 5, 101,
		0, 0, 254, 255, 5, 97, 0, 0, 255, 256, 5, 116, 0, 0, 256, 38, 1, 0, 0,
		0, 257, 258, 5, 105, 0, 0, 258, 259, 5, 110, 0, 0, 259, 40, 1, 0, 0, 0,
		260, 261, 5, 103, 0, 0, 261, 262, 5, 117, 0, 0, 262, 263, 5, 97, 0, 0,
		263, 264, 5, 114, 0, 0, 264, 265, 5, 100, 0, 0, 265, 42, 1, 0, 0, 0, 266,
		267, 5, 73, 0, 0, 267, 268, 5, 110, 0, 0, 268, 269, 5, 116, 0, 0, 269,
		44, 1, 0, 0, 0, 270, 271, 5, 70, 0, 0, 271, 272, 5, 108, 0, 0, 272, 273,
		5, 111, 0, 0, 273, 274, 5, 97, 0, 0, 274, 275, 5, 116, 0, 0, 275, 46, 1,
		0, 0, 0, 276, 277, 5, 66, 0, 0, 277, 278, 5, 111, 0, 0, 278, 279, 5, 111,
		0, 0, 279, 280, 5, 108, 0, 0, 280, 48, 1, 0, 0, 0, 281, 282, 5, 83, 0,
		0, 282, 283, 5, 116, 0, 0, 283, 284, 5, 114, 0, 0, 284, 285, 5, 105, 0,
		0, 285, 286, 5, 110, 0, 0, 286, 287, 5, 103, 0, 0, 287, 50, 1, 0, 0, 0,
		288, 289, 5, 67, 0, 0, 289, 290, 5, 104, 0, 0, 290, 291, 5, 97, 0, 0, 291,
		292, 5, 114, 0, 0, 292, 293, 5, 97, 0, 0, 293, 294, 5, 99, 0, 0, 294, 295,
		5, 116, 0, 0, 295, 296, 5, 101, 0, 0, 296, 297, 5, 114, 0, 0, 297, 52,
		1, 0, 0, 0, 298, 299, 5, 110, 0, 0, 299, 300, 5, 105, 0, 0, 300, 301, 5,
		108, 0, 0, 301, 54, 1, 0, 0, 0, 302, 303, 5, 46, 0, 0, 303, 304, 5, 46,
		0, 0, 304, 305, 5, 46, 0, 0, 305, 56, 1, 0, 0, 0, 306, 308, 7, 2, 0, 0,
		307, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309,
		310, 1, 0, 0, 0, 310, 58, 1, 0, 0, 0, 311, 313, 7, 2, 0, 0, 312, 311, 1,
		0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0, 0,
		0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 319, 5, 46, 0, 0, 318,
		320, 7, 2, 0, 0, 319, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 319,
		1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 60, 1, 0, 0, 0, 323, 324, 5, 116,
		0, 0, 324, 325, 5, 114, 0, 0, 325, 326, 5, 117, 0, 0, 326, 333, 5, 101,
		0, 0, 327, 328, 5, 102, 0, 0, 328, 329, 5, 97, 0, 0, 329, 330, 5, 108,
		0, 0, 330, 331, 5, 115, 0, 0, 331, 333, 5, 101, 0, 0, 332, 323, 1, 0, 0,
		0, 332, 327, 1, 0, 0, 0, 333, 62, 1, 0, 0, 0, 334, 340, 5, 34, 0, 0, 335,
		339, 8, 3, 0, 0, 336, 337, 5, 92, 0, 0, 337, 339, 7, 4, 0, 0, 338, 335,
		1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0,
		0, 0, 340, 341, 1, 0, 0, 0, 341, 343, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0,
		343, 344, 5, 34, 0, 0, 344, 64, 1, 0, 0, 0, 345, 351, 5, 39, 0, 0, 346,
		350, 8, 5, 0, 0, 347, 348, 5, 92, 0, 0, 348, 350, 7, 4, 0, 0, 349, 346,
		1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0,
		0, 0, 351, 352, 1, 0, 0, 0, 352, 354, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0,
		354, 355, 5, 39, 0, 0, 355, 66, 1, 0, 0, 0, 356, 360, 7, 6, 0, 0, 357,
		359, 7, 7, 0, 0, 358, 357, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358,
		1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 68, 1, 0, 0, 0, 362, 360, 1, 0,
		0, 0, 363, 364, 5, 45, 0, 0, 364, 365, 5, 62, 0, 0, 365, 70, 1, 0, 0, 0,
		366, 367, 5, 61, 0, 0, 367, 368, 5, 61, 0, 0, 368, 72, 1, 0, 0, 0, 369,
		370, 5, 33, 0, 0, 370, 371, 5, 61, 0, 0, 371, 74, 1, 0, 0, 0, 372, 373,
		5, 60, 0, 0, 373, 76, 1, 0, 0, 0, 374, 375, 5, 62, 0, 0, 375, 78, 1, 0,
		0, 0, 376, 377, 5, 60, 0, 0, 377, 378, 5, 61, 0, 0, 378, 80, 1, 0, 0, 0,
		379, 380, 5, 62, 0, 0, 380, 381, 5, 61, 0, 0, 381, 82, 1, 0, 0, 0, 382,
		383, 5, 61, 0, 0, 383, 84, 1, 0, 0, 0, 384, 385, 5, 42, 0, 0, 385, 386,
		5, 61, 0, 0, 386, 86, 1, 0, 0, 0, 387, 388, 5, 47, 0, 0, 388, 389, 5, 61,
		0, 0, 389, 88, 1, 0, 0, 0, 390, 391, 5, 43, 0, 0, 391, 392, 5, 61, 0, 0,
		392, 90, 1, 0, 0, 0, 393, 394, 5, 45, 0, 0, 394, 395, 5, 61, 0, 0, 395,
		92, 1, 0, 0, 0, 396, 397, 5, 37, 0, 0, 397, 398, 5, 61, 0, 0, 398, 94,
		1, 0, 0, 0, 399, 400, 5, 42, 0, 0, 400, 96, 1, 0, 0, 0, 401, 402, 5, 47,
		0, 0, 402, 98, 1, 0, 0, 0, 403, 404, 5, 43, 0, 0, 404, 100, 1, 0, 0, 0,
		405, 406, 5, 45, 0, 0, 406, 102, 1, 0, 0, 0, 407, 408, 5, 37, 0, 0, 408,
		104, 1, 0, 0, 0, 409, 410, 5, 38, 0, 0, 410, 411, 5, 38, 0, 0, 411, 106,
		1, 0, 0, 0, 412, 413, 5, 124, 0, 0, 413, 414, 5, 124, 0, 0, 414, 108, 1,
		0, 0, 0, 415, 416, 5, 33, 0, 0, 416, 110, 1, 0, 0, 0, 417, 418, 5, 63,
		0, 0, 418, 112, 1, 0, 0, 0, 419, 420, 5, 40, 0, 0, 420, 114, 1, 0, 0, 0,
		421, 422, 5, 41, 0, 0, 422, 116, 1, 0, 0, 0, 423, 424, 5, 123, 0, 0, 424,
		118, 1, 0, 0, 0, 425, 426, 5, 125, 0, 0, 426, 120, 1, 0, 0, 0, 427, 428,
		5, 91, 0, 0, 428, 122, 1, 0, 0, 0, 429, 430, 5, 93, 0, 0, 430, 124, 1,
		0, 0, 0, 431, 432, 5, 92, 0, 0, 432, 126, 1, 0, 0, 0, 433, 434, 5, 44,
		0, 0, 434, 128, 1, 0, 0, 0, 435, 436, 5, 59, 0, 0, 436, 130, 1, 0, 0, 0,
		437, 438, 5, 58, 0, 0, 438, 132, 1, 0, 0, 0, 439, 440, 5, 46, 0, 0, 440,
		134, 1, 0, 0, 0, 13, 0, 138, 148, 159, 309, 314, 321, 332, 338, 340, 349,
		351, 360, 1, 6, 0, 0,
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
	SwiftLexerKw_IF           = 8
	SwiftLexerKw_ELSE         = 9
	SwiftLexerKw_SWITCH       = 10
	SwiftLexerKw_CASE         = 11
	SwiftLexerKw_DEFAULT      = 12
	SwiftLexerKw_FOR          = 13
	SwiftLexerKw_WHILE        = 14
	SwiftLexerKw_BREAK        = 15
	SwiftLexerKw_CONTINUE     = 16
	SwiftLexerKw_RETURN       = 17
	SwiftLexerKw_DO           = 18
	SwiftLexerKw_REPEAT       = 19
	SwiftLexerKw_IN           = 20
	SwiftLexerKw_GUARD        = 21
	SwiftLexerKw_INT          = 22
	SwiftLexerKw_FLOAT        = 23
	SwiftLexerKw_BOOL         = 24
	SwiftLexerKw_STRING       = 25
	SwiftLexerKw_CHAR         = 26
	SwiftLexerKw_NIL          = 27
	SwiftLexerKw_RANGE        = 28
	SwiftLexerINT             = 29
	SwiftLexerFLOAT           = 30
	SwiftLexerBOOL            = 31
	SwiftLexerSTRING          = 32
	SwiftLexerCHAR            = 33
	SwiftLexerID              = 34
	SwiftLexerOp_ARROW        = 35
	SwiftLexerOp_EQ           = 36
	SwiftLexerOp_NEQ          = 37
	SwiftLexerOp_LT           = 38
	SwiftLexerOp_GT           = 39
	SwiftLexerOp_LE           = 40
	SwiftLexerOp_GE           = 41
	SwiftLexerOp_ASSIGN       = 42
	SwiftLexerOp_MUL_ASSIGN   = 43
	SwiftLexerOp_DIV_ASSIGN   = 44
	SwiftLexerOp_PLUS_ASSIGN  = 45
	SwiftLexerOp_MINUS_ASSIGN = 46
	SwiftLexerOp_MOD_ASSIGN   = 47
	SwiftLexerOp_MUL          = 48
	SwiftLexerOp_DIV          = 49
	SwiftLexerOp_PLUS         = 50
	SwiftLexerOp_MINUS        = 51
	SwiftLexerOp_MOD          = 52
	SwiftLexerOp_AND          = 53
	SwiftLexerOp_OR           = 54
	SwiftLexerOp_NOT          = 55
	SwiftLexerOp_TERNARY      = 56
	SwiftLexerLPAREN          = 57
	SwiftLexerRPAREN          = 58
	SwiftLexerLBRACE          = 59
	SwiftLexerRBRACE          = 60
	SwiftLexerLBRACKET        = 61
	SwiftLexerRBRACKET        = 62
	SwiftLexerBACKSLASH       = 63
	SwiftLexerCOMMA           = 64
	SwiftLexerSEMICOLON       = 65
	SwiftLexerCOLON           = 66
	SwiftLexerDOT             = 67
)
