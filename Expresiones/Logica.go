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

func (this Logica) Interpretar(tabla *TS.TablaSimbolos) interface{} {
	izq := this.Op_izq.Interpretar(tabla)
	if reflect.TypeOf(izq).Name() == "Excepcion" {
		return izq
	}

	if this.Op_der != nil {
		der := this.Op_der.(Abstract.Instruccion).Interpretar(tabla)
		if reflect.TypeOf(der).Name() == "Excepcion" {
			return der
		}
		//AND
		if this.Operador == TS.AND {
			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(bool) && obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool)
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion And", this.Fila, this.Columna}
		}
		//OR
		if this.Operador == TS.OR {
			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(bool) || obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool)
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Or", this.Fila, this.Columna}
		}
	}
	//NOT
	if this.Operador == TS.NOT {
		if this.Op_izq.GetTipo() == TS.BOOLEAN {
			return !obtenerVal(this.Op_izq.GetTipo(), izq).(bool)
		}
		return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Not", this.Fila, this.Columna}
	}

	return nil
}

func (this Logica) GetTipo() TS.TIPO {
	return this.Tipo
}

func NewLogica(operador TS.OperadorAritmetico, op_izq Abstract.Instruccion, op_der interface{}, fila int, columna int) Logica {
	return Logica{Operador: operador, Op_izq: op_izq, Op_der: op_der, Fila: fila, Columna: columna, Tipo: TS.BOOLEAN}
}
