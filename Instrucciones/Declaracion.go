package Instrucciones

import (
	"fmt"
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Declaracion struct {
	Identificador string
	Expresion     Abstract.Instruccion
	Fila          int
	Columna       int
}

func (this Declaracion) Interpretar(table *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	value := this.Expresion.Interpretar(table, Funciones)
	tipo := this.Expresion.GetTipo()

	if reflect.TypeOf(value).Name() == "Excepcion" {
		return value
	}

	simbolo := TS.NewSimbolo(this.Identificador, this.Fila, this.Columna, value, tipo)

	result := table.SetTabla(simbolo)

	if reflect.TypeOf(result).Name() == "Excepcion" {
		return result
	}

	return 0
}

func (this Declaracion) GetTipo() TS.TIPO {
	return TS.INSTRUCCION
}

func (this Declaracion) SetTipo(tipo TS.TIPO) {
	fmt.Println("No se debe usar")
}

func NewDeclaracion(identificador string, expresion Abstract.Instruccion, fila int, columna int) Declaracion {
	return Declaracion{Identificador: identificador, Expresion: expresion, Fila: fila, Columna: columna}
}
