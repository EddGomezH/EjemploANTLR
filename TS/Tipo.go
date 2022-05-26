package TS

type TIPO int64

const (
	ENTERO      = 1
	FLOAT       = 2
	BOOLEAN     = 3
	CARACTER    = 4
	CADENA      = 5
	INSTRUCCION = 6
)

type OperadorAritmetico int64

const (
	MAS    = 1
	MENOS  = 2
	POR    = 3
	DIV    = 4
	POT    = 5
	MOD    = 7
	UMENOS = 8
)
