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
		"'return'", "'do'", "'repeat'", "'in'", "'Int'", "'Float'", "'Bool'",
		"'String'", "'Character'", "'nil'", "'...'", "", "", "", "", "", "",
		"'->'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'='", "'*='",
		"'/='", "'+='", "'-='", "'%='", "'*'", "'/'", "'+'", "'-'", "'%'", "'&&'",
		"'||'", "'!'", "'?'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'\\'",
		"','", "';'", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_REPEAT", "Kw_IN", "Kw_INT", "Kw_FLOAT", "Kw_BOOL", "Kw_STRING",
		"Kw_CHAR", "Kw_NIL", "Kw_RANGE", "INT", "FLOAT", "BOOL", "STRING", "CHAR",
		"ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT", "Op_GT", "Op_LE", "Op_GE",
		"Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN", "Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN",
		"Op_MOD_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS", "Op_MINUS", "Op_MOD",
		"Op_AND", "Op_OR", "Op_NOT", "Op_TERNARY", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"program", "block", "statement", "variableType", "variableDeclaration",
		"variableAssignment", "ifStatement", "ifTail", "elseStatement", "whileStatement",
		"switchStatement", "switchCase", "switchDefault", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 187, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 1, 5,
		1, 33, 8, 1, 10, 1, 12, 1, 36, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		43, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 54,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 61, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 69, 8, 4, 3, 4, 71, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 5, 6, 80, 8, 6, 10, 6, 12, 6, 83, 9, 6, 1, 6, 3, 6, 86,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		109, 8, 10, 10, 10, 12, 10, 112, 9, 10, 1, 10, 3, 10, 115, 8, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 124, 8, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 130, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3,
		13, 147, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 182, 8, 13, 10, 13, 12, 13, 185, 9,
		13, 1, 13, 0, 1, 26, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 0, 8, 1, 0, 21, 26, 1, 0, 4, 5, 2, 0, 41, 41, 44, 45, 1, 0, 47, 48,
		1, 0, 49, 50, 2, 0, 38, 38, 40, 40, 2, 0, 37, 37, 39, 39, 1, 0, 35, 36,
		206, 0, 28, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 44, 1,
		0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 72, 1, 0, 0, 0, 12, 76, 1, 0, 0, 0, 14,
		87, 1, 0, 0, 0, 16, 93, 1, 0, 0, 0, 18, 98, 1, 0, 0, 0, 20, 104, 1, 0,
		0, 0, 22, 118, 1, 0, 0, 0, 24, 125, 1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28,
		29, 3, 2, 1, 0, 29, 30, 5, 0, 0, 1, 30, 1, 1, 0, 0, 0, 31, 33, 3, 4, 2,
		0, 32, 31, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35,
		1, 0, 0, 0, 35, 3, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 43, 3, 10, 5, 0,
		38, 43, 3, 8, 4, 0, 39, 43, 3, 12, 6, 0, 40, 43, 3, 18, 9, 0, 41, 43, 3,
		20, 10, 0, 42, 37, 1, 0, 0, 0, 42, 38, 1, 0, 0, 0, 42, 39, 1, 0, 0, 0,
		42, 40, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 5, 1, 0, 0, 0, 44, 45, 7, 0,
		0, 0, 45, 7, 1, 0, 0, 0, 46, 47, 7, 1, 0, 0, 47, 48, 5, 33, 0, 0, 48, 49,
		5, 65, 0, 0, 49, 50, 3, 6, 3, 0, 50, 51, 5, 41, 0, 0, 51, 53, 3, 26, 13,
		0, 52, 54, 5, 64, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 71,
		1, 0, 0, 0, 55, 56, 7, 1, 0, 0, 56, 57, 5, 33, 0, 0, 57, 58, 5, 41, 0,
		0, 58, 60, 3, 26, 13, 0, 59, 61, 5, 64, 0, 0, 60, 59, 1, 0, 0, 0, 60, 61,
		1, 0, 0, 0, 61, 71, 1, 0, 0, 0, 62, 63, 7, 1, 0, 0, 63, 64, 5, 33, 0, 0,
		64, 65, 5, 65, 0, 0, 65, 66, 3, 6, 3, 0, 66, 68, 5, 55, 0, 0, 67, 69, 5,
		64, 0, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70,
		46, 1, 0, 0, 0, 70, 55, 1, 0, 0, 0, 70, 62, 1, 0, 0, 0, 71, 9, 1, 0, 0,
		0, 72, 73, 5, 33, 0, 0, 73, 74, 7, 2, 0, 0, 74, 75, 3, 26, 13, 0, 75, 11,
		1, 0, 0, 0, 76, 81, 3, 14, 7, 0, 77, 78, 5, 9, 0, 0, 78, 80, 3, 14, 7,
		0, 79, 77, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82,
		1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 86, 3, 16, 8, 0,
		85, 84, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 13, 1, 0, 0, 0, 87, 88, 5,
		8, 0, 0, 88, 89, 3, 26, 13, 0, 89, 90, 5, 58, 0, 0, 90, 91, 3, 2, 1, 0,
		91, 92, 5, 59, 0, 0, 92, 15, 1, 0, 0, 0, 93, 94, 5, 9, 0, 0, 94, 95, 5,
		58, 0, 0, 95, 96, 3, 2, 1, 0, 96, 97, 5, 59, 0, 0, 97, 17, 1, 0, 0, 0,
		98, 99, 5, 14, 0, 0, 99, 100, 3, 26, 13, 0, 100, 101, 5, 58, 0, 0, 101,
		102, 3, 2, 1, 0, 102, 103, 5, 59, 0, 0, 103, 19, 1, 0, 0, 0, 104, 105,
		5, 10, 0, 0, 105, 106, 3, 26, 13, 0, 106, 110, 5, 58, 0, 0, 107, 109, 3,
		22, 11, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0,
		0, 0, 110, 111, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0,
		113, 115, 3, 24, 12, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 117, 5, 59, 0, 0, 117, 21, 1, 0, 0, 0, 118, 119,
		5, 11, 0, 0, 119, 120, 3, 26, 13, 0, 120, 121, 5, 65, 0, 0, 121, 123, 3,
		2, 1, 0, 122, 124, 5, 15, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0,
		0, 124, 23, 1, 0, 0, 0, 125, 126, 5, 12, 0, 0, 126, 127, 5, 65, 0, 0, 127,
		129, 3, 2, 1, 0, 128, 130, 5, 15, 0, 0, 129, 128, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 130, 25, 1, 0, 0, 0, 131, 132, 6, 13, -1, 0, 132, 133, 5, 50,
		0, 0, 133, 147, 3, 26, 13, 19, 134, 135, 5, 54, 0, 0, 135, 147, 3, 26,
		13, 12, 136, 137, 5, 56, 0, 0, 137, 138, 3, 26, 13, 0, 138, 139, 5, 57,
		0, 0, 139, 147, 1, 0, 0, 0, 140, 147, 5, 28, 0, 0, 141, 147, 5, 33, 0,
		0, 142, 147, 5, 29, 0, 0, 143, 147, 5, 31, 0, 0, 144, 147, 5, 32, 0, 0,
		145, 147, 5, 30, 0, 0, 146, 131, 1, 0, 0, 0, 146, 134, 1, 0, 0, 0, 146,
		136, 1, 0, 0, 0, 146, 140, 1, 0, 0, 0, 146, 141, 1, 0, 0, 0, 146, 142,
		1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 145, 1, 0,
		0, 0, 147, 183, 1, 0, 0, 0, 148, 149, 10, 18, 0, 0, 149, 150, 7, 3, 0,
		0, 150, 182, 3, 26, 13, 19, 151, 152, 10, 17, 0, 0, 152, 153, 7, 4, 0,
		0, 153, 182, 3, 26, 13, 18, 154, 155, 10, 16, 0, 0, 155, 156, 5, 51, 0,
		0, 156, 182, 3, 26, 13, 17, 157, 158, 10, 15, 0, 0, 158, 159, 7, 5, 0,
		0, 159, 182, 3, 26, 13, 16, 160, 161, 10, 14, 0, 0, 161, 162, 7, 6, 0,
		0, 162, 182, 3, 26, 13, 15, 163, 164, 10, 13, 0, 0, 164, 165, 7, 7, 0,
		0, 165, 182, 3, 26, 13, 14, 166, 167, 10, 11, 0, 0, 167, 168, 5, 55, 0,
		0, 168, 169, 3, 26, 13, 0, 169, 170, 5, 65, 0, 0, 170, 171, 3, 26, 13,
		12, 171, 182, 1, 0, 0, 0, 172, 173, 10, 10, 0, 0, 173, 174, 5, 52, 0, 0,
		174, 182, 3, 26, 13, 11, 175, 176, 10, 9, 0, 0, 176, 177, 5, 53, 0, 0,
		177, 182, 3, 26, 13, 10, 178, 179, 10, 8, 0, 0, 179, 180, 5, 27, 0, 0,
		180, 182, 3, 26, 13, 9, 181, 148, 1, 0, 0, 0, 181, 151, 1, 0, 0, 0, 181,
		154, 1, 0, 0, 0, 181, 157, 1, 0, 0, 0, 181, 160, 1, 0, 0, 0, 181, 163,
		1, 0, 0, 0, 181, 166, 1, 0, 0, 0, 181, 172, 1, 0, 0, 0, 181, 175, 1, 0,
		0, 0, 181, 178, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 27, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 15, 34,
		42, 53, 60, 68, 70, 81, 85, 110, 114, 123, 129, 146, 181, 183,
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
	SwiftParserKw_INT          = 21
	SwiftParserKw_FLOAT        = 22
	SwiftParserKw_BOOL         = 23
	SwiftParserKw_STRING       = 24
	SwiftParserKw_CHAR         = 25
	SwiftParserKw_NIL          = 26
	SwiftParserKw_RANGE        = 27
	SwiftParserINT             = 28
	SwiftParserFLOAT           = 29
	SwiftParserBOOL            = 30
	SwiftParserSTRING          = 31
	SwiftParserCHAR            = 32
	SwiftParserID              = 33
	SwiftParserOp_ARROW        = 34
	SwiftParserOp_EQ           = 35
	SwiftParserOp_NEQ          = 36
	SwiftParserOp_LT           = 37
	SwiftParserOp_GT           = 38
	SwiftParserOp_LE           = 39
	SwiftParserOp_GE           = 40
	SwiftParserOp_ASSIGN       = 41
	SwiftParserOp_MUL_ASSIGN   = 42
	SwiftParserOp_DIV_ASSIGN   = 43
	SwiftParserOp_PLUS_ASSIGN  = 44
	SwiftParserOp_MINUS_ASSIGN = 45
	SwiftParserOp_MOD_ASSIGN   = 46
	SwiftParserOp_MUL          = 47
	SwiftParserOp_DIV          = 48
	SwiftParserOp_PLUS         = 49
	SwiftParserOp_MINUS        = 50
	SwiftParserOp_MOD          = 51
	SwiftParserOp_AND          = 52
	SwiftParserOp_OR           = 53
	SwiftParserOp_NOT          = 54
	SwiftParserOp_TERNARY      = 55
	SwiftParserLPAREN          = 56
	SwiftParserRPAREN          = 57
	SwiftParserLBRACE          = 58
	SwiftParserRBRACE          = 59
	SwiftParserLBRACKET        = 60
	SwiftParserRBRACKET        = 61
	SwiftParserBACKSLASH       = 62
	SwiftParserCOMMA           = 63
	SwiftParserSEMICOLON       = 64
	SwiftParserCOLON           = 65
	SwiftParserDOT             = 66
)

// SwiftParser rules.
const (
	SwiftParserRULE_program             = 0
	SwiftParserRULE_block               = 1
	SwiftParserRULE_statement           = 2
	SwiftParserRULE_variableType        = 3
	SwiftParserRULE_variableDeclaration = 4
	SwiftParserRULE_variableAssignment  = 5
	SwiftParserRULE_ifStatement         = 6
	SwiftParserRULE_ifTail              = 7
	SwiftParserRULE_elseStatement       = 8
	SwiftParserRULE_whileStatement      = 9
	SwiftParserRULE_switchStatement     = 10
	SwiftParserRULE_switchCase          = 11
	SwiftParserRULE_switchDefault       = 12
	SwiftParserRULE_expr                = 13
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
		p.SetState(28)
		p.Block()
	}
	{
		p.SetState(29)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589952304) != 0 {
		{
			p.SetState(31)
			p.Statement()
		}

		p.SetState(36)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.VariableAssignment()
		}

	case SwiftParserKw_LET, SwiftParserKw_VAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.VariableDeclaration()
		}

	case SwiftParserKw_IF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.IfStatement()
		}

	case SwiftParserKw_WHILE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.WhileStatement()
		}

	case SwiftParserKw_SWITCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(41)
			p.SwitchStatement()
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
	Kw_NIL() antlr.TerminalNode

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

func (s *VariableTypeContext) Kw_NIL() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_NIL, 0)
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
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132120576) != 0) {
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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)

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
			p.SetState(47)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.VariableType()
		}
		{
			p.SetState(50)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.expr(0)
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(52)
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
			p.SetState(55)

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
			p.SetState(56)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.expr(0)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(59)
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
			p.SetState(62)

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
			p.SetState(63)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.VariableType()
		}
		{
			p.SetState(66)
			p.Match(SwiftParserOp_TERNARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftParserSEMICOLON {
			{
				p.SetState(67)
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
	p.EnterRule(localctx, 10, SwiftParserRULE_variableAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VariableAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54975581388800) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VariableAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(74)
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
	p.EnterRule(localctx, 12, SwiftParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.IfTail()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(77)
				p.Match(SwiftParserKw_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(78)
				p.IfTail()
			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_ELSE {
		{
			p.SetState(84)
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
	p.EnterRule(localctx, 14, SwiftParserRULE_ifTail)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(SwiftParserKw_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.expr(0)
	}
	{
		p.SetState(89)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
		p.Block()
	}
	{
		p.SetState(91)
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
	p.EnterRule(localctx, 16, SwiftParserRULE_elseStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(SwiftParserKw_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Block()
	}
	{
		p.SetState(96)
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
	p.EnterRule(localctx, 18, SwiftParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(SwiftParserKw_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.expr(0)
	}
	{
		p.SetState(100)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Block()
	}
	{
		p.SetState(102)
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
	p.EnterRule(localctx, 20, SwiftParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(SwiftParserKw_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.expr(0)
	}
	{
		p.SetState(106)
		p.Match(SwiftParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserKw_CASE {
		{
			p.SetState(107)
			p.SwitchCase()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_DEFAULT {
		{
			p.SetState(113)
			p.SwitchDefault()
		}

	}
	{
		p.SetState(116)
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
	p.EnterRule(localctx, 22, SwiftParserRULE_switchCase)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(SwiftParserKw_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.expr(0)
	}
	{
		p.SetState(120)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Block()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(122)
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
	p.EnterRule(localctx, 24, SwiftParserRULE_switchDefault)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(SwiftParserKw_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(SwiftParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Block()
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftParserKw_BREAK {
		{
			p.SetState(128)
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, SwiftParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserOp_MINUS:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(132)
			p.Match(SwiftParserOp_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.expr(19)
		}

	case SwiftParserOp_NOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(134)
			p.Match(SwiftParserOp_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)

			var _x = p.expr(12)

			localctx.(*NotExprContext).right = _x
		}

	case SwiftParserLPAREN:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.expr(0)
		}
		{
			p.SetState(138)
			p.Match(SwiftParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserINT:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(140)
			p.Match(SwiftParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(141)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserFLOAT:
		localctx = NewFloatExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(142)
			p.Match(SwiftParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserSTRING:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.Match(SwiftParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserCHAR:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(SwiftParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserBOOL:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(145)
			p.Match(SwiftParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(149)

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
					p.SetState(150)

					var _x = p.expr(19)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(152)

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
					p.SetState(153)

					var _x = p.expr(18)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 3:
				localctx = NewArithmeticExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(155)

					var _m = p.Match(SwiftParserOp_MOD)

					localctx.(*ArithmeticExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(156)

					var _x = p.expr(17)

					localctx.(*ArithmeticExprContext).right = _x
				}

			case 4:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(158)

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
					p.SetState(159)

					var _x = p.expr(16)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 5:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(161)

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
					p.SetState(162)

					var _x = p.expr(15)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 6:
				localctx = NewComparasionExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparasionExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(164)

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
					p.SetState(165)

					var _x = p.expr(14)

					localctx.(*ComparasionExprContext).right = _x
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*TernaryExprContext).condition = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(167)
					p.Match(SwiftParserOp_TERNARY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(168)

					var _x = p.expr(0)

					localctx.(*TernaryExprContext).cTrue = _x
				}
				{
					p.SetState(169)
					p.Match(SwiftParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(170)

					var _x = p.expr(12)

					localctx.(*TernaryExprContext).cFalse = _x
				}

			case 8:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(173)

					var _m = p.Match(SwiftParserOp_AND)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(174)

					var _x = p.expr(11)

					localctx.(*LogicalExprContext).right = _x
				}

			case 9:
				localctx = NewLogicalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(176)

					var _m = p.Match(SwiftParserOp_OR)

					localctx.(*LogicalExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(177)

					var _x = p.expr(10)

					localctx.(*LogicalExprContext).right = _x
				}

			case 10:
				localctx = NewRangeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RangeExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(178)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(179)
					p.Match(SwiftParserKw_RANGE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(180)

					var _x = p.expr(9)

					localctx.(*RangeExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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

func (p *SwiftParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
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
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
