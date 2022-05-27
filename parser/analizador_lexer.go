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
		"RPRINT", "RVAR", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE",
		"RTRUE", "RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD",
		"MAYOR", "MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO",
		"OR", "AND", "NOT", "ID", "IGUAL", "COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 185, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		4, 5, 77, 8, 5, 11, 5, 12, 5, 78, 1, 6, 4, 6, 82, 8, 6, 11, 6, 12, 6, 83,
		1, 6, 1, 6, 4, 6, 88, 8, 6, 11, 6, 12, 6, 89, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 107,
		8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 115, 8, 10, 10, 10,
		12, 10, 118, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 3, 25, 157,
		8, 25, 1, 25, 5, 25, 160, 8, 25, 10, 25, 12, 25, 163, 9, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 4, 27, 171, 8, 27, 11, 27, 12, 27, 172,
		3, 27, 175, 8, 27, 1, 27, 1, 27, 1, 28, 4, 28, 180, 8, 28, 11, 28, 12,
		28, 181, 1, 28, 1, 28, 0, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 1, 0, 6, 1, 0, 48, 57, 2, 0, 65, 90, 97,
		122, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32, 32, 196, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 65, 1, 0, 0, 0, 5, 69, 1, 0,
		0, 0, 7, 71, 1, 0, 0, 0, 9, 73, 1, 0, 0, 0, 11, 76, 1, 0, 0, 0, 13, 81,
		1, 0, 0, 0, 15, 91, 1, 0, 0, 0, 17, 96, 1, 0, 0, 0, 19, 102, 1, 0, 0, 0,
		21, 110, 1, 0, 0, 0, 23, 121, 1, 0, 0, 0, 25, 123, 1, 0, 0, 0, 27, 125,
		1, 0, 0, 0, 29, 127, 1, 0, 0, 0, 31, 129, 1, 0, 0, 0, 33, 131, 1, 0, 0,
		0, 35, 133, 1, 0, 0, 0, 37, 135, 1, 0, 0, 0, 39, 138, 1, 0, 0, 0, 41, 141,
		1, 0, 0, 0, 43, 144, 1, 0, 0, 0, 45, 147, 1, 0, 0, 0, 47, 150, 1, 0, 0,
		0, 49, 153, 1, 0, 0, 0, 51, 156, 1, 0, 0, 0, 53, 164, 1, 0, 0, 0, 55, 166,
		1, 0, 0, 0, 57, 179, 1, 0, 0, 0, 59, 60, 5, 112, 0, 0, 60, 61, 5, 114,
		0, 0, 61, 62, 5, 105, 0, 0, 62, 63, 5, 110, 0, 0, 63, 64, 5, 116, 0, 0,
		64, 2, 1, 0, 0, 0, 65, 66, 5, 118, 0, 0, 66, 67, 5, 97, 0, 0, 67, 68, 5,
		114, 0, 0, 68, 4, 1, 0, 0, 0, 69, 70, 5, 59, 0, 0, 70, 6, 1, 0, 0, 0, 71,
		72, 5, 40, 0, 0, 72, 8, 1, 0, 0, 0, 73, 74, 5, 41, 0, 0, 74, 10, 1, 0,
		0, 0, 75, 77, 7, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 12, 1, 0, 0, 0, 80, 82, 7, 0, 0, 0,
		81, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 5, 46, 0, 0, 86, 88, 7, 0, 0, 0, 87,
		86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 14, 1, 0, 0, 0, 91, 92, 5, 116, 0, 0, 92, 93, 5, 114, 0, 0, 93,
		94, 5, 117, 0, 0, 94, 95, 5, 101, 0, 0, 95, 16, 1, 0, 0, 0, 96, 97, 5,
		102, 0, 0, 97, 98, 5, 97, 0, 0, 98, 99, 5, 108, 0, 0, 99, 100, 5, 115,
		0, 0, 100, 101, 5, 101, 0, 0, 101, 18, 1, 0, 0, 0, 102, 106, 5, 39, 0,
		0, 103, 107, 7, 1, 0, 0, 104, 107, 3, 57, 28, 0, 105, 107, 7, 0, 0, 0,
		106, 103, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107,
		108, 1, 0, 0, 0, 108, 109, 5, 39, 0, 0, 109, 20, 1, 0, 0, 0, 110, 116,
		5, 34, 0, 0, 111, 115, 7, 1, 0, 0, 112, 115, 3, 57, 28, 0, 113, 115, 7,
		0, 0, 0, 114, 111, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0,
		0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		119, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 5, 34, 0, 0, 120, 22,
		1, 0, 0, 0, 121, 122, 5, 43, 0, 0, 122, 24, 1, 0, 0, 0, 123, 124, 5, 45,
		0, 0, 124, 26, 1, 0, 0, 0, 125, 126, 5, 42, 0, 0, 126, 28, 1, 0, 0, 0,
		127, 128, 5, 47, 0, 0, 128, 30, 1, 0, 0, 0, 129, 130, 5, 37, 0, 0, 130,
		32, 1, 0, 0, 0, 131, 132, 5, 62, 0, 0, 132, 34, 1, 0, 0, 0, 133, 134, 5,
		60, 0, 0, 134, 36, 1, 0, 0, 0, 135, 136, 5, 62, 0, 0, 136, 137, 5, 61,
		0, 0, 137, 38, 1, 0, 0, 0, 138, 139, 5, 60, 0, 0, 139, 140, 5, 61, 0, 0,
		140, 40, 1, 0, 0, 0, 141, 142, 5, 61, 0, 0, 142, 143, 5, 61, 0, 0, 143,
		42, 1, 0, 0, 0, 144, 145, 5, 33, 0, 0, 145, 146, 5, 61, 0, 0, 146, 44,
		1, 0, 0, 0, 147, 148, 5, 124, 0, 0, 148, 149, 5, 124, 0, 0, 149, 46, 1,
		0, 0, 0, 150, 151, 5, 38, 0, 0, 151, 152, 5, 38, 0, 0, 152, 48, 1, 0, 0,
		0, 153, 154, 5, 33, 0, 0, 154, 50, 1, 0, 0, 0, 155, 157, 7, 2, 0, 0, 156,
		155, 1, 0, 0, 0, 157, 161, 1, 0, 0, 0, 158, 160, 7, 3, 0, 0, 159, 158,
		1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 52, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 165, 5, 61, 0, 0,
		165, 54, 1, 0, 0, 0, 166, 167, 5, 47, 0, 0, 167, 168, 5, 47, 0, 0, 168,
		174, 1, 0, 0, 0, 169, 171, 8, 4, 0, 0, 170, 169, 1, 0, 0, 0, 171, 172,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0,
		0, 0, 174, 170, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0,
		176, 177, 6, 27, 0, 0, 177, 56, 1, 0, 0, 0, 178, 180, 7, 5, 0, 0, 179,
		178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 6, 28, 0, 0, 184, 58, 1, 0,
		0, 0, 13, 0, 78, 83, 89, 106, 114, 116, 156, 159, 161, 172, 174, 181, 1,
		6, 0, 0,
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
	analizadorLexerRVAR       = 2
	analizadorLexerPUNTOCOMA  = 3
	analizadorLexerPARA       = 4
	analizadorLexerPARC       = 5
	analizadorLexerENTERO     = 6
	analizadorLexerFLOTANTE   = 7
	analizadorLexerRTRUE      = 8
	analizadorLexerRFALSE     = 9
	analizadorLexerCHAR       = 10
	analizadorLexerSTRING     = 11
	analizadorLexerMAS        = 12
	analizadorLexerMENOS      = 13
	analizadorLexerMUL        = 14
	analizadorLexerDIV        = 15
	analizadorLexerMOD        = 16
	analizadorLexerMAYOR      = 17
	analizadorLexerMENOR      = 18
	analizadorLexerMAYORIGUAL = 19
	analizadorLexerMENORIGUAL = 20
	analizadorLexerIGUALIGUAL = 21
	analizadorLexerDISTINTO   = 22
	analizadorLexerOR         = 23
	analizadorLexerAND        = 24
	analizadorLexerNOT        = 25
	analizadorLexerID         = 26
	analizadorLexerIGUAL      = 27
	analizadorLexerCOMMENT    = 28
	analizadorLexerWS         = 29
)
