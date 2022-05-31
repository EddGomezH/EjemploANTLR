package Expresiones

import (
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Logica struct {
	Operador TS.OperadorAritmetico
	Op_izq   Abstract.Instruccion
	Op_der   interface{}
	Fila     int
	Columna  int
	Tipo     TS.TIPO
}

func (this Logica) Interpretar(tabla *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	izq := this.Op_izq.Interpretar(tabla, Funciones)
	if reflect.TypeOf(izq).Name() == "Excepcion" {
		return izq
	}

	if this.Op_der != nil {
		der := this.Op_der.(Abstract.Instruccion).Interpretar(tabla, Funciones)
		if reflect.TypeOf(der).Name() == "Excepcion" {
			return der
		}
		//AND
		if this.Operador == TS.AND {
			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return obtenerVal(obtenerTipo(izq), izq).(bool) && obtenerVal(obtenerTipo(der), der).(bool)
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion And", this.Fila, this.Columna}
		}
		//OR
		if this.Operador == TS.OR {
			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return obtenerVal(obtenerTipo(izq), izq).(bool) || obtenerVal(obtenerTipo(der), der).(bool)
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Or", this.Fila, this.Columna}
		}
	}
	//NOT
	if this.Operador == TS.NOT {
		if obtenerTipo(izq) == TS.BOOLEAN {
			return !obtenerVal(obtenerTipo(izq), izq).(bool)
		}
		return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Not", this.Fila, this.Columna}
	}

	return nil
}

func (this Logica) GetTipo() TS.TIPO {
	return this.Tipo
}

func (this Logica) SetTipo(tipo TS.TIPO) {
	this.Tipo = tipo
}

func NewLogica(operador TS.OperadorAritmetico, op_izq Abstract.Instruccion, op_der interface{}, fila int, columna int) Logica {
	return Logica{Operador: operador, Op_izq: op_izq, Op_der: op_der, Fila: fila, Columna: columna, Tipo: TS.BOOLEAN}
}
