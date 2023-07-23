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
		"", "'*'", "'+'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "MUL", "ADD", "LB", "RB", "DIGIT", "WS",
	}
	staticData.RuleNames = []string{
		"l", "e", "t", "f",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 41, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 18, 8, 1, 10, 1, 12, 1,
		21, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 29, 8, 2, 10, 2, 12,
		2, 32, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 39, 8, 3, 1, 3, 0, 2,
		2, 4, 4, 0, 2, 4, 6, 0, 0, 39, 0, 8, 1, 0, 0, 0, 2, 11, 1, 0, 0, 0, 4,
		22, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1,
		10, 1, 1, 0, 0, 0, 11, 12, 6, 1, -1, 0, 12, 13, 3, 4, 2, 0, 13, 19, 1,
		0, 0, 0, 14, 15, 10, 2, 0, 0, 15, 16, 5, 2, 0, 0, 16, 18, 3, 4, 2, 0, 17,
		14, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0,
		0, 20, 3, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22, 23, 6, 2, -1, 0, 23, 24,
		3, 6, 3, 0, 24, 30, 1, 0, 0, 0, 25, 26, 10, 2, 0, 0, 26, 27, 5, 1, 0, 0,
		27, 29, 3, 6, 3, 0, 28, 25, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 5, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33,
		34, 5, 3, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 4, 0, 0, 36, 39, 1, 0, 0,
		0, 37, 39, 5, 5, 0, 0, 38, 33, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 7, 1,
		0, 0, 0, 3, 19, 30, 38,
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
	SwiftParserEOF   = antlr.TokenEOF
	SwiftParserMUL   = 1
	SwiftParserADD   = 2
	SwiftParserLB    = 3
	SwiftParserRB    = 4
	SwiftParserDIGIT = 5
	SwiftParserWS    = 6
)

// SwiftParser rules.
const (
	SwiftParserRULE_l = 0
	SwiftParserRULE_e = 1
	SwiftParserRULE_t = 2
	SwiftParserRULE_f = 3
)

// ILContext is an interface to support dynamic dispatch.
type ILContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	E() IEContext
	EOF() antlr.TerminalNode

	// IsLContext differentiates from other interfaces.
	IsLContext()
}

type LContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLContext() *LContext {
	var p = new(LContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_l
	return p
}

func InitEmptyLContext(p *LContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_l
}

func (*LContext) IsLContext() {}

func NewLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LContext {
	var p = new(LContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_l

	return p
}

func (s *LContext) GetParser() antlr.Parser { return s.parser }

func (s *LContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *LContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftParserEOF, 0)
}

func (s *LContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterL(s)
	}
}

func (s *LContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitL(s)
	}
}

func (p *SwiftParser) L() (localctx ILContext) {
	localctx = NewLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftParserRULE_l)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.e(0)
	}
	{
		p.SetState(9)
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

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyAll(ctx *EContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassTContext struct {
	EContext
}

func NewPassTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassTContext {
	var p = new(PassTContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *PassTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassTContext) T() ITContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *PassTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterPassT(s)
	}
}

func (s *PassTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitPassT(s)
	}
}

type SumContext struct {
	EContext
}

func NewSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumContext {
	var p = new(SumContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftParserADD, 0)
}

func (s *SumContext) T() ITContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitSum(s)
	}
}

func (p *SwiftParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *SwiftParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, SwiftParserRULE_e, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPassTContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(12)
		p.t(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(19)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSumContext(p, NewEContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_e)
			p.SetState(14)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(15)
				p.Match(SwiftParserADD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(16)
				p.t(0)
			}

		}
		p.SetState(21)
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_t
	return p
}

func InitEmptyTContext(p *TContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_t
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) CopyAll(ctx *TContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassFContext struct {
	TContext
}

func NewPassFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassFContext {
	var p = new(PassFContext)

	InitEmptyTContext(&p.TContext)
	p.parser = parser
	p.CopyAll(ctx.(*TContext))

	return p
}

func (s *PassFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassFContext) F() IFContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *PassFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterPassF(s)
	}
}

func (s *PassFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitPassF(s)
	}
}

type MulContext struct {
	TContext
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	InitEmptyTContext(&p.TContext)
	p.parser = parser
	p.CopyAll(ctx.(*TContext))

	return p
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) T() ITContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *MulContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftParserMUL, 0)
}

func (s *MulContext) F() IFContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitMul(s)
	}
}

func (p *SwiftParser) T() (localctx ITContext) {
	return p.t(0)
}

func (p *SwiftParser) t(_p int) (localctx ITContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, SwiftParserRULE_t, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPassFContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(23)
		p.F()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(30)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMulContext(p, NewTContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, SwiftParserRULE_t)
			p.SetState(25)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(26)
				p.Match(SwiftParserMUL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(27)
				p.F()
			}

		}
		p.SetState(32)
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFContext is an interface to support dynamic dispatch.
type IFContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFContext differentiates from other interfaces.
	IsFContext()
}

type FContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_f
	return p
}

func InitEmptyFContext(p *FContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftParserRULE_f
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) CopyAll(ctx *FContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassEContext struct {
	FContext
}

func NewPassEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassEContext {
	var p = new(PassEContext)

	InitEmptyFContext(&p.FContext)
	p.parser = parser
	p.CopyAll(ctx.(*FContext))

	return p
}

func (s *PassEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassEContext) LB() antlr.TerminalNode {
	return s.GetToken(SwiftParserLB, 0)
}

func (s *PassEContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *PassEContext) RB() antlr.TerminalNode {
	return s.GetToken(SwiftParserRB, 0)
}

func (s *PassEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterPassE(s)
	}
}

func (s *PassEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitPassE(s)
	}
}

type DigitContext struct {
	FContext
}

func NewDigitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DigitContext {
	var p = new(DigitContext)

	InitEmptyFContext(&p.FContext)
	p.parser = parser
	p.CopyAll(ctx.(*FContext))

	return p
}

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(SwiftParserDIGIT, 0)
}

func (s *DigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.EnterDigit(s)
	}
}

func (s *DigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftListener); ok {
		listenerT.ExitDigit(s)
	}
}

func (p *SwiftParser) F() (localctx IFContext) {
	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftParserRULE_f)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftParserLB:
		localctx = NewPassEContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(SwiftParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.e(0)
		}
		{
			p.SetState(35)
			p.Match(SwiftParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftParserDIGIT:
		localctx = NewDigitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(SwiftParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	case 1:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	case 2:
		var t *TContext = nil
		if localctx != nil {
			t = localctx.(*TContext)
		}
		return p.T_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftParser) T_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
