package Expresiones

import (
	"strconv"

	"github.com/emivnajera/TS"
)

type Primitivos struct {
	Valor   string
	Tipo    TS.TIPO
	Fila    int
	Columna int
}

func (this Primitivos) Interpretar(table *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	if this.Tipo == TS.ENTERO {
		intvar, err := strconv.Atoi(this.Valor)
		if err == nil {
			return intvar
		}
	}
	if this.Tipo == TS.FLOAT {
		floatvar, err := strconv.ParseFloat(this.Valor, 64)
		if err == nil {
			return floatvar
		}
	}
	if this.Tipo == TS.BOOLEAN {
		boolvar, err := strconv.ParseBool(this.Valor)
		if err == nil {
			return boolvar
		}
	}
	if this.Tipo == TS.CARACTER {
		return this.Valor
	}
	if this.Tipo == TS.CADENA {
		return this.Valor
	}

	return TS.Excepcion{"Semantico", "Tipo de Dato Invalido", this.Fila, this.Columna}
}

func (this Primitivos) GetTipo() TS.TIPO {
	return this.Tipo
}

func NewPrimitivo(valor string, tipo TS.TIPO, fila int, columna int) Primitivos {
	return Primitivos{Valor: valor, Tipo: tipo, Fila: fila, Columna: columna}
}
