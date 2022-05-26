package Expresiones

import (
	"fmt"
	"math"
	"reflect"

	"github.com/emivnajera/Abstract"
	"github.com/emivnajera/TS"
)

type Aritmetica struct {
	Operador TS.OperadorAritmetico
	Op_izq   Abstract.Instruccion
	Op_der   interface{}
	Fila     int
	Columna  int
	Tipo     TS.TIPO
}

func (this Aritmetica) Interpretar(tabla *TS.TablaSimbolos) interface{} {
	izq := this.Op_izq.Interpretar(tabla)
	if reflect.TypeOf(izq).Name() == "Excepcion" {
		return izq
	}

	if this.Op_der != nil {
		der := this.Op_der.(Abstract.Instruccion).Interpretar(tabla)
		if reflect.TypeOf(der).Name() == "Excepcion" {
			return der
		}

		//SUMA
		if this.Operador == TS.MAS {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) + boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				this.Tipo = TS.CADENA
				return fmt.Sprintf("%d", obtenerVal(this.Op_izq.GetTipo(), izq).(int)) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) + float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) + boolToFloat(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				this.Tipo = TS.CADENA
				return fmt.Sprintf("%v", obtenerVal(this.Op_izq.GetTipo(), izq).(float64)) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return boolToFloat(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) + boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + fmt.Sprintf("%v", obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.CADENA
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + fmt.Sprintf("%v", obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64))
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.CADENA
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + fmt.Sprintf("%v", obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				this.Tipo = TS.CADENA
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CADENA) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CARACTER) {
				this.Tipo = TS.CADENA
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CARACTER) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CARACTER) {
				this.Tipo = TS.CADENA
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			if (this.Op_izq.GetTipo() == TS.CARACTER) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.CADENA) {
				this.Tipo = TS.CADENA
				return obtenerVal(this.Op_izq.GetTipo(), izq).(string) + obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(string)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Suma -", this.Fila, this.Columna}
		}

		if this.Operador == TS.MENOS {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) - obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) - obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) - boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) - float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) - obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) - boolToFloat(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) - obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return boolToFloat(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) - obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.BOOLEAN) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(this.Op_izq.GetTipo(), izq).(bool)) - boolToInt(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(bool))
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Resta -", this.Fila, this.Columna}
		}

		if this.Operador == TS.POR {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(this.Op_izq.GetTipo(), izq).(int) * obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) * obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) * float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) * obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Multiplicacion -", this.Fila, this.Columna}

		}

		if this.Operador == TS.DIV {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) / float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)) / obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) / float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(this.Op_izq.GetTipo(), izq).(float64) / obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Division -", this.Fila, this.Columna}
		}

		if this.Operador == TS.MOD {
			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return math.Mod(float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)), float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)))
			}

			if (this.Op_izq.GetTipo() == TS.ENTERO) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return math.Mod(float64(obtenerVal(this.Op_izq.GetTipo(), izq).(int)), obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return math.Mod(obtenerVal(this.Op_izq.GetTipo(), izq).(float64), float64(obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(int)))
			}

			if (this.Op_izq.GetTipo() == TS.FLOAT) && (this.Op_der.(Abstract.Instruccion).GetTipo() == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return math.Mod(obtenerVal(this.Op_izq.GetTipo(), izq).(float64), obtenerVal(this.Op_der.(Abstract.Instruccion).GetTipo(), der).(float64))
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Modulo -", this.Fila, this.Columna}
		}

		return TS.Excepcion{"Semantico", "Tipo Erroneo de Operacion", this.Fila, this.Columna}

	}

	if this.Operador == TS.MENOS {
		if this.Op_izq.GetTipo() == TS.ENTERO {
			this.Tipo = TS.ENTERO
			return -obtenerVal(this.Op_izq.GetTipo(), izq).(int)
		}

		if this.Op_izq.GetTipo() == TS.FLOAT {
			this.Tipo = TS.FLOAT
			return -obtenerVal(this.Op_izq.GetTipo(), izq).(float64)
		}

		return TS.Excepcion{"Semantico", "Tipo Erroneo Para Negacion Unaria", this.Fila, this.Columna}
	}

	return nil
}

func (this Aritmetica) GetTipo() TS.TIPO {
	return this.Tipo
}

func obtenerVal(tipo TS.TIPO, val interface{}) interface{} {
	if tipo == TS.ENTERO {
		return val.(int)
	}
	if tipo == TS.FLOAT {
		return val.(float64)
	}
	if tipo == TS.BOOLEAN {
		return val.(bool)
	}
	return fmt.Sprintf("%v", val)
}

func boolToInt(valor bool) int {
	if valor == true {
		return 1
	} else {
		return 0
	}
}

func boolToFloat(valor bool) float64 {
	if valor == true {
		return float64(1)
	} else {
		return float64(0)
	}
}

func NewAritmetica(operador TS.OperadorAritmetico, op_izq Abstract.Instruccion, op_der interface{}, fila int, columna int) Aritmetica {
	return Aritmetica{Operador: operador, Op_izq: op_izq, Op_der: op_der, Fila: fila, Columna: columna}
}
