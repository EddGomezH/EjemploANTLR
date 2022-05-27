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
		"RPRINT", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE", "RTRUE",
		"RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD", "MAYOR",
		"MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO", "OR",
		"AND", "NOT", "COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 164, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 4, 4, 67, 8, 4, 11, 4, 12, 4, 68, 1, 5, 4, 5, 72, 8, 5, 11, 5, 12,
		5, 73, 1, 5, 1, 5, 4, 5, 78, 8, 5, 11, 5, 12, 5, 79, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 97, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 105, 8, 9, 10,
		9, 12, 9, 108, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 4, 24, 150, 8, 24, 11, 24, 12, 24, 151, 3, 24, 154, 8, 24, 1, 24,
		1, 24, 1, 25, 4, 25, 159, 8, 25, 11, 25, 12, 25, 160, 1, 25, 1, 25, 0,
		0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 1, 0, 4, 1, 0,
		48, 57, 2, 0, 65, 90, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13,
		32, 32, 174, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 1, 53,
		1, 0, 0, 0, 3, 59, 1, 0, 0, 0, 5, 61, 1, 0, 0, 0, 7, 63, 1, 0, 0, 0, 9,
		66, 1, 0, 0, 0, 11, 71, 1, 0, 0, 0, 13, 81, 1, 0, 0, 0, 15, 86, 1, 0, 0,
		0, 17, 92, 1, 0, 0, 0, 19, 100, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 113,
		1, 0, 0, 0, 25, 115, 1, 0, 0, 0, 27, 117, 1, 0, 0, 0, 29, 119, 1, 0, 0,
		0, 31, 121, 1, 0, 0, 0, 33, 123, 1, 0, 0, 0, 35, 125, 1, 0, 0, 0, 37, 128,
		1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 134, 1, 0, 0, 0, 43, 137, 1, 0, 0,
		0, 45, 140, 1, 0, 0, 0, 47, 143, 1, 0, 0, 0, 49, 145, 1, 0, 0, 0, 51, 158,
		1, 0, 0, 0, 53, 54, 5, 112, 0, 0, 54, 55, 5, 114, 0, 0, 55, 56, 5, 105,
		0, 0, 56, 57, 5, 110, 0, 0, 57, 58, 5, 116, 0, 0, 58, 2, 1, 0, 0, 0, 59,
		60, 5, 59, 0, 0, 60, 4, 1, 0, 0, 0, 61, 62, 5, 40, 0, 0, 62, 6, 1, 0, 0,
		0, 63, 64, 5, 41, 0, 0, 64, 8, 1, 0, 0, 0, 65, 67, 7, 0, 0, 0, 66, 65,
		1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0,
		69, 10, 1, 0, 0, 0, 70, 72, 7, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 73, 1,
		0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75,
		77, 5, 46, 0, 0, 76, 78, 7, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 79, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82,
		5, 116, 0, 0, 82, 83, 5, 114, 0, 0, 83, 84, 5, 117, 0, 0, 84, 85, 5, 101,
		0, 0, 85, 14, 1, 0, 0, 0, 86, 87, 5, 102, 0, 0, 87, 88, 5, 97, 0, 0, 88,
		89, 5, 108, 0, 0, 89, 90, 5, 115, 0, 0, 90, 91, 5, 101, 0, 0, 91, 16, 1,
		0, 0, 0, 92, 96, 5, 39, 0, 0, 93, 97, 7, 1, 0, 0, 94, 97, 3, 51, 25, 0,
		95, 97, 7, 0, 0, 0, 96, 93, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 95, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 5, 39, 0, 0, 99, 18, 1, 0, 0, 0, 100,
		106, 5, 34, 0, 0, 101, 105, 7, 1, 0, 0, 102, 105, 3, 51, 25, 0, 103, 105,
		7, 0, 0, 0, 104, 101, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 103, 1, 0,
		0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		107, 109, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 5, 34, 0, 0, 110,
		20, 1, 0, 0, 0, 111, 112, 5, 43, 0, 0, 112, 22, 1, 0, 0, 0, 113, 114, 5,
		45, 0, 0, 114, 24, 1, 0, 0, 0, 115, 116, 5, 42, 0, 0, 116, 26, 1, 0, 0,
		0, 117, 118, 5, 47, 0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5, 37, 0, 0, 120,
		30, 1, 0, 0, 0, 121, 122, 5, 62, 0, 0, 122, 32, 1, 0, 0, 0, 123, 124, 5,
		60, 0, 0, 124, 34, 1, 0, 0, 0, 125, 126, 5, 62, 0, 0, 126, 127, 5, 61,
		0, 0, 127, 36, 1, 0, 0, 0, 128, 129, 5, 60, 0, 0, 129, 130, 5, 61, 0, 0,
		130, 38, 1, 0, 0, 0, 131, 132, 5, 61, 0, 0, 132, 133, 5, 61, 0, 0, 133,
		40, 1, 0, 0, 0, 134, 135, 5, 33, 0, 0, 135, 136, 5, 61, 0, 0, 136, 42,
		1, 0, 0, 0, 137, 138, 5, 124, 0, 0, 138, 139, 5, 124, 0, 0, 139, 44, 1,
		0, 0, 0, 140, 141, 5, 38, 0, 0, 141, 142, 5, 38, 0, 0, 142, 46, 1, 0, 0,
		0, 143, 144, 5, 33, 0, 0, 144, 48, 1, 0, 0, 0, 145, 146, 5, 47, 0, 0, 146,
		147, 5, 47, 0, 0, 147, 153, 1, 0, 0, 0, 148, 150, 8, 2, 0, 0, 149, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0,
		0, 0, 152, 154, 1, 0, 0, 0, 153, 149, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 156, 6, 24, 0, 0, 156, 50, 1, 0, 0, 0, 157,
		159, 7, 3, 0, 0, 158, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 158,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 6, 25,
		0, 0, 163, 52, 1, 0, 0, 0, 10, 0, 68, 73, 79, 96, 104, 106, 151, 153, 160,
		1, 6, 0, 0,
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
	analizadorLexerRPRINT     = 1
	analizadorLexerPUNTOCOMA  = 2
	analizadorLexerPARA       = 3
	analizadorLexerPARC       = 4
	analizadorLexerENTERO     = 5
	analizadorLexerFLOTANTE   = 6
	analizadorLexerRTRUE      = 7
	analizadorLexerRFALSE     = 8
	analizadorLexerCHAR       = 9
	analizadorLexerSTRING     = 10
	analizadorLexerMAS        = 11
	analizadorLexerMENOS      = 12
	analizadorLexerMUL        = 13
	analizadorLexerDIV        = 14
	analizadorLexerMOD        = 15
	analizadorLexerMAYOR      = 16
	analizadorLexerMENOR      = 17
	analizadorLexerMAYORIGUAL = 18
	analizadorLexerMENORIGUAL = 19
	analizadorLexerIGUALIGUAL = 20
	analizadorLexerDISTINTO   = 21
	analizadorLexerOR         = 22
	analizadorLexerAND        = 23
	analizadorLexerNOT        = 24
	analizadorLexerCOMMENT    = 25
	analizadorLexerWS         = 26
)
