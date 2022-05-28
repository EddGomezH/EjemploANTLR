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
		"", "'print'", "'var'", "';'", "'('", "')'", "", "", "'true'", "'false'",
		"", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'>='", "'<='",
		"'=='", "'!='", "'||'", "'&&'", "'!'", "", "'='",
	}
	staticData.symbolicNames = []string{
		"", "RPRINT", "RVAR", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE",
		"RTRUE", "RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD",
		"MAYOR", "MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO",
		"OR", "AND", "NOT", "ID", "IGUAL", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"start", "instruccion", "declaracion", "asignacion", "expresion", "imprimir",
		"finins",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 154, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 19, 8, 0, 10,
		0, 12, 0, 22, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 35, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 73, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 5, 4, 140, 8, 4, 10, 4, 12, 4, 143, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 1, 8, 7, 0, 2, 4, 6, 8, 10, 12, 0,
		0, 170, 0, 14, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 43,
		1, 0, 0, 0, 8, 72, 1, 0, 0, 0, 10, 144, 1, 0, 0, 0, 12, 151, 1, 0, 0, 0,
		14, 20, 6, 0, -1, 0, 15, 16, 3, 2, 1, 0, 16, 17, 6, 0, -1, 0, 17, 19, 1,
		0, 0, 0, 18, 15, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20,
		21, 1, 0, 0, 0, 21, 23, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24, 6, 0, -1,
		0, 24, 1, 1, 0, 0, 0, 25, 26, 3, 10, 5, 0, 26, 27, 6, 1, -1, 0, 27, 35,
		1, 0, 0, 0, 28, 29, 3, 4, 2, 0, 29, 30, 6, 1, -1, 0, 30, 35, 1, 0, 0, 0,
		31, 32, 3, 6, 3, 0, 32, 33, 6, 1, -1, 0, 33, 35, 1, 0, 0, 0, 34, 25, 1,
		0, 0, 0, 34, 28, 1, 0, 0, 0, 34, 31, 1, 0, 0, 0, 35, 3, 1, 0, 0, 0, 36,
		37, 5, 2, 0, 0, 37, 38, 5, 26, 0, 0, 38, 39, 5, 27, 0, 0, 39, 40, 3, 8,
		4, 0, 40, 41, 3, 12, 6, 0, 41, 42, 6, 2, -1, 0, 42, 5, 1, 0, 0, 0, 43,
		44, 5, 26, 0, 0, 44, 45, 5, 27, 0, 0, 45, 46, 3, 8, 4, 0, 46, 47, 3, 12,
		6, 0, 47, 48, 6, 3, -1, 0, 48, 7, 1, 0, 0, 0, 49, 50, 6, 4, -1, 0, 50,
		51, 5, 6, 0, 0, 51, 73, 6, 4, -1, 0, 52, 53, 5, 7, 0, 0, 53, 73, 6, 4,
		-1, 0, 54, 55, 5, 8, 0, 0, 55, 73, 6, 4, -1, 0, 56, 57, 5, 9, 0, 0, 57,
		73, 6, 4, -1, 0, 58, 59, 5, 10, 0, 0, 59, 73, 6, 4, -1, 0, 60, 61, 5, 11,
		0, 0, 61, 73, 6, 4, -1, 0, 62, 63, 5, 13, 0, 0, 63, 64, 3, 8, 4, 16, 64,
		65, 6, 4, -1, 0, 65, 73, 1, 0, 0, 0, 66, 67, 5, 25, 0, 0, 67, 68, 3, 8,
		4, 4, 68, 69, 6, 4, -1, 0, 69, 73, 1, 0, 0, 0, 70, 71, 5, 26, 0, 0, 71,
		73, 6, 4, -1, 0, 72, 49, 1, 0, 0, 0, 72, 52, 1, 0, 0, 0, 72, 54, 1, 0,
		0, 0, 72, 56, 1, 0, 0, 0, 72, 58, 1, 0, 0, 0, 72, 60, 1, 0, 0, 0, 72, 62,
		1, 0, 0, 0, 72, 66, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 141, 1, 0, 0, 0,
		74, 75, 10, 15, 0, 0, 75, 76, 5, 14, 0, 0, 76, 77, 3, 8, 4, 16, 77, 78,
		6, 4, -1, 0, 78, 140, 1, 0, 0, 0, 79, 80, 10, 14, 0, 0, 80, 81, 5, 15,
		0, 0, 81, 82, 3, 8, 4, 15, 82, 83, 6, 4, -1, 0, 83, 140, 1, 0, 0, 0, 84,
		85, 10, 13, 0, 0, 85, 86, 5, 16, 0, 0, 86, 87, 3, 8, 4, 14, 87, 88, 6,
		4, -1, 0, 88, 140, 1, 0, 0, 0, 89, 90, 10, 12, 0, 0, 90, 91, 5, 12, 0,
		0, 91, 92, 3, 8, 4, 13, 92, 93, 6, 4, -1, 0, 93, 140, 1, 0, 0, 0, 94, 95,
		10, 11, 0, 0, 95, 96, 5, 13, 0, 0, 96, 97, 3, 8, 4, 12, 97, 98, 6, 4, -1,
		0, 98, 140, 1, 0, 0, 0, 99, 100, 10, 10, 0, 0, 100, 101, 5, 22, 0, 0, 101,
		102, 3, 8, 4, 11, 102, 103, 6, 4, -1, 0, 103, 140, 1, 0, 0, 0, 104, 105,
		10, 9, 0, 0, 105, 106, 5, 17, 0, 0, 106, 107, 3, 8, 4, 10, 107, 108, 6,
		4, -1, 0, 108, 140, 1, 0, 0, 0, 109, 110, 10, 8, 0, 0, 110, 111, 5, 21,
		0, 0, 111, 112, 3, 8, 4, 9, 112, 113, 6, 4, -1, 0, 113, 140, 1, 0, 0, 0,
		114, 115, 10, 7, 0, 0, 115, 116, 5, 18, 0, 0, 116, 117, 3, 8, 4, 8, 117,
		118, 6, 4, -1, 0, 118, 140, 1, 0, 0, 0, 119, 120, 10, 6, 0, 0, 120, 121,
		5, 19, 0, 0, 121, 122, 3, 8, 4, 7, 122, 123, 6, 4, -1, 0, 123, 140, 1,
		0, 0, 0, 124, 125, 10, 5, 0, 0, 125, 126, 5, 20, 0, 0, 126, 127, 3, 8,
		4, 6, 127, 128, 6, 4, -1, 0, 128, 140, 1, 0, 0, 0, 129, 130, 10, 3, 0,
		0, 130, 131, 5, 23, 0, 0, 131, 132, 3, 8, 4, 4, 132, 133, 6, 4, -1, 0,
		133, 140, 1, 0, 0, 0, 134, 135, 10, 2, 0, 0, 135, 136, 5, 24, 0, 0, 136,
		137, 3, 8, 4, 3, 137, 138, 6, 4, -1, 0, 138, 140, 1, 0, 0, 0, 139, 74,
		1, 0, 0, 0, 139, 79, 1, 0, 0, 0, 139, 84, 1, 0, 0, 0, 139, 89, 1, 0, 0,
		0, 139, 94, 1, 0, 0, 0, 139, 99, 1, 0, 0, 0, 139, 104, 1, 0, 0, 0, 139,
		109, 1, 0, 0, 0, 139, 114, 1, 0, 0, 0, 139, 119, 1, 0, 0, 0, 139, 124,
		1, 0, 0, 0, 139, 129, 1, 0, 0, 0, 139, 134, 1, 0, 0, 0, 140, 143, 1, 0,
		0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 9, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 144, 145, 5, 1, 0, 0, 145, 146, 5, 4, 0, 0, 146, 147,
		3, 8, 4, 0, 147, 148, 5, 5, 0, 0, 148, 149, 3, 12, 6, 0, 149, 150, 6, 5,
		-1, 0, 150, 11, 1, 0, 0, 0, 151, 152, 5, 3, 0, 0, 152, 13, 1, 0, 0, 0,
		5, 20, 34, 72, 139, 141,
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
	analizadorParserRVAR       = 2
	analizadorParserPUNTOCOMA  = 3
	analizadorParserPARA       = 4
	analizadorParserPARC       = 5
	analizadorParserENTERO     = 6
	analizadorParserFLOTANTE   = 7
	analizadorParserRTRUE      = 8
	analizadorParserRFALSE     = 9
	analizadorParserCHAR       = 10
	analizadorParserSTRING     = 11
	analizadorParserMAS        = 12
	analizadorParserMENOS      = 13
	analizadorParserMUL        = 14
	analizadorParserDIV        = 15
	analizadorParserMOD        = 16
	analizadorParserMAYOR      = 17
	analizadorParserMENOR      = 18
	analizadorParserMAYORIGUAL = 19
	analizadorParserMENORIGUAL = 20
	analizadorParserIGUALIGUAL = 21
	analizadorParserDISTINTO   = 22
	analizadorParserOR         = 23
	analizadorParserAND        = 24
	analizadorParserNOT        = 25
	analizadorParserID         = 26
	analizadorParserIGUAL      = 27
	analizadorParserCOMMENT    = 28
	analizadorParserWS         = 29
)

// analizadorParser rules.
const (
	analizadorParserRULE_start       = 0
	analizadorParserRULE_instruccion = 1
	analizadorParserRULE_declaracion = 2
	analizadorParserRULE_asignacion  = 3
	analizadorParserRULE_expresion   = 4
	analizadorParserRULE_imprimir    = 5
	analizadorParserRULE_finins      = 6
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<analizadorParserRPRINT)|(1<<analizadorParserRVAR)|(1<<analizadorParserID))) != 0 {
		{
			p.SetState(15)

			var _x = p.Instruccion()

			localctx.(*StartContext)._instruccion = _x
		}
		instrucciones = append(instrucciones, localctx.(*StartContext).Get_instruccion().GetNodo())

		p.SetState(22)
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

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	nodo         Abstract.Instruccion
	_imprimir    IImprimirContext
	_declaracion IDeclaracionContext
	_asignacion  IAsignacionContext
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

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

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

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case analizadorParserRPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)

			var _x = p.Imprimir()

			localctx.(*InstruccionContext)._imprimir = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_imprimir().GetNodo()

	case analizadorParserRVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_declaracion().GetNodo()

	case analizadorParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_asignacion().GetNodo()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RVAR returns the _RVAR token.
	Get_RVAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_RVAR sets the _RVAR token.
	Set_RVAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	nodo       Abstract.Instruccion
	_RVAR      antlr.Token
	_ID        antlr.Token
	_expresion IExpresionContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_RVAR() antlr.Token { return s._RVAR }

func (s *DeclaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclaracionContext) Set_RVAR(v antlr.Token) { s._RVAR = v }

func (s *DeclaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclaracionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclaracionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *DeclaracionContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *DeclaracionContext) RVAR() antlr.TerminalNode {
	return s.GetToken(analizadorParserRVAR, 0)
}

func (s *DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(analizadorParserID, 0)
}

func (s *DeclaracionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(analizadorParserIGUAL, 0)
}

func (s *DeclaracionContext) Expresion() IExpresionContext {
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

func (s *DeclaracionContext) Finins() IFininsContext {
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

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *analizadorParser) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, analizadorParserRULE_declaracion)

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
		p.SetState(36)

		var _m = p.Match(analizadorParserRVAR)

		localctx.(*DeclaracionContext)._RVAR = _m
	}
	{
		p.SetState(37)

		var _m = p.Match(analizadorParserID)

		localctx.(*DeclaracionContext)._ID = _m
	}
	{
		p.SetState(38)
		p.Match(analizadorParserIGUAL)
	}
	{
		p.SetState(39)

		var _x = p.expresion(0)

		localctx.(*DeclaracionContext)._expresion = _x
	}
	{
		p.SetState(40)
		p.Finins()
	}
	localctx.(*DeclaracionContext).nodo = Instrucciones.NewDeclaracion((func() string {
		if localctx.(*DeclaracionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*DeclaracionContext).Get_ID().GetText()
		}
	}()), localctx.(*DeclaracionContext).Get_expresion().GetNodo(), (func() int {
		if localctx.(*DeclaracionContext).Get_RVAR() == nil {
			return 0
		} else {
			return localctx.(*DeclaracionContext).Get_RVAR().GetLine()
		}
	}()), (func() int {
		if localctx.(*DeclaracionContext).Get_RVAR() == nil {
			return 0
		} else {
			return localctx.(*DeclaracionContext).Get_RVAR().GetColumn()
		}
	}()))

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	nodo       Abstract.Instruccion
	_ID        antlr.Token
	_expresion IExpresionContext
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignacionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *AsignacionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AsignacionContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *AsignacionContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(analizadorParserID, 0)
}

func (s *AsignacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(analizadorParserIGUAL, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
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

func (s *AsignacionContext) Finins() IFininsContext {
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

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *analizadorParser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, analizadorParserRULE_asignacion)

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
		p.SetState(43)

		var _m = p.Match(analizadorParserID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(44)
		p.Match(analizadorParserIGUAL)
	}
	{
		p.SetState(45)

		var _x = p.expresion(0)

		localctx.(*AsignacionContext)._expresion = _x
	}
	{
		p.SetState(46)
		p.Finins()
	}
	localctx.(*AsignacionContext).nodo = Instrucciones.NewAsignacion((func() string {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*AsignacionContext).Get_expresion().GetNodo(), (func() int {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetColumn()
		}
	}()))

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

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

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

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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
	_ID         antlr.Token
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

func (s *ExpresionContext) Get_ID() antlr.Token { return s._ID }

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

func (s *ExpresionContext) Set_ID(v antlr.Token) { s._ID = v }

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

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(analizadorParserID, 0)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, analizadorParserRULE_expresion, _p)

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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case analizadorParserENTERO:
		{
			p.SetState(50)

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
			p.SetState(52)

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
			p.SetState(54)

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
			p.SetState(56)

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
			p.SetState(58)

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
			p.SetState(60)

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
			p.SetState(62)

			var _m = p.Match(analizadorParserMENOS)

			localctx.(*ExpresionContext)._MENOS = _m
		}
		{
			p.SetState(63)

			var _x = p.expresion(16)

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
			p.SetState(66)

			var _m = p.Match(analizadorParserNOT)

			localctx.(*ExpresionContext)._NOT = _m
		}
		{
			p.SetState(67)

			var _x = p.expresion(4)

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

	case analizadorParserID:
		{
			p.SetState(70)

			var _m = p.Match(analizadorParserID)

			localctx.(*ExpresionContext)._ID = _m
		}
		localctx.(*ExpresionContext).nodo = Expresiones.NewIdentificador((func() string {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetColumn()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(75)

					var _m = p.Match(analizadorParserMUL)

					localctx.(*ExpresionContext)._MUL = _m
				}
				{
					p.SetState(76)

					var _x = p.expresion(16)

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
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(80)

					var _m = p.Match(analizadorParserDIV)

					localctx.(*ExpresionContext)._DIV = _m
				}
				{
					p.SetState(81)

					var _x = p.expresion(15)

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
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(85)

					var _m = p.Match(analizadorParserMOD)

					localctx.(*ExpresionContext)._MOD = _m
				}
				{
					p.SetState(86)

					var _x = p.expresion(14)

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
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(90)

					var _m = p.Match(analizadorParserMAS)

					localctx.(*ExpresionContext)._MAS = _m
				}
				{
					p.SetState(91)

					var _x = p.expresion(13)

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
				p.SetState(94)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(95)

					var _m = p.Match(analizadorParserMENOS)

					localctx.(*ExpresionContext)._MENOS = _m
				}
				{
					p.SetState(96)

					var _x = p.expresion(12)

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
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(100)

					var _m = p.Match(analizadorParserDISTINTO)

					localctx.(*ExpresionContext)._DISTINTO = _m
				}
				{
					p.SetState(101)

					var _x = p.expresion(11)

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
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(105)

					var _m = p.Match(analizadorParserMAYOR)

					localctx.(*ExpresionContext)._MAYOR = _m
				}
				{
					p.SetState(106)

					var _x = p.expresion(10)

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
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(110)

					var _m = p.Match(analizadorParserIGUALIGUAL)

					localctx.(*ExpresionContext)._IGUALIGUAL = _m
				}
				{
					p.SetState(111)

					var _x = p.expresion(9)

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
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(115)

					var _m = p.Match(analizadorParserMENOR)

					localctx.(*ExpresionContext)._MENOR = _m
				}
				{
					p.SetState(116)

					var _x = p.expresion(8)

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
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(120)

					var _m = p.Match(analizadorParserMAYORIGUAL)

					localctx.(*ExpresionContext)._MAYORIGUAL = _m
				}
				{
					p.SetState(121)

					var _x = p.expresion(7)

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
				p.SetState(124)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(125)

					var _m = p.Match(analizadorParserMENORIGUAL)

					localctx.(*ExpresionContext)._MENORIGUAL = _m
				}
				{
					p.SetState(126)

					var _x = p.expresion(6)

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
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(130)

					var _m = p.Match(analizadorParserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(131)

					var _x = p.expresion(4)

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
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(135)

					var _m = p.Match(analizadorParserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(136)

					var _x = p.expresion(3)

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
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 10, analizadorParserRULE_imprimir)

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
		p.SetState(144)

		var _m = p.Match(analizadorParserRPRINT)

		localctx.(*ImprimirContext)._RPRINT = _m
	}
	{
		p.SetState(145)
		p.Match(analizadorParserPARA)
	}
	{
		p.SetState(146)

		var _x = p.expresion(0)

		localctx.(*ImprimirContext)._expresion = _x
	}
	{
		p.SetState(147)
		p.Match(analizadorParserPARC)
	}
	{
		p.SetState(148)
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
	p.EnterRule(localctx, 12, analizadorParserRULE_finins)

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
		p.SetState(151)
		p.Match(analizadorParserPUNTOCOMA)
	}

	return localctx
}

func (p *analizadorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
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
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
