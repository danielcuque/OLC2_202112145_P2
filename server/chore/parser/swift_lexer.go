// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

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
		"'return'", "'do'", "'repeat'", "'in'", "'Int'", "'Double'", "'Bool'",
		"'String'", "'nil'", "'...'", "", "", "", "", "", "", "'->'", "'=='",
		"'!='", "'<'", "'>'", "'<='", "'>='", "'='", "'*='", "'/='", "'+='",
		"'-='", "'%='", "'*'", "'/'", "'+'", "'-'", "'%'", "'&&'", "'||'", "'!'",
		"'?'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'\\'", "','", "';'",
		"':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_INT", "Kw_DOUBLE", "Kw_BOOL", "Kw_STRING",
		"Kw_NIL", "Kw_RANGE", "INT", "DOUBLE", "BOOL", "STRING", "CHAR", "ID",
		"Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_ASSIGN",
		"Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN",
		"Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", "Op_MOD",
		"Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_INT", "Kw_DOUBLE", "Kw_BOOL", "Kw_STRING",
		"Kw_NIL", "Kw_RANGE", "INT", "DOUBLE", "BOOL", "STRING", "CHAR", "ID",
		"Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE", "Op_ASSIGN",
		"Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN",
		"Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", "Op_MOD",
		"Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 65, 422, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		62, 2, 63, 7, 63, 2, 64, 7, 64, 1, 0, 4, 0, 133, 8, 0, 11, 0, 12, 0, 134,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 143, 8, 1, 10, 1, 12, 1, 146,
		9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 154, 8, 2, 10, 2, 12, 2,
		157, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 4, 26, 289, 8, 26, 11, 26, 12, 26, 290, 1, 27, 5, 27, 294, 8, 27,
		10, 27, 12, 27, 297, 9, 27, 1, 27, 1, 27, 4, 27, 301, 8, 27, 11, 27, 12,
		27, 302, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		3, 28, 314, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 320, 8, 29, 10, 29,
		12, 29, 323, 9, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 331,
		8, 30, 10, 30, 12, 30, 334, 9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 340,
		8, 31, 10, 31, 12, 31, 343, 9, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1,
		50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 155,
		0, 65, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46,
		93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109,
		55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125,
		63, 127, 64, 129, 65, 1, 0, 8, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2,
		0, 10, 10, 13, 13, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92,
		3, 0, 10, 10, 13, 13, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 3,
		0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 433,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
		1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0,
		93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0,
		0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1,
		0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0,
		115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0,
		0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129,
		1, 0, 0, 0, 1, 132, 1, 0, 0, 0, 3, 138, 1, 0, 0, 0, 5, 149, 1, 0, 0, 0,
		7, 163, 1, 0, 0, 0, 9, 167, 1, 0, 0, 0, 11, 171, 1, 0, 0, 0, 13, 176, 1,
		0, 0, 0, 15, 183, 1, 0, 0, 0, 17, 186, 1, 0, 0, 0, 19, 191, 1, 0, 0, 0,
		21, 198, 1, 0, 0, 0, 23, 203, 1, 0, 0, 0, 25, 211, 1, 0, 0, 0, 27, 215,
		1, 0, 0, 0, 29, 221, 1, 0, 0, 0, 31, 227, 1, 0, 0, 0, 33, 236, 1, 0, 0,
		0, 35, 243, 1, 0, 0, 0, 37, 246, 1, 0, 0, 0, 39, 253, 1, 0, 0, 0, 41, 256,
		1, 0, 0, 0, 43, 260, 1, 0, 0, 0, 45, 267, 1, 0, 0, 0, 47, 272, 1, 0, 0,
		0, 49, 279, 1, 0, 0, 0, 51, 283, 1, 0, 0, 0, 53, 288, 1, 0, 0, 0, 55, 295,
		1, 0, 0, 0, 57, 313, 1, 0, 0, 0, 59, 315, 1, 0, 0, 0, 61, 326, 1, 0, 0,
		0, 63, 337, 1, 0, 0, 0, 65, 344, 1, 0, 0, 0, 67, 347, 1, 0, 0, 0, 69, 350,
		1, 0, 0, 0, 71, 353, 1, 0, 0, 0, 73, 355, 1, 0, 0, 0, 75, 357, 1, 0, 0,
		0, 77, 360, 1, 0, 0, 0, 79, 363, 1, 0, 0, 0, 81, 365, 1, 0, 0, 0, 83, 368,
		1, 0, 0, 0, 85, 371, 1, 0, 0, 0, 87, 374, 1, 0, 0, 0, 89, 377, 1, 0, 0,
		0, 91, 380, 1, 0, 0, 0, 93, 382, 1, 0, 0, 0, 95, 384, 1, 0, 0, 0, 97, 386,
		1, 0, 0, 0, 99, 388, 1, 0, 0, 0, 101, 390, 1, 0, 0, 0, 103, 393, 1, 0,
		0, 0, 105, 396, 1, 0, 0, 0, 107, 398, 1, 0, 0, 0, 109, 400, 1, 0, 0, 0,
		111, 402, 1, 0, 0, 0, 113, 404, 1, 0, 0, 0, 115, 406, 1, 0, 0, 0, 117,
		408, 1, 0, 0, 0, 119, 410, 1, 0, 0, 0, 121, 412, 1, 0, 0, 0, 123, 414,
		1, 0, 0, 0, 125, 416, 1, 0, 0, 0, 127, 418, 1, 0, 0, 0, 129, 420, 1, 0,
		0, 0, 131, 133, 7, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136,
		137, 6, 0, 0, 0, 137, 2, 1, 0, 0, 0, 138, 139, 5, 47, 0, 0, 139, 140, 5,
		47, 0, 0, 140, 144, 1, 0, 0, 0, 141, 143, 8, 1, 0, 0, 142, 141, 1, 0, 0,
		0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		147, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147, 148, 6, 1, 0, 0, 148, 4, 1,
		0, 0, 0, 149, 150, 5, 47, 0, 0, 150, 151, 5, 42, 0, 0, 151, 155, 1, 0,
		0, 0, 152, 154, 9, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0,
		155, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157,
		155, 1, 0, 0, 0, 158, 159, 5, 42, 0, 0, 159, 160, 5, 47, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 162, 6, 2, 0, 0, 162, 6, 1, 0, 0, 0, 163, 164, 5, 108,
		0, 0, 164, 165, 5, 101, 0, 0, 165, 166, 5, 116, 0, 0, 166, 8, 1, 0, 0,
		0, 167, 168, 5, 118, 0, 0, 168, 169, 5, 97, 0, 0, 169, 170, 5, 114, 0,
		0, 170, 10, 1, 0, 0, 0, 171, 172, 5, 102, 0, 0, 172, 173, 5, 117, 0, 0,
		173, 174, 5, 110, 0, 0, 174, 175, 5, 99, 0, 0, 175, 12, 1, 0, 0, 0, 176,
		177, 5, 115, 0, 0, 177, 178, 5, 116, 0, 0, 178, 179, 5, 114, 0, 0, 179,
		180, 5, 117, 0, 0, 180, 181, 5, 99, 0, 0, 181, 182, 5, 116, 0, 0, 182,
		14, 1, 0, 0, 0, 183, 184, 5, 105, 0, 0, 184, 185, 5, 102, 0, 0, 185, 16,
		1, 0, 0, 0, 186, 187, 5, 101, 0, 0, 187, 188, 5, 108, 0, 0, 188, 189, 5,
		115, 0, 0, 189, 190, 5, 101, 0, 0, 190, 18, 1, 0, 0, 0, 191, 192, 5, 115,
		0, 0, 192, 193, 5, 119, 0, 0, 193, 194, 5, 105, 0, 0, 194, 195, 5, 116,
		0, 0, 195, 196, 5, 99, 0, 0, 196, 197, 5, 104, 0, 0, 197, 20, 1, 0, 0,
		0, 198, 199, 5, 99, 0, 0, 199, 200, 5, 97, 0, 0, 200, 201, 5, 115, 0, 0,
		201, 202, 5, 101, 0, 0, 202, 22, 1, 0, 0, 0, 203, 204, 5, 100, 0, 0, 204,
		205, 5, 101, 0, 0, 205, 206, 5, 102, 0, 0, 206, 207, 5, 97, 0, 0, 207,
		208, 5, 117, 0, 0, 208, 209, 5, 108, 0, 0, 209, 210, 5, 116, 0, 0, 210,
		24, 1, 0, 0, 0, 211, 212, 5, 102, 0, 0, 212, 213, 5, 111, 0, 0, 213, 214,
		5, 114, 0, 0, 214, 26, 1, 0, 0, 0, 215, 216, 5, 119, 0, 0, 216, 217, 5,
		104, 0, 0, 217, 218, 5, 105, 0, 0, 218, 219, 5, 108, 0, 0, 219, 220, 5,
		101, 0, 0, 220, 28, 1, 0, 0, 0, 221, 222, 5, 98, 0, 0, 222, 223, 5, 114,
		0, 0, 223, 224, 5, 101, 0, 0, 224, 225, 5, 97, 0, 0, 225, 226, 5, 107,
		0, 0, 226, 30, 1, 0, 0, 0, 227, 228, 5, 99, 0, 0, 228, 229, 5, 111, 0,
		0, 229, 230, 5, 110, 0, 0, 230, 231, 5, 116, 0, 0, 231, 232, 5, 105, 0,
		0, 232, 233, 5, 110, 0, 0, 233, 234, 5, 117, 0, 0, 234, 235, 5, 101, 0,
		0, 235, 32, 1, 0, 0, 0, 236, 237, 5, 114, 0, 0, 237, 238, 5, 101, 0, 0,
		238, 239, 5, 116, 0, 0, 239, 240, 5, 117, 0, 0, 240, 241, 5, 114, 0, 0,
		241, 242, 5, 110, 0, 0, 242, 34, 1, 0, 0, 0, 243, 244, 5, 100, 0, 0, 244,
		245, 5, 111, 0, 0, 245, 36, 1, 0, 0, 0, 246, 247, 5, 114, 0, 0, 247, 248,
		5, 101, 0, 0, 248, 249, 5, 112, 0, 0, 249, 250, 5, 101, 0, 0, 250, 251,
		5, 97, 0, 0, 251, 252, 5, 116, 0, 0, 252, 38, 1, 0, 0, 0, 253, 254, 5,
		105, 0, 0, 254, 255, 5, 110, 0, 0, 255, 40, 1, 0, 0, 0, 256, 257, 5, 73,
		0, 0, 257, 258, 5, 110, 0, 0, 258, 259, 5, 116, 0, 0, 259, 42, 1, 0, 0,
		0, 260, 261, 5, 68, 0, 0, 261, 262, 5, 111, 0, 0, 262, 263, 5, 117, 0,
		0, 263, 264, 5, 98, 0, 0, 264, 265, 5, 108, 0, 0, 265, 266, 5, 101, 0,
		0, 266, 44, 1, 0, 0, 0, 267, 268, 5, 66, 0, 0, 268, 269, 5, 111, 0, 0,
		269, 270, 5, 111, 0, 0, 270, 271, 5, 108, 0, 0, 271, 46, 1, 0, 0, 0, 272,
		273, 5, 83, 0, 0, 273, 274, 5, 116, 0, 0, 274, 275, 5, 114, 0, 0, 275,
		276, 5, 105, 0, 0, 276, 277, 5, 110, 0, 0, 277, 278, 5, 103, 0, 0, 278,
		48, 1, 0, 0, 0, 279, 280, 5, 110, 0, 0, 280, 281, 5, 105, 0, 0, 281, 282,
		5, 108, 0, 0, 282, 50, 1, 0, 0, 0, 283, 284, 5, 46, 0, 0, 284, 285, 5,
		46, 0, 0, 285, 286, 5, 46, 0, 0, 286, 52, 1, 0, 0, 0, 287, 289, 7, 2, 0,
		0, 288, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290,
		291, 1, 0, 0, 0, 291, 54, 1, 0, 0, 0, 292, 294, 7, 2, 0, 0, 293, 292, 1,
		0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0,
		0, 296, 298, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 298, 300, 5, 46, 0, 0, 299,
		301, 7, 2, 0, 0, 300, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 300,
		1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 56, 1, 0, 0, 0, 304, 305, 5, 116,
		0, 0, 305, 306, 5, 114, 0, 0, 306, 307, 5, 117, 0, 0, 307, 314, 5, 101,
		0, 0, 308, 309, 5, 102, 0, 0, 309, 310, 5, 97, 0, 0, 310, 311, 5, 108,
		0, 0, 311, 312, 5, 115, 0, 0, 312, 314, 5, 101, 0, 0, 313, 304, 1, 0, 0,
		0, 313, 308, 1, 0, 0, 0, 314, 58, 1, 0, 0, 0, 315, 321, 5, 34, 0, 0, 316,
		320, 8, 3, 0, 0, 317, 318, 5, 92, 0, 0, 318, 320, 7, 4, 0, 0, 319, 316,
		1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0, 321, 319, 1, 0,
		0, 0, 321, 322, 1, 0, 0, 0, 322, 324, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0,
		324, 325, 5, 34, 0, 0, 325, 60, 1, 0, 0, 0, 326, 332, 5, 39, 0, 0, 327,
		331, 8, 5, 0, 0, 328, 329, 5, 92, 0, 0, 329, 331, 7, 4, 0, 0, 330, 327,
		1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 330, 1, 0,
		0, 0, 332, 333, 1, 0, 0, 0, 333, 335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0,
		335, 336, 5, 39, 0, 0, 336, 62, 1, 0, 0, 0, 337, 341, 7, 6, 0, 0, 338,
		340, 7, 7, 0, 0, 339, 338, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339,
		1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 64, 1, 0, 0, 0, 343, 341, 1, 0,
		0, 0, 344, 345, 5, 45, 0, 0, 345, 346, 5, 62, 0, 0, 346, 66, 1, 0, 0, 0,
		347, 348, 5, 61, 0, 0, 348, 349, 5, 61, 0, 0, 349, 68, 1, 0, 0, 0, 350,
		351, 5, 33, 0, 0, 351, 352, 5, 61, 0, 0, 352, 70, 1, 0, 0, 0, 353, 354,
		5, 60, 0, 0, 354, 72, 1, 0, 0, 0, 355, 356, 5, 62, 0, 0, 356, 74, 1, 0,
		0, 0, 357, 358, 5, 60, 0, 0, 358, 359, 5, 61, 0, 0, 359, 76, 1, 0, 0, 0,
		360, 361, 5, 62, 0, 0, 361, 362, 5, 61, 0, 0, 362, 78, 1, 0, 0, 0, 363,
		364, 5, 61, 0, 0, 364, 80, 1, 0, 0, 0, 365, 366, 5, 42, 0, 0, 366, 367,
		5, 61, 0, 0, 367, 82, 1, 0, 0, 0, 368, 369, 5, 47, 0, 0, 369, 370, 5, 61,
		0, 0, 370, 84, 1, 0, 0, 0, 371, 372, 5, 43, 0, 0, 372, 373, 5, 61, 0, 0,
		373, 86, 1, 0, 0, 0, 374, 375, 5, 45, 0, 0, 375, 376, 5, 61, 0, 0, 376,
		88, 1, 0, 0, 0, 377, 378, 5, 37, 0, 0, 378, 379, 5, 61, 0, 0, 379, 90,
		1, 0, 0, 0, 380, 381, 5, 42, 0, 0, 381, 92, 1, 0, 0, 0, 382, 383, 5, 47,
		0, 0, 383, 94, 1, 0, 0, 0, 384, 385, 5, 43, 0, 0, 385, 96, 1, 0, 0, 0,
		386, 387, 5, 45, 0, 0, 387, 98, 1, 0, 0, 0, 388, 389, 5, 37, 0, 0, 389,
		100, 1, 0, 0, 0, 390, 391, 5, 38, 0, 0, 391, 392, 5, 38, 0, 0, 392, 102,
		1, 0, 0, 0, 393, 394, 5, 124, 0, 0, 394, 395, 5, 124, 0, 0, 395, 104, 1,
		0, 0, 0, 396, 397, 5, 33, 0, 0, 397, 106, 1, 0, 0, 0, 398, 399, 5, 63,
		0, 0, 399, 108, 1, 0, 0, 0, 400, 401, 5, 40, 0, 0, 401, 110, 1, 0, 0, 0,
		402, 403, 5, 41, 0, 0, 403, 112, 1, 0, 0, 0, 404, 405, 5, 123, 0, 0, 405,
		114, 1, 0, 0, 0, 406, 407, 5, 125, 0, 0, 407, 116, 1, 0, 0, 0, 408, 409,
		5, 91, 0, 0, 409, 118, 1, 0, 0, 0, 410, 411, 5, 93, 0, 0, 411, 120, 1,
		0, 0, 0, 412, 413, 5, 92, 0, 0, 413, 122, 1, 0, 0, 0, 414, 415, 5, 44,
		0, 0, 415, 124, 1, 0, 0, 0, 416, 417, 5, 59, 0, 0, 417, 126, 1, 0, 0, 0,
		418, 419, 5, 58, 0, 0, 419, 128, 1, 0, 0, 0, 420, 421, 5, 46, 0, 0, 421,
		130, 1, 0, 0, 0, 13, 0, 134, 144, 155, 290, 295, 302, 313, 319, 321, 330,
		332, 341, 1, 6, 0, 0,
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
	l.GrammarFileName = "Swift.g4"
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
	SwiftLexerKw_INT          = 21
	SwiftLexerKw_DOUBLE       = 22
	SwiftLexerKw_BOOL         = 23
	SwiftLexerKw_STRING       = 24
	SwiftLexerKw_NIL          = 25
	SwiftLexerKw_RANGE        = 26
	SwiftLexerINT             = 27
	SwiftLexerDOUBLE          = 28
	SwiftLexerBOOL            = 29
	SwiftLexerSTRING          = 30
	SwiftLexerCHAR            = 31
	SwiftLexerID              = 32
	SwiftLexerOp_ARROW        = 33
	SwiftLexerOp_EQ           = 34
	SwiftLexerOp_NEQ          = 35
	SwiftLexerOp_LT           = 36
	SwiftLexerOp_GT           = 37
	SwiftLexerOp_LE           = 38
	SwiftLexerOp_GE           = 39
	SwiftLexerOp_ASSIGN       = 40
	SwiftLexerOp_MUL_ASSIGN   = 41
	SwiftLexerOp_DIV_ASSIGN   = 42
	SwiftLexerOp_PLUS_ASSIGN  = 43
	SwiftLexerOp_MINUS_ASSIGN = 44
	SwiftLexerOp_MOD_ASSIGN   = 45
	SwiftLexerOp_MUL          = 46
	SwiftLexerOp_DIV          = 47
	SwiftLexerOp_PLUS         = 48
	SwiftLexerOp_MINUS        = 49
	SwiftLexerOp_MOD          = 50
	SwiftLexerOp_AND          = 51
	SwiftLexerOp_OR           = 52
	SwiftLexerOp_NOT          = 53
	SwiftLexerOp_TERNARY      = 54
	SwiftLexerLPAREN          = 55
	SwiftLexerRPAREN          = 56
	SwiftLexerLBRACE          = 57
	SwiftLexerRBRACE          = 58
	SwiftLexerLBRACKET        = 59
	SwiftLexerRBRACKET        = 60
	SwiftLexerBACKSLASH       = 61
	SwiftLexerCOMMA           = 62
	SwiftLexerSEMICOLON       = 63
	SwiftLexerCOLON           = 64
	SwiftLexerDOT             = 65
)