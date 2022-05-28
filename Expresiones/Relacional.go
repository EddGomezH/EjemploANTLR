package Expresiones

import (
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Relacional struct {
	Operador TS.OperadorAritmetico
	Op_izq   Abstract.Instruccion
	Op_der   interface{}
	Fila     int
	Columna  int
	Tipo     TS.TIPO
}

func (this Relacional) Interpretar(tabla *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	izq := this.Op_izq.Interpretar(tabla, Funciones)
	if reflect.TypeOf(izq).Name() == "Excepcion" {
		return izq
	}

	if this.Op_der != nil {
		der := this.Op_der.(Abstract.Instruccion).Interpretar(tabla, Funciones)
		if reflect.TypeOf(der).Name() == "Excepcion" {
			return der
		}

		//MENORQUE
		if this.Operador == TS.MENORQUE {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) < obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) < obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) < float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) < obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) < boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Menor Que", this.Fila, this.Columna}
		}

		//MAYORQUE
		if this.Operador == TS.MAYORQUE {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) > obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) > obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) > float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) > obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) > boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Mayor Que", this.Fila, this.Columna}
		}

		//MAYORQUE
		if this.Operador == TS.MAYORIGUAL {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) >= obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) >= obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) >= float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) >= obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) >= boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Mayor Igual Que", this.Fila, this.Columna}
		}

		//MENORIGUAL
		if this.Operador == TS.MENORIGUAL {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) <= obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) <= obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) <= float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) <= obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) <= boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Menor Igual Que", this.Fila, this.Columna}
		}

		//DIFERENTE
		if this.Operador == TS.DIFERENTE {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) != float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(bool) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool)
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CARACTER) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CARACTER) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CARACTER) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CARACTER) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) != obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Diferente Que", this.Fila, this.Columna}
		}

		//IGUAL
		if this.Operador == TS.IGUALIGUAL {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) == float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(bool) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool)
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CARACTER) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CARACTER) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CARACTER) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CARACTER) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) == obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Diferente Que", this.Fila, this.Columna}
		}

		return TS.Excepcion{"Semantico", "Tipo Erroneo de Operacion", this.Fila, this.Columna}

	}

	return nil
}

func (this Relacional) GetTipo() TS.TIPO {
	return this.Tipo
}

func NewRelacional(operador TS.OperadorAritmetico, op_izq Abstract.Instruccion, op_der interface{}, fila int, columna int) Relacional {
	return Relacional{Operador: operador, Op_izq: op_izq, Op_der: op_der, Fila: fila, Columna: columna, Tipo: TS.BOOLEAN}
}
