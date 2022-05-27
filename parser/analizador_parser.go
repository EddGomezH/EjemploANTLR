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
		"", "'print'", "';'", "'('", "')'", "", "", "'true'", "'false'", "",
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='", "'<='",
		"'=='", "'!='", "'||'", "'&&'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE",
		"RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD", "MAYOR",
		"MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO", "OR",
		"AND", "NOT", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"start", "instruccion", "expresion", "imprimir", "finins",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 127, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 15, 8, 0, 10, 0, 12, 0, 18, 9, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 46, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 113, 8, 2, 10,
		2, 12, 2, 116, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 0, 1, 4, 5, 0, 2, 4, 6, 8, 0, 0, 142, 0, 10, 1, 0, 0, 0, 2, 21,
		1, 0, 0, 0, 4, 45, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 124, 1, 0, 0, 0,
		10, 16, 6, 0, -1, 0, 11, 12, 3, 2, 1, 0, 12, 13, 6, 0, -1, 0, 13, 15, 1,
		0, 0, 0, 14, 11, 1, 0, 0, 0, 15, 18, 1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16,
		17, 1, 0, 0, 0, 17, 19, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 19, 20, 6, 0, -1,
		0, 20, 1, 1, 0, 0, 0, 21, 22, 3, 6, 3, 0, 22, 23, 6, 1, -1, 0, 23, 3, 1,
		0, 0, 0, 24, 25, 6, 2, -1, 0, 25, 26, 5, 5, 0, 0, 26, 46, 6, 2, -1, 0,
		27, 28, 5, 6, 0, 0, 28, 46, 6, 2, -1, 0, 29, 30, 5, 7, 0, 0, 30, 46, 6,
		2, -1, 0, 31, 32, 5, 8, 0, 0, 32, 46, 6, 2, -1, 0, 33, 34, 5, 9, 0, 0,
		34, 46, 6, 2, -1, 0, 35, 36, 5, 10, 0, 0, 36, 46, 6, 2, -1, 0, 37, 38,
		5, 12, 0, 0, 38, 39, 3, 4, 2, 15, 39, 40, 6, 2, -1, 0, 40, 46, 1, 0, 0,
		0, 41, 42, 5, 24, 0, 0, 42, 43, 3, 4, 2, 3, 43, 44, 6, 2, -1, 0, 44, 46,
		1, 0, 0, 0, 45, 24, 1, 0, 0, 0, 45, 27, 1, 0, 0, 0, 45, 29, 1, 0, 0, 0,
		45, 31, 1, 0, 0, 0, 45, 33, 1, 0, 0, 0, 45, 35, 1, 0, 0, 0, 45, 37, 1,
		0, 0, 0, 45, 41, 1, 0, 0, 0, 46, 114, 1, 0, 0, 0, 47, 48, 10, 14, 0, 0,
		48, 49, 5, 13, 0, 0, 49, 50, 3, 4, 2, 15, 50, 51, 6, 2, -1, 0, 51, 113,
		1, 0, 0, 0, 52, 53, 10, 13, 0, 0, 53, 54, 5, 14, 0, 0, 54, 55, 3, 4, 2,
		14, 55, 56, 6, 2, -1, 0, 56, 113, 1, 0, 0, 0, 57, 58, 10, 12, 0, 0, 58,
		59, 5, 15, 0, 0, 59, 60, 3, 4, 2, 13, 60, 61, 6, 2, -1, 0, 61, 113, 1,
		0, 0, 0, 62, 63, 10, 11, 0, 0, 63, 64, 5, 11, 0, 0, 64, 65, 3, 4, 2, 12,
		65, 66, 6, 2, -1, 0, 66, 113, 1, 0, 0, 0, 67, 68, 10, 10, 0, 0, 68, 69,
		5, 12, 0, 0, 69, 70, 3, 4, 2, 11, 70, 71, 6, 2, -1, 0, 71, 113, 1, 0, 0,
		0, 72, 73, 10, 9, 0, 0, 73, 74, 5, 21, 0, 0, 74, 75, 3, 4, 2, 10, 75, 76,
		6, 2, -1, 0, 76, 113, 1, 0, 0, 0, 77, 78, 10, 8, 0, 0, 78, 79, 5, 16, 0,
		0, 79, 80, 3, 4, 2, 9, 80, 81, 6, 2, -1, 0, 81, 113, 1, 0, 0, 0, 82, 83,
		10, 7, 0, 0, 83, 84, 5, 20, 0, 0, 84, 85, 3, 4, 2, 8, 85, 86, 6, 2, -1,
		0, 86, 113, 1, 0, 0, 0, 87, 88, 10, 6, 0, 0, 88, 89, 5, 17, 0, 0, 89, 90,
		3, 4, 2, 7, 90, 91, 6, 2, -1, 0, 91, 113, 1, 0, 0, 0, 92, 93, 10, 5, 0,
		0, 93, 94, 5, 18, 0, 0, 94, 95, 3, 4, 2, 6, 95, 96, 6, 2, -1, 0, 96, 113,
		1, 0, 0, 0, 97, 98, 10, 4, 0, 0, 98, 99, 5, 19, 0, 0, 99, 100, 3, 4, 2,
		5, 100, 101, 6, 2, -1, 0, 101, 113, 1, 0, 0, 0, 102, 103, 10, 2, 0, 0,
		103, 104, 5, 22, 0, 0, 104, 105, 3, 4, 2, 3, 105, 106, 6, 2, -1, 0, 106,
		113, 1, 0, 0, 0, 107, 108, 10, 1, 0, 0, 108, 109, 5, 23, 0, 0, 109, 110,
		3, 4, 2, 2, 110, 111, 6, 2, -1, 0, 111, 113, 1, 0, 0, 0, 112, 47, 1, 0,
		0, 0, 112, 52, 1, 0, 0, 0, 112, 57, 1, 0, 0, 0, 112, 62, 1, 0, 0, 0, 112,
		67, 1, 0, 0, 0, 112, 72, 1, 0, 0, 0, 112, 77, 1, 0, 0, 0, 112, 82, 1, 0,
		0, 0, 112, 87, 1, 0, 0, 0, 112, 92, 1, 0, 0, 0, 112, 97, 1, 0, 0, 0, 112,
		102, 1, 0, 0, 0, 112, 107, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 5, 1, 0, 0, 0, 116, 114, 1, 0, 0,
		0, 117, 118, 5, 1, 0, 0, 118, 119, 5, 3, 0, 0, 119, 120, 3, 4, 2, 0, 120,
		121, 5, 4, 0, 0, 121, 122, 3, 8, 4, 0, 122, 123, 6, 3, -1, 0, 123, 7, 1,
		0, 0, 0, 124, 125, 5, 2, 0, 0, 125, 9, 1, 0, 0, 0, 4, 16, 45, 112, 114,
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
	analizadorParserEOF        = antlr.TokenEOF
	analizadorParserRPRINT     = 1
	analizadorParserPUNTOCOMA  = 2
	analizadorParserPARA       = 3
	analizadorParserPARC       = 4
	analizadorParserENTERO     = 5
	analizadorParserFLOTANTE   = 6
	analizadorParserRTRUE      = 7
	analizadorParserRFALSE     = 8
	analizadorParserCHAR       = 9
	analizadorParserSTRING     = 10
	analizadorParserMAS        = 11
	analizadorParserMENOS      = 12
	analizadorParserMUL        = 13
	analizadorParserDIV        = 14
	analizadorParserMOD        = 15
	analizadorParserMAYOR      = 16
	analizadorParserMENOR      = 17
	analizadorParserMAYORIGUAL = 18
	analizadorParserMENORIGUAL = 19
	analizadorParserIGUALIGUAL = 20
	analizadorParserDISTINTO   = 21
	analizadorParserOR         = 22
	analizadorParserAND        = 23
	analizadorParserNOT        = 24
	analizadorParserCOMMENT    = 25
	analizadorParserWS         = 26
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

	// Get_MENOS returns the _MENOS token.
	Get_MENOS() antlr.Token

	// Get_NOT returns the _NOT token.
	Get_NOT() antlr.Token

	// Get_MUL returns the _MUL token.
	Get_MUL() antlr.Token

	// Get_DIV returns the _DIV token.
	Get_DIV() antlr.Token

	// Get_MOD returns the _MOD token.
	Get_MOD() antlr.Token

	// Get_MAS returns the _MAS token.
	Get_MAS() antlr.Token

	// Get_DISTINTO returns the _DISTINTO token.
	Get_DISTINTO() antlr.Token

	// Get_MAYOR returns the _MAYOR token.
	Get_MAYOR() antlr.Token

	// Get_IGUALIGUAL returns the _IGUALIGUAL token.
	Get_IGUALIGUAL() antlr.Token

	// Get_MENOR returns the _MENOR token.
	Get_MENOR() antlr.Token

	// Get_MAYORIGUAL returns the _MAYORIGUAL token.
	Get_MAYORIGUAL() antlr.Token

	// Get_MENORIGUAL returns the _MENORIGUAL token.
	Get_MENORIGUAL() antlr.Token

	// Get_OR returns the _OR token.
	Get_OR() antlr.Token

	// Get_AND returns the _AND token.
	Get_AND() antlr.Token

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

	// Set_MENOS sets the _MENOS token.
	Set_MENOS(antlr.Token)

	// Set_NOT sets the _NOT token.
	Set_NOT(antlr.Token)

	// Set_MUL sets the _MUL token.
	Set_MUL(antlr.Token)

	// Set_DIV sets the _DIV token.
	Set_DIV(antlr.Token)

	// Set_MOD sets the _MOD token.
	Set_MOD(antlr.Token)

	// Set_MAS sets the _MAS token.
	Set_MAS(antlr.Token)

	// Set_DISTINTO sets the _DISTINTO token.
	Set_DISTINTO(antlr.Token)

	// Set_MAYOR sets the _MAYOR token.
	Set_MAYOR(antlr.Token)

	// Set_IGUALIGUAL sets the _IGUALIGUAL token.
	Set_IGUALIGUAL(antlr.Token)

	// Set_MENOR sets the _MENOR token.
	Set_MENOR(antlr.Token)

	// Set_MAYORIGUAL sets the _MAYORIGUAL token.
	Set_MAYORIGUAL(antlr.Token)

	// Set_MENORIGUAL sets the _MENORIGUAL token.
	Set_MENORIGUAL(antlr.Token)

	// Set_OR sets the _OR token.
	Set_OR(antlr.Token)

	// Set_AND sets the _AND token.
	Set_AND(antlr.Token)

	// GetOpi returns the opi rule contexts.
	GetOpi() IExpresionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetOpd returns the opd rule contexts.
	GetOpd() IExpresionContext

	// SetOpi sets the opi rule contexts.
	SetOpi(IExpresionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetOpd sets the opd rule contexts.
	SetOpd(IExpresionContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	nodo        Abstract.Instruccion
	opi         IExpresionContext
	_ENTERO     antlr.Token
	_FLOTANTE   antlr.Token
	_RTRUE      antlr.Token
	_RFALSE     antlr.Token
	_CHAR       antlr.Token
	_STRING     antlr.Token
	_MENOS      antlr.Token
	_expresion  IExpresionContext
	_NOT        antlr.Token
	_MUL        antlr.Token
	opd         IExpresionContext
	_DIV        antlr.Token
	_MOD        antlr.Token
	_MAS        antlr.Token
	_DISTINTO   antlr.Token
	_MAYOR      antlr.Token
	_IGUALIGUAL antlr.Token
	_MENOR      antlr.Token
	_MAYORIGUAL antlr.Token
	_MENORIGUAL antlr.Token
	_OR         antlr.Token
	_AND        antlr.Token
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

func (s *ExpresionContext) Get_MENOS() antlr.Token { return s._MENOS }

func (s *ExpresionContext) Get_NOT() antlr.Token { return s._NOT }

func (s *ExpresionContext) Get_MUL() antlr.Token { return s._MUL }

func (s *ExpresionContext) Get_DIV() antlr.Token { return s._DIV }

func (s *ExpresionContext) Get_MOD() antlr.Token { return s._MOD }

func (s *ExpresionContext) Get_MAS() antlr.Token { return s._MAS }

func (s *ExpresionContext) Get_DISTINTO() antlr.Token { return s._DISTINTO }

func (s *ExpresionContext) Get_MAYOR() antlr.Token { return s._MAYOR }

func (s *ExpresionContext) Get_IGUALIGUAL() antlr.Token { return s._IGUALIGUAL }

func (s *ExpresionContext) Get_MENOR() antlr.Token { return s._MENOR }

func (s *ExpresionContext) Get_MAYORIGUAL() antlr.Token { return s._MAYORIGUAL }

func (s *ExpresionContext) Get_MENORIGUAL() antlr.Token { return s._MENORIGUAL }

func (s *ExpresionContext) Get_OR() antlr.Token { return s._OR }

func (s *ExpresionContext) Get_AND() antlr.Token { return s._AND }

func (s *ExpresionContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ExpresionContext) Set_FLOTANTE(v antlr.Token) { s._FLOTANTE = v }

func (s *ExpresionContext) Set_RTRUE(v antlr.Token) { s._RTRUE = v }

func (s *ExpresionContext) Set_RFALSE(v antlr.Token) { s._RFALSE = v }

func (s *ExpresionContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *ExpresionContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExpresionContext) Set_MENOS(v antlr.Token) { s._MENOS = v }

func (s *ExpresionContext) Set_NOT(v antlr.Token) { s._NOT = v }

func (s *ExpresionContext) Set_MUL(v antlr.Token) { s._MUL = v }

func (s *ExpresionContext) Set_DIV(v antlr.Token) { s._DIV = v }

func (s *ExpresionContext) Set_MOD(v antlr.Token) { s._MOD = v }

func (s *ExpresionContext) Set_MAS(v antlr.Token) { s._MAS = v }

func (s *ExpresionContext) Set_DISTINTO(v antlr.Token) { s._DISTINTO = v }

func (s *ExpresionContext) Set_MAYOR(v antlr.Token) { s._MAYOR = v }

func (s *ExpresionContext) Set_IGUALIGUAL(v antlr.Token) { s._IGUALIGUAL = v }

func (s *ExpresionContext) Set_MENOR(v antlr.Token) { s._MENOR = v }

func (s *ExpresionContext) Set_MAYORIGUAL(v antlr.Token) { s._MAYORIGUAL = v }

func (s *ExpresionContext) Set_MENORIGUAL(v antlr.Token) { s._MENORIGUAL = v }

func (s *ExpresionContext) Set_OR(v antlr.Token) { s._OR = v }

func (s *ExpresionContext) Set_AND(v antlr.Token) { s._AND = v }

func (s *ExpresionContext) GetOpi() IExpresionContext { return s.opi }

func (s *ExpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionContext) GetOpd() IExpresionContext { return s.opd }

func (s *ExpresionContext) SetOpi(v IExpresionContext) { s.opi = v }

func (s *ExpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionContext) SetOpd(v IExpresionContext) { s.opd = v }

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

func (s *ExpresionContext) MENOS() antlr.TerminalNode {
	return s.GetToken(analizadorParserMENOS, 0)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ExpresionContext) NOT() antlr.TerminalNode {
	return s.GetToken(analizadorParserNOT, 0)
}

func (s *ExpresionContext) MUL() antlr.TerminalNode {
	return s.GetToken(analizadorParserMUL, 0)
}

func (s *ExpresionContext) DIV() antlr.TerminalNode {
	return s.GetToken(analizadorParserDIV, 0)
}

func (s *ExpresionContext) MOD() antlr.TerminalNode {
	return s.GetToken(analizadorParserMOD, 0)
}

func (s *ExpresionContext) MAS() antlr.TerminalNode {
	return s.GetToken(analizadorParserMAS, 0)
}

func (s *ExpresionContext) DISTINTO() antlr.TerminalNode {
	return s.GetToken(analizadorParserDISTINTO, 0)
}

func (s *ExpresionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(analizadorParserMAYOR, 0)
}

func (s *ExpresionContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(analizadorParserIGUALIGUAL, 0)
}

func (s *ExpresionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(analizadorParserMENOR, 0)
}

func (s *ExpresionContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(analizadorParserMAYORIGUAL, 0)
}

func (s *ExpresionContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(analizadorParserMENORIGUAL, 0)
}

func (s *ExpresionContext) OR() antlr.TerminalNode {
	return s.GetToken(analizadorParserOR, 0)
}

func (s *ExpresionContext) AND() antlr.TerminalNode {
	return s.GetToken(analizadorParserAND, 0)
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
	return p.expresion(0)
}

func (p *analizadorParser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, analizadorParserRULE_expresion, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case analizadorParserENTERO:
		{
			p.SetState(25)

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
		{
			p.SetState(27)

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
		{
			p.SetState(29)

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
		{
			p.SetState(31)

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
		{
			p.SetState(33)

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
		{
			p.SetState(35)

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

	case analizadorParserMENOS:
		{
			p.SetState(37)

			var _m = p.Match(analizadorParserMENOS)

			localctx.(*ExpresionContext)._MENOS = _m
		}
		{
			p.SetState(38)

			var _x = p.expresion(15)

			localctx.(*ExpresionContext)._expresion = _x
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewAritmetica(TS.MENOS, localctx.(*ExpresionContext).Get_expresion().GetNodo(), nil, (func() int {
			if localctx.(*ExpresionContext).Get_MENOS() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_MENOS().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_MENOS() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_MENOS().GetColumn()
			}
		}()))

	case analizadorParserNOT:
		{
			p.SetState(41)

			var _m = p.Match(analizadorParserNOT)

			localctx.(*ExpresionContext)._NOT = _m
		}
		{
			p.SetState(42)

			var _x = p.expresion(3)

			localctx.(*ExpresionContext)._expresion = _x
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewLogica(TS.NOT, localctx.(*ExpresionContext).Get_expresion().GetNodo(), localctx.(*ExpresionContext).Get_expresion().GetNodo(), (func() int {
			if localctx.(*ExpresionContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_NOT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_NOT().GetColumn()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(48)

					var _m = p.Match(analizadorParserMUL)

					localctx.(*ExpresionContext)._MUL = _m
				}
				{
					p.SetState(49)

					var _x = p.expresion(15)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewAritmetica(TS.POR, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MUL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MUL().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MUL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MUL().GetColumn()
					}
				}()))

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(53)

					var _m = p.Match(analizadorParserDIV)

					localctx.(*ExpresionContext)._DIV = _m
				}
				{
					p.SetState(54)

					var _x = p.expresion(14)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewAritmetica(TS.DIV, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_DIV() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_DIV().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_DIV() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_DIV().GetColumn()
					}
				}()))

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(57)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(58)

					var _m = p.Match(analizadorParserMOD)

					localctx.(*ExpresionContext)._MOD = _m
				}
				{
					p.SetState(59)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewAritmetica(TS.MOD, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MOD() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MOD().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MOD() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MOD().GetColumn()
					}
				}()))

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(63)

					var _m = p.Match(analizadorParserMAS)

					localctx.(*ExpresionContext)._MAS = _m
				}
				{
					p.SetState(64)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewAritmetica(TS.MAS, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MAS() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MAS().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MAS() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MAS().GetColumn()
					}
				}()))

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(68)

					var _m = p.Match(analizadorParserMENOS)

					localctx.(*ExpresionContext)._MENOS = _m
				}
				{
					p.SetState(69)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewAritmetica(TS.MENOS, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MENOS() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MENOS().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MENOS() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MENOS().GetColumn()
					}
				}()))

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(73)

					var _m = p.Match(analizadorParserDISTINTO)

					localctx.(*ExpresionContext)._DISTINTO = _m
				}
				{
					p.SetState(74)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewRelacional(TS.DIFERENTE, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_DISTINTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_DISTINTO().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_DISTINTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_DISTINTO().GetColumn()
					}
				}()))

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(78)

					var _m = p.Match(analizadorParserMAYOR)

					localctx.(*ExpresionContext)._MAYOR = _m
				}
				{
					p.SetState(79)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewRelacional(TS.MAYORQUE, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MAYOR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MAYOR().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MAYOR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MAYOR().GetColumn()
					}
				}()))

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(83)

					var _m = p.Match(analizadorParserIGUALIGUAL)

					localctx.(*ExpresionContext)._IGUALIGUAL = _m
				}
				{
					p.SetState(84)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewRelacional(TS.IGUALIGUAL, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_IGUALIGUAL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_IGUALIGUAL().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_IGUALIGUAL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_IGUALIGUAL().GetColumn()
					}
				}()))

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(88)

					var _m = p.Match(analizadorParserMENOR)

					localctx.(*ExpresionContext)._MENOR = _m
				}
				{
					p.SetState(89)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewRelacional(TS.MENORQUE, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MENOR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MENOR().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MENOR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MENOR().GetColumn()
					}
				}()))

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(93)

					var _m = p.Match(analizadorParserMAYORIGUAL)

					localctx.(*ExpresionContext)._MAYORIGUAL = _m
				}
				{
					p.SetState(94)

					var _x = p.expresion(6)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewRelacional(TS.MAYORIGUAL, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MAYORIGUAL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MAYORIGUAL().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MAYORIGUAL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MAYORIGUAL().GetColumn()
					}
				}()))

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(98)

					var _m = p.Match(analizadorParserMENORIGUAL)

					localctx.(*ExpresionContext)._MENORIGUAL = _m
				}
				{
					p.SetState(99)

					var _x = p.expresion(5)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewRelacional(TS.MENORIGUAL, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_MENORIGUAL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MENORIGUAL().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_MENORIGUAL() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_MENORIGUAL().GetColumn()
					}
				}()))

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(103)

					var _m = p.Match(analizadorParserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(104)

					var _x = p.expresion(3)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewLogica(TS.OR, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetColumn()
					}
				}()))

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(108)

					var _m = p.Match(analizadorParserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(109)

					var _x = p.expresion(2)

					localctx.(*ExpresionContext).opd = _x
					localctx.(*ExpresionContext)._expresion = _x
				}
				localctx.(*ExpresionContext).nodo = Expresiones.NewLogica(TS.AND, localctx.(*ExpresionContext).GetOpi().GetNodo(), localctx.(*ExpresionContext).GetOpd().GetNodo(), (func() int {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetColumn()
					}
				}()))

			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
		p.SetState(117)

		var _m = p.Match(analizadorParserRPRINT)

		localctx.(*ImprimirContext)._RPRINT = _m
	}
	{
		p.SetState(118)
		p.Match(analizadorParserPARA)
	}
	{
		p.SetState(119)

		var _x = p.expresion(0)

		localctx.(*ImprimirContext)._expresion = _x
	}
	{
		p.SetState(120)
		p.Match(analizadorParserPARC)
	}
	{
		p.SetState(121)
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
		p.SetState(124)
		p.Match(analizadorParserPUNTOCOMA)
	}

	return localctx
}

func (p *analizadorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *analizadorParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
