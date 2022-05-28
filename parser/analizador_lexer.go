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
		"'=='", "'!='", "'||'", "'&&'", "'!'", "'func'", "", "'='", "'{'", "'}'",
		"','",
	}
	staticData.symbolicNames = []string{
		"", "RPRINT", "RVAR", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE",
		"RTRUE", "RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD",
		"MAYOR", "MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO",
		"OR", "AND", "NOT", "RFUNC", "ID", "IGUAL", "LLAVEA", "LLAVEC", "COMA",
		"COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"RPRINT", "RVAR", "PUNTOCOMA", "PARA", "PARC", "ENTERO", "FLOTANTE",
		"RTRUE", "RFALSE", "CHAR", "STRING", "MAS", "MENOS", "MUL", "DIV", "MOD",
		"MAYOR", "MENOR", "MAYORIGUAL", "MENORIGUAL", "IGUALIGUAL", "DISTINTO",
		"OR", "AND", "NOT", "RFUNC", "ID", "IGUAL", "LLAVEA", "LLAVEC", "COMA",
		"COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 204, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 4, 5, 85, 8, 5, 11,
		5, 12, 5, 86, 1, 6, 4, 6, 90, 8, 6, 11, 6, 12, 6, 91, 1, 6, 1, 6, 4, 6,
		96, 8, 6, 11, 6, 12, 6, 97, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 115, 8, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 123, 8, 10, 10, 10, 12, 10, 126,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 3, 26, 170, 8, 26, 1, 26, 5, 26, 173, 8, 26, 10, 26, 12, 26, 176,
		9, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 4, 31, 190, 8, 31, 11, 31, 12, 31, 191, 3, 31, 194, 8,
		31, 1, 31, 1, 31, 1, 32, 4, 32, 199, 8, 32, 11, 32, 12, 32, 200, 1, 32,
		1, 32, 0, 0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 1, 0, 6, 1, 0, 48,
		57, 2, 0, 65, 90, 97, 122, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13, 32,
		32, 215, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 1, 67, 1, 0, 0, 0, 3,
		73, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 79, 1, 0, 0, 0, 9, 81, 1, 0, 0, 0,
		11, 84, 1, 0, 0, 0, 13, 89, 1, 0, 0, 0, 15, 99, 1, 0, 0, 0, 17, 104, 1,
		0, 0, 0, 19, 110, 1, 0, 0, 0, 21, 118, 1, 0, 0, 0, 23, 129, 1, 0, 0, 0,
		25, 131, 1, 0, 0, 0, 27, 133, 1, 0, 0, 0, 29, 135, 1, 0, 0, 0, 31, 137,
		1, 0, 0, 0, 33, 139, 1, 0, 0, 0, 35, 141, 1, 0, 0, 0, 37, 143, 1, 0, 0,
		0, 39, 146, 1, 0, 0, 0, 41, 149, 1, 0, 0, 0, 43, 152, 1, 0, 0, 0, 45, 155,
		1, 0, 0, 0, 47, 158, 1, 0, 0, 0, 49, 161, 1, 0, 0, 0, 51, 163, 1, 0, 0,
		0, 53, 169, 1, 0, 0, 0, 55, 177, 1, 0, 0, 0, 57, 179, 1, 0, 0, 0, 59, 181,
		1, 0, 0, 0, 61, 183, 1, 0, 0, 0, 63, 185, 1, 0, 0, 0, 65, 198, 1, 0, 0,
		0, 67, 68, 5, 112, 0, 0, 68, 69, 5, 114, 0, 0, 69, 70, 5, 105, 0, 0, 70,
		71, 5, 110, 0, 0, 71, 72, 5, 116, 0, 0, 72, 2, 1, 0, 0, 0, 73, 74, 5, 118,
		0, 0, 74, 75, 5, 97, 0, 0, 75, 76, 5, 114, 0, 0, 76, 4, 1, 0, 0, 0, 77,
		78, 5, 59, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 40, 0, 0, 80, 8, 1, 0, 0,
		0, 81, 82, 5, 41, 0, 0, 82, 10, 1, 0, 0, 0, 83, 85, 7, 0, 0, 0, 84, 83,
		1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0,
		87, 12, 1, 0, 0, 0, 88, 90, 7, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		95, 5, 46, 0, 0, 94, 96, 7, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0,
		0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 14, 1, 0, 0, 0, 99, 100,
		5, 116, 0, 0, 100, 101, 5, 114, 0, 0, 101, 102, 5, 117, 0, 0, 102, 103,
		5, 101, 0, 0, 103, 16, 1, 0, 0, 0, 104, 105, 5, 102, 0, 0, 105, 106, 5,
		97, 0, 0, 106, 107, 5, 108, 0, 0, 107, 108, 5, 115, 0, 0, 108, 109, 5,
		101, 0, 0, 109, 18, 1, 0, 0, 0, 110, 114, 5, 39, 0, 0, 111, 115, 7, 1,
		0, 0, 112, 115, 3, 65, 32, 0, 113, 115, 7, 0, 0, 0, 114, 111, 1, 0, 0,
		0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		117, 5, 39, 0, 0, 117, 20, 1, 0, 0, 0, 118, 124, 5, 34, 0, 0, 119, 123,
		7, 1, 0, 0, 120, 123, 3, 65, 32, 0, 121, 123, 7, 0, 0, 0, 122, 119, 1,
		0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0,
		0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 127, 128, 5, 34, 0, 0, 128, 22, 1, 0, 0, 0, 129, 130,
		5, 43, 0, 0, 130, 24, 1, 0, 0, 0, 131, 132, 5, 45, 0, 0, 132, 26, 1, 0,
		0, 0, 133, 134, 5, 42, 0, 0, 134, 28, 1, 0, 0, 0, 135, 136, 5, 47, 0, 0,
		136, 30, 1, 0, 0, 0, 137, 138, 5, 37, 0, 0, 138, 32, 1, 0, 0, 0, 139, 140,
		5, 62, 0, 0, 140, 34, 1, 0, 0, 0, 141, 142, 5, 60, 0, 0, 142, 36, 1, 0,
		0, 0, 143, 144, 5, 62, 0, 0, 144, 145, 5, 61, 0, 0, 145, 38, 1, 0, 0, 0,
		146, 147, 5, 60, 0, 0, 147, 148, 5, 61, 0, 0, 148, 40, 1, 0, 0, 0, 149,
		150, 5, 61, 0, 0, 150, 151, 5, 61, 0, 0, 151, 42, 1, 0, 0, 0, 152, 153,
		5, 33, 0, 0, 153, 154, 5, 61, 0, 0, 154, 44, 1, 0, 0, 0, 155, 156, 5, 124,
		0, 0, 156, 157, 5, 124, 0, 0, 157, 46, 1, 0, 0, 0, 158, 159, 5, 38, 0,
		0, 159, 160, 5, 38, 0, 0, 160, 48, 1, 0, 0, 0, 161, 162, 5, 33, 0, 0, 162,
		50, 1, 0, 0, 0, 163, 164, 5, 102, 0, 0, 164, 165, 5, 117, 0, 0, 165, 166,
		5, 110, 0, 0, 166, 167, 5, 99, 0, 0, 167, 52, 1, 0, 0, 0, 168, 170, 7,
		2, 0, 0, 169, 168, 1, 0, 0, 0, 170, 174, 1, 0, 0, 0, 171, 173, 7, 3, 0,
		0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174,
		175, 1, 0, 0, 0, 175, 54, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 178, 5,
		61, 0, 0, 178, 56, 1, 0, 0, 0, 179, 180, 5, 123, 0, 0, 180, 58, 1, 0, 0,
		0, 181, 182, 5, 125, 0, 0, 182, 60, 1, 0, 0, 0, 183, 184, 5, 44, 0, 0,
		184, 62, 1, 0, 0, 0, 185, 186, 5, 47, 0, 0, 186, 187, 5, 47, 0, 0, 187,
		193, 1, 0, 0, 0, 188, 190, 8, 4, 0, 0, 189, 188, 1, 0, 0, 0, 190, 191,
		1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0,
		0, 0, 193, 189, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0,
		195, 196, 6, 31, 0, 0, 196, 64, 1, 0, 0, 0, 197, 199, 7, 5, 0, 0, 198,
		197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201,
		1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 6, 32, 0, 0, 203, 66, 1, 0,
		0, 0, 13, 0, 86, 91, 97, 114, 122, 124, 169, 172, 174, 191, 193, 200, 1,
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
	analizadorLexerRFUNC      = 26
	analizadorLexerID         = 27
	analizadorLexerIGUAL      = 28
	analizadorLexerLLAVEA     = 29
	analizadorLexerLLAVEC     = 30
	analizadorLexerCOMA       = 31
	analizadorLexerCOMMENT    = 32
	analizadorLexerWS         = 33
)
