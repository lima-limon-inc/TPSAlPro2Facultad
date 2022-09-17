package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func nodoCrear[T any](datoArgumento T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = datoArgumento
	nodo.prox = nil
	return nodo
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() == true {
		panic("La cola esta vacia")
	}
	return cola.primero.dato

}
func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodoNuevo := nodoCrear(elem) //Creo un nodoNuevo con el dato pasado
	switch cola.EstaVacia() {
	case true:
		cola.primero = nodoNuevo
	case false:
		cola.ultimo.prox = nodoNuevo //Hago que el ultimo elemento de la lista apunte a ese nuevo elemento
	}
	cola.ultimo = nodoNuevo //Hago que el puntero ultimo apunte a la direccion de nodoNuevo. nodoNuevo es un puntero a un nodo

}

//func (cola *colaEnlazada[T]) Desencolar() T {}
