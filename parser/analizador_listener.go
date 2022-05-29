// Code generated from analizador.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // analizador

import "github.com/antlr/antlr4/runtime/Go/antlr"

// analizadorListener is a complete listener for a parse tree produced by analizadorParser.
type analizadorListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterFuncion is called when entering the funcion production.
	EnterFuncion(c *FuncionContext)

	// EnterParametroll is called when entering the parametroll production.
	EnterParametroll(c *ParametrollContext)

	// EnterParametrolls is called when entering the parametrolls production.
	EnterParametrolls(c *ParametrollsContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterImprimir is called when entering the imprimir production.
	EnterImprimir(c *ImprimirContext)

	// EnterFinins is called when entering the finins production.
	EnterFinins(c *FininsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitFuncion is called when exiting the funcion production.
	ExitFuncion(c *FuncionContext)

	// ExitParametroll is called when exiting the parametroll production.
	ExitParametroll(c *ParametrollContext)

	// ExitParametrolls is called when exiting the parametrolls production.
	ExitParametrolls(c *ParametrollsContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitImprimir is called when exiting the imprimir production.
	ExitImprimir(c *ImprimirContext)

	// ExitFinins is called when exiting the finins production.
	ExitFinins(c *FininsContext)
}
