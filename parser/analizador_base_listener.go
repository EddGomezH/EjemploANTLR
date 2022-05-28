// Code generated from analizador.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // analizador

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseanalizadorListener is a complete listener for a parse tree produced by analizadorParser.
type BaseanalizadorListener struct{}

var _ analizadorListener = &BaseanalizadorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseanalizadorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseanalizadorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseanalizadorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseanalizadorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseanalizadorListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseanalizadorListener) ExitStart(ctx *StartContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseanalizadorListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseanalizadorListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BaseanalizadorListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BaseanalizadorListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseanalizadorListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseanalizadorListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseanalizadorListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseanalizadorListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterImprimir is called when production imprimir is entered.
func (s *BaseanalizadorListener) EnterImprimir(ctx *ImprimirContext) {}

// ExitImprimir is called when production imprimir is exited.
func (s *BaseanalizadorListener) ExitImprimir(ctx *ImprimirContext) {}

// EnterFinins is called when production finins is entered.
func (s *BaseanalizadorListener) EnterFinins(ctx *FininsContext) {}

// ExitFinins is called when production finins is exited.
func (s *BaseanalizadorListener) ExitFinins(ctx *FininsContext) {}
