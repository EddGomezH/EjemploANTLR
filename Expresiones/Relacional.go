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
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(int) < obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) < obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) < float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) < obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) < boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Menor Que", this.Fila, this.Columna}
		}

		//MAYORQUE
		if this.Operador == TS.MAYORQUE {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(int) > obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) > obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) > float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) > obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) > boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Mayor Que", this.Fila, this.Columna}
		}

		//MAYORQUE
		if this.Operador == TS.MAYORIGUAL {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(int) >= obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) >= obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) >= float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) >= obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) >= boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Mayor Igual Que", this.Fila, this.Columna}
		}

		//MENORIGUAL
		if this.Operador == TS.MENORIGUAL {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(int) <= obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) <= obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) <= float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) <= obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) <= boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}
			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Menor Igual Que", this.Fila, this.Columna}
		}

		//DIFERENTE
		if this.Operador == TS.DIFERENTE {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(int) != obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) != obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) != float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) != obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return obtenerVal(obtenerTipo(izq), izq).(bool) != obtenerVal(obtenerTipo(der), der).(bool)
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.CADENA) {
				return obtenerVal(obtenerTipo(izq), izq).(string) != obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.CARACTER) {
				return obtenerVal(obtenerTipo(izq), izq).(string) != obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CARACTER) && (obtenerTipo(der) == TS.CARACTER) {
				return obtenerVal(obtenerTipo(izq), izq).(string) != obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CARACTER) && (obtenerTipo(der) == TS.CADENA) {
				return obtenerVal(obtenerTipo(izq), izq).(string) != obtenerVal(obtenerTipo(der), der).(string)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo para Operacion Diferente Que", this.Fila, this.Columna}
		}

		//IGUAL
		if this.Operador == TS.IGUALIGUAL {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(int) == obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) == obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) == float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				return obtenerVal(obtenerTipo(izq), izq).(float64) == obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				return obtenerVal(obtenerTipo(izq), izq).(bool) == obtenerVal(obtenerTipo(der), der).(bool)
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.CADENA) {
				return obtenerVal(obtenerTipo(izq), izq).(string) == obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.CARACTER) {
				return obtenerVal(obtenerTipo(izq), izq).(string) == obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CARACTER) && (obtenerTipo(der) == TS.CARACTER) {
				return obtenerVal(obtenerTipo(izq), izq).(string) == obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CARACTER) && (obtenerTipo(der) == TS.CADENA) {
				return obtenerVal(obtenerTipo(izq), izq).(string) == obtenerVal(obtenerTipo(der), der).(string)
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

func (this Relacional) SetTipo(tipo TS.TIPO) {
	this.Tipo = tipo
}

func NewRelacional(operador TS.OperadorAritmetico, op_izq Abstract.Instruccion, op_der interface{}, fila int, columna int) Relacional {
	return Relacional{Operador: operador, Op_izq: op_izq, Op_der: op_der, Fila: fila, Columna: columna, Tipo: TS.BOOLEAN}
}
