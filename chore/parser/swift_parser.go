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
		"'return'", "'do'", "'in'", "'Int'", "'Double'", "'Bool'", "'String'",
		"", "", "", "", "", "", "'->'", "'=='", "'!='", "'<'", "'>'", "'<='",
		"'>='", "'='", "'*='", "'/='", "'+='", "'-='", "'*'", "'/'", "'+'",
		"'-'", "'%'", "'&&'", "'||'", "'!'", "'('", "')'", "'{'", "'}'", "'['",
		"']'", "'\\'", "','", "';'", "':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "COMMENT", "BLOCK_COMMENT", "Kw_LET", "Kw_VAR", "Kw_FUNC",
		"Kw_STRUCT", "Kw_IF", "Kw_ELSE", "Kw_SWITCH", "Kw_CASE", "Kw_DEFAULT",
		"Kw_FOR", "Kw_WHILE", "Kw_BREAK", "Kw_CONTINUE", "Kw_RETURN", "Kw_DO",
		"Kw_IN", "Kw_INT", "Kw_DOUBLE", "Kw_BOOL", "Kw_STRING", "INT", "DOUBLE",
		"BOOL", "STRING", "CHAR", "ID", "Op_ARROW", "Op_EQ", "Op_NEQ", "Op_LT",
		"Op_GT", "Op_LE", "Op_GE", "Op_ASSIGN", "Op_MUL_ASSIGN", "Op_DIV_ASSIGN",
		"Op_PLUS_ASSIGN", "Op_MINUS_ASSIGN", "Op_MUL", "Op_DIV", "Op_PLUS",
		"Op_MINUS", "Op_MOD", "Op_AND", "Op_OR", "Op_NOT", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "BACKSLASH", "COMMA", "SEMICOLON",
		"COLON", "DOT",
	}
	staticData.RuleNames = []string{
		"program", "block", "statement", "declaration", "data_type", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 60, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 17, 8, 1, 10, 1, 12, 1, 20,
		9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 30, 8, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 41, 8, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5,
		55, 8, 5, 10, 5, 12, 5, 58, 9, 5, 1, 5, 0, 1, 10, 6, 0, 2, 4, 6, 8, 10,
		0, 5, 1, 0, 4, 5, 1, 0, 20, 23, 1, 0, 33, 36, 1, 0, 42, 43, 1, 0, 44, 45,
		61, 0, 12, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 21, 1, 0, 0, 0, 6, 23, 1,
		0, 0, 0, 8, 31, 1, 0, 0, 0, 10, 40, 1, 0, 0, 0, 12, 13, 3, 2, 1, 0, 13,
		14, 5, 0, 0, 1, 14, 1, 1, 0, 0, 0, 15, 17, 3, 4, 2, 0, 16, 15, 1, 0, 0,
		0, 17, 20, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 3, 1,
		0, 0, 0, 20, 18, 1, 0, 0, 0, 21, 22, 3, 6, 3, 0, 22, 5, 1, 0, 0, 0, 23,
		24, 7, 0, 0, 0, 24, 29, 5, 29, 0, 0, 25, 26, 5, 59, 0, 0, 26, 30, 3, 8,
		4, 0, 27, 28, 5, 37, 0, 0, 28, 30, 3, 10, 5, 0, 29, 25, 1, 0, 0, 0, 29,
		27, 1, 0, 0, 0, 30, 7, 1, 0, 0, 0, 31, 32, 7, 1, 0, 0, 32, 9, 1, 0, 0,
		0, 33, 34, 6, 5, -1, 0, 34, 41, 5, 24, 0, 0, 35, 41, 5, 29, 0, 0, 36, 37,
		5, 50, 0, 0, 37, 38, 3, 10, 5, 0, 38, 39, 5, 51, 0, 0, 39, 41, 1, 0, 0,
		0, 40, 33, 1, 0, 0, 0, 40, 35, 1, 0, 0, 0, 40, 36, 1, 0, 0, 0, 41, 56,
		1, 0, 0, 0, 42, 43, 10, 7, 0, 0, 43, 44, 7, 2, 0, 0, 44, 55, 3, 10, 5,
		8, 45, 46, 10, 6, 0, 0, 46, 47, 7, 3, 0, 0, 47, 55, 3, 10, 5, 7, 48, 49,
		10, 5, 0, 0, 49, 50, 7, 4, 0, 0, 50, 55, 3, 10, 5, 6, 51, 52, 10, 4, 0,
		0, 52, 53, 5, 46, 0, 0, 53, 55, 3, 10, 5, 5, 54, 42, 1, 0, 0, 0, 54, 45,
		1, 0, 0, 0, 54, 48, 1, 0, 0, 0, 54, 51, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0,
		56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 11, 1, 0, 0, 0, 58, 56, 1,
		0, 0, 0, 5, 18, 29, 40, 54, 56,
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
	SwiftParserKw_IN           = 19
	SwiftParserKw_INT          = 20
	SwiftParserKw_DOUBLE       = 21
	SwiftParserKw_BOOL         = 22
	SwiftParserKw_STRING       = 23
	SwiftParserINT             = 24
	SwiftParserDOUBLE          = 25
	SwiftParserBOOL            = 26
	SwiftParserSTRING          = 27
	SwiftParserCHAR            = 28
	SwiftParserID              = 29
	SwiftParserOp_ARROW        = 30
	SwiftParserOp_EQ           = 31
	SwiftParserOp_NEQ          = 32
	SwiftParserOp_LT           = 33
	SwiftParserOp_GT           = 34
	SwiftParserOp_LE           = 35
	SwiftParserOp_GE           = 36
	SwiftParserOp_ASSIGN       = 37
	SwiftParserOp_MUL_ASSIGN   = 38
	SwiftParserOp_DIV_ASSIGN   = 39
	SwiftParserOp_PLUS_ASSIGN  = 40
	SwiftParserOp_MINUS_ASSIGN = 41
	SwiftParserOp_MUL          = 42
	SwiftParserOp_DIV          = 43
	SwiftParserOp_PLUS         = 44
	SwiftParserOp_MINUS        = 45
	SwiftParserOp_MOD          = 46
	SwiftParserOp_AND          = 47
	SwiftParserOp_OR           = 48
	SwiftParserOp_NOT          = 49
	SwiftParserLPAREN          = 50
	SwiftParserRPAREN          = 51
	SwiftParserLBRACE          = 52
	SwiftParserRBRACE          = 53
	SwiftParserLBRACKET        = 54
	SwiftParserRBRACKET        = 55
	SwiftParserBACKSLASH       = 56
	SwiftParserCOMMA           = 57
	SwiftParserSEMICOLON       = 58
	SwiftParserCOLON           = 59
	SwiftParserDOT             = 60
)

// SwiftParser rules.
const (
	SwiftParserRULE_program     = 0
	SwiftParserRULE_block       = 1
	SwiftParserRULE_statement   = 2
	SwiftParserRULE_declaration = 3
	SwiftParserRULE_data_type   = 4
	SwiftParserRULE_expr        = 5
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
		p.SetState(12)
		p.Block()
	}
	{
		p.SetState(13)
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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftParserKw_LET || _la == SwiftParserKw_VAR {
		{
			p.SetState(15)
			p.Statement()
		}

		p.SetState(20)
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
	Declaration() IDeclarationContext

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

func (s *StatementContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.Declaration()
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

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Kw_VAR() antlr.TerminalNode
	Kw_LET() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Data_type() IData_typeContext
	Op_ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *DeclarationContext) Kw_VAR() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_VAR, 0)
}

func (s *DeclarationContext) Kw_LET() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_LET, 0)
}

func (s *DeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftParserCOLON, 0)
}

func (s *DeclarationContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *DeclarationContext) Op_ASSIGN() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_ASSIGN, 0)
}

func (s *DeclarationContext) Expr() IExprContext {
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

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftParserRULE_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SwiftParserKw_LET || _la == SwiftParserKw_VAR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(24)
		p.Match(SwiftParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserCOLON:
		{
			p.SetState(25)
			p.Match(SwiftParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)
			p.Data_type()
		}

	case SwiftParserOp_ASSIGN:
		{
			p.SetState(27)
			p.Match(SwiftParserOp_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
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

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Kw_INT() antlr.TerminalNode
	Kw_DOUBLE() antlr.TerminalNode
	Kw_STRING() antlr.TerminalNode
	Kw_BOOL() antlr.TerminalNode

	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_data_type
	return p
}

func InitEmptyData_typeContext(p *Data_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_data_type
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) Kw_INT() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_INT, 0)
}

func (s *Data_typeContext) Kw_DOUBLE() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_DOUBLE, 0)
}

func (s *Data_typeContext) Kw_STRING() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_STRING, 0)
}

func (s *Data_typeContext) Kw_BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftParserKw_BOOL, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitData_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SwiftParser) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftParserRULE_data_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
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

type OpContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpContext {
	var p = new(OpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpContext) GetOp() antlr.Token { return s.op }

func (s *OpContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpContext) GetLeft() IExprContext { return s.left }

func (s *OpContext) GetRight() IExprContext { return s.right }

func (s *OpContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpContext) SetRight(v IExprContext) { s.right = v }

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) AllExpr() []IExprContext {
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

func (s *OpContext) Expr(i int) IExprContext {
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

func (s *OpContext) Op_LT() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_LT, 0)
}

func (s *OpContext) Op_GT() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_GT, 0)
}

func (s *OpContext) Op_LE() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_LE, 0)
}

func (s *OpContext) Op_GE() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_GE, 0)
}

func (s *OpContext) Op_MUL() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MUL, 0)
}

func (s *OpContext) Op_DIV() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_DIV, 0)
}

func (s *OpContext) Op_PLUS() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_PLUS, 0)
}

func (s *OpContext) Op_MINUS() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MINUS, 0)
}

func (s *OpContext) Op_MOD() antlr.TerminalNode {
	return s.GetToken(SwiftParserOp_MOD, 0)
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DigitContext struct {
	ExprContext
	digit antlr.Token
}

func NewDigitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DigitContext {
	var p = new(DigitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DigitContext) GetDigit() antlr.Token { return s.digit }

func (s *DigitContext) SetDigit(v antlr.Token) { s.digit = v }

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftParserINT, 0)
}

func (s *DigitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitDigit(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	ExprContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftParserID, 0)
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenContext struct {
	ExprContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserLPAREN, 0)
}

func (s *ParenContext) Expr() IExprContext {
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

func (s *ParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SwiftParserRPAREN, 0)
}

func (s *ParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SwiftVisitor:
		return t.VisitParen(s)

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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SwiftParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserINT:
		localctx = NewDigitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(34)

			var _m = p.Match(SwiftParserINT)

			localctx.(*DigitContext).digit = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(SwiftParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserLPAREN:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.Match(SwiftParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.expr(0)
		}
		{
			p.SetState(38)
			p.Match(SwiftParserRPAREN)
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(43)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&128849018880) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(44)

					var _x = p.expr(8)

					localctx.(*OpContext).right = _x
				}

			case 2:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(46)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_MUL || _la == SwiftParserOp_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(47)

					var _x = p.expr(7)

					localctx.(*OpContext).right = _x
				}

			case 3:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(49)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftParserOp_PLUS || _la == SwiftParserOp_MINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(50)

					var _x = p.expr(6)

					localctx.(*OpContext).right = _x
				}

			case 4:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_expr)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(52)
					p.Match(SwiftParserOp_MOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(53)

					var _x = p.expr(5)

					localctx.(*OpContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	case 5:
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
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
