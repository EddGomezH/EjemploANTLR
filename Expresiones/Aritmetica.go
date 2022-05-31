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

func (this Aritmetica) Interpretar(tabla *TS.TablaSimbolos, Funciones *[]interface{}) interface{} {
	izq := this.Op_izq.Interpretar(tabla, Funciones)
	if reflect.TypeOf(izq).Name() == "Excepcion" {
		return izq
	}

	this.Op_izq.SetTipo(obtenerTipo(izq))

	if this.Op_der != nil {
		der := this.Op_der.(Abstract.Instruccion).Interpretar(tabla, Funciones)
		this.Op_der.(Abstract.Instruccion).SetTipo(obtenerTipo(der))
		if reflect.TypeOf(der).Name() == "Excepcion" {
			return der
		}

		//SUMA
		if this.Operador == TS.MAS {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(obtenerTipo(izq), izq).(int) + obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) + obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return obtenerVal(obtenerTipo(izq), izq).(int) + boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.CADENA) {
				this.Tipo = TS.CADENA
				return fmt.Sprintf("%d", obtenerVal(obtenerTipo(izq), izq).(int)) + obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) + float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) + obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) + boolToFloat(obtenerVal(obtenerTipo(der), der).(bool))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.CADENA) {
				this.Tipo = TS.CADENA
				return fmt.Sprintf("%v", obtenerVal(obtenerTipo(izq), izq).(float64)) + obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) + obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return boolToFloat(obtenerVal(obtenerTipo(izq), izq).(bool)) + obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) + boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(obtenerTipo(izq), izq).(string) + fmt.Sprintf("%v", obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.CADENA
				return obtenerVal(obtenerTipo(izq), izq).(string) + fmt.Sprintf("%v", obtenerVal(obtenerTipo(der), der).(float64))
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.CADENA
				return obtenerVal(obtenerTipo(izq), izq).(string) + fmt.Sprintf("%v", obtenerVal(obtenerTipo(der), der).(bool))
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.CADENA) {
				this.Tipo = TS.CADENA
				return obtenerVal(obtenerTipo(izq), izq).(string) + obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CADENA) && (obtenerTipo(der) == TS.CARACTER) {
				this.Tipo = TS.CADENA
				return obtenerVal(obtenerTipo(izq), izq).(string) + obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CARACTER) && (obtenerTipo(der) == TS.CARACTER) {
				this.Tipo = TS.CADENA
				return obtenerVal(obtenerTipo(izq), izq).(string) + obtenerVal(obtenerTipo(der), der).(string)
			}

			if (obtenerTipo(izq) == TS.CARACTER) && (obtenerTipo(der) == TS.CADENA) {
				this.Tipo = TS.CADENA
				return obtenerVal(obtenerTipo(izq), izq).(string) + obtenerVal(obtenerTipo(der), der).(string)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Suma -", this.Fila, this.Columna}
		}

		if this.Operador == TS.MENOS {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(obtenerTipo(izq), izq).(int) - obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) - obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return obtenerVal(obtenerTipo(izq), izq).(int) - boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) - float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) - obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) - boolToFloat(obtenerVal(obtenerTipo(der), der).(bool))
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) - obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return boolToFloat(obtenerVal(obtenerTipo(izq), izq).(bool)) - obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.BOOLEAN) && (obtenerTipo(der) == TS.BOOLEAN) {
				this.Tipo = TS.ENTERO
				return boolToInt(obtenerVal(obtenerTipo(izq), izq).(bool)) - boolToInt(obtenerVal(obtenerTipo(der), der).(bool))
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Resta -", this.Fila, this.Columna}
		}

		if this.Operador == TS.POR {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.ENTERO
				return obtenerVal(obtenerTipo(izq), izq).(int) * obtenerVal(obtenerTipo(der), der).(int)
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) * obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) * float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) * obtenerVal(obtenerTipo(der), der).(float64)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Multiplicacion -", this.Fila, this.Columna}

		}

		if this.Operador == TS.DIV {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) / float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return float64(obtenerVal(obtenerTipo(izq), izq).(int)) / obtenerVal(obtenerTipo(der), der).(float64)
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) / float64(obtenerVal(obtenerTipo(der), der).(int))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return obtenerVal(obtenerTipo(izq), izq).(float64) / obtenerVal(obtenerTipo(der), der).(float64)
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Division -", this.Fila, this.Columna}
		}

		if this.Operador == TS.MOD {
			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return math.Mod(float64(obtenerVal(obtenerTipo(izq), izq).(int)), float64(obtenerVal(obtenerTipo(der), der).(int)))
			}

			if (obtenerTipo(izq) == TS.ENTERO) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return math.Mod(float64(obtenerVal(obtenerTipo(izq), izq).(int)), obtenerVal(obtenerTipo(der), der).(float64))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.ENTERO) {
				this.Tipo = TS.FLOAT
				return math.Mod(obtenerVal(obtenerTipo(izq), izq).(float64), float64(obtenerVal(obtenerTipo(der), der).(int)))
			}

			if (obtenerTipo(izq) == TS.FLOAT) && (obtenerTipo(der) == TS.FLOAT) {
				this.Tipo = TS.FLOAT
				return math.Mod(obtenerVal(obtenerTipo(izq), izq).(float64), obtenerVal(obtenerTipo(der), der).(float64))
			}

			return TS.Excepcion{"Semantico", "Tipo Erroneo Para Modulo -", this.Fila, this.Columna}
		}

		return TS.Excepcion{"Semantico", "Tipo Erroneo de Operacion", this.Fila, this.Columna}

	}

	if this.Operador == TS.MENOS {
		if obtenerTipo(izq) == TS.ENTERO {
			this.Tipo = TS.ENTERO
			return -obtenerVal(obtenerTipo(izq), izq).(int)
		}

		if obtenerTipo(izq) == TS.FLOAT {
			this.Tipo = TS.FLOAT
			return -obtenerVal(obtenerTipo(izq), izq).(float64)
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

func obtenerTipo(valor interface{}) TS.TIPO {
	if reflect.TypeOf(valor).Name() == "int" {
		return TS.ENTERO
	}

	if reflect.TypeOf(valor).Name() == "float64" {
		return TS.FLOAT
	}

	if reflect.TypeOf(valor).Name() == "string" {
		return TS.CADENA
	}

	if reflect.TypeOf(valor).Name() == "bool" {
		return TS.BOOLEAN
	}
	return TS.INSTRUCCION
}

func (this Aritmetica) SetTipo(tipo TS.TIPO) {
	this.Tipo = tipo
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
