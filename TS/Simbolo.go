package TS

type Simbolo struct {
	Id      string
	Fila    int
	Columna int
	Valor   interface{}
	Tipo    TIPO
}

func NewSimbolo(id string, fila int, columna int, valor interface{}, tipo TIPO) Simbolo {
	return Simbolo{Id: id, Fila: fila, Columna: columna, Valor: valor, Tipo: tipo}
}

func (this Simbolo) GetTipo() TIPO {
	return this.Tipo
}

func (this Simbolo) GetValor() interface{} {
	return this.Valor
}
