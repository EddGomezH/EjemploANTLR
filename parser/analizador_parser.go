// Code generated from analizador.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // analizador

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "github.com/emivnajera/Abstract"
import "github.com/emivnajera/Expresiones"
import "github.com/emivnajera/Instrucciones"
import "github.com/emivnajera/TS"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type analizadorParser struct {
	*antlr.BaseParser
}

var analizadorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func analizadorParserInit() {
	staticData := &analizadorParserStaticData
	staticData.literalNames = []string{
		"", "'print'", "';'", "'('", "')'", "", "", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE",
		"RFALSE", "CHAR", "STRING", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"start", "instruccion", "expresion", "imprimir", "finins",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 48, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 15, 8, 0, 10, 0, 12, 0, 18, 9, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 0, 48, 0, 10, 1,
		0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 45,
		1, 0, 0, 0, 10, 16, 6, 0, -1, 0, 11, 12, 3, 2, 1, 0, 12, 13, 6, 0, -1,
		0, 13, 15, 1, 0, 0, 0, 14, 11, 1, 0, 0, 0, 15, 18, 1, 0, 0, 0, 16, 14,
		1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 19, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0,
		19, 20, 6, 0, -1, 0, 20, 1, 1, 0, 0, 0, 21, 22, 3, 6, 3, 0, 22, 23, 6,
		1, -1, 0, 23, 3, 1, 0, 0, 0, 24, 25, 5, 5, 0, 0, 25, 37, 6, 2, -1, 0, 26,
		27, 5, 6, 0, 0, 27, 37, 6, 2, -1, 0, 28, 29, 5, 7, 0, 0, 29, 37, 6, 2,
		-1, 0, 30, 31, 5, 8, 0, 0, 31, 37, 6, 2, -1, 0, 32, 33, 5, 9, 0, 0, 33,
		37, 6, 2, -1, 0, 34, 35, 5, 10, 0, 0, 35, 37, 6, 2, -1, 0, 36, 24, 1, 0,
		0, 0, 36, 26, 1, 0, 0, 0, 36, 28, 1, 0, 0, 0, 36, 30, 1, 0, 0, 0, 36, 32,
		1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0, 38, 39, 5, 1, 0, 0,
		39, 40, 5, 3, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 5, 4, 0, 0, 42, 43, 3,
		8, 4, 0, 43, 44, 6, 3, -1, 0, 44, 7, 1, 0, 0, 0, 45, 46, 5, 2, 0, 0, 46,
		9, 1, 0, 0, 0, 2, 16, 36,
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

// analizadorParserInit initializes any static state used to implement analizadorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewanalizadorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AnalizadorParserInit() {
	staticData := &analizadorParserStaticData
	staticData.once.Do(analizadorParserInit)
}

// NewanalizadorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewanalizadorParser(input antlr.TokenStream) *analizadorParser {
	AnalizadorParserInit()
	this := new(analizadorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &analizadorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "analizador.g4"

	return this
}

// analizadorParser tokens.
const (
	analizadorParserEOF       = antlr.TokenEOF
	analizadorParserRPRINT    = 1
	analizadorParserPUNTOCOMA = 2
	analizadorParserPARA      = 3
	analizadorParserPARC      = 4
	analizadorParserENTERO    = 5
	analizadorParserFLOTANTE  = 6
	analizadorParserRTRUE     = 7
	analizadorParserRFALSE    = 8
	analizadorParserCHAR      = 9
	analizadorParserSTRING    = 10
	analizadorParserCOMMENT   = 11
	analizadorParserWS        = 12
)

// analizadorParser rules.
const (
	analizadorParserRULE_start       = 0
	analizadorParserRULE_instruccion = 1
	analizadorParserRULE_expresion   = 2
	analizadorParserRULE_imprimir    = 3
	analizadorParserRULE_finins      = 4
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_instruccion IInstruccionContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *StartContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *StartContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
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

	return t.(IInstruccionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *analizadorParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, analizadorParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	instrucciones := []Abstract.Instruccion{}
	TSGlobal := TS.TablaSimbolos{}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == analizadorParserRPRINT {
		{
			p.SetState(11)

			var _x = p.Instruccion()

			localctx.(*StartContext)._instruccion = _x
		}
		instrucciones = append(instrucciones, localctx.(*StartContext).Get_instruccion().GetNodo())

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	for _, n := range instrucciones {
		n.Interpretar(&TSGlobal)
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_imprimir returns the _imprimir rule contexts.
	Get_imprimir() IImprimirContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	nodo      Abstract.Instruccion
	_imprimir IImprimirContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_imprimir() IImprimirContext { return s._imprimir }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *InstruccionContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *InstruccionContext) Imprimir() IImprimirContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImprimirContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImprimirContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *analizadorParser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, analizadorParserRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)

		var _x = p.Imprimir()

		localctx.(*InstruccionContext)._imprimir = _x
	}
	localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_imprimir().GetNodo()

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_FLOTANTE returns the _FLOTANTE token.
	Get_FLOTANTE() antlr.Token

	// Get_RTRUE returns the _RTRUE token.
	Get_RTRUE() antlr.Token

	// Get_RFALSE returns the _RFALSE token.
	Get_RFALSE() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_FLOTANTE sets the _FLOTANTE token.
	Set_FLOTANTE(antlr.Token)

	// Set_RTRUE sets the _RTRUE token.
	Set_RTRUE(antlr.Token)

	// Set_RFALSE sets the _RFALSE token.
	Set_RFALSE(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	nodo      Abstract.Instruccion
	_ENTERO   antlr.Token
	_FLOTANTE antlr.Token
	_RTRUE    antlr.Token
	_RFALSE   antlr.Token
	_CHAR     antlr.Token
	_STRING   antlr.Token
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *ExpresionContext) Get_FLOTANTE() antlr.Token { return s._FLOTANTE }

func (s *ExpresionContext) Get_RTRUE() antlr.Token { return s._RTRUE }

func (s *ExpresionContext) Get_RFALSE() antlr.Token { return s._RFALSE }

func (s *ExpresionContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *ExpresionContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExpresionContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ExpresionContext) Set_FLOTANTE(v antlr.Token) { s._FLOTANTE = v }

func (s *ExpresionContext) Set_RTRUE(v antlr.Token) { s._RTRUE = v }

func (s *ExpresionContext) Set_RFALSE(v antlr.Token) { s._RFALSE = v }

func (s *ExpresionContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *ExpresionContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExpresionContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *ExpresionContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *ExpresionContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(analizadorParserENTERO, 0)
}

func (s *ExpresionContext) FLOTANTE() antlr.TerminalNode {
	return s.GetToken(analizadorParserFLOTANTE, 0)
}

func (s *ExpresionContext) RTRUE() antlr.TerminalNode {
	return s.GetToken(analizadorParserRTRUE, 0)
}

func (s *ExpresionContext) RFALSE() antlr.TerminalNode {
	return s.GetToken(analizadorParserRFALSE, 0)
}

func (s *ExpresionContext) CHAR() antlr.TerminalNode {
	return s.GetToken(analizadorParserCHAR, 0)
}

func (s *ExpresionContext) STRING() antlr.TerminalNode {
	return s.GetToken(analizadorParserSTRING, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *analizadorParser) Expresion() (localctx IExpresionContext) {
	this := p
	_ = this

	localctx = NewExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, analizadorParserRULE_expresion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case analizadorParserENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)

			var _m = p.Match(analizadorParserENTERO)

			localctx.(*ExpresionContext)._ENTERO = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewPrimitivo((func() string {
			if localctx.(*ExpresionContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_ENTERO().GetText()
			}
		}()), TS.ENTERO, (func() int {
			if localctx.(*ExpresionContext).Get_ENTERO() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_ENTERO().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_ENTERO() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_ENTERO().GetColumn()
			}
		}()))

	case analizadorParserFLOTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)

			var _m = p.Match(analizadorParserFLOTANTE)

			localctx.(*ExpresionContext)._FLOTANTE = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewPrimitivo((func() string {
			if localctx.(*ExpresionContext).Get_FLOTANTE() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_FLOTANTE().GetText()
			}
		}()), TS.FLOAT, (func() int {
			if localctx.(*ExpresionContext).Get_FLOTANTE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_FLOTANTE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_FLOTANTE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_FLOTANTE().GetColumn()
			}
		}()))

	case analizadorParserRTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)

			var _m = p.Match(analizadorParserRTRUE)

			localctx.(*ExpresionContext)._RTRUE = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewPrimitivo((func() string {
			if localctx.(*ExpresionContext).Get_RTRUE() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_RTRUE().GetText()
			}
		}()), TS.BOOLEAN, (func() int {
			if localctx.(*ExpresionContext).Get_RTRUE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_RTRUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_RTRUE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_RTRUE().GetColumn()
			}
		}()))

	case analizadorParserRFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)

			var _m = p.Match(analizadorParserRFALSE)

			localctx.(*ExpresionContext)._RFALSE = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewPrimitivo((func() string {
			if localctx.(*ExpresionContext).Get_RFALSE() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_RFALSE().GetText()
			}
		}()), TS.BOOLEAN, (func() int {
			if localctx.(*ExpresionContext).Get_RFALSE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_RFALSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_RFALSE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_RFALSE().GetColumn()
			}
		}()))

	case analizadorParserCHAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)

			var _m = p.Match(analizadorParserCHAR)

			localctx.(*ExpresionContext)._CHAR = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewPrimitivo(strings.Replace((func() string {
			if localctx.(*ExpresionContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CHAR().GetText()
			}
		}()), "'", "", 2), TS.CARACTER, (func() int {
			if localctx.(*ExpresionContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_CHAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_CHAR().GetColumn()
			}
		}()))

	case analizadorParserSTRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(34)

			var _m = p.Match(analizadorParserSTRING)

			localctx.(*ExpresionContext)._STRING = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewPrimitivo(strings.Replace((func() string {
			if localctx.(*ExpresionContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_STRING().GetText()
			}
		}()), "\"", "", -1), TS.CADENA, (func() int {
			if localctx.(*ExpresionContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_STRING().GetColumn()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImprimirContext is an interface to support dynamic dispatch.
type IImprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RPRINT returns the _RPRINT token.
	Get_RPRINT() antlr.Token

	// Set_RPRINT sets the _RPRINT token.
	Set_RPRINT(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsImprimirContext differentiates from other interfaces.
	IsImprimirContext()
}

type ImprimirContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	nodo       Abstract.Instruccion
	_RPRINT    antlr.Token
	_expresion IExpresionContext
}

func NewEmptyImprimirContext() *ImprimirContext {
	var p = new(ImprimirContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_imprimir
	return p
}

func (*ImprimirContext) IsImprimirContext() {}

func NewImprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprimirContext {
	var p = new(ImprimirContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_imprimir

	return p
}

func (s *ImprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprimirContext) Get_RPRINT() antlr.Token { return s._RPRINT }

func (s *ImprimirContext) Set_RPRINT(v antlr.Token) { s._RPRINT = v }

func (s *ImprimirContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ImprimirContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ImprimirContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *ImprimirContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *ImprimirContext) RPRINT() antlr.TerminalNode {
	return s.GetToken(analizadorParserRPRINT, 0)
}

func (s *ImprimirContext) PARA() antlr.TerminalNode {
	return s.GetToken(analizadorParserPARA, 0)
}

func (s *ImprimirContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ImprimirContext) PARC() antlr.TerminalNode {
	return s.GetToken(analizadorParserPARC, 0)
}

func (s *ImprimirContext) Finins() IFininsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFininsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFininsContext)
}

func (s *ImprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImprimirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterImprimir(s)
	}
}

func (s *ImprimirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitImprimir(s)
	}
}

func (p *analizadorParser) Imprimir() (localctx IImprimirContext) {
	this := p
	_ = this

	localctx = NewImprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, analizadorParserRULE_imprimir)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)

		var _m = p.Match(analizadorParserRPRINT)

		localctx.(*ImprimirContext)._RPRINT = _m
	}
	{
		p.SetState(39)
		p.Match(analizadorParserPARA)
	}
	{
		p.SetState(40)

		var _x = p.Expresion()

		localctx.(*ImprimirContext)._expresion = _x
	}
	{
		p.SetState(41)
		p.Match(analizadorParserPARC)
	}
	{
		p.SetState(42)
		p.Finins()
	}
	localctx.(*ImprimirContext).nodo = Instrucciones.NewImprimir(localctx.(*ImprimirContext).Get_expresion().GetNodo(), (func() int {
		if localctx.(*ImprimirContext).Get_RPRINT() == nil {
			return 0
		} else {
			return localctx.(*ImprimirContext).Get_RPRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*ImprimirContext).Get_RPRINT() == nil {
			return 0
		} else {
			return localctx.(*ImprimirContext).Get_RPRINT().GetColumn()
		}
	}()))

	return localctx
}

// IFininsContext is an interface to support dynamic dispatch.
type IFininsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFininsContext differentiates from other interfaces.
	IsFininsContext()
}

type FininsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFininsContext() *FininsContext {
	var p = new(FininsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_finins
	return p
}

func (*FininsContext) IsFininsContext() {}

func NewFininsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FininsContext {
	var p = new(FininsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_finins

	return p
}

func (s *FininsContext) GetParser() antlr.Parser { return s.parser }

func (s *FininsContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(analizadorParserPUNTOCOMA, 0)
}

func (s *FininsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FininsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FininsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterFinins(s)
	}
}

func (s *FininsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitFinins(s)
	}
}

func (p *analizadorParser) Finins() (localctx IFininsContext) {
	this := p
	_ = this

	localctx = NewFininsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, analizadorParserRULE_finins)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(analizadorParserPUNTOCOMA)
	}

	return localctx
}
