package admmemoria

import (
	"administracionmemoria/administrador"
	"fmt"
)

type Persona struct {
	nombre    string
	hij_mayor *Persona
	hij_menor *Persona
	xadre     *Persona
}

// CrearPersona devuelve un puntero a una nueva Persona, con el nombre y el padre/madre indicado (que podría ser nil)
func CrearPersona(nombre string, xadre *Persona) *Persona {
	if xadre != nil && xadre.hij_mayor != nil && xadre.hij_menor != nil {
		panic("En este modelo sólo permitimos hasta 2 hijos")
	}

	per := administrador.PedirMemoria[Persona]()
	(*per).nombre = nombre
	if xadre == nil {
		return per
	}
	(*per).xadre = xadre

	if xadre.hij_mayor == nil {
		xadre.hij_mayor = per
	} else {
		xadre.hij_menor = per
	}

	return per
}

// Imprimir imprime a todos los miembros de la familia
func (per *Persona) Imprimir() {
	if per == nil {
		return
	}
	fmt.Println(per.nombre)
	per.hij_mayor.Imprimir()
	per.hij_menor.Imprimir()
}

// Destruir libera la memoria (simulada) de esta Persona y todos sus descendientes
func (per *Persona) Destruir() {
	if per ==  nil {
	return //Si llegamos hasta aca, signfica que la cadena "murio" (es decir que no hay mas hijos)
	}
	per.hij_menor.Destruir() //Destruimos primero el hijo menor, porque ese puede llegar a no existir
	per.hij_mayor.Destruir() //El hijo mayor nos garantiza que avancemos en el arbol. 
	administrador.LiberarMemoria[Persona](per)

}
