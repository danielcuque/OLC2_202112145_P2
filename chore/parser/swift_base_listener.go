// Code generated from Swift.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swift

import "github.com/antlr4-go/antlr/v4"

// BaseSwiftListener is a complete listener for a parse tree produced by SwiftParser.
type BaseSwiftListener struct{}

var _ SwiftListener = &BaseSwiftListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterL is called when production l is entered.
func (s *BaseSwiftListener) EnterL(ctx *LContext) {}

// ExitL is called when production l is exited.
func (s *BaseSwiftListener) ExitL(ctx *LContext) {}

// EnterPassT is called when production PassT is entered.
func (s *BaseSwiftListener) EnterPassT(ctx *PassTContext) {}

// ExitPassT is called when production PassT is exited.
func (s *BaseSwiftListener) ExitPassT(ctx *PassTContext) {}

// EnterSum is called when production Sum is entered.
func (s *BaseSwiftListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production Sum is exited.
func (s *BaseSwiftListener) ExitSum(ctx *SumContext) {}

// EnterPassF is called when production PassF is entered.
func (s *BaseSwiftListener) EnterPassF(ctx *PassFContext) {}

// ExitPassF is called when production PassF is exited.
func (s *BaseSwiftListener) ExitPassF(ctx *PassFContext) {}

// EnterMul is called when production Mul is entered.
func (s *BaseSwiftListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production Mul is exited.
func (s *BaseSwiftListener) ExitMul(ctx *MulContext) {}

// EnterPassE is called when production PassE is entered.
func (s *BaseSwiftListener) EnterPassE(ctx *PassEContext) {}

// ExitPassE is called when production PassE is exited.
func (s *BaseSwiftListener) ExitPassE(ctx *PassEContext) {}

// EnterDigit is called when production Digit is entered.
func (s *BaseSwiftListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production Digit is exited.
func (s *BaseSwiftListener) ExitDigit(ctx *DigitContext) {}
