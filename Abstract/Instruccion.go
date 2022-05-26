package Abstract

import (
	"github.com/emivnajera/TS"
)

type Instruccion interface {
	GetTipo() TS.TIPO
	Interpretar(table *TS.TablaSimbolos) interface{}
}
