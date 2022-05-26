// Code generated from analizador.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type analizadorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var analizadorlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func analizadorlexerLexerInit() {
	staticData := &analizadorlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'print'", "';'", "'('", "')'", "", "", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE",
		"RFALSE", "CHAR", "STRING", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE",
		"RFALSE", "CHAR", "STRING", "COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 102, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 4, 4, 39, 8, 4, 11, 4, 12, 4, 40, 1, 5, 4,
		5, 44, 8, 5, 11, 5, 12, 5, 45, 1, 5, 1, 5, 4, 5, 50, 8, 5, 11, 5, 12, 5,
		51, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 69, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 77, 8, 9, 10, 9, 12, 9, 80, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 4, 10, 88, 8, 10, 11, 10, 12, 10, 89, 3, 10, 92, 8, 10, 1, 10, 1,
		10, 1, 11, 4, 11, 97, 8, 11, 11, 11, 12, 11, 98, 1, 11, 1, 11, 0, 0, 12,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 1, 0, 4, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 2, 0, 10, 10, 13,
		13, 3, 0, 9, 10, 13, 13, 32, 32, 112, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 31, 1,
		0, 0, 0, 5, 33, 1, 0, 0, 0, 7, 35, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0, 11, 43,
		1, 0, 0, 0, 13, 53, 1, 0, 0, 0, 15, 58, 1, 0, 0, 0, 17, 64, 1, 0, 0, 0,
		19, 72, 1, 0, 0, 0, 21, 83, 1, 0, 0, 0, 23, 96, 1, 0, 0, 0, 25, 26, 5,
		112, 0, 0, 26, 27, 5, 114, 0, 0, 27, 28, 5, 105, 0, 0, 28, 29, 5, 110,
		0, 0, 29, 30, 5, 116, 0, 0, 30, 2, 1, 0, 0, 0, 31, 32, 5, 59, 0, 0, 32,
		4, 1, 0, 0, 0, 33, 34, 5, 40, 0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 41, 0,
		0, 36, 8, 1, 0, 0, 0, 37, 39, 7, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1,
		0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 10, 1, 0, 0, 0, 42,
		44, 7, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0,
		0, 45, 46, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 5, 46, 0, 0, 48, 50,
		7, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5,
		114, 0, 0, 55, 56, 5, 117, 0, 0, 56, 57, 5, 101, 0, 0, 57, 14, 1, 0, 0,
		0, 58, 59, 5, 102, 0, 0, 59, 60, 5, 97, 0, 0, 60, 61, 5, 108, 0, 0, 61,
		62, 5, 115, 0, 0, 62, 63, 5, 101, 0, 0, 63, 16, 1, 0, 0, 0, 64, 68, 5,
		39, 0, 0, 65, 69, 7, 1, 0, 0, 66, 69, 3, 23, 11, 0, 67, 69, 7, 0, 0, 0,
		68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1,
		0, 0, 0, 70, 71, 5, 39, 0, 0, 71, 18, 1, 0, 0, 0, 72, 78, 5, 34, 0, 0,
		73, 77, 7, 1, 0, 0, 74, 77, 3, 23, 11, 0, 75, 77, 7, 0, 0, 0, 76, 73, 1,
		0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78,
		76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0,
		0, 81, 82, 5, 34, 0, 0, 82, 20, 1, 0, 0, 0, 83, 84, 5, 47, 0, 0, 84, 85,
		5, 47, 0, 0, 85, 91, 1, 0, 0, 0, 86, 88, 8, 2, 0, 0, 87, 86, 1, 0, 0, 0,
		88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1,
		0, 0, 0, 91, 87, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		94, 6, 10, 0, 0, 94, 22, 1, 0, 0, 0, 95, 97, 7, 3, 0, 0, 96, 95, 1, 0,
		0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100,
		1, 0, 0, 0, 100, 101, 6, 11, 0, 0, 101, 24, 1, 0, 0, 0, 10, 0, 40, 45,
		51, 68, 76, 78, 89, 91, 98, 1, 6, 0, 0,
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

// analizadorLexerInit initializes any static state used to implement analizadorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewanalizadorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AnalizadorLexerInit() {
	staticData := &analizadorlexerLexerStaticData
	staticData.once.Do(analizadorlexerLexerInit)
}

// NewanalizadorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewanalizadorLexer(input antlr.CharStream) *analizadorLexer {
	AnalizadorLexerInit()
	l := new(analizadorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &analizadorlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "analizador.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// analizadorLexer tokens.
const (
	analizadorLexerRPRINT    = 1
	analizadorLexerPUNTOCOMA = 2
	analizadorLexerPARA      = 3
	analizadorLexerPARC      = 4
	analizadorLexerENTERO    = 5
	analizadorLexerFLOTANTE  = 6
	analizadorLexerRTRUE     = 7
	analizadorLexerRFALSE    = 8
	analizadorLexerCHAR      = 9
	analizadorLexerSTRING    = 10
	analizadorLexerCOMMENT   = 11
	analizadorLexerWS        = 12
)
