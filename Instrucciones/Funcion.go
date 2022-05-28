package Instrucciones

import (
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Funcion struct {
	Identificador string
	Instrucciones []Abstract.Instruccion
	Parametros    []string
	Fila          int
	Columna       int
	tipo          TS.TIPO
}

func (this Funcion) Interpretar(table *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	nuevaTabla := TS.NewTabla(table)
	for _, n := range this.Instrucciones {
		n.Interpretar(&nuevaTabla, Funciones)

		if reflect.TypeOf(n).Name() == "Excepcion" {
			return n
		}
	}
	return nil
}

func (this Funcion) GetTipo() TS.TIPO {
	return this.tipo
}

func NewFuncion(identificador string, instrucciones []Abstract.Instruccion, parametros []string, fila int, columna int) Funcion {
	return Funcion{Identificador: identificador, Instrucciones: instrucciones, Parametros: parametros, Fila: fila, Columna: columna}
}
