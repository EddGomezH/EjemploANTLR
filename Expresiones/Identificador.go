package Expresiones

import (
	"fmt"

	"github.com/emivnajera/TS"
)

type Identificador struct {
	Identificador string
	Fila          int
	Columna       int
	Tipo          TS.TIPO
}

func (this Identificador) Interpretar(tabla *TS.TablaSimbolos) interface{} {
	simbolo := tabla.GetTabla(this.Identificador)

	if simbolo == nil {
		return TS.Excepcion{"Semantico", "Variable No Existe", this.Fila, this.Columna}
	}

	this.Tipo = simbolo.(TS.Simbolo).GetTipo()

	fmt.Println("Valor: ", simbolo.(TS.Simbolo).Valor)

	return simbolo.(TS.Simbolo).Valor
}

func (this Identificador) GetTipo() TS.TIPO {
	return this.Tipo
}

func NewIdentificador(identificador string, fila int, columna int) Identificador {
	return Identificador{Identificador: identificador, Fila: fila, Columna: columna}
}
