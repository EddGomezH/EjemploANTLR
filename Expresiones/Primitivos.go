package Expresiones

import (
	"strconv"

	"github.com/emivnajera/TS"
)

type Primitivos struct {
	valor   string
	tipo    TS.TIPO
	fila    int
	columna int
}

func (this Primitivos) Interpretar(table *TS.TablaSimbolos) interface{} {
	if this.tipo == TS.ENTERO {
		intvar, err := strconv.Atoi(this.valor)
		if err == nil {
			return intvar
		}
	}
	if this.tipo == TS.FLOAT {
		floatvar, err := strconv.ParseFloat(this.valor, 64)
		if err == nil {
			return floatvar
		}
	}
	if this.tipo == TS.BOOLEAN {
		boolvar, err := strconv.ParseBool(this.valor)
		if err == nil {
			return boolvar
		}
	}
	if this.tipo == TS.CARACTER {
		return this.valor
	}
	if this.tipo == TS.CADENA {
		return this.valor
	}

	return TS.Excepcion{"Semantico", "Tipo de Dato Invalido", this.fila, this.columna}
}

func NewPrimitivo(Valor string, Tipo TS.TIPO, Fila int, Columna int) Primitivos {
	return Primitivos{valor: Valor, tipo: Tipo, fila: Fila, columna: Columna}
}
