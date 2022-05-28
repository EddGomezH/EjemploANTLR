package Instrucciones

import (
	"fmt"
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Asignacion struct {
	Identificador string
	Expresion     Abstract.Instruccion
	Fila          int
	Columna       int
	tipo          TS.TIPO
}

func (this Asignacion) Interpretar(table *TS.TablaSimbolos) interface{} {
	value := this.Expresion.Interpretar(table)
	if reflect.TypeOf(value).Name() == "Excepcion" {
		return value
	}
	this.tipo = this.Expresion.GetTipo()

	result := table.ActualizarTabla(this.Identificador, value, this.Fila, this.Columna)

	if reflect.TypeOf(result).Name() == "Excepcion" {
		fmt.Println(result)
		return value
	}

	return nil
}

func (this Asignacion) GetTipo() TS.TIPO {
	return this.tipo
}

func NewAsignacion(identificador string, expresion Abstract.Instruccion, fila int, columna int) Asignacion {
	return Asignacion{Identificador: identificador, Expresion: expresion, Fila: fila, Columna: columna}
}
