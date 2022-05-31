package Instrucciones

import (
	"fmt"
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Llamada struct {
	Identificador string
	Parametros    []Abstract.Instruccion
	Fila          int
	Columna       int
	Tipo          TS.TIPO
}

func (this Llamada) Interpretar(table *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	var result Funcion
	encontrado := false
	for _, n := range *Funciones {
		if n.(Funcion).Identificador == this.Identificador {
			result = n.(Funcion)
			encontrado = true
		}
	}

	if !encontrado {
		return TS.Excepcion{"Semantico", "No Se Encontro la Funcion", this.Fila, this.Columna}
	}
	nuevaTabla := TS.NewTabla(table)
	if len(result.Parametros) == len(this.Parametros) {
		contador := 0
		for _, e := range this.Parametros {
			resultExpresion := e.Interpretar(table, Funciones)
			if reflect.TypeOf(resultExpresion).Name() == "Excepcion" {
				return resultExpresion
			}

			simbolo := TS.NewSimbolo(result.Parametros[contador], this.Fila, this.Columna, resultExpresion, e.GetTipo())
			resultTabla := nuevaTabla.SetTabla(simbolo)
			if reflect.TypeOf(resultTabla).Name() == "Excepcion" {
				return resultTabla
			}
			contador++
		}
		value := result.Interpretar(&nuevaTabla, Funciones)

		return value

	} else {
		fmt.Println(TS.Excepcion{"Semantico", "Numero Incorrecto de Atributos", this.Fila, this.Columna})
	}
	return nil
}

func (this Llamada) GetTipo() TS.TIPO {
	return this.Tipo
}

func (this Llamada) SetTipo(tipo TS.TIPO) {
	this.Tipo = tipo
}

func NewLlamada(identificador string, parametros []Abstract.Instruccion, fila int, columna int) Llamada {
	return Llamada{Identificador: identificador, Parametros: parametros, Fila: fila, Columna: columna}
}
