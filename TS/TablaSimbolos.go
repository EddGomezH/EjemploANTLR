package TS

type TablaSimbolos struct {
	Tabla    []Simbolo
	anterior *TablaSimbolos
}

func (this *TablaSimbolos) SetTabla(sim Simbolo) interface{} {
	for _, n := range this.Tabla {
		if sim.Id == n.Id {
			return Excepcion{"Semantico", "Variable Ya Existe", sim.Fila, sim.Columna}
		}
	}
	this.Tabla = append(this.Tabla, sim)
	return 0
}

func (this TablaSimbolos) GetTabla(id string) interface{} {
	for _, n := range this.Tabla {
		if id == n.Id {
			return n
		}
	}
	this = *this.anterior
	return nil
}

func (this *TablaSimbolos) ActualizarTabla(id string, value interface{}) interface{} {
	contador := 0
	for this.Tabla != nil {
		for _, n := range this.Tabla {
			if id == n.Id {
				this.Tabla[contador].Valor = value
				return nil
			}
			contador++
		}
		this = this.anterior
	}
	return nil
}

func NewTabla(Anterior TablaSimbolos) TablaSimbolos {
	return TablaSimbolos{anterior: &Anterior}
}
