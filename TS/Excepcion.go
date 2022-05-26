package TS

import "strconv"

type Excepcion struct {
	Tipo        string
	Descripcion string
	Fila        int
	Columna     int
}

func (this Excepcion) ToString() string {
	cad := this.Tipo + " - " + this.Descripcion + ", En " + strconv.Itoa(this.Fila) + ":" + strconv.Itoa(this.Columna)
	return cad
}
