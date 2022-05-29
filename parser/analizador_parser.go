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
import "reflect"

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
		"start", "instruccion", "declaracion", "asignacion", "parametro", "parametros",
		"funcion", "parametroll", "parametrolls", "llamada", "expresion", "imprimir",
		"finins",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 252, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 31, 8, 0,
		10, 0, 12, 0, 34, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 53, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 77, 8, 5,
		10, 5, 12, 5, 80, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 91, 8, 6, 10, 6, 12, 6, 94, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 108, 8, 6, 10, 6, 12,
		6, 111, 9, 6, 1, 6, 1, 6, 1, 6, 3, 6, 116, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 127, 8, 8, 10, 8, 12, 8, 130, 9, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 146, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 171, 8, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 238, 8, 10, 10, 10, 12, 10, 241, 9, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 0, 1, 20, 13, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 0, 270, 0, 26, 1, 0, 0, 0,
		2, 52, 1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 67, 1, 0, 0,
		0, 10, 70, 1, 0, 0, 0, 12, 115, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 120,
		1, 0, 0, 0, 18, 145, 1, 0, 0, 0, 20, 170, 1, 0, 0, 0, 22, 242, 1, 0, 0,
		0, 24, 249, 1, 0, 0, 0, 26, 32, 6, 0, -1, 0, 27, 28, 3, 2, 1, 0, 28, 29,
		6, 0, -1, 0, 29, 31, 1, 0, 0, 0, 30, 27, 1, 0, 0, 0, 31, 34, 1, 0, 0, 0,
		32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 32, 1,
		0, 0, 0, 35, 36, 6, 0, -1, 0, 36, 1, 1, 0, 0, 0, 37, 38, 3, 22, 11, 0,
		38, 39, 6, 1, -1, 0, 39, 53, 1, 0, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 6,
		1, -1, 0, 42, 53, 1, 0, 0, 0, 43, 44, 3, 6, 3, 0, 44, 45, 6, 1, -1, 0,
		45, 53, 1, 0, 0, 0, 46, 47, 3, 12, 6, 0, 47, 48, 6, 1, -1, 0, 48, 53, 1,
		0, 0, 0, 49, 50, 3, 18, 9, 0, 50, 51, 6, 1, -1, 0, 51, 53, 1, 0, 0, 0,
		52, 37, 1, 0, 0, 0, 52, 40, 1, 0, 0, 0, 52, 43, 1, 0, 0, 0, 52, 46, 1,
		0, 0, 0, 52, 49, 1, 0, 0, 0, 53, 3, 1, 0, 0, 0, 54, 55, 5, 2, 0, 0, 55,
		56, 5, 27, 0, 0, 56, 57, 5, 28, 0, 0, 57, 58, 3, 20, 10, 0, 58, 59, 3,
		24, 12, 0, 59, 60, 6, 2, -1, 0, 60, 5, 1, 0, 0, 0, 61, 62, 5, 27, 0, 0,
		62, 63, 5, 28, 0, 0, 63, 64, 3, 20, 10, 0, 64, 65, 3, 24, 12, 0, 65, 66,
		6, 3, -1, 0, 66, 7, 1, 0, 0, 0, 67, 68, 5, 27, 0, 0, 68, 69, 6, 4, -1,
		0, 69, 9, 1, 0, 0, 0, 70, 71, 3, 8, 4, 0, 71, 78, 6, 5, -1, 0, 72, 73,
		5, 31, 0, 0, 73, 74, 3, 8, 4, 0, 74, 75, 6, 5, -1, 0, 75, 77, 1, 0, 0,
		0, 76, 72, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79,
		1, 0, 0, 0, 79, 11, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 26, 0, 0,
		82, 83, 6, 6, -1, 0, 83, 84, 5, 27, 0, 0, 84, 85, 5, 4, 0, 0, 85, 86, 5,
		5, 0, 0, 86, 92, 5, 29, 0, 0, 87, 88, 3, 2, 1, 0, 88, 89, 6, 6, -1, 0,
		89, 91, 1, 0, 0, 0, 90, 87, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1,
		0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95,
		96, 5, 30, 0, 0, 96, 116, 6, 6, -1, 0, 97, 98, 5, 26, 0, 0, 98, 99, 6,
		6, -1, 0, 99, 100, 5, 27, 0, 0, 100, 101, 5, 4, 0, 0, 101, 102, 3, 10,
		5, 0, 102, 103, 5, 5, 0, 0, 103, 109, 5, 29, 0, 0, 104, 105, 3, 2, 1, 0,
		105, 106, 6, 6, -1, 0, 106, 108, 1, 0, 0, 0, 107, 104, 1, 0, 0, 0, 108,
		111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112,
		1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 5, 30, 0, 0, 113, 114, 6, 6,
		-1, 0, 114, 116, 1, 0, 0, 0, 115, 81, 1, 0, 0, 0, 115, 97, 1, 0, 0, 0,
		116, 13, 1, 0, 0, 0, 117, 118, 3, 20, 10, 0, 118, 119, 6, 7, -1, 0, 119,
		15, 1, 0, 0, 0, 120, 121, 3, 14, 7, 0, 121, 128, 6, 8, -1, 0, 122, 123,
		5, 31, 0, 0, 123, 124, 3, 14, 7, 0, 124, 125, 6, 8, -1, 0, 125, 127, 1,
		0, 0, 0, 126, 122, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0,
		0, 128, 129, 1, 0, 0, 0, 129, 17, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131,
		132, 5, 27, 0, 0, 132, 133, 6, 9, -1, 0, 133, 134, 5, 4, 0, 0, 134, 135,
		5, 5, 0, 0, 135, 136, 3, 24, 12, 0, 136, 137, 6, 9, -1, 0, 137, 146, 1,
		0, 0, 0, 138, 139, 5, 27, 0, 0, 139, 140, 5, 4, 0, 0, 140, 141, 3, 16,
		8, 0, 141, 142, 5, 5, 0, 0, 142, 143, 3, 24, 12, 0, 143, 144, 6, 9, -1,
		0, 144, 146, 1, 0, 0, 0, 145, 131, 1, 0, 0, 0, 145, 138, 1, 0, 0, 0, 146,
		19, 1, 0, 0, 0, 147, 148, 6, 10, -1, 0, 148, 149, 5, 6, 0, 0, 149, 171,
		6, 10, -1, 0, 150, 151, 5, 7, 0, 0, 151, 171, 6, 10, -1, 0, 152, 153, 5,
		8, 0, 0, 153, 171, 6, 10, -1, 0, 154, 155, 5, 9, 0, 0, 155, 171, 6, 10,
		-1, 0, 156, 157, 5, 10, 0, 0, 157, 171, 6, 10, -1, 0, 158, 159, 5, 11,
		0, 0, 159, 171, 6, 10, -1, 0, 160, 161, 5, 13, 0, 0, 161, 162, 3, 20, 10,
		16, 162, 163, 6, 10, -1, 0, 163, 171, 1, 0, 0, 0, 164, 165, 5, 25, 0, 0,
		165, 166, 3, 20, 10, 4, 166, 167, 6, 10, -1, 0, 167, 171, 1, 0, 0, 0, 168,
		169, 5, 27, 0, 0, 169, 171, 6, 10, -1, 0, 170, 147, 1, 0, 0, 0, 170, 150,
		1, 0, 0, 0, 170, 152, 1, 0, 0, 0, 170, 154, 1, 0, 0, 0, 170, 156, 1, 0,
		0, 0, 170, 158, 1, 0, 0, 0, 170, 160, 1, 0, 0, 0, 170, 164, 1, 0, 0, 0,
		170, 168, 1, 0, 0, 0, 171, 239, 1, 0, 0, 0, 172, 173, 10, 15, 0, 0, 173,
		174, 5, 14, 0, 0, 174, 175, 3, 20, 10, 16, 175, 176, 6, 10, -1, 0, 176,
		238, 1, 0, 0, 0, 177, 178, 10, 14, 0, 0, 178, 179, 5, 15, 0, 0, 179, 180,
		3, 20, 10, 15, 180, 181, 6, 10, -1, 0, 181, 238, 1, 0, 0, 0, 182, 183,
		10, 13, 0, 0, 183, 184, 5, 16, 0, 0, 184, 185, 3, 20, 10, 14, 185, 186,
		6, 10, -1, 0, 186, 238, 1, 0, 0, 0, 187, 188, 10, 12, 0, 0, 188, 189, 5,
		12, 0, 0, 189, 190, 3, 20, 10, 13, 190, 191, 6, 10, -1, 0, 191, 238, 1,
		0, 0, 0, 192, 193, 10, 11, 0, 0, 193, 194, 5, 13, 0, 0, 194, 195, 3, 20,
		10, 12, 195, 196, 6, 10, -1, 0, 196, 238, 1, 0, 0, 0, 197, 198, 10, 10,
		0, 0, 198, 199, 5, 22, 0, 0, 199, 200, 3, 20, 10, 11, 200, 201, 6, 10,
		-1, 0, 201, 238, 1, 0, 0, 0, 202, 203, 10, 9, 0, 0, 203, 204, 5, 17, 0,
		0, 204, 205, 3, 20, 10, 10, 205, 206, 6, 10, -1, 0, 206, 238, 1, 0, 0,
		0, 207, 208, 10, 8, 0, 0, 208, 209, 5, 21, 0, 0, 209, 210, 3, 20, 10, 9,
		210, 211, 6, 10, -1, 0, 211, 238, 1, 0, 0, 0, 212, 213, 10, 7, 0, 0, 213,
		214, 5, 18, 0, 0, 214, 215, 3, 20, 10, 8, 215, 216, 6, 10, -1, 0, 216,
		238, 1, 0, 0, 0, 217, 218, 10, 6, 0, 0, 218, 219, 5, 19, 0, 0, 219, 220,
		3, 20, 10, 7, 220, 221, 6, 10, -1, 0, 221, 238, 1, 0, 0, 0, 222, 223, 10,
		5, 0, 0, 223, 224, 5, 20, 0, 0, 224, 225, 3, 20, 10, 6, 225, 226, 6, 10,
		-1, 0, 226, 238, 1, 0, 0, 0, 227, 228, 10, 3, 0, 0, 228, 229, 5, 23, 0,
		0, 229, 230, 3, 20, 10, 4, 230, 231, 6, 10, -1, 0, 231, 238, 1, 0, 0, 0,
		232, 233, 10, 2, 0, 0, 233, 234, 5, 24, 0, 0, 234, 235, 3, 20, 10, 3, 235,
		236, 6, 10, -1, 0, 236, 238, 1, 0, 0, 0, 237, 172, 1, 0, 0, 0, 237, 177,
		1, 0, 0, 0, 237, 182, 1, 0, 0, 0, 237, 187, 1, 0, 0, 0, 237, 192, 1, 0,
		0, 0, 237, 197, 1, 0, 0, 0, 237, 202, 1, 0, 0, 0, 237, 207, 1, 0, 0, 0,
		237, 212, 1, 0, 0, 0, 237, 217, 1, 0, 0, 0, 237, 222, 1, 0, 0, 0, 237,
		227, 1, 0, 0, 0, 237, 232, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 21, 1, 0, 0, 0, 241, 239, 1, 0,
		0, 0, 242, 243, 5, 1, 0, 0, 243, 244, 5, 4, 0, 0, 244, 245, 3, 20, 10,
		0, 245, 246, 5, 5, 0, 0, 246, 247, 3, 24, 12, 0, 247, 248, 6, 11, -1, 0,
		248, 23, 1, 0, 0, 0, 249, 250, 5, 3, 0, 0, 250, 25, 1, 0, 0, 0, 11, 32,
		52, 78, 92, 109, 115, 128, 145, 170, 237, 239,
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
	analizadorParserRFUNC      = 26
	analizadorParserID         = 27
	analizadorParserIGUAL      = 28
	analizadorParserLLAVEA     = 29
	analizadorParserLLAVEC     = 30
	analizadorParserCOMA       = 31
	analizadorParserCOMMENT    = 32
	analizadorParserWS         = 33
)

// analizadorParser rules.
const (
	analizadorParserRULE_start        = 0
	analizadorParserRULE_instruccion  = 1
	analizadorParserRULE_declaracion  = 2
	analizadorParserRULE_asignacion   = 3
	analizadorParserRULE_parametro    = 4
	analizadorParserRULE_parametros   = 5
	analizadorParserRULE_funcion      = 6
	analizadorParserRULE_parametroll  = 7
	analizadorParserRULE_parametrolls = 8
	analizadorParserRULE_llamada      = 9
	analizadorParserRULE_expresion    = 10
	analizadorParserRULE_imprimir     = 11
	analizadorParserRULE_finins       = 12
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
	var Funciones []interface{}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<analizadorParserRPRINT)|(1<<analizadorParserRVAR)|(1<<analizadorParserRFUNC)|(1<<analizadorParserID))) != 0 {
		{
			p.SetState(27)

			var _x = p.Instruccion()

			localctx.(*StartContext)._instruccion = _x
		}
		instrucciones = append(instrucciones, localctx.(*StartContext).Get_instruccion().GetNodo())

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	for _, n := range instrucciones {
		if reflect.TypeOf(n).Name() == "Funcion" {
			Funciones = append(Funciones, n.(Instrucciones.Funcion))
		} else {
			n.Interpretar(&TSGlobal, &Funciones)
		}
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

	// Get_funcion returns the _funcion rule contexts.
	Get_funcion() IFuncionContext

	// Get_llamada returns the _llamada rule contexts.
	Get_llamada() ILlamadaContext

	// Set_imprimir sets the _imprimir rule contexts.
	Set_imprimir(IImprimirContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_funcion sets the _funcion rule contexts.
	Set_funcion(IFuncionContext)

	// Set_llamada sets the _llamada rule contexts.
	Set_llamada(ILlamadaContext)

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
	_funcion     IFuncionContext
	_llamada     ILlamadaContext
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

func (s *InstruccionContext) Get_funcion() IFuncionContext { return s._funcion }

func (s *InstruccionContext) Get_llamada() ILlamadaContext { return s._llamada }

func (s *InstruccionContext) Set_imprimir(v IImprimirContext) { s._imprimir = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_funcion(v IFuncionContext) { s._funcion = v }

func (s *InstruccionContext) Set_llamada(v ILlamadaContext) { s._llamada = v }

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

func (s *InstruccionContext) Funcion() IFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionContext)
}

func (s *InstruccionContext) Llamada() ILlamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaContext)
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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)

			var _x = p.Imprimir()

			localctx.(*InstruccionContext)._imprimir = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_imprimir().GetNodo()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_declaracion().GetNodo()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_asignacion().GetNodo()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)

			var _x = p.Funcion()

			localctx.(*InstruccionContext)._funcion = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_funcion().GetNodo()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(49)

			var _x = p.Llamada()

			localctx.(*InstruccionContext)._llamada = _x
		}
		localctx.(*InstruccionContext).nodo = localctx.(*InstruccionContext).Get_llamada().GetNodo()

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
		p.SetState(54)

		var _m = p.Match(analizadorParserRVAR)

		localctx.(*DeclaracionContext)._RVAR = _m
	}
	{
		p.SetState(55)

		var _m = p.Match(analizadorParserID)

		localctx.(*DeclaracionContext)._ID = _m
	}
	{
		p.SetState(56)
		p.Match(analizadorParserIGUAL)
	}
	{
		p.SetState(57)

		var _x = p.expresion(0)

		localctx.(*DeclaracionContext)._expresion = _x
	}
	{
		p.SetState(58)
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
		p.SetState(61)

		var _m = p.Match(analizadorParserID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(62)
		p.Match(analizadorParserIGUAL)
	}
	{
		p.SetState(63)

		var _x = p.expresion(0)

		localctx.(*AsignacionContext)._expresion = _x
	}
	{
		p.SetState(64)
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

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetId returns the id attribute.
	GetId() string

	// SetId sets the id attribute.
	SetId(string)

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     string
	_ID    antlr.Token
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_parametro
	return p
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) Get_ID() antlr.Token { return s._ID }

func (s *ParametroContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParametroContext) GetId() string { return s.id }

func (s *ParametroContext) SetId(v string) { s.id = v }

func (s *ParametroContext) ID() antlr.TerminalNode {
	return s.GetToken(analizadorParserID, 0)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (p *analizadorParser) Parametro() (localctx IParametroContext) {
	this := p
	_ = this

	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, analizadorParserRULE_parametro)

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
		p.SetState(67)

		var _m = p.Match(analizadorParserID)

		localctx.(*ParametroContext)._ID = _m
	}
	localctx.(*ParametroContext).id = (func() string {
		if localctx.(*ParametroContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ParametroContext).Get_ID().GetText()
		}
	}())

	return localctx
}

// IParametrosContext is an interface to support dynamic dispatch.
type IParametrosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_parametro returns the _parametro rule contexts.
	Get_parametro() IParametroContext

	// Set_parametro sets the _parametro rule contexts.
	Set_parametro(IParametroContext)

	// GetLista returns the lista attribute.
	GetLista() []string

	// SetLista sets the lista attribute.
	SetLista([]string)

	// IsParametrosContext differentiates from other interfaces.
	IsParametrosContext()
}

type ParametrosContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	lista      []string
	_parametro IParametroContext
}

func NewEmptyParametrosContext() *ParametrosContext {
	var p = new(ParametrosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_parametros
	return p
}

func (*ParametrosContext) IsParametrosContext() {}

func NewParametrosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosContext {
	var p = new(ParametrosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_parametros

	return p
}

func (s *ParametrosContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosContext) Get_parametro() IParametroContext { return s._parametro }

func (s *ParametrosContext) Set_parametro(v IParametroContext) { s._parametro = v }

func (s *ParametrosContext) GetLista() []string { return s.lista }

func (s *ParametrosContext) SetLista(v []string) { s.lista = v }

func (s *ParametrosContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *ParametrosContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
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

	return t.(IParametroContext)
}

func (s *ParametrosContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(analizadorParserCOMA)
}

func (s *ParametrosContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(analizadorParserCOMA, i)
}

func (s *ParametrosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterParametros(s)
	}
}

func (s *ParametrosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitParametros(s)
	}
}

func (p *analizadorParser) Parametros() (localctx IParametrosContext) {
	this := p
	_ = this

	localctx = NewParametrosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, analizadorParserRULE_parametros)
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
	{
		p.SetState(70)

		var _x = p.Parametro()

		localctx.(*ParametrosContext)._parametro = _x
	}
	localctx.(*ParametrosContext).lista = append(localctx.(*ParametrosContext).lista, localctx.(*ParametrosContext).Get_parametro().GetId())
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == analizadorParserCOMA {
		{
			p.SetState(72)
			p.Match(analizadorParserCOMA)
		}
		{
			p.SetState(73)

			var _x = p.Parametro()

			localctx.(*ParametrosContext)._parametro = _x
		}
		localctx.(*ParametrosContext).lista = append(localctx.(*ParametrosContext).lista, localctx.(*ParametrosContext).Get_parametro().GetId())

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncionContext is an interface to support dynamic dispatch.
type IFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RFUNC returns the _RFUNC token.
	Get_RFUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_RFUNC sets the _RFUNC token.
	Set_RFUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Get_parametros returns the _parametros rule contexts.
	Get_parametros() IParametrosContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// Set_parametros sets the _parametros rule contexts.
	Set_parametros(IParametrosContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsFuncionContext differentiates from other interfaces.
	IsFuncionContext()
}

type FuncionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	nodo         Abstract.Instruccion
	_RFUNC       antlr.Token
	_ID          antlr.Token
	_instruccion IInstruccionContext
	_parametros  IParametrosContext
}

func NewEmptyFuncionContext() *FuncionContext {
	var p = new(FuncionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_funcion
	return p
}

func (*FuncionContext) IsFuncionContext() {}

func NewFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionContext {
	var p = new(FuncionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_funcion

	return p
}

func (s *FuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionContext) Get_RFUNC() antlr.Token { return s._RFUNC }

func (s *FuncionContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncionContext) Set_RFUNC(v antlr.Token) { s._RFUNC = v }

func (s *FuncionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncionContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *FuncionContext) Get_parametros() IParametrosContext { return s._parametros }

func (s *FuncionContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *FuncionContext) Set_parametros(v IParametrosContext) { s._parametros = v }

func (s *FuncionContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *FuncionContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *FuncionContext) RFUNC() antlr.TerminalNode {
	return s.GetToken(analizadorParserRFUNC, 0)
}

func (s *FuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(analizadorParserID, 0)
}

func (s *FuncionContext) PARA() antlr.TerminalNode {
	return s.GetToken(analizadorParserPARA, 0)
}

func (s *FuncionContext) PARC() antlr.TerminalNode {
	return s.GetToken(analizadorParserPARC, 0)
}

func (s *FuncionContext) LLAVEA() antlr.TerminalNode {
	return s.GetToken(analizadorParserLLAVEA, 0)
}

func (s *FuncionContext) LLAVEC() antlr.TerminalNode {
	return s.GetToken(analizadorParserLLAVEC, 0)
}

func (s *FuncionContext) AllInstruccion() []IInstruccionContext {
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

func (s *FuncionContext) Instruccion(i int) IInstruccionContext {
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

func (s *FuncionContext) Parametros() IParametrosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrosContext)
}

func (s *FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterFuncion(s)
	}
}

func (s *FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitFuncion(s)
	}
}

func (p *analizadorParser) Funcion() (localctx IFuncionContext) {
	this := p
	_ = this

	localctx = NewFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, analizadorParserRULE_funcion)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)

			var _m = p.Match(analizadorParserRFUNC)

			localctx.(*FuncionContext)._RFUNC = _m
		}
		instrucciones := []Abstract.Instruccion{}
		parametros := []string{}
		{
			p.SetState(83)

			var _m = p.Match(analizadorParserID)

			localctx.(*FuncionContext)._ID = _m
		}
		{
			p.SetState(84)
			p.Match(analizadorParserPARA)
		}
		{
			p.SetState(85)
			p.Match(analizadorParserPARC)
		}
		{
			p.SetState(86)
			p.Match(analizadorParserLLAVEA)
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<analizadorParserRPRINT)|(1<<analizadorParserRVAR)|(1<<analizadorParserRFUNC)|(1<<analizadorParserID))) != 0 {
			{
				p.SetState(87)

				var _x = p.Instruccion()

				localctx.(*FuncionContext)._instruccion = _x
			}
			instrucciones = append(instrucciones, localctx.(*FuncionContext).Get_instruccion().GetNodo())

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(95)
			p.Match(analizadorParserLLAVEC)
		}
		localctx.(*FuncionContext).nodo = Instrucciones.NewFuncion((func() string {
			if localctx.(*FuncionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncionContext).Get_ID().GetText()
			}
		}()), instrucciones, parametros, (func() int {
			if localctx.(*FuncionContext).Get_RFUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncionContext).Get_RFUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncionContext).Get_RFUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncionContext).Get_RFUNC().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)

			var _m = p.Match(analizadorParserRFUNC)

			localctx.(*FuncionContext)._RFUNC = _m
		}
		instrucciones := []Abstract.Instruccion{}
		{
			p.SetState(99)

			var _m = p.Match(analizadorParserID)

			localctx.(*FuncionContext)._ID = _m
		}
		{
			p.SetState(100)
			p.Match(analizadorParserPARA)
		}
		{
			p.SetState(101)

			var _x = p.Parametros()

			localctx.(*FuncionContext)._parametros = _x
		}
		{
			p.SetState(102)
			p.Match(analizadorParserPARC)
		}
		{
			p.SetState(103)
			p.Match(analizadorParserLLAVEA)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<analizadorParserRPRINT)|(1<<analizadorParserRVAR)|(1<<analizadorParserRFUNC)|(1<<analizadorParserID))) != 0 {
			{
				p.SetState(104)

				var _x = p.Instruccion()

				localctx.(*FuncionContext)._instruccion = _x
			}
			instrucciones = append(instrucciones, localctx.(*FuncionContext).Get_instruccion().GetNodo())

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(analizadorParserLLAVEC)
		}
		localctx.(*FuncionContext).nodo = Instrucciones.NewFuncion((func() string {
			if localctx.(*FuncionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncionContext).Get_ID().GetText()
			}
		}()), instrucciones, localctx.(*FuncionContext).Get_parametros().GetLista(), (func() int {
			if localctx.(*FuncionContext).Get_RFUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncionContext).Get_RFUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncionContext).Get_RFUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncionContext).Get_RFUNC().GetColumn()
			}
		}()))

	}

	return localctx
}

// IParametrollContext is an interface to support dynamic dispatch.
type IParametrollContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsParametrollContext differentiates from other interfaces.
	IsParametrollContext()
}

type ParametrollContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	nodo       Abstract.Instruccion
	_expresion IExpresionContext
}

func NewEmptyParametrollContext() *ParametrollContext {
	var p = new(ParametrollContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_parametroll
	return p
}

func (*ParametrollContext) IsParametrollContext() {}

func NewParametrollContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrollContext {
	var p = new(ParametrollContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_parametroll

	return p
}

func (s *ParametrollContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrollContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ParametrollContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ParametrollContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *ParametrollContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *ParametrollContext) Expresion() IExpresionContext {
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

func (s *ParametrollContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrollContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrollContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterParametroll(s)
	}
}

func (s *ParametrollContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitParametroll(s)
	}
}

func (p *analizadorParser) Parametroll() (localctx IParametrollContext) {
	this := p
	_ = this

	localctx = NewParametrollContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, analizadorParserRULE_parametroll)

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

		var _x = p.expresion(0)

		localctx.(*ParametrollContext)._expresion = _x
	}
	localctx.(*ParametrollContext).nodo = localctx.(*ParametrollContext).Get_expresion().GetNodo()

	return localctx
}

// IParametrollsContext is an interface to support dynamic dispatch.
type IParametrollsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_parametroll returns the _parametroll rule contexts.
	Get_parametroll() IParametrollContext

	// Set_parametroll sets the _parametroll rule contexts.
	Set_parametroll(IParametrollContext)

	// GetLista returns the lista attribute.
	GetLista() []Abstract.Instruccion

	// SetLista sets the lista attribute.
	SetLista([]Abstract.Instruccion)

	// IsParametrollsContext differentiates from other interfaces.
	IsParametrollsContext()
}

type ParametrollsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        []Abstract.Instruccion
	_parametroll IParametrollContext
}

func NewEmptyParametrollsContext() *ParametrollsContext {
	var p = new(ParametrollsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_parametrolls
	return p
}

func (*ParametrollsContext) IsParametrollsContext() {}

func NewParametrollsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrollsContext {
	var p = new(ParametrollsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_parametrolls

	return p
}

func (s *ParametrollsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrollsContext) Get_parametroll() IParametrollContext { return s._parametroll }

func (s *ParametrollsContext) Set_parametroll(v IParametrollContext) { s._parametroll = v }

func (s *ParametrollsContext) GetLista() []Abstract.Instruccion { return s.lista }

func (s *ParametrollsContext) SetLista(v []Abstract.Instruccion) { s.lista = v }

func (s *ParametrollsContext) AllParametroll() []IParametrollContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametrollContext); ok {
			len++
		}
	}

	tst := make([]IParametrollContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametrollContext); ok {
			tst[i] = t.(IParametrollContext)
			i++
		}
	}

	return tst
}

func (s *ParametrollsContext) Parametroll(i int) IParametrollContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrollContext); ok {
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

	return t.(IParametrollContext)
}

func (s *ParametrollsContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(analizadorParserCOMA)
}

func (s *ParametrollsContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(analizadorParserCOMA, i)
}

func (s *ParametrollsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrollsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrollsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterParametrolls(s)
	}
}

func (s *ParametrollsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitParametrolls(s)
	}
}

func (p *analizadorParser) Parametrolls() (localctx IParametrollsContext) {
	this := p
	_ = this

	localctx = NewParametrollsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, analizadorParserRULE_parametrolls)
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
	{
		p.SetState(120)

		var _x = p.Parametroll()

		localctx.(*ParametrollsContext)._parametroll = _x
	}
	localctx.(*ParametrollsContext).lista = append(localctx.(*ParametrollsContext).lista, localctx.(*ParametrollsContext).Get_parametroll().GetNodo())
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == analizadorParserCOMA {
		{
			p.SetState(122)
			p.Match(analizadorParserCOMA)
		}
		{
			p.SetState(123)

			var _x = p.Parametroll()

			localctx.(*ParametrollsContext)._parametroll = _x
		}
		localctx.(*ParametrollsContext).lista = append(localctx.(*ParametrollsContext).lista, localctx.(*ParametrollsContext).Get_parametroll().GetNodo())

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILlamadaContext is an interface to support dynamic dispatch.
type ILlamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_parametrolls returns the _parametrolls rule contexts.
	Get_parametrolls() IParametrollsContext

	// Set_parametrolls sets the _parametrolls rule contexts.
	Set_parametrolls(IParametrollsContext)

	// GetNodo returns the nodo attribute.
	GetNodo() Abstract.Instruccion

	// SetNodo sets the nodo attribute.
	SetNodo(Abstract.Instruccion)

	// IsLlamadaContext differentiates from other interfaces.
	IsLlamadaContext()
}

type LlamadaContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	nodo          Abstract.Instruccion
	_ID           antlr.Token
	_parametrolls IParametrollsContext
}

func NewEmptyLlamadaContext() *LlamadaContext {
	var p = new(LlamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = analizadorParserRULE_llamada
	return p
}

func (*LlamadaContext) IsLlamadaContext() {}

func NewLlamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaContext {
	var p = new(LlamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = analizadorParserRULE_llamada

	return p
}

func (s *LlamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaContext) Get_ID() antlr.Token { return s._ID }

func (s *LlamadaContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *LlamadaContext) Get_parametrolls() IParametrollsContext { return s._parametrolls }

func (s *LlamadaContext) Set_parametrolls(v IParametrollsContext) { s._parametrolls = v }

func (s *LlamadaContext) GetNodo() Abstract.Instruccion { return s.nodo }

func (s *LlamadaContext) SetNodo(v Abstract.Instruccion) { s.nodo = v }

func (s *LlamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(analizadorParserID, 0)
}

func (s *LlamadaContext) PARA() antlr.TerminalNode {
	return s.GetToken(analizadorParserPARA, 0)
}

func (s *LlamadaContext) PARC() antlr.TerminalNode {
	return s.GetToken(analizadorParserPARC, 0)
}

func (s *LlamadaContext) Finins() IFininsContext {
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

func (s *LlamadaContext) Parametrolls() IParametrollsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametrollsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametrollsContext)
}

func (s *LlamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.EnterLlamada(s)
	}
}

func (s *LlamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(analizadorListener); ok {
		listenerT.ExitLlamada(s)
	}
}

func (p *analizadorParser) Llamada() (localctx ILlamadaContext) {
	this := p
	_ = this

	localctx = NewLlamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, analizadorParserRULE_llamada)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)

			var _m = p.Match(analizadorParserID)

			localctx.(*LlamadaContext)._ID = _m
		}
		parametros := []Abstract.Instruccion{}
		{
			p.SetState(133)
			p.Match(analizadorParserPARA)
		}
		{
			p.SetState(134)
			p.Match(analizadorParserPARC)
		}
		{
			p.SetState(135)
			p.Finins()
		}
		localctx.(*LlamadaContext).nodo = Instrucciones.NewLlamada((func() string {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetText()
			}
		}()), parametros, (func() int {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)

			var _m = p.Match(analizadorParserID)

			localctx.(*LlamadaContext)._ID = _m
		}
		{
			p.SetState(139)
			p.Match(analizadorParserPARA)
		}
		{
			p.SetState(140)

			var _x = p.Parametrolls()

			localctx.(*LlamadaContext)._parametrolls = _x
		}
		{
			p.SetState(141)
			p.Match(analizadorParserPARC)
		}
		{
			p.SetState(142)
			p.Finins()
		}
		localctx.(*LlamadaContext).nodo = Instrucciones.NewLlamada((func() string {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetText()
			}
		}()), localctx.(*LlamadaContext).Get_parametrolls().GetLista(), (func() int {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*LlamadaContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*LlamadaContext).Get_ID().GetColumn()
			}
		}()))

	}

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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, analizadorParserRULE_expresion, _p)

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
	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case analizadorParserENTERO:
		{
			p.SetState(148)

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
			p.SetState(150)

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
			p.SetState(152)

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
			p.SetState(154)

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
			p.SetState(156)

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
			p.SetState(158)

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
			p.SetState(160)

			var _m = p.Match(analizadorParserMENOS)

			localctx.(*ExpresionContext)._MENOS = _m
		}
		{
			p.SetState(161)

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
			p.SetState(164)

			var _m = p.Match(analizadorParserNOT)

			localctx.(*ExpresionContext)._NOT = _m
		}
		{
			p.SetState(165)

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
			p.SetState(168)

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
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).opi = _prevctx
				p.PushNewRecursionContext(localctx, _startState, analizadorParserRULE_expresion)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(173)

					var _m = p.Match(analizadorParserMUL)

					localctx.(*ExpresionContext)._MUL = _m
				}
				{
					p.SetState(174)

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
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(178)

					var _m = p.Match(analizadorParserDIV)

					localctx.(*ExpresionContext)._DIV = _m
				}
				{
					p.SetState(179)

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
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(183)

					var _m = p.Match(analizadorParserMOD)

					localctx.(*ExpresionContext)._MOD = _m
				}
				{
					p.SetState(184)

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
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(188)

					var _m = p.Match(analizadorParserMAS)

					localctx.(*ExpresionContext)._MAS = _m
				}
				{
					p.SetState(189)

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
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(193)

					var _m = p.Match(analizadorParserMENOS)

					localctx.(*ExpresionContext)._MENOS = _m
				}
				{
					p.SetState(194)

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
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(198)

					var _m = p.Match(analizadorParserDISTINTO)

					localctx.(*ExpresionContext)._DISTINTO = _m
				}
				{
					p.SetState(199)

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
				p.SetState(202)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(203)

					var _m = p.Match(analizadorParserMAYOR)

					localctx.(*ExpresionContext)._MAYOR = _m
				}
				{
					p.SetState(204)

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
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(208)

					var _m = p.Match(analizadorParserIGUALIGUAL)

					localctx.(*ExpresionContext)._IGUALIGUAL = _m
				}
				{
					p.SetState(209)

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
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(213)

					var _m = p.Match(analizadorParserMENOR)

					localctx.(*ExpresionContext)._MENOR = _m
				}
				{
					p.SetState(214)

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
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(218)

					var _m = p.Match(analizadorParserMAYORIGUAL)

					localctx.(*ExpresionContext)._MAYORIGUAL = _m
				}
				{
					p.SetState(219)

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
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(223)

					var _m = p.Match(analizadorParserMENORIGUAL)

					localctx.(*ExpresionContext)._MENORIGUAL = _m
				}
				{
					p.SetState(224)

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
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(228)

					var _m = p.Match(analizadorParserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(229)

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
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(233)

					var _m = p.Match(analizadorParserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(234)

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
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, analizadorParserRULE_imprimir)

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
		p.SetState(242)

		var _m = p.Match(analizadorParserRPRINT)

		localctx.(*ImprimirContext)._RPRINT = _m
	}
	{
		p.SetState(243)
		p.Match(analizadorParserPARA)
	}
	{
		p.SetState(244)

		var _x = p.expresion(0)

		localctx.(*ImprimirContext)._expresion = _x
	}
	{
		p.SetState(245)
		p.Match(analizadorParserPARC)
	}
	{
		p.SetState(246)
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
	p.EnterRule(localctx, 24, analizadorParserRULE_finins)

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
		p.SetState(249)
		p.Match(analizadorParserPUNTOCOMA)
	}

	return localctx
}

func (p *analizadorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
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
