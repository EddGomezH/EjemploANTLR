package Abstract

import (
	"github.com/emivnajera/TS"
)

type Instruccion interface {
	Interpretar(table *TS.TablaSimbolos) interface{}
}
