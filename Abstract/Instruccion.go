package Abstract

import (
	"github.com/emivnajera/TS"
)

type Instruccion interface {
	GetTipo() TS.TIPO
	SetTipo(tipo TS.TIPO)
	Interpretar(table *TS.TablaSimbolos, Funciones *[]interface{}) interface{}
}
